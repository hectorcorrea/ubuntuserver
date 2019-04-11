# Apache

Install Apache 

```
$ sudo apt update
$ sudo apt install apache2
```

Configure the [Uncomplicated Firewall](https://wiki.ubuntu.com/UncomplicatedFirewall) to allow Apache and SSH on port 22

```
$ sudo ufw app list
$ sudo ufw allow 'Apache Full'
$ sudo ufw status

$ cat /etc/ufw/ufw.conf
# will show ENABLED=no, after the next command and a reboot should show "yes"

$ ufw allow 22/tcp
$ sudo ufw enable
```

Check that Apache is running
```
$ sudo systemctl status apache2
$ curl localhost:80
```


## install mod_proxy

I certainly need this to enable `mod_proxy`:
```
sudo a2enmod proxy
sudo a2enmod proxy_http
# sudo a2enmod proxy_html
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

* https://www.digitalocean.com/community/tutorials/how-to-install-the-apache-web-server-on-ubuntu-18-04

* https://httpd.apache.org/docs/2.4/howto/reverse_proxy.html

* https://www.digitalocean.com/community/tutorials/how-to-use-apache-http-server-as-reverse-proxy-using-mod_proxy-extension

* https://www.digitalocean.com/community/tutorials/how-to-use-apache-as-a-reverse-proxy-with-mod_proxy-on-ubuntu-16-04

* http://httpd.apache.org/docs/2.2/mod/core.html#namevirtualhost

* http://httpd.apache.org/docs/2.2/vhosts/examples.html#proxy