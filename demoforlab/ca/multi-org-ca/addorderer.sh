. setclient.sh orderer admin
fabric-ca-client register --id.name orderer --id.type orderer --id.affiliation orderer --id.secret pw
. setclient orderer orderer
mkdir -p $FABRIC_CA_CLIENT_HOME
cp /vagrant/demoforlab/config/multi-org-ca/yaml.0/identities/orderer/fabric-ca-client-config.yaml /vagrant/demoforlab/ca/multi-org-ca/client/orderer/orderer
fabric-ca-client enroll -u http://orderer:pw@localhost:7054
./add-admincerts.sh orderer orderer