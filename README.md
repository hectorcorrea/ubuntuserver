# Setting up an Ubuntu Server
Steps to setup an Ubuntu 20 Server to host my websites.

The shell files (*.sh) contain only the executable commands that need to be run if you don't care about the explanations.

Files under `./etc/` are samples ready to be copied to a real server under `/etc/` to get Systemd and Apache configured. These files contain most of configuration settings needed, except passwords and other sensitive information.


## Instructions
SSH into the box to test, clone this repo, and then execute the shell files as needed. For example:

```
ssh ubuntu_box
git clone https://github.com/hectorcorrea/ubuntuserver.git
cd ubuntuserver
```

## Hello world
Setting up a "hello world" web site and turning it into a service:

```
chmod u+x ./01_helloworld_service.sh
sudo ./01_helloworld_service.sh
curl localhost:9001
curl localhost:9001/file/
curl localhost:9001/hello/
```

## Setting up Apache
See [apache.md](https://github.com/hectorcorrea/ubuntuserver/blob/main/apache.md) for instructions on how to setup Apache or run the following scrip:

```
chmod u+x ./02_apache.sh
sudo ./02_apache.sh
```

## Hello World test (optional) 
To wire Apache to the "hello world" website:

```
cp ./etc/apache2/sites-enabled/000-helloworld.conf /etc/apache2/sites-enabled/000-helloworld.conf
systemctl restart apache2
```

To test it, `cURL` port `80` and you should get responsed from the "hello world" service:

```
curl localhost:80 
curl localhost:80/file/ 
curl localhost:80/hello/ 
```

To disable the "hello world" service:

```
chmod u+x 03_bye_helloworld_service.sh
./03_bye_helloworld_service.sh
```

## Site: HectorCorrea.com
Deploy the code for HectorCorrea.com *and then* run the following script to set it up as a service and wire it to Apache. Remember to set the ENV variables needed for production (e.g. BLOG_USR)

```
sudo ./04_hc.sh
```

## Other useful information

* [SystemD](https://github.com/hectorcorrea/ubuntuserver/tree/main/systemd.md)
* [Certificates](https://github.com/hectorcorrea/ubuntuserver/tree/main/certs.md)