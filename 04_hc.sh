#!/usr/bin/env bash
# Sets up the systemd script for hc's service and wires it up to Apache
#

# Service 
cp ./etc/systemd/system/hc.service /etc/systemd/system/hc.service
chmod 664 /etc/systemd/system/hc.service

sudo systemctl daemon-reload

sudo systemctl start hc
sudo systemctl status hc

sudo systemctl enable hc

# Apache
cp ./etc/apache2/sites-enabled/001-hc.conf /etc/apache2/sites-enabled/001-hc.conf
systemctl restart apache2
