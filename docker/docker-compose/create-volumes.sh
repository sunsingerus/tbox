#!/bin/bash

CUR_DIR="$(cd "$(dirname "${BASH_SOURCE[0]}")" >/dev/null 2>&1 && pwd)"
source "${CUR_DIR}/config.sh"

mkdir -p ${VOLUMES_ROOT}/mysql/{data,logs}
sudo chown 1001:1001 ${VOLUMES_ROOT}/mysql/{data,logs}

mkdir -p ${VOLUMES_ROOT}/postgresql/data
