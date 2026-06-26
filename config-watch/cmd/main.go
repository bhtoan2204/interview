package main

import (
	"context"
	"log"
	"os"
	"os/signal"
	"syscall"
	"time"

	"github.com/hashicorp/consul/api"
)

func main() {
	address := os.Getenv("CONSUL_ADDRESS")
	if address == "" {
		address = "localhost:8500"
	}

	key := "TOPIC"

	config := api.DefaultConfig()
	config.Address = address

	client, err := api.NewClient(config)
	if err != nil {
		log.Fatalf("create Consul client failed: %v", err)
	}

	ctx, cancel := signal.NotifyContext(
		context.Background(),
		os.Interrupt,
		syscall.SIGTERM,
	)
	defer cancel()

	log.Printf("watching Consul key=%q address=%q", key, address)

	watchKey(ctx, client, key)
}

func watchKey(ctx context.Context, client *api.Client, key string) {
	var waitIndex uint64
	var previousValue string
	var initialized bool

	for {
		select {
		case <-ctx.Done():
			log.Println("watcher stopped")
			return
		default:
		}

		pair, meta, err := client.KV().Get(key, &api.QueryOptions{
			WaitIndex: waitIndex,
			WaitTime:  5 * time.Minute,
		})

		if err != nil {
			log.Printf("read Consul failed: %v", err)

			select {
			case <-ctx.Done():
				return
			case <-time.After(3 * time.Second):
				continue
			}
		}

		if meta == nil {
			continue
		}

		// Request timeout nhưng dữ liệu không đổi.
		if waitIndex != 0 && meta.LastIndex == waitIndex {
			continue
		}

		waitIndex = meta.LastIndex

		if pair == nil {
			if initialized {
				log.Printf("config deleted: key=%q", key)
			} else {
				log.Printf("config not found: key=%q", key)
			}

			previousValue = ""
			initialized = true
			continue
		}

		currentValue := string(pair.Value)

		if !initialized {
			log.Printf(
				"initial config: key=%q value=%q modifyIndex=%d",
				key,
				currentValue,
				pair.ModifyIndex,
			)

			previousValue = currentValue
			initialized = true
			continue
		}

		if currentValue != previousValue {
			log.Printf(
				"config changed: key=%q old=%q new=%q modifyIndex=%d",
				key,
				previousValue,
				currentValue,
				pair.ModifyIndex,
			)

			previousValue = currentValue
		}
	}
}
