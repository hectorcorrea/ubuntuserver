#!/usr/bin/env bash
# Installs apache and the modules required for ProxyPass and a sample
# configuration to proxy pass our helloworld sample.

sudo apt update
sudo apt install apache2

sudo ufw app list
sudo ufw allow 'Apache Full'
sudo ufw status
ufw allow 22/tcp
sudo ufw enable

sudo systemctl status apache2
curl localhost:80

sudo a2enmod proxy
sudo a2enmod proxy_http
# sudo a2enmod proxy_html

# Wires Apache to our helloworld site
cp ./etc/apache2/sites-enabled/000-default.conf /etc/apache2/sites-enabled/000-default.conf

systemctl restart apache2
#systemctl status apache2.service
#/etc/init.d/apache2 reload
#cat /var/log/apache2/error.log 
