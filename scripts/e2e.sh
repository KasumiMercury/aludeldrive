#!/usr/bin/env bash
set -euo pipefail

version=$1
addr=127.0.0.1:18080
dir=$(mktemp -d)
printf 'listen: %s\n' "$addr" > "$dir/aludel.yaml"

target/aludel serve -config "$dir/aludel.yaml" > "$dir/server.log" 2>&1 &
pid=$!
trap 'kill "$pid" 2>/dev/null || true; rm -rf "$dir"' EXIT

for _ in $(seq 50); do
  curl -sf -o /dev/null -X POST -H 'Content-Type: application/json' -d '{}' \
    "http://$addr/aludel.v1.ServerService/GetServerInfo" && break
  sleep 0.1
done

out=$(target/debug/aludel-cli --server "http://$addr" version 2>&1)
echo "$out"
test "$out" = "$(printf 'client %s\nserver %s' "$version" "$version")"
