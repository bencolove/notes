svg
object url referring to svg
img.src = object url
draw img into svg

canvas.toDataUri


$(document).ready(function() {
    var canvas = document.getElementById("canvas");
    var ctx = canvas.getContext("2d");
    var data = "<svg xmlns='http://www.w3.org/2000/svg' width='200' height='200'>" +
        "<style>" + 
          "td {background-color:yellow;" +
          "padding: 10px;}" + 
        "</style>" +         
        "<foreignObject width='100%' height='100%'>" + $("#mytable").html() +
        "</foreignObject>" +
        "</svg>";
  
  console.log(data)
  
    var DOMURL = self.URL || self.webkitURL || self;
    var img = new Image();
  img.crossOrigin = "Anonymous";
    var svg = new Blob([data], {
        type: "image/svg+xml;charset=utf-8"
    });
    var url = DOMURL.createObjectURL(svg);
    img.onload = function() {
        ctx.drawImage(img, 0, 0);
        DOMURL.revokeObjectURL(url);
      console.log(canvas.toDataURL('image/png'))
    };
    img.src = url;  
  console.log(url)
  
});

>Problem: the SVG element security constraints  
**`NOT`** *CORS*, the reason for *`tainted`* canvas but sensitive contents within a **`SVG`** like `foreignObject`, external resources endangering it.

So tainting it by locking export methods is what they do and complain.

>Solution1: draw HTML elements within SVG  
This is to parse inline HTML elements and CSS then re-draw them within a SVG, like `html2canvas`

>Solution2: use data url `data:image/svg+xml;` directly instead of blobURL  
```javascript
var canvas = document.getElementById('c');
var ctx = canvas.getContext('2d');

var data = '<svg xmlns="http://www.w3.org/2000/svg" width="200" height="200">' +
  '<foreignObject width="100%" height="100%">' +
  '<div xmlns="http://www.w3.org/1999/xhtml" style="font-size:40px">' +
  '<em>I</em> like ' +
  '<span style="color:white; text-shadow:0 0 2px blue;">' +
  'beer</span>' +
  '</div>' +
  '</foreignObject>' +
  '</svg>';


var img = new Image();
// instead of a blobURL, if we use a dataURL, chrome seems happy...
var url = 'data:image/svg+xml; charset=utf8, ' + encodeURIComponent(data);

img.onload = function() {
  c.width = this.width;
  c.height = this.height;
  ctx.drawImage(img, 0, 0);
  var dataURL = canvas.toDataURL();
  console.log(dataURL);
};

img.src = url;

```

[workaround]: https://stackoverflow.com/questions/40897039/problems-with-getting-canvas-datauri-from-svg-with-foreignobject