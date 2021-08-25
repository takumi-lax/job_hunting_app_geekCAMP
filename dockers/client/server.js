var express = require('express');
var serveStatic = require("serve-static")
var errorhandler = require('errorhandler');
const path = require('path');
app = express();
app.use(serveStatic(path.join(__dirname, 'dist')));
const port = process.env.PORT || 8080;
// app.get("/",function(req,res){
//     res.sendFile(__dirname+"/index.html");
  
//   });

app.listen(port);