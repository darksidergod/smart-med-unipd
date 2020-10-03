#!/bin/bash
# Starts the specified orderer
# Usage: ./orderer-start.sh  ORDERER_NUMBER 
if [ "$1" == "" ]; then
    echo "Please provide orderer number e.g., 1, 2 ..."
else
    if [ "$1" == "1" ]; then 
        ORDERER_IDENTITY="orderer.acme.com"
    else
        ORDERER_IDENTITY="orderer$1.acme.com"
    fi
    docker-compose  -f ./config/docker-compose-base.yaml    \
                    -f ./tls/docker-compose-tls.yaml        \
                    -f ./raft/docker-compose-raft-45.yaml   \
                    -f ./raft/docker-compose-raft.yaml start $ORDERER_IDENTITY
fi