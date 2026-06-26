#!/usr/bin/env bash
# Generates a tiny PKI for the mesh:
#   ca.crt          - the trust anchor both sidecars validate against
#   server.{crt,key} - identity presented by server-envoy (SAN: server.mesh.local)
#   client.{crt,key} - identity presented by client-envoy (SAN: client.mesh.local)
#
# mTLS = both sides present a cert and validate the peer against the CA *and*
# against the expected SPIFFE-style identity (the SAN), not just "is it signed".
set -euo pipefail

cd "$(dirname "$0")"

DAYS=3650

echo ">> generating CA"
openssl req -x509 -newkey rsa:2048 -nodes \
  -keyout ca.key -out ca.crt -days "$DAYS" \
  -subj "/CN=mesh-ca"

gen_leaf() {
  local name="$1" san="$2"
  echo ">> generating leaf cert: $name (SAN=$san)"
  openssl req -newkey rsa:2048 -nodes \
    -keyout "$name.key" -out "$name.csr" \
    -subj "/CN=$san"

  # Sign with the CA, copying the SAN into the issued cert via an ext file
  # (openssl x509 does not carry over extensions otherwise).
  printf "subjectAltName=DNS:%s\n" "$san" > "$name-ext.cnf"
  openssl x509 -req -in "$name.csr" \
    -CA ca.crt -CAkey ca.key -CAcreateserial \
    -out "$name.crt" -days "$DAYS" \
    -extfile "$name-ext.cnf"

  rm -f "$name.csr" "$name-ext.cnf"
}

gen_leaf server server.mesh.local
gen_leaf client client.mesh.local

rm -f ca.srl

# The Envoy container runs as UID 101 (envoy). Under rootless podman the host
# user maps to container-root, so a default 0600 key is unreadable by UID 101
# and Envoy reports "Failed to load incomplete private key". Make the demo keys
# world-readable so the sidecars can load them. (Demo-only; never do this for
# real private keys.)
chmod 644 *.key

echo ">> done. files:"
ls -1 *.crt *.key
