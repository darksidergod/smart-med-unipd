=========================================================
# Option #1
# Installation instruction for VirtualBox Express Install
# ONLY FOR VirtualBoxUsers
=========================================================
1. Open Vagrantfile & ensure that the box is set appropriately
    config.vm.box = "acloudfan/hlfdev2.0-0"

    COMMENT out the following line 
    config.vm.box = "bento/ubuntu-18.04"

2. Execute on host machine
    > vagrant up

3. Initialize the VM by executing the script & you are done !!!
    Log into the VM & change directory to vagrant/setup

    > vagrant ssh
    > cd   /vagrant/setup
    > ./init-vexpress.sh

4. Validate the setup
    Log out of the VM and 
    > exit

    Log back in
    > vagrant ssh
    > cd /vagrant/setup
    > ./validate-prereqs.sh


================================================
# Option #2
# Installation instruction for Standard Install
================================================
1. Open Vagrantfile & ensure that the box is set appropriately
    config.vm.box = "bento/ubuntu-18.04"

    COMMENT out the following line 
    config.vm.box = "acloudfan/hlfdev2.0-0"

2. Execute on host machine
    > vagrant up

Install the Tools & Fabric
PS: Use these instructons for NATIVE install on Mac/Ubuntu
    - Execute on terminal prommpt
==========================================================
Log into the VM
> vagrant ssh
> cd /vagrant/setup
> chmod 755 *.sh

1. Install Docker
sudo  ./docker.sh
exit

* Log back in to the VM & validate docker 

docker info

2. Install GoLang
sudo  ./go.sh
exit 

* Log back in to the VM & check GoLang version

go version

3. Setup Fabric
sudo -E   ./fabric-setup.sh
exit

* Log out & Log back in to the VM & check GoLang version
vagrant ssh

orderer version
peer version

4. Setup Fabric CA
cd /vagrant/setup
sudo -E  ./caserver-setup.sh

* Log back in to the VM & check GoLang version

fabric-ca-client version
fabric-ca-server version

5. Install the JQ tool
sudo ./jq.sh

6. Validate the setup
    ./validate-prereqs.sh

===============================================================
# Issues?
# Run the following script and share it with Raj in course Q&A
===============================================================
Log into the VM
> vagrant ssh
> cd /vagrant/setup
> ./validate-all.sh  > report.txt

Copy & paste the content of report.txt to Q&A post or email.