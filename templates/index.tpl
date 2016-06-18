<!doctype html>
<html class="no-js" lang="en">
<head>
  <link rel="apple-touch-icon" sizes="57x57" href="/static/img/icon/apple-icon-57x57.png">
  <link rel="apple-touch-icon" sizes="60x60" href="/static/img/icon/apple-icon-60x60.png">
  <link rel="apple-touch-icon" sizes="72x72" href="/static/img/icon/apple-icon-72x72.png">
  <link rel="apple-touch-icon" sizes="76x76" href="/static/img/icon/apple-icon-76x76.png">
  <link rel="apple-touch-icon" sizes="114x114" href="/static/img/icon/apple-icon-114x114.png">
  <link rel="apple-touch-icon" sizes="120x120" href="/static/img/icon/apple-icon-120x120.png">
  <link rel="apple-touch-icon" sizes="144x144" href="/static/img/icon/apple-icon-144x144.png">
  <link rel="apple-touch-icon" sizes="152x152" href="/static/img/icon/apple-icon-152x152.png">
  <link rel="apple-touch-icon" sizes="180x180" href="/static/img//apple-icon-180x180.png">
  <link rel="icon" type="image/png" sizes="192x192" href="/static/img/icon/android-icon-192x192.png">
  <link rel="icon" type="image/png" sizes="32x32" href="/static/img/icon/favicon-32x32.png">
  <link rel="icon" type="image/png" sizes="96x96" href="/static/img/icon/favicon-96x96.png">
  <link rel="icon" type="image/png" sizes="16x16" href="/static/img/icon/favicon-16x16.png">
  <link rel="manifest" href="/static/img/icon/manifest.json">
  <meta name="msapplication-TileColor" content="#ffffff">
  <meta name="msapplication-TileImage" content="/static/img/icon/ms-icon-144x144.png">
  <meta name="theme-color" content="#ffffff">
  <link rel="shortcut icon" href="/static/img/icon/favicon.ico" type="image/x-icon">
  <link rel="icon" href="/static/img/icon/favicon.ico" type="image/x-icon">
  <meta charset="utf-8">
  <meta http-equiv="x-ua-compatible" content="ie=edge">
  <title>Socketizer</title>
  <meta name="description" content="">
  <meta name="viewport" content="width=device-width, initial-scale=1">
  <style>
    body {
      font-family: "Helvetica Neue", Helvetica, Arial, sans-serif;
    }

    h1 {
      font-size: 2.8rem;
    }

    code {
      font-size: 90%;
      line-height: 3.5em;
      font-family: Monaco, Consolas, "Andale Mono", "DejaVu Sans Mono", monospace;

      display: inline;
      color: #555555;
      padding: 1em 1em;
      background: #f4f4f4;
    }
  </style>
</head>
<body>
<!--[if lt IE 8]>
<p class="browserupgrade">You are using an <strong>outdated</strong> browser.
                          Please <a href="http://browsehappy.com/">upgrade your browser</a> to improve your experience.
</p>
<![endif]-->

<!-- Add your site or application content here -->
<h1><img src="/static/img/logo.png" style="width:80px;height:auto;"> <span
    style="line-height: 80px;vertical-align: top">
  Socketizer</span>
</h1>
<p>Welcome to being live!</p>

<h2>Info <button onclick="getInfo()">Refresh</button></h2>
<p id="info"></p>

<h2>Websocket Message: </h2>
<code id="message"></code>

<script src="/static/js/websocket.js"></script>
<script>
  var getInfo = function () {
    var request = new XMLHttpRequest();
    request.addEventListener('load', function () {
      var el = document.querySelector('#info');
      var msg = JSON.parse(this.response);
      el.innerHTML = '<strong>Active Domains Now: </strong>' + msg.DomainCount + '<br>';
      el.innerHTML += '<strong>Total Connected Clients: </strong>' + msg.ClientSub + '<br>';
      el.innerHTML += '<strong>Domain List: </strong>' + msg.DomainList + '<br>';
    });
    request.open('GET', 'http://localhost:8080/service/api/v1/pool-info', true);
    request.send();
  };
  getInfo();
</script>
<!-- Google Analytics: change UA-XXXXX-Y to be your site's ID. -->
<script>
  window.ga = function () {
    ga.q.push(arguments)
  };
  ga.q = [];
  ga.l = +new Date;
  ga('create', 'UA-XXXXX-Y', 'auto');
  ga('send', 'pageview')
</script>
<script src="https://www.google-analytics.com/analytics.js" async defer></script>
</body>
</html>
