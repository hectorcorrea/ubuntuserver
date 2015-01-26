Setting up an Ubuntu Server
================================
Steps to setup an Ubuntu Server to host Node.js and Ruby on Rails web sites. The following tools are installed and configured:

* Nginx (as the front end web server)

* Node.js (from source)
  * Small node.js web site use as a demo.
  * Config file to hook it to Nginx.

  TODO
  * Upstart script for Node.js web site.

* Ruby (from source via ruby-install)
  * chruby to select the current Ruby (Ruby, gems, paths)
  * Small Sinatra web site to use as a demo.
  * Uses thin as the app server.
  * Config file to hook it to Nginx.

  * A Rails web site (currently using a copy of one of my own sites)
  * Uses thin as the app server.
  * Config file to hook it to Nginx.

  TODO
  * Upstart script for Rails web site.
