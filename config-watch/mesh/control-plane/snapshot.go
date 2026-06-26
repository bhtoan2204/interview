package main

import (
	"time"

	clusterv3 "github.com/envoyproxy/go-control-plane/envoy/config/cluster/v3"
	corev3 "github.com/envoyproxy/go-control-plane/envoy/config/core/v3"
	endpointv3 "github.com/envoyproxy/go-control-plane/envoy/config/endpoint/v3"
	listenerv3 "github.com/envoyproxy/go-control-plane/envoy/config/listener/v3"
	routev3 "github.com/envoyproxy/go-control-plane/envoy/config/route/v3"
	routerv3 "github.com/envoyproxy/go-control-plane/envoy/extensions/filters/http/router/v3"
	hcm "github.com/envoyproxy/go-control-plane/envoy/extensions/filters/network/http_connection_manager/v3"
	tlsv3 "github.com/envoyproxy/go-control-plane/envoy/extensions/transport_sockets/tls/v3"
	matcherv3 "github.com/envoyproxy/go-control-plane/envoy/type/matcher/v3"
	"github.com/envoyproxy/go-control-plane/pkg/cache/types"
	cachev3 "github.com/envoyproxy/go-control-plane/pkg/cache/v3"
	resourcev3 "github.com/envoyproxy/go-control-plane/pkg/resource/v3"
	"google.golang.org/protobuf/proto"
	"google.golang.org/protobuf/types/known/anypb"
	"google.golang.org/protobuf/types/known/durationpb"
	"google.golang.org/protobuf/types/known/wrapperspb"
)

// MeshConfig is the dynamic state the control plane pulls from Consul KV and
// turns into Envoy resources. Anything here can change at runtime and the new
// snapshot is pushed to the sidecars without a restart.
type MeshConfig struct {
	Tag          string // injected as the x-mesh-tag request header by server-envoy
	UpstreamHost string // DNS name client-envoy dials for the server sidecar
	MeshPort     uint32 // mTLS port server-envoy listens on
}

const (
	nodeServer = "server-envoy"
	nodeClient = "client-envoy"

	certDir = "/etc/envoy/certs"
)

// buildSnapshots produces one snapshot per Envoy node for the given config
// version. The version string changes whenever Consul changes, which is how
// Envoy knows to apply the update.
func buildSnapshots(version string, cfg MeshConfig) (map[string]*cachev3.Snapshot, error) {
	out := map[string]*cachev3.Snapshot{}

	// ---- server-envoy: ingress, terminates mTLS, injects x-mesh-tag --------
	tagHeader := []*corev3.HeaderValueOption{{
		Header:       &corev3.HeaderValue{Key: "x-mesh-tag", Value: cfg.Tag},
		AppendAction: corev3.HeaderValueOption_OVERWRITE_IF_EXISTS_OR_ADD,
	}}
	serverSnap, err := cachev3.NewSnapshot(version, map[resourcev3.Type][]types.Resource{
		resourcev3.ListenerType: {
			makeListener("ingress", cfg.MeshPort,
				httpConnManager("ingress_http", "local_app", tagHeader),
				downstreamTLS()),
		},
		resourcev3.ClusterType: {
			makeCluster("local_app", "127.0.0.1", 8080, clusterv3.Cluster_STATIC, nil),
		},
	})
	if err != nil {
		return nil, err
	}
	out[nodeServer] = serverSnap

	// ---- client-envoy: egress, originates mTLS to the server sidecar -------
	clientSnap, err := cachev3.NewSnapshot(version, map[resourcev3.Type][]types.Resource{
		resourcev3.ListenerType: {
			makeListener("egress", 9191,
				httpConnManager("egress_http", "server_mesh", nil),
				nil),
		},
		resourcev3.ClusterType: {
			makeCluster("server_mesh", cfg.UpstreamHost, cfg.MeshPort,
				clusterv3.Cluster_STRICT_DNS, upstreamTLS("server.mesh.local")),
		},
	})
	if err != nil {
		return nil, err
	}
	out[nodeClient] = clientSnap

	return out, nil
}

func makeListener(name string, port uint32, manager *anypb.Any, ts *corev3.TransportSocket) *listenerv3.Listener {
	fc := &listenerv3.FilterChain{
		Filters: []*listenerv3.Filter{{
			Name:       "envoy.filters.network.http_connection_manager",
			ConfigType: &listenerv3.Filter_TypedConfig{TypedConfig: manager},
		}},
	}
	if ts != nil {
		fc.TransportSocket = ts
	}
	return &listenerv3.Listener{
		Name:         name,
		Address:      socketAddress("0.0.0.0", port),
		FilterChains: []*listenerv3.FilterChain{fc},
	}
}

