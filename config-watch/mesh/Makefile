# Use `docker` if you have it, otherwise fall back to podman.
COMPOSE ?= podman-compose
GOFLAGS  = CGO_ENABLED=0 GOOS=linux GOARCH=amd64

.PHONY: all certs build up down logs clean curl

all: certs build up

# 1. mint the mesh PKI (CA + server + client identities)
certs:
	bash certs/gen-certs.sh

# 2. compile the apps + control plane as static linux binaries (-> alpine images)
build:
	cd server-app && $(GOFLAGS) go build -o server-app .
	cd client-app && $(GOFLAGS) go build -o client-app .
	cd control-plane && $(GOFLAGS) go build -o control-plane .

# 3. bring the mesh up (builds app images, pulls envoy)
up:
	$(COMPOSE) up --build -d

down:
	$(COMPOSE) down

# tail client-app: you should see JSON responses proxied back through mTLS
logs:
	$(COMPOSE) logs -f client-app

# hit the mesh entrypoint directly from the host
curl:
	curl -s localhost:9191/ ; echo

# hot-reload demo: change the tag in Consul KV; the control plane re-pushes
# config and server-envoy starts stamping the new x-mesh-tag with no restart.
#   make tag TAG=v2
TAG ?= v2
tag:
	curl -s -X PUT localhost:8500/v1/kv/mesh/tag -d '$(TAG)' ; echo " <- set mesh/tag=$(TAG)"

clean: down
	rm -f certs/*.crt certs/*.key certs/*.srl
	rm -f server-app/server-app client-app/client-app control-plane/control-plane
