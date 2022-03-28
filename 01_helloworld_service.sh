#!/usr/bin/env bash
# Sets up the systemd script for our helloworld test service
#
cp ./etc/systemd/system/helloworld.service /etc/systemd/system/helloworld.service
chmod 664 /etc/systemd/system/helloworld.service

sudo systemctl daemon-reload

sudo systemctl start helloworld
sudo systemctl status helloworld

sudo systemctl enable helloworld
