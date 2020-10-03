#!/bin/bash
#Adds the Orderer to /etc/hosts. 
#You may manually add the hosts to /etc/hosts
source ../setup/manage_hosts.sh 

HOSTNAME=orderer2.acme.com
removehost $HOSTNAME            &> /dev/null
addhost $HOSTNAME

HOSTNAME=orderer3.acme.com
removehost $HOSTNAME            &> /dev/null
addhost $HOSTNAME

HOSTNAME=orderer4.acme.com
removehost $HOSTNAME            &> /dev/null
addhost $HOSTNAME

HOSTNAME=orderer5.acme.com
removehost $HOSTNAME            &> /dev/null
addhost $HOSTNAME