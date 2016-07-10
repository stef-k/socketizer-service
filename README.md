## Socketizer Service

The Websockets and API server behind Socketizer service.

### Endpoints

Values inside curly braces state some parameter passed to URL

#### Websockets

##### WordPress
 
 ```wss://service.socketizer.com/service/wordpress/live/{hostname}```
 
 Websockets endpoint for WordPress clients

#### API

##### WordPress
 
 ```/service/api/v1/wordpress/cmd/client/refresh/post/``` 
 
 a WordPress post (post, comment, bbRpess, WooCommerce) has been updated
