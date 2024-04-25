#!/bin/bash

CUR_DIR="$(cd "$(dirname "${BASH_SOURCE[0]}")" >/dev/null 2>&1 && pwd)"
cd "${CUR_DIR}"

cd ./envoy || { echo "unable to find envoy"; exit 1; }
./build-image.sh
cd - || { echo "unable to return back"; exit 100; }
