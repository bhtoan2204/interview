# Expects ./control-plane (a static linux/amd64 binary) built by `make build`.
FROM alpine:3.20
COPY control-plane /control-plane
ENTRYPOINT ["/control-plane"]
