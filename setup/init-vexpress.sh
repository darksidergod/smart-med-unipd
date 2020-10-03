#!/bin/bash
# This script needs to be used ONLY for "Virtual Box Express Install"
if [ ! -d "$HOME/goccbup/gocc" ]; then 
    echo "This script works only with the 'Virtual Box Express' install !!!"
    exit
fi


source ./to_absolute_path.sh

GOPATH=$PWD/../gopath
to-absolute-path $GOPATH
GOPATH=$ABS_PATH


echo "Setting GOPATH=$GOPATH"



echo "export GOPATH=$GOPATH" >> ~/.profile
echo "======== Updated .profile with GOROOT/GOPATH/PATH===="

echo "export GOROOT=/usr/local/go" >> ~/.bashrc
echo "export GOPATH=$GOPATH" >> ~/.bashrc
echo "======== Updated .profile with GOROOT/GOPATH/PATH===="

# UPDATED Dec 15, 2019
GOCACHE="$HOME/.go-cache"
echo "export GOCACHE=$HOME/.go-cache" >> ~/.bashrc
mkdir -p $GOCACHE
chown -R $USER $GOCACHE

mkdir -p $HOME/packages
chown -R $USER $HOME/packages
chmod a+rwx $HOME/packages

./update_etc_hosts.sh

# Updating the shim
LOCT=$HOME/goccbup/gocc
# copy the pkg
echo "=============================="
echo "Setting up Shim & Go Binaries"
mkdir -p $GOPATH/pkg
cp -r $LOCT/pkg $GOPATH
mkdir -p $GOPATH/bin
cp -r $LOCT/bin $GOPATH

echo "Setting up Fabric & Golang Libraries"
cp -r $LOCT/src/github.com $GOPATH/src
cp -r $LOCT/src/golang.org $GOPATH/src
cp -r $LOCT/src/google.golang.org $GOPATH/src

echo "=============================="

rm -rf $HOME/*

echo "DONE."
echo "You MUST log out of VM and Log back in."
