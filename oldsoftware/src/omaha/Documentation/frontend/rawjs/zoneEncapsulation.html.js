
  $(document).ready(function () {
    $('#something').imgAreaSelect({
        handles: true,
        onSelectEnd: function(){

        	var information = [];
        	var selection = document.getElementsByClassName("imgareaselect-selection");
        	var parent = $(selection).parent();
        	information.push(parent.offset().left);
        	information.push(parent.offset().top);
        	information.push(parent.width());
        	information.push(parent.height());

        	console.log(information);

        	var check_height = information[1] + information[3];
        	var check_width = information[0] + information[2];

        	console.log(check_height, check_width);
        }
    });
}); 


var canvas = document.getElementById('c');
canvas.width = $(window).width();
canvas.height = .98 * $(window).height();
var ctx = canvas.getContext('2d');
var img = document.createElement('IMG');

img.onload = function() {
    var OwnCanv = new fabric.Canvas('c', {
        isDrawingMode: true
    });

    var imgInstance = new fabric.Image(img, {
        left: (window.innerWidth - img.width / 2) / 2,
        top: (window.innerHeight - img.height / 2) / 2,
        width: img.width / 2,
        height: img.height / 2,
    });
    OwnCanv.add(imgInstance);

    OwnCanv.freeDrawingBrush.color = "#3f51b5"
    OwnCanv.freeDrawingBrush.width = 4

    OwnCanv.on('path:created', function(options) {
        var path = options.path;
        OwnCanv.isDrawingMode = false;
        OwnCanv.clipTo = function(ctx) {
            path.render(ctx);
            console.log(path);
        };
        OwnCanv.add(imgInstance);
    });
}

img.src = "map.png";
