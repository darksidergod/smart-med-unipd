RAFT
====
- The Orderers MUST be enabled for TLS
- The Peers MUST be enabled for TLS
- Configtx requires a profile for the Genesis with Orderer type=etcdraft
  Consenters list is provided with 
  + TLS Certificates
  + Orderer endpoint address
- The orderer.yaml MUST have the Cluster setup which requires the TLS certificates
============================
# Launch & Test RAFT network
============================
1. Launch the RAFT based network (2 peer + 3 orderers)
./init-setup.sh  raft

2. Check out the orderer logs to find out the leader
./raft/orderer-logs.sh  #orderer-number#

2. Run the tests agains the network using the chaincode operations
./test-all.sh tls

# Needed if you will be connecting using the host names
# Setup the orderer hosts
sudo ./raft/add-orderer-hosts.sh 
cat /etc/hosts

===============================================
# Exercise : Expand the RAFT network
# Raft Testing
================================================
1. ./init-setup.sh raft                  Launches the 3 Orderer + 2 Peer

2. Look for the leader in the network by cheking the Logs

    Note the Orderer# that is the leader

3. Update the ./raft/docker-compose-raft-45.yaml
    Add the properties for Orderer for RAFT
    Look for "<<< SETUP " 

    Solution: raft/solution/docker-compose-raft-45.yaml

4. Launch Orderer#4 & Orderer#5

    ./raft/launch-45.sh       # Launches Orderer4.acme.com & Orderer5.acme.com

   * Check logs for Orderer #4 and #5 to ensure there are NO errors
       ./raft/orderer-logs.sh  4
       ./raft/orderer-logs.sh  5

       Did these Orderer identify the Leader in the network?

5. Kill the Leader & Look for the Leader 
  ./raft/orderer-stop.sh  Orderer#
  < Now there are 4 Orderers>

6. Look for the leader in the network by cheking the Logs

    Note the Orderer# that is the leader

7. Kill the Leader & Look for the new Leader in Orderer Logs

  ./raft/orderer-stop.sh  Orderer#
  < Now there are 3 Orderers>

8. Kill the Leader & Look for the Leader 
  ./raft/orderer-stop.sh  Orderer#
  < Now there are 2 Orderers - Your network is Unstable !!! >

9. Start  any 1 stopped orderer & network will recover
  ./raft/orderer-start.sh  Orderer#

10. Confirm that leader is elected
  ./raft/orderer-logs.sh  #     < Use the Orderer# that is up >

11. Test with chaincode
  ./test-all.sh  raft





Test
====
./clean.sh all
cd raft

cryptogen generate --config=crypto-config.yaml --output=../config/crypto-config

echo    "====>Generating the genesis block"
export FABRIC_CFG_PATH=$PWD
configtxgen -outputBlock  ./orderer/airlinegenesis.block -channelID ordererchannel  -profile AirlineOrdererGenesis

echo    "====>Generating the channel create tx"
configtxgen -outputCreateChannelTx  airlinechannel.tx -channelID airlinechannel  -profile AirlineChannel

echo "********************* Launching with TLS enabled *******************"
docker-compose -f ./docker-compose-base.yaml  up   orderer.acme.com  -d
docker-compose -f ./docker-compose-base.yaml  up   orderer2.acme.com  -d
docker-compose -f ./docker-compose-base.yaml  up   orderer3.acme.com  -d


docker  rm   $(docker ps -a -q)     &> /dev/null
docker volume prune -f
docker network prune -f