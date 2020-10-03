#!/bin/bash
# Launches couchdb docker container
# If the container is already up then you may see an error which is OK
# To manually kill the container you may use:   docker rm -f couchdb

# 1. Start the container

# v1.3-3
# Pull version 2.2.0 of CouchDB - latest version 2.3.0 is not compatible with 
docker run  -e COUCHDB_USER=user -e COUCHDB_PASSWORD=password --name=couchdb -p 5984:5984 -d couchdb:2.2.0

# 2. Verify that the container is up
# Updated the time from 5s to 8s
sleep 8s

curl localhost:5984

echo "Do you see a welcome message?"
echo "If you don't see the message - try the following command in a few seconds:"
echo "curl localhost:5984"
