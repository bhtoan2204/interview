// client-app is the downstream caller. It only ever speaks plaintext HTTP to
// its own sidecar (client-envoy) on localhost:9191. It has no idea the call is
// upgraded to mTLS and routed to a remote service — that is the whole point of
// a service mesh: the application code stays oblivious to transport security.
package main

import (
	"context"
	"io"
	"log"
	"net/http"
	"os"
	"os/signal"
	"syscall"
	"time"
)

func main() {
	target := os.Getenv("TARGET_URL")
	if target == "" {
		target = "http://localhost:9191/"
	}

	interval := 3 * time.Second

	ctx, cancel := signal.NotifyContext(
		context.Background(),
		os.Interrupt,
		syscall.SIGTERM,
	)
	defer cancel()

	client := &http.Client{Timeout: 5 * time.Second}

	log.Printf("client-app calling %s every %s", target, interval)

	ticker := time.NewTicker(interval)
	defer ticker.Stop()

	for {
		call(ctx, client, target)

		select {
		case <-ctx.Done():
			log.Println("client-app stopped")
			return
		case <-ticker.C:
		}
	}
}

func call(ctx context.Context, client *http.Client, url string) {
	req, err := http.NewRequestWithContext(ctx, http.MethodGet, url, nil)
	if err != nil {
		log.Printf("build request failed: %v", err)
		return
	}

	resp, err := client.Do(req)
	if err != nil {
		log.Printf("request failed (mesh not ready yet?): %v", err)
		return
	}
	defer resp.Body.Close()

	body, _ := io.ReadAll(resp.Body)
	log.Printf("status=%d body=%s", resp.StatusCode, string(body))
}
