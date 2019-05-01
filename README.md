# Setting up an Ubuntu Server
Steps to setup an Ubuntu 18 Server to host my websites.

The markdown files (*.md) are explanations and examples on what each of the major components (System.d, Apache, MySQL) that we'll configure on the server do and notes on some not-so-obvious configuration procedures.

The shell files (*.sh) contain only the executable commands that need to be run if you don't care about the explanations.

Files under `./etc/` are samples ready to be copied to a real server under `/etc/` to get Systemd and Apache configured. These files contain most of configuration settings needed, except passwords and other sensitive information.


## Instructions
SSH into the box to test, clone this repo, and then execute the shell files
as needed. For example:

```
ssh ubuntu_box
git clone https://github.com/hectorcorrea/ubuntuserver.git
cd ubuntuserver

chmod u+x ./01_helloworld_service
./01_helloworld_service
```

