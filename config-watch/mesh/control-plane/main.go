// control-plane is an xDS server that turns Consul KV into live Envoy config.
//
// It serves Listeners and Clusters to both sidecars over ADS (gRPC), and runs a
// Consul blocking-query watch on the "mesh/" key prefix. Whenever a key changes
// it rebuilds the snapshot and pushes it — the Envoys hot-reload with no restart.
package main

import (
	"context"
	"log"
	"net"
	"os"
	"os/signal"
	"strconv"
	"syscall"
	"time"

	clusterservice "github.com/envoyproxy/go-control-plane/envoy/service/cluster/v3"
	discovery "github.com/envoyproxy/go-control-plane/envoy/service/discovery/v3"
	listenerservice "github.com/envoyproxy/go-control-plane/envoy/service/listener/v3"
	cachev3 "github.com/envoyproxy/go-control-plane/pkg/cache/v3"
	serverv3 "github.com/envoyproxy/go-control-plane/pkg/server/v3"
	"github.com/hashicorp/consul/api"
	"google.golang.org/grpc"
)

const (
	keyPrefix       = "mesh/"
	keyTag          = "mesh/tag"
	defaultTag      = "v1"
	defaultUpstream = "server-envoy"
	defaultMeshPort = 8443
)

func main() {
	consulAddr := envOr("CONSUL_ADDRESS", "consul:8500")
	grpcAddr := envOr("XDS_LISTEN", ":18000")
	upstreamHost := envOr("UPSTREAM_HOST", defaultUpstream)

	cfg := api.DefaultConfig()
	cfg.Address = consulAddr
	consul, err := api.NewClient(cfg)
	if err != nil {
		log.Fatalf("create Consul client failed: %v", err)
	}

	ctx, cancel := signal.NotifyContext(context.Background(), os.Interrupt, syscall.SIGTERM)
	defer cancel()

	snapCache := cachev3.NewSnapshotCache(true, cachev3.IDHash{}, xdsLogger{})

	// Serve ADS over gRPC.
	srv := serverv3.NewServer(ctx, snapCache, nil)
	grpcServer := grpc.NewServer()
	discovery.RegisterAggregatedDiscoveryServiceServer(grpcServer, srv)
	clusterservice.RegisterClusterDiscoveryServiceServer(grpcServer, srv)
	listenerservice.RegisterListenerDiscoveryServiceServer(grpcServer, srv)

	lis, err := net.Listen("tcp", grpcAddr)
	if err != nil {
		log.Fatalf("listen %s failed: %v", grpcAddr, err)
	}
	go func() {
		log.Printf("xDS server listening on %s (consul=%s)", grpcAddr, consulAddr)
		if err := grpcServer.Serve(lis); err != nil {
			log.Printf("grpc server stopped: %v", err)
		}
	}()

	watchConsul(ctx, consul, snapCache, upstreamHost)

	grpcServer.GracefulStop()
	log.Println("control-plane stopped")
}

// watchConsul blocks until ctx is cancelled, pushing a fresh snapshot on every
// change to the mesh/ prefix. This is the same Consul blocking-query pattern as
// the repo's existing KV watcher, repurposed to drive Envoy config.
func watchConsul(ctx context.Context, consul *api.Client, snapCache cachev3.SnapshotCache, upstreamHost string) {
	seedTag(consul)

	var waitIndex uint64
	for {
		select {
		case <-ctx.Done():
			return
		default:
		}

		pairs, meta, err := consul.KV().List(keyPrefix, &api.QueryOptions{
			WaitIndex: waitIndex,
			WaitTime:  5 * time.Minute,
		})
		if err != nil {
			log.Printf("consul read failed: %v", err)
			select {
			case <-ctx.Done():
				return
			case <-time.After(3 * time.Second):
				continue
			}
		}
		if meta == nil || meta.LastIndex == waitIndex {
			continue // timeout, nothing changed
		}
		waitIndex = meta.LastIndex

		cfg := MeshConfig{
			Tag:          valueOr(pairs, keyTag, defaultTag),
			UpstreamHost: upstreamHost,
			MeshPort:     defaultMeshPort,
		}

		version := strconv.FormatUint(meta.LastIndex, 10)
		snaps, err := buildSnapshots(version, cfg)
		if err != nil {
			log.Printf("build snapshot failed: %v", err)
			continue
		}

		for node, snap := range snaps {
			if err := snap.Consistent(); err != nil {
				log.Printf("inconsistent snapshot for %s: %v", node, err)
				continue
			}
			if err := snapCache.SetSnapshot(ctx, node, snap); err != nil {
				log.Printf("set snapshot for %s failed: %v", node, err)
			}
		}
		log.Printf("pushed snapshot version=%s tag=%q upstream=%s", version, cfg.Tag, upstreamHost)
	}
}

// seedTag writes a default value so the demo works on a fresh Consul.
func seedTag(consul *api.Client) {
	pair, _, err := consul.KV().Get(keyTag, nil)
	if err != nil {
		log.Printf("seed check failed: %v", err)
		return
	}
	if pair == nil {
		if _, err := consul.KV().Put(&api.KVPair{Key: keyTag, Value: []byte(defaultTag)}, nil); err != nil {
			log.Printf("seed %s failed: %v", keyTag, err)
			return
		}
		log.Printf("seeded %s=%q", keyTag, defaultTag)
	}
}

func valueOr(pairs api.KVPairs, key, fallback string) string {
	for _, p := range pairs {
		if p.Key == key && len(p.Value) > 0 {
			return string(p.Value)
		}
	}
	return fallback
}

func envOr(key, fallback string) string {
	if v := os.Getenv(key); v != "" {
		return v
	}
	return fallback
}

// xdsLogger adapts the go-control-plane logger to the standard log package.
type xdsLogger struct{}

func (xdsLogger) Debugf(string, ...any) {}
func (xdsLogger) Infof(f string, a ...any)  { log.Printf("[xds] "+f, a...) }
func (xdsLogger) Warnf(f string, a ...any)  { log.Printf("[xds][warn] "+f, a...) }
func (xdsLogger) Errorf(f string, a ...any) { log.Printf("[xds][error] "+f, a...) }
