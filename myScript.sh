#!/bin/bash


  # echo
  # echo "##########################################################"
  # echo "##### Generate certificates using cryptogen tool #########"
  # echo "##########################################################"

  # tools/cryptogen generate --output = "./artifacts/channel/crypto-config" --config=./artifacts/channel/cryptogen.yaml

  echo "##########################################################"
  echo "#########  Generating Orderer Genesis block ##############"
  echo "##########################################################"

  tools/configtxgen -profile TwoOrgsOrdererGenesis --configPath=./artifacts/channel/ -outputBlock ./artifacts/channel/genesis.block

  echo "#################################################################"
  echo "### Generating channel configuration transaction 'channel.tx' ###"
  echo "#################################################################"
 
  tools/configtxgen -profile ThreeOrgsChannel --configPath=./artifacts/channel/ -outputCreateChannelTx ./artifacts/channel/mychannel.tx -channelID "mychannel"


  echo "#################################################################"
  echo "#######    Generating anchor peer update for Org1MSP   ##########"
  echo "#################################################################"

  tools/configtxgen -profile ThreeOrgsChannel --configPath=./artifacts/channel/ -outputAnchorPeersUpdate ./artifacts/channel/Org1MSPanchors.tx -channelID "mychannel" -asOrg Org1MSP


  echo "#################################################################"
  echo "#######    Generating anchor peer update for Org2MSP   ##########"
  echo "#################################################################"
  
  tools/configtxgen -profile ThreeOrgsChannel --configPath=./artifacts/channel/ -outputAnchorPeersUpdate ./artifacts/channel/Org2MSPanchors.tx -channelID "mychannel" -asOrg Org2MSP

  
  echo "#################################################################"
  echo "#######    Generating anchor peer update for Org3MSP   ##########"
  echo "#################################################################"
  
  tools/configtxgen -profile ThreeOrgsChannel --configPath=./artifacts/channel/ -outputAnchorPeersUpdate ./artifacts/channel/Org3MSPanchors.tx -channelID "mychannel" -asOrg Org3MSP




