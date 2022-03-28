# Systemd
Ubuntu uses `systemd` to handle programs that must be run on boot (aka services). 

To install a service create a service file under like `/etc/systemd/system/helloworld.service` and set its permissions with `chmod 664 /etc/systemd/system/helloworld.service`

See `./etc/systemd/system` in this repo for a few examples.


## Interacting with the service
Once the service file has been created you can force it to reloaded with 

```
$ sudo systemctl daemon-reload
```

The basic operations to start and stop can be executed as follows:

```
$ sudo systemctl start helloworld
$ sudo systemctl status helloworld
$ sudo systemctl stop helloworld
```

To view the system.d log for a given service you can run

```
$ journalctl -u helloworld 
# journalctl -f -u helloworld

```


## Starting the service on reboot
In order for the service to be started on reboot **you must run the following command** (this creates a symlink file that informs system.d to restart it on boot)

```
$ sudo systemctl enable helloworld
```


## References
* [Systemd Service File example](https://www.shellhacks.com/systemd-service-file-example/)
* [Understanding Systemd Units and Unit Files](https://www.digitalocean.com/community/tutorials/understanding-systemd-units-and-unit-files)
* Good read but not sure I need to do this: [Integration of a Go service with systemd: readiness & liveness](https://vincent.bernat.ch/en/blog/2017-systemd-golang) by Vincent Bernat 
* Starting a Go executable as a service via systemd: https://stackoverflow.com/questions/37297714/why-systemd-cannot-start-golang-web-app-no-answer
* [systemd: Type=simple and avoiding forking considered harmful?](https://www.lucas-nussbaum.net/blog/?p=877)
* [Using environment variables](https://coreos.com/os/docs/latest/using-environment-variables-in-systemd-units.html)
