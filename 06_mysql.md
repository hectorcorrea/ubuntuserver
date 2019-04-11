# MySQL

```
sudo apt update
sudo apt install mysql-server
```

```
sudo mysql_secure_installation
# Password plugin: n
# Enter new password: *-*-*
# Remove anonymous: y
# Disallow root login remotely: y
# Remove test db: y
# Reload settings: y
```

Test it
```
mysql
```


Check the status (it should be running)
```
systemctl status mysql.service
```


Manually start it
```
sudo systemctl start mysql

```


## Local user

(As root) create a local user and grant it enough rights for testing access:

```
$ mysql 
mysql> CREATE USER 'hellouser'@'localhost' IDENTIFIED BY 'the-password';
mysqlCREATE DATABASE hellodb;
mysqlGRANT SELECT,INSERT,UPDATE,DELETE,CREATE,DROP ON hellodb.* TO 'hellouser'@'localhost';
mysqlEXIT;
```

Login as our new `hellouser` and test again:
```
$ mysql -u hellouser -p
[enter password]
mysql> USE hellodb;
```



## References
* https://www.digitalocean.com/community/tutorials/how-to-install-mysql-on-ubuntu-18-04






