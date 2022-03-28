#!/usr/bin/env bash
# Sets up the systemd script for hc's service and wires it up to Apache
#

# Service 
cp ./etc/systemd/system/cooking.service /etc/systemd/system/cooking.service
chmod 664 /etc/systemd/system/cooking.service

sudo systemctl daemon-reload

sudo systemctl start cooking
sudo systemctl status cooking

sudo systemctl enable cooking

# Apache
cp ./etc/apache2/sites-enabled/002-cooking.conf /etc/apache2/sites-enabled/002-cooking.conf
systemctl restart apache2
