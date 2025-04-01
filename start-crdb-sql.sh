#!/bin/zsh

set -euxo pipefail


if command -v cockroach &> /dev/null; then
  cockroach sql --insecure
else
  echo "You don't have cockroachdb installed. Check your PATH."
fi
