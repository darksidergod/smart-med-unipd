#!/bin/bash
#Sets the environment variable ORDERER_ADDRESS
#MUST be used with '.' or 'source'
#Usage:   .  ./raft/set-orderer-address.sh  ORDERER_NUMBER
#         source   ./raft/set-orderer-address.sh  ORDERER_NUMBER

if [ "$1" == "" ]; then
    echo "Please provide orderer number e.g., 1, 2 ..."
else
    if [ "$1" == "1" ]; then 
        export ORDERER_ADDRESS="orderer.acme.com:7050"
    else

        export ORDERER_ADDRESS="orderer$1.acme.com:8050"
    fi
fi