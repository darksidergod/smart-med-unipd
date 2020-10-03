#!/bin/bash
#Launches the orderer 4 & 5
#Ensure the orderer 1, 2 & 3 are already up using the ./init-set.sh raft

docker-compose -f ./config/docker-compose-base.yaml \
               -f ./tls/docker-compose-tls.yaml     \
               -f ./raft/docker-compose-raft.yaml   \
               -f ./raft/docker-compose-raft-45.yaml up -d \
               orderer4.acme.com orderer5.acme.com  