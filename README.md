## Socketizer Service

The Websockets and API server behind Socketizer service.

### Endpoints

Values inside curly braces state some parameter passed to URL

#### Websockets

##### WordPress

 Websockets endpoint for WordPress clients

```
 wss://service.socketizer.com/service/wordpress/live/{hostname}
```


#### API

##### WordPress
 
  a WordPress post (post, comment, bbRpess, WooCommerce) has been updated
 
 ```
 /service/api/v1/wordpress/cmd/client/refresh/post/
 ``` 
