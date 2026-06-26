# Expects ./client-app (a static linux/amd64 binary) built by `make build`.
FROM alpine:3.20
COPY client-app /client-app
ENTRYPOINT ["/client-app"]
