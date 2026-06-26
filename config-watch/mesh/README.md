# Envoy service mesh demo (mTLS sidecars)

A minimal, working service mesh built from two app + Envoy sidecar pairs. The
applications speak plain HTTP; the **sidecars do mutual TLS** between each other.

```
client-app ──localhost:9191──▶ client-envoy ══ mTLS ══▶ server-envoy ──localhost:8080──▶ server-app
           (plain HTTP)        (egress proxy)  (encrypted) (ingress proxy)  (plain HTTP)
```

## Why "localhost"?

Each app **shares a network namespace with its sidecar**
(`network_mode: service:<sidecar>` in `docker-compose.yml`). That is what makes
it a real sidecar: `client-app` reaches `client-envoy` on `localhost:9191`, and
`server-envoy` reaches `server-app` on `localhost:8080`. Only the two Envoys talk
over the (encrypted) network. The apps never see TLS — the mesh owns it.

## Components

| Service        | Role            | Listens on        | Talks to                       |
|----------------|-----------------|-------------------|--------------------------------|
| `client-app`   | caller (Go)     | —                 | `localhost:9191` (its sidecar) |
| `client-envoy` | egress proxy    | `:9191` plaintext | `server-envoy:8443` via mTLS   |
| `server-envoy` | ingress proxy   | `:8443` mTLS      | `localhost:8080` (its app)     |
| `server-app`   | upstream (Go)   | `:8080`           | —                              |

## PKI (`certs/`)

`gen-certs.sh` mints a tiny CA and two leaf identities:

- `server.mesh.local` — presented by `server-envoy`
- `client.mesh.local` — presented by `client-envoy`

Each sidecar validates the peer against the CA **and** pins the peer's SAN
(`match_typed_subject_alt_names`), i.e. SPIFFE-style identity, not just "signed
by our CA". `server-envoy` sets `require_client_certificate: true`, so a caller
without a valid client cert is rejected at the TLS handshake.

> Demo keys are `chmod 644` so the Envoy container (UID 101) can read them under
> rootless podman. Never do that with real keys.

## Run it

```bash
make all        # certs + build binaries + bring the mesh up
make logs       # watch client-app receive 200s proxied back over mTLS
```

Or step by step:

```bash
make certs      # generate the PKI
make build      # compile the two Go apps to static linux binaries
make up         # podman-compose up --build -d
```

`client-app` calls its sidecar every 3s and logs the JSON response.

## Prove mTLS is enforced

```bash
# from the host, through the mesh entrypoint:
curl localhost:9191/

# bypass the client sidecar and hit server-envoy directly.
# Without a client cert -> rejected:
podman run --rm --network mesh_default -v ./certs:/c:ro curlimages/curl \
  --connect-to server.mesh.local:8443:server-envoy:8443 \
  https://server.mesh.local:8443/ --cacert /c/ca.crt
# -> curl: (55) Broken pipe   (server-envoy requires a client cert)

# With a valid client cert -> 200 + JSON:
podman run --rm --network mesh_default -v ./certs:/c:ro curlimages/curl \
  --connect-to server.mesh.local:8443:server-envoy:8443 \
  https://server.mesh.local:8443/ --cacert /c/ca.crt \
  --cert /c/client.crt --key /c/client.key
```

## Admin / debugging

- `localhost:9901` — server-envoy admin (config_dump, stats, clusters)
- `localhost:9902` — client-envoy admin
- `make logs` / `podman logs <service>`
- `make down` to stop, `make clean` to also delete certs and binaries

> Uses `podman-compose`. If you use Docker, set `COMPOSE=docker compose` (e.g.
> `make up COMPOSE="docker compose"`).
