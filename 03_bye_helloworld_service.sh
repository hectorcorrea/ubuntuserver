#!/usr/bin/env bash
# Stops and removes the hello world service
sudo systemctl stop helloworld
rm /etc/systemd/system/helloworld.service
sudo systemctl daemon-reload
