Take a look at the Rails logs.

```
tail log/production.log
```


Take a look at the Nginx logs to see how the request is being handled
at the web server level.

```
tail /var/log/nginx/access.log
tail /var/log/nginx/error.log
```


Test if the Ruby site is running. This should return the HTML of the site.

```
curl localhost:5100
```


If everything else fails, delete the `thin` internal files and reboot

```
rm tmp/pids/thin.5100.pid
rm tmp/pids/thin.5101.pid
reboot
```
