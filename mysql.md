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

Make sure to configure the database to use UTF-8 as the character set, by default MySQL uses Latin 1.

```
$ mysql -u hellouser -p
mysql> USE hellodb;
mysql> SELECT @@character_set_database, @@collation_database;
mysql> ALTER DATABASE hellodb CHARACTER SET utf8 COLLATE utf8_general_ci;
```


## References
* https://www.digitalocean.com/community/tutorials/how-to-install-mysql-on-ubuntu-18-04

