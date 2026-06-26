// server-app is the upstream service. It listens on plaintext :8080 and is only
// ever reached through its own sidecar (server-envoy) over localhost, so it
// never needs to know anything about TLS — the mesh handles that for it.
package main

import (
	"encoding/json"
	"log"
	"net/http"
	"os"
	"time"
)

func main() {
	addr := os.Getenv("LISTEN_ADDR")
	if addr == "" {
		addr = "127.0.0.1:8080"
	}

	hostname, _ := os.Hostname()

	mux := http.NewServeMux()
	mux.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		// Envoy injects these headers; surfacing them proves traffic flowed
		// through the mesh rather than hitting the app directly.
		resp := map[string]any{
			"message":         "hello from server-app",
			"served_by":       hostname,
			"path":            r.URL.Path,
			"method":          r.Method,
			"time":            time.Now().Format(time.RFC3339),
			"x-forwarded-for": r.Header.Get("x-forwarded-for"),
			"x-request-id":    r.Header.Get("x-request-id"),
			// injected by server-envoy from Consul KV (mesh/tag) — proves the
			// xDS hot-reload landed.
			"x-mesh-tag": r.Header.Get("x-mesh-tag"),
		}

		w.Header().Set("Content-Type", "application/json")
		_ = json.NewEncoder(w).Encode(resp)

		log.Printf("handled %s %s from %s", r.Method, r.URL.Path, r.RemoteAddr)
	})

	log.Printf("server-app listening on %s", addr)
	if err := http.ListenAndServe(addr, mux); err != nil {
		log.Fatalf("server-app failed: %v", err)
	}
}
