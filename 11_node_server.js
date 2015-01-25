// A basic node.js server
// save this as server.js and run it via node server.js
//
var http = require("http");
var port = 3000;
var serverUrl = "localhost";

var server = http.createServer(function(req, res) {
  console.log("Request: " + req.url);
  var now = new Date();
  var html = "<p>Hello World, the time is <b>" + now + "</b>.</p>";
  res.end(html);
});

console.log("Listening at http://" + serverUrl + ":" + port);

server.listen(port, serverUrl);