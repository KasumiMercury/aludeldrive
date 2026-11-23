#!/usr/bin/env bash
set -euo pipefail

if git diff --quiet HEAD^1 HEAD -- "$@"; then
  echo "run=false" >> "$GITHUB_OUTPUT"
else
  echo "run=true" >> "$GITHUB_OUTPUT"
fi