func httpConnManager(statPrefix, cluster string, reqHeaders []*corev3.HeaderValueOption) *anypb.Any {
	routeConfig := &routev3.RouteConfiguration{
		Name: "local_route",
		VirtualHosts: []*routev3.VirtualHost{{
			Name:    "vh",
			Domains: []string{"*"},
			Routes: []*routev3.Route{{
				Match:  &routev3.RouteMatch{PathSpecifier: &routev3.RouteMatch_Prefix{Prefix: "/"}},
				Action: &routev3.Route_Route{Route: &routev3.RouteAction{ClusterSpecifier: &routev3.RouteAction_Cluster{Cluster: cluster}}},
			}},
			RequestHeadersToAdd: reqHeaders,
		}},
	}
	manager := &hcm.HttpConnectionManager{
		StatPrefix:     statPrefix,
		CodecType:      hcm.HttpConnectionManager_AUTO,
		RouteSpecifier: &hcm.HttpConnectionManager_RouteConfig{RouteConfig: routeConfig},
		HttpFilters: []*hcm.HttpFilter{{
			Name:       "envoy.filters.http.router",
			ConfigType: &hcm.HttpFilter_TypedConfig{TypedConfig: mustAny(&routerv3.Router{})},
		}},
	}
	return mustAny(manager)
}

func makeCluster(name, host string, port uint32, dt clusterv3.Cluster_DiscoveryType, ts *corev3.TransportSocket) *clusterv3.Cluster {
	c := &clusterv3.Cluster{
		Name:                 name,
		ConnectTimeout:       durationpb.New(time.Second),
		ClusterDiscoveryType: &clusterv3.Cluster_Type{Type: dt},
		LoadAssignment: &endpointv3.ClusterLoadAssignment{
			ClusterName: name,
			Endpoints: []*endpointv3.LocalityLbEndpoints{{
				LbEndpoints: []*endpointv3.LbEndpoint{{
					HostIdentifier: &endpointv3.LbEndpoint_Endpoint{Endpoint: &endpointv3.Endpoint{
						Address: socketAddress(host, port),
					}},
				}},
			}},
		},
	}
	if ts != nil {
		c.TransportSocket = ts
	}
	return c
}

func downstreamTLS() *corev3.TransportSocket {
	ctx := &tlsv3.DownstreamTlsContext{
		RequireClientCertificate: wrapperspb.Bool(true),
		CommonTlsContext: commonTLS(
			certDir+"/server.crt", certDir+"/server.key", "client.mesh.local"),
	}
	return transportSocket(ctx)
}

func upstreamTLS(sni string) *corev3.TransportSocket {
	ctx := &tlsv3.UpstreamTlsContext{
		Sni: sni,
		CommonTlsContext: commonTLS(
			certDir+"/client.crt", certDir+"/client.key", "server.mesh.local"),
	}
	return transportSocket(ctx)
}

// commonTLS presents a cert/key and validates the peer against the CA, pinning
// its SAN — the mTLS identity check, identical to the previous static config.
func commonTLS(cert, key, sanMatch string) *tlsv3.CommonTlsContext {
	return &tlsv3.CommonTlsContext{
		TlsCertificates: []*tlsv3.TlsCertificate{{
			CertificateChain: fileDataSource(cert),
			PrivateKey:       fileDataSource(key),
		}},
		ValidationContextType: &tlsv3.CommonTlsContext_ValidationContext{
			ValidationContext: &tlsv3.CertificateValidationContext{
				TrustedCa: fileDataSource(certDir + "/ca.crt"),
				MatchTypedSubjectAltNames: []*tlsv3.SubjectAltNameMatcher{{
					SanType: tlsv3.SubjectAltNameMatcher_DNS,
					Matcher: &matcherv3.StringMatcher{MatchPattern: &matcherv3.StringMatcher_Exact{Exact: sanMatch}},
				}},
			},
		},
	}
}

func transportSocket(cfg proto.Message) *corev3.TransportSocket {
	return &corev3.TransportSocket{
		Name:       "envoy.transport_sockets.tls",
		ConfigType: &corev3.TransportSocket_TypedConfig{TypedConfig: mustAny(cfg)},
	}
}

func socketAddress(addr string, port uint32) *corev3.Address {
	return &corev3.Address{Address: &corev3.Address_SocketAddress{SocketAddress: &corev3.SocketAddress{
		Address:       addr,
		PortSpecifier: &corev3.SocketAddress_PortValue{PortValue: port},
	}}}
}

func fileDataSource(path string) *corev3.DataSource {
	return &corev3.DataSource{Specifier: &corev3.DataSource_Filename{Filename: path}}
}

func mustAny(m proto.Message) *anypb.Any {
	a, err := anypb.New(m)
	if err != nil {
		panic(err)
	}
	return a
}
