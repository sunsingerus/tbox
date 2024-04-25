#!/bin/bash

echo "--- Generate Artifacts --- "
./dev/run_code_generator.sh
echo "--- Format Code --- "
./dev/format_unformatted_sources.sh
echo "--- Build --- "
go mod tidy
go mod vendor
go build ./pkg/...
echo "--- Done ---"
