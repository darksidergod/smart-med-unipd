#!/bin/bash

echo "============== Validate Docker =========="
docker version
docker images
echo "========== Validate Docker Compose ======"
docker-compose  version

echo "============== Validate version =========="
go version

echo "============== GOPATH =========="
echo $GOPATH

echo "============== Fabric =========="
peer version
orderer version

echo "============== Fabric CA ======="
fabric-ca-client version
fabric-ca-server version
