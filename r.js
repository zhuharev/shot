var system = require('system');
var args = system.args;


var url = "http://"+args[1],
filename = args[2]+".png",
timeout = args[3],
page = require('webpage').create();

page.viewportSize = { width: 1200, height: 700 };
page.clipRect = { top: 0, left: 0, width: 1200, height: 700 };
var f = false;

page.onLoadFinished = function(status) {
  console.log('Status: ' + status);
   setTimeout(function(){
    page.render(filename);
    phantom.exit();
  }, 15000)
};

page.onResourceReceived = function(response) {
if (response.url === url && !f) {
  setTimeout(function(){
    page.render(filename);
    phantom.exit();
  }, 15000)
  f = true
}
};

page.open(url, function() {
    
});