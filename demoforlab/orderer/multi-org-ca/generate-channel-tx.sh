# Generates the orderer | generate the airline channel transaction

# export ORDERER_GENERAL_LOGLEVEL=debug
export FABRIC_LOGGING_SPEC=INFO
export FABRIC_CFG_PATH=$PWD

function usage {
    echo "./generate-channel-tx.sh "
    echo "     Creates the channel.tx for the channel healthcarechannel"
}

echo    'Writing application channel with channelID healthcarechannel..'
configtxgen -profile HealthcareChannel -outputCreateChannelTx ./healthcare-channel.tx -channelID healthcarechannel
echo    'Launching orderer using the launch script. :)'

