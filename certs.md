# Installing SSL certificates

Install `certbot` as per the instructions in https://certbot.eff.org/instructions. For Ubuntu 20 the commands are:

```
sudo snap install core
sudo snap refresh core
sudo snap install --classic certbot
sudo ln -s /snap/bin/certbot /usr/bin/certbot
sudo certbot --apache
```

There is no need to run a cron job to renew the certificates, this is now handled automatically. You can test it via:

```
sudo certbot renew --dry-run
```

and you can view the "scheduled task" via:

```
systemctl list-timers
```

## References
* https://certbot.eff.org/lets-encrypt/ubuntubionic-apache
* https://eff-certbot.readthedocs.io/en/stable/install.html
