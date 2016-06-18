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
  var debug = true;
  var websocketSupported = false;
  var serviceUrl = 'localhost:8080';
  var connectTo = 'ws://' + serviceUrl + '/service/live/' + socketizer.host;

  // this
  var self = {};
  self.attempts = 1;
  // public variables
  var pub = {};
  pub.connected = false;

  self.debug = function (msg) {
    if (debug) {
      console.debug(msg);
    }
  };

  pub.init = function () {
    self.check();
    self.connect();
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
    self.debug('connection closed, will try to reconnect');
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

    self.debug(msg);
    if (msg.hasOwnProperty('Data')) {
      if (msg.Data.hasOwnProperty('message')) {
        console.log('got message: ', msg.Data.message);
      } else if (msg.Data.hasOwnProperty('cmd')) {
        if (msg.Data.cmd === 'refreshPost') {
          if (msg.Data.hasOwnProperty('postUrl')) {
            postUrl = msg.Data.postUrl;
            // check if we are in current post page
            if (postUrl === currentPage) {
              console.log('refreshing post: ', postUrl);
              $('#main').load(postUrl + ' #main > *');
              return false;
            } else if ( currentPage === postsPage ) {
              // we are in central blog page, show we refresh all posts
              console.log('fetching all posts');
              $('#main').load(socketizer.postsPage + ' #main > *');
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
