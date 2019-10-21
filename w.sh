#!/bin/bash

# workspace:
#   left: .
#   right: watchexec --exts go ./w.sh

set -e    # fail on error

echo "---[start]-----------------------------------"

go test ./...
go build -o ./tfl

echo "---       -----------------------------------"

TFLINT_LOG=debug ./tfl
