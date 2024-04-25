#!/bin/bash

TAG="envoy:dev"

echo "Build envoy image"
docker build -t "${TAG}" -f ./Dockerfile . || { echo "unable to build"; exit 1; }

echo "Check envoy is listed in docker"
docker image ls "${TAG}"
