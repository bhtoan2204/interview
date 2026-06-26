# Expects ./server-app (a static linux/amd64 binary) built by `make build`.
FROM alpine:3.20
COPY server-app /server-app
ENTRYPOINT ["/server-app"]
