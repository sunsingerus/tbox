#!/bin/bash

# Docker image builder

# Source configuration
CUR_DIR="$(cd "$(dirname "${BASH_SOURCE[0]}")" >/dev/null 2>&1 && pwd)"
source "${CUR_DIR}/go_build_config.sh"

# Externally configurable build-dependent options
TAG="${TAG:-sunsingerus/tbox:dev}"
DOCKERHUB_LOGIN="${DOCKERHUB_LOGIN}"
DOCKERHUB_PUBLISH="${DOCKERHUB_PUBLISH:-yes}"
DOCKERFILE="${SRC_ROOT}/dockerfile/tbox/Dockerfile"
MINIKUBE="${MINIKUBE:-no}"

TAG="${TAG}" \
DOCKERHUB_LOGIN="${DOCKERHUB_LOGIN}" \
DOCKERHUB_PUBLISH="${DOCKERHUB_PUBLISH}" \
DOCKERFILE="${DOCKERFILE}" \
MINIKUBE="${MINIKUBE}" \
"${CUR_DIR}/image_build_universal.sh"
