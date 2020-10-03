# source setclient caserver admin
# mkdir -p $FABRIC_CA_CLIENT_HOME
# cp /vagrant/demoforlab/config/multi-org-ca/yaml.0/fabric-ca-client-config.yaml $PWD/client/caserver/admin
fabric-ca-client enroll -u http://admin:pw@localhost:7054
./register-enroll-admins.sh
./setup-org-msp.sh acme
./setup-org-msp.sh budget
./setup-org-msp.sh orderer
. setclient orderer admin
cd ../../orderer/multi-org-ca
./register-enroll-orderer.sh
./generate-genesis.sh
./generate-channel-tx.sh

