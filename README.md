# Setting up an Ubuntu Server
Steps to setup an Ubuntu 19 Server to host my websites.

The shell files (*.sh) contain only the executable commands that need to be run if you don't care about the explanations.

Files under `./etc/` are samples ready to be copied to a real server under `/etc/` to get Systemd and Apache configured. These files contain most of configuration settings needed, except passwords and other sensitive information.


# Instructions
SSH into the box to test, clone this repo, and then execute the shell files as needed. For example:

```
$ ssh ubuntu_box
$ git clone https://github.com/hectorcorrea/ubuntuserver.git
$ cd ubuntuserver

$ chmod u+x ./01_helloworld_service.sh
$ sudo ./01_helloworld_service.sh
$ curl localhost:9001
$ curl localhost:9001/file/
$ curl localhost:9001/hello/

$ chmod u+x ./02_apache.sh
$ sudo ./02_apache.sh
```


# Systemd
Ubuntu 18 uses `systemd` instead of `upstart` to handle services (programs that must be run on boot). This is a good guide to [get started](https://wiki.ubuntu.com/SystemdForUpstartUsers) on the conversion. 

To install a service create a service file under like `/etc/systemd/system/helloworld.service`
and set its permissions with `chmod 664 /etc/systemd/system/helloworld.service`

See `./etc/systemd/system` in this repo for a few examples.


## Interacting with the service
Once the service file has been created you can force it to reloaded with 

```
$ sudo systemctl daemon-reload
```

The basic operations to start and stop can be executed as follows:

```
$ sudo systemctl start helloworld
$ sudo systemctl status helloworld
$ sudo systemctl stop helloworld
```

To view the system.d log for a given service you can run

```
$ journalctl -u helloworld 
```


## Starting the service on reboot
In order for the service to be started on reboot **you must run the following command** (this creates a symlink file that informs system.d to restart it on boot)

```
$ sudo systemctl enable helloworld
```


## References
* [Systemd Service File example](https://www.shellhacks.com/systemd-service-file-example/)
* [Understanding Systemd Units and Unit Files](https://www.digitalocean.com/community/tutorials/understanding-systemd-units-and-unit-files)
* Good read but not sure I need to do this: [Integration of a Go service with systemd: readiness & liveness](https://vincent.bernat.ch/en/blog/2017-systemd-golang) by Vincent Bernat 
* Starting a Go executable as a service via systemd: https://stackoverflow.com/questions/37297714/why-systemd-cannot-start-golang-web-app-no-answer
* [systemd: Type=simple and avoiding forking considered harmful?](https://www.lucas-nussbaum.net/blog/?p=877)
* [Using environment variables](https://coreos.com/os/docs/latest/using-environment-variables-in-systemd-units.html)


# Installing and configuring Apache

Install Apache 

```
$ sudo apt update
$ sudo apt install apache2
```

Configure the [Uncomplicated Firewall](https://wiki.ubuntu.com/UncomplicatedFirewall) to allow Apache on port 80 and 443 (SSL)

```
# will show inactive
$ sudo ufw status

$ sudo ufw app list
$ sudo ufw allow 'Apache Full'
```


It's important **very important** to also allow SSH on port 22

```
$ ufw allow 22/tcp
```


Enable UFW to start on reboot:
```
$ sudo ufw enable

$ cat /etc/ufw/ufw.conf
# will show ENABLED=yes
```

Check that Apache is running
```
$ sudo systemctl status apache2
$ curl localhost:80
```


## Install mod_proxy
In order to use Apache as a reverse proxy (instead of using Nginx) I need to enable `mod_proxy`:

```
$ sudo a2enmod proxy
$ sudo a2enmod proxy_http
```

Edit file `/etc/apache2/sites-enabled/000-default.conf` and put this 
inside the VirtualHost section:

```
    ProxyPreserveHost On
    ProxyPass "/"  "http://localhost:9001/"
    ProxyPassReverse "/" "http://localhost:9001/"
```

Restart apache and check 

```
$ systemctl restart apache2
$ systemctl status apache2.service
$ /etc/init.d/apache2 reload

$ cat /var/log/apache2/error.log 
```


To check the syntax 

```
$ apache2ctl -S 
```

It seems that DocumentRoot and Redirect rules are ignored when I use ProxyPass, people recomment 
using `mod_rewrite` instead of ProxyPass to work around this or [excluding paths](https://stackoverflow.com/a/16262697/446681) from ProxyPass via the `!` syntax as in the following example:

```
    ProxyPass /specialfolder !
    ProxyPass / http://localhost:8080/tomcat-webapp/
    ProxyPassReverse / http://localhost:8080/tomcat-webapp/
    Alias /specialfolder /var/www/specialfolder
```

## References

* According to [this](https://ubuntuforums.org/showthread.php?t=1209173) "Apache Full" in `ufw` enables Apache for ports 80 and 443, whereas "Apache" only enables port 80, and "Apache Secure" only enables port 443.
* https://www.digitalocean.com/community/tutorials/how-to-install-the-apache-web-server-on-ubuntu-18-04
* https://httpd.apache.org/docs/2.4/howto/reverse_proxy.html
* https://www.digitalocean.com/community/tutorials/how-to-use-apache-http-server-as-reverse-proxy-using-mod_proxy-extension
* https://www.digitalocean.com/community/tutorials/how-to-use-apache-as-a-reverse-proxy-with-mod_proxy-on-ubuntu-16-04
* http://httpd.apache.org/docs/2.2/mod/core.html#namevirtualhost
* http://httpd.apache.org/docs/2.2/vhosts/examples.html#proxy


# Installing and configuring MySQL

Install MySQL:

```
$ sudo apt update
$ sudo apt install mysql-server
```

Secure the installation:

```
$ sudo mysql_secure_installation
# Password plugin: n
# Enter new password: *-*-*
# Remove anonymous: y
# Disallow root login remotely: y
# Remove test db: y
# Reload settings: y
```

Test it

```
$ mysql
```


Check the status (it should be running)

```
$ systemctl status mysql.service
```


Manually start it

```
$ sudo systemctl start mysql
```


## Creating a MySQL local user

(As root) create a local user and grant it enough rights for testing access:

```
$ mysql 
mysql> CREATE USER 'hellouser'@'localhost' IDENTIFIED BY 'the-password';


mysql> CREATE DATABASE hellodb;
mysql> GRANT SELECT,INSERT,UPDATE,DELETE,CREATE,DROP ON hellodb.* TO 'hellouser'@'localhost';
mysql> EXIT;
```

Login as our new `hellouser` and test again:

```
$ mysql -u hellouser -p
[enter password]
mysql> USE hellodb;
```


## References
* https://www.digitalocean.com/community/tutorials/how-to-install-mysql-on-ubuntu-18-04


# Installing SSL certificates

Install `certbot`
```
$ sudo apt-get update
$ sudo apt-get install software-properties-common
$ sudo add-apt-repository universe
$ sudo add-apt-repository ppa:certbot/certbot
$ sudo apt-get update
$ sudo apt-get install certbot python-certbot-apache 
```

generate the certificates and configure Apache to serve the SSL traffic:
```
$ sudo certbot --apache

   Your cert will expire on 2019-07-30. To obtain a new or tweaked
   version of this certificate in the future, simply run certbot again
   with the "certonly" option. To non-interactively renew *all* of
   your certificates, run "certbot renew"
```

## References
* https://certbot.eff.org/lets-encrypt/ubuntubionic-apache
