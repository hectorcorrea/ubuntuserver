# Apache

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
sudo a2enmod proxy
sudo a2enmod proxy_http
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
systemctl restart apache2
systemctl status apache2.service
/etc/init.d/apache2 reload

cat /var/log/apache2/error.log 

```


## References

* According to [this](https://ubuntuforums.org/showthread.php?t=1209173) "Apache Full" in `ufw` enables Apache for ports 80 and 443, whereas "Apache" only enables port 80, and "Apache Secure" only enables port 443.

* https://www.digitalocean.com/community/tutorials/how-to-install-the-apache-web-server-on-ubuntu-18-04

* https://httpd.apache.org/docs/2.4/howto/reverse_proxy.html

* https://www.digitalocean.com/community/tutorials/how-to-use-apache-http-server-as-reverse-proxy-using-mod_proxy-extension

* https://www.digitalocean.com/community/tutorials/how-to-use-apache-as-a-reverse-proxy-with-mod_proxy-on-ubuntu-16-04

* http://httpd.apache.org/docs/2.2/mod/core.html#namevirtualhost

* http://httpd.apache.org/docs/2.2/vhosts/examples.html#proxy