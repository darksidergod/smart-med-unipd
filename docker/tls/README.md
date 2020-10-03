=======================
# Testing the TLS Setup
=======================

1. Launch the setup with TLS enabled
./clean.sh all
./init-setup.sh   tls

2. Test the setup



==============
# Manual Steps
==============
peer channel create -c airlinechannel -f ./config/airlinechannel.tx -o $ORDERER_ADDRESS --tls true --cafile $ORDERER_CA_ROOTFILE

--outputBlock config/airlinechannel.block

peer channel join   -b ./airlinechannel.block -o $ORDERER_ADDRESS 