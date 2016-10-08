# Socketizer Service

The Websockets and API server behind Socketizer service.

The whole project is based uppon 3 sub-projects:

 * [Socketizer](https://github.com/stef-k/socketizer) which is the front-end, that showcases the service, registers new users, made with Python and Django
 * Socketizer-Service (this repository) which is the WebSockets server, responsible for pushing live updates to WordPress sites, made with Go
 * [Socketizer-WordPress](https://github.com/stef-k/socketizer-wordpress) which is the WordPress plugin, responsible to call the websocket server API, made with PHP

## Configuration of the Server

OS: Ubuntu 16.04 LTS

### Supervisor

This service uses Supervisor for process management

 * Install supervisor
 * Add a SystemD config file `supervisor.service` as the following
 
```
[Unit]
Description=Supervisord Service

[Service]
Restart=on-failure
RestartSec=5s
User=root
ExecStart=/usr/bin/supervisord -n -c /etc/supervisor/supervisord.conf

[Install]
WantedBy=multi-user.target

```
 
 * Create a configuration directory `mkdir /etc/supervisor`
 * Echo the supervisor config to /etc/supervisor directory `echo_supervisord_conf > /etc/supervisor/supervisor.conf`
 * Change the [include] directive to include configuration files from /etc/supervisor/conf.d such as `files = /etc/supervisor/conf.d/*.conf `

### Raise OS open file limits

Raise to 100.000 

edit `/etc/security/limits.conf` and add the following at the end of the file
 
```
 *    soft nofile 100000
 *    hard nofile 100000
 root soft nofile 100000
 root hard nofile 100000
 ```

edit `/etc/pam.d/common-session` and add the following at the end of the file

`session required pam_limits.so`

edit `/etc/pam.d/common-session-noninteractive` and add the following at the end of the file

`session required pam_limits.so`

### Raise Supervisor Limits

Raise to 100000

edit `/etc/supervisor/supervisord.conf` and set `minfds=` as following

`minfds=100000`

To pick up the changes you must log out - relogin

### Raise Nginx Limits

Raise to 100000

edit `/etc/nginx/nginx.conf` in two places

1.  add `worker_rlimit_nofile 100000;` after `worker_processes` (usually near the top of the file
2.  at the events block edit `worker_connections` set `worker_connections 100000`;


## Endpoints

Values inside curly braces state some parameter passed to URL

### Websockets

#### WordPress

 Websockets endpoint for WordPress clients

```
 wss://service.socketizer.com/service/wordpress/live/{hostname}
```


### API

#### WordPress
 
  a WordPress post (post, comment, bbRpess, WooCommerce) has been updated
 
 ```
 /service/api/v1/wordpress/cmd/client/refresh/post/
 ``` 
