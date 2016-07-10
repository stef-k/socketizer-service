# Socketizer Service

The Websockets and API server behind Socketizer service.

## Configuration of the Server

OS: Ubuntu 16.04 LTS

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
