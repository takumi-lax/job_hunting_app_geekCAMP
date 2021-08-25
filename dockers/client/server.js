const express = require('express');
const serveStatic = require("serve-static")
var errorhandler = require('errorhandler');
const path = require('path');
app = express();
// app.use(serveStatic(path.join(__dirname, 'dist')));
app.use(express.static(__dirname + '/public'));
const port = process.env.PORT || 80;
app.get("/",function(req,res){
    res.sendFile(__dirname+"/index.html");
  
  });
app.listen(port);