#!/usr/bin/env sh
set -eu
base="${NOTARY_BASE_URL:-http://127.0.0.1:8090}"
curl -fsS "$base/healthz" >/dev/null
curl -fsS -X POST "$base/api/v1/signers" -H 'content-type: application/json' -d '{"id":"smoke-a","name":"a","certificate_digest":"ca","key_version":"v1"}' >/dev/null || true
curl -fsS -X POST "$base/api/v1/signers" -H 'content-type: application/json' -d '{"id":"smoke-b","name":"b","certificate_digest":"cb","key_version":"v1"}' >/dev/null || true
claim=$(curl -fsS -X POST "$base/api/v1/timestamp-requests" -H 'content-type: application/json' -d '{"digest":"aaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaa","algorithm":"SHA-256","kind":"smoke","policy_version":"v1","idempotency_key":"smoke-1"}')
id=$(printf '%s' "$claim" | sed -n 's/.*"ID":"\([^"]*\)".*/\1/p')
curl -fsS -X POST "$base/api/v1/timestamp-requests/$id/notarize" >/dev/null
echo "smoke ok: $id"
