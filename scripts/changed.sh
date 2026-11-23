#!/usr/bin/env bash
set -euo pipefail

base=HEAD^1
if [[ "${GITHUB_EVENT_NAME:-}" == push ]]; then
  base=$(jq -r .before "$GITHUB_EVENT_PATH")
  if [[ "$base" =~ ^0+$ ]] || ! git fetch --quiet --depth=1 origin "$base"; then
    echo "run=true" >> "$GITHUB_OUTPUT"
    exit 0
  fi
fi

if git diff --quiet "$base" HEAD -- "$@"; then
  echo "run=false" >> "$GITHUB_OUTPUT"
else
  echo "run=true" >> "$GITHUB_OUTPUT"
fi
