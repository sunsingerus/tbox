#!/bin/bash

# Build platform-specific binary
# Do not forget to update version

# Source configuration
CUR_DIR="$(cd "$(dirname "${BASH_SOURCE[0]}")" >/dev/null 2>&1 && pwd)"
source "${CUR_DIR}/go_build_config.sh"




# Prepare modules
GO111MODULE=on go mod tidy
GO111MODULE=on go mod "${MODULES_DIR}"

if [[ -z "${OUTPUT_BIN}" ]]; then
    # Not specified, requires default value
    OUTPUT_BIN="${SERVICE_BIN}"
fi

if [[ -z "${MAIN_SRC_FILE}" ]]; then
    # Not specified, requires default value
    MAIN_SRC_FILE="${SRC_ROOT}/cmd/service/main.go"
fi

GOOS=linux
GOARCH=amd64

if CGO_ENABLED=0 GO111MODULE=on GOOS="${GOOS}" GOARCH="${GOARCH}" go build \
    -mod="${MODULES_DIR}" \
    -a \
    -ldflags " \
        -X ${REPO}/pkg/version.Version=${VERSION} \
        -X ${REPO}/pkg/version.GitSHA=${GIT_SHA}  \
        -X ${REPO}/pkg/version.BuiltAt=${NOW}     \
    " \
    -o "${OUTPUT_BIN}" \
    "${MAIN_SRC_FILE}"
then
    echo "Build OK"
else
    echo "WARNING! BUILD FAILED!"
    echo "Check logs for details"
    exit 1
fi
