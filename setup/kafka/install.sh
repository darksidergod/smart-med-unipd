#!/bin/bash
# Installs Kafka under the $HOME folder
# Location of binaries => $HOME/kafka/bin

# Update
sudo apt-get update

# Install JRE
sudo apt-get --assume-yes install default-jre

# Install Zookeepr
sudo apt-get --assume-yes install zookeeperd
# Stop zookeeper and prevent it from auto starting
sudo service zookeeper stop
sudo systemctl remove zookeeper

# Create a temp directory
mkdir $HOME/kafka
cd $HOME/kafka

# Fabric 2.x Update: Updated Kafka to v2.1
wget http://www-us.apache.org/dist/kafka/2.4.0/kafka_2.13-2.4.0.tgz

mv *.tgz  kafka.tgz

# Untar & delete tar file
tar -xvzf ./kafka.tgz --strip 1
rm kafka.tgz

