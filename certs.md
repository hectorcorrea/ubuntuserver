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
