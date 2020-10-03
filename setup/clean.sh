#!/bin/bash
#SCript cleans up the setup to its initial state by running cleaner in all folders

ROOT_FOLDER=$PWD/..

#1. cryptogen
cd $ROOT_FOLDER/cryptogen
./clean.sh all

#2. configtx
cd $ROOT_FOLDER/configtx
./clean.sh all

#3. orderer
cd $ROOT_FOLDER/orderer/simple-two-org
./clean.sh all

cd $ROOT_FOLDER/orderer/multi-org
./clean.sh all

cd $ROOT_FOLDER/orderer/multi-org-ca
./clean.sh all

#4. peer
cd $ROOT_FOLDER/peer/simple-two-org
./clean.sh all

cd $ROOT_FOLDER/peer/multi-org
./clean.sh all

cd $ROOT_FOLDER/peer/multi-org-ca
./clean.sh all

#5. ca
cd $ROOT_FOLDER/ca
./clean.sh all

cd $ROOT_FOLDER/ca/multi-org-ca
./clean.sh all

#6. cloud
cd $ROOT_FOLDER/cloud
./clean.sh all
rm bins/peer/peer
rm bins/orderer/orderer

#7. docker
cd $ROOT_FOLDER/docker
./clean.sh all
cp $ROOT_FOLDER/docker/raft/solution/docker-compose-raft-45-exercise.yaml  $ROOT_FOLDER/docker/raft/docker-compose-raft-45.yaml

#8. k8s
cd $ROOT_FOLDER/k8s
./clean-all.sh

# Cleanup fabric sample
rm -rf $ROOT_FOLDER/fabric-samples

# Cleanup $GOPATH/pkg
rm -rf $GOPATH/pkg

# cleanup $GOPATH/src/ all except chaincode_example02 
rm -rf $GOPATH/src/g*

