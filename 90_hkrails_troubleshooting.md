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


See if thin is running

```
ps aux | grep thin
```



Test if the Ruby site is running. This should return the HTML of the site.

```
curl localhost:5100
```


If everything else fails, delete the `thin.pid` internal files and reboot.
The values inside these .pid files usually matches with the PIDs that
`ps aux | grep` thin returned.

```
rm tmp/pids/thin.5100.pid
rm tmp/pids/thin.5101.pid
reboot
```
