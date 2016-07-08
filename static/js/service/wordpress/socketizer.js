'use strict';
var Socketizer = Socketizer || {};
/**
 * Module: Socketizer.main
 * @description
 * @namespace
 * @memberof Socketizer.main
 */
Socketizer.main = (function ($) {
  // private variables
  var websocketSupported = false;
  // var serviceUrl = '192.168.1.3:8080';
  // var connectTo = 'ws://' + serviceUrl + '/service/wordpress/live/' + socketizer.host;
  var connectTo = 'wss://service/wordpress/live/' + socketizer.host;

  // this
  var self = {};
  self.attempts = 1;
  // public variables
  var pub = {};
  pub.connected = false;

  pub.init = function () {
    self.healthyClose();
    self.check();
    self.connect();
  };

  self.healthyClose = function () {
    window.onbeforeunload = function () {
      if (pub.connected) {
        self.socket.close(1000);
      }
    }
  };

  /**
   * Check for websocket support
   */
  self.check = function () {
    if (window.WebSocket) {
      websocketSupported = true
    } else {
      pub.connected = 'Socketizer: Please update to a modern web browser';
    }
  };

  /**
   * Connect to websocket server
   */
  self.connect = function () {
    if (websocketSupported) {
      self.socket = new WebSocket(connectTo);
      self.socket.onopen = function () {
        pub.connected = true;
      };
      self.socket.onmessage = function (e) {
        self.receive(e);
      };
      self.socket.onclose = function () {
        pub.connected = false;
        self.close();
      };
    }
  };

  /**
   * Handle close event
   */
  self.close = function () {
    var time = self.generateInterval(self.attempts);
    setTimeout(function () {
      // increase the attempts by 1
      self.attempts++;

      // try to reconnect
      self.connect();
    }, time);
  };

  // based on http://blog.johnryding.com/post/78544969349/how-to-reconnect-web-sockets-in-a-realtime-web-app
  self.generateInterval = function (k) {
    var maxInterval = (Math.pow(2, k) - 1) * 1000;

    if (maxInterval > 30 * 1000) {
      maxInterval = 30 * 1000; // If the generated interval is more than 30 seconds, truncate it down to 30 seconds.
    }

    // generate the interval to a random number between 0 and the maxInterval determined from above
    return Math.random() * maxInterval;
  };

  // Used when a client needs an update when a new post or comment is published
  // this will avoid clients making requrests at the same time
  self.smallInterval = function () {
    return Math.random(1, 1500);
  };

  /**
   * Handle incoming messages
   */
  self.receive = function (e) {

    // received message
    var msg = JSON.parse(e.data);
    // get blog url
    var postsPage = socketizer.postsPage;
    var currentPage = window.location.href;
    var postUrl;

    if (msg.hasOwnProperty('Data')) {
      if (msg.Data.hasOwnProperty('cmd')) {
        if (msg.Data.cmd === 'refreshPost') {
          postUrl = msg.Data.postUrl;
          var pageForPosts = msg.Data.pageForPosts;
          var postsPageIsHomePage = pageForPosts.replace(/^https?:\/\//, '') === msg.Data.host;
          var selector = '#post-' + msg.Data.postId;
          var postExists = $(selector).length === 1;
          currentPage = currentPage.split('#comment-')[0];

          // if there is a new comment
          if (msg.Data.what === 'comment') {
            // if we are in single post page
            if (postUrl === currentPage && postExists) {
              setTimeout(function () {
                $('body').load(postUrl);
              }, self.smallInterval());
              return false;
            }
          } else if (msg.Data.what === 'product') { // if there is a new woocommerce product sale
            selector = '#product-' + msg.Data.postId;
            postExists = $(selector).length === 1;
            // if we are in single post page
            if (postUrl === currentPage && postExists) {
              setTimeout(function () {
                $('body').load(postUrl);
              }, self.smallInterval());
              return false;
            }
          } else if (msg.Data.what === 'bb_reply') { // new reply
            if (postUrl === currentPage) {
              setTimeout(function () {
                $('body').load(postUrl);
              }, self.smallInterval());
              return false;
            }
          } else if (msg.Data.what === 'bb_topic') { // new forum
            if (postUrl === currentPage) {
              setTimeout(function () {
                $('body').load(postUrl);
              }, self.smallInterval());
              return false;
            }
          } else if (msg.Data.what === 'bb_forum') {
            console.log(msg.Data.what);
            if (postUrl === currentPage) {
              setTimeout(function () {
                $('body').load(postUrl);
              }, self.smallInterval());
              return false;
            }
          } else if (msg.Data.what === 'post') { // if there is a new post
            // if in single post page
            if (postUrl === currentPage && postExists) {
              setTimeout(function () {
                $('body').load(postUrl);
              }, self.smallInterval());
              return false;
            } else if (currentPage === postsPage && postExists) { // if in all posts page (recent posts)
              setTimeout(function () {
                $('body').load(socketizer.postsPage);
              }, self.smallInterval());
              return false;
            } else if (postsPageIsHomePage && postExists) { // if landing page is posts page
              setTimeout(function () {
                $('body').load(pageForPosts);
              }, self.smallInterval());
              return false;
            }
          }
        }
      }
    }
  };
  return pub;
})(jQuery);
document.addEventListener('DOMContentLoaded', Socketizer.main.init);
