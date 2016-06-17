
var timingDataArray = [];
var defaultWidth = 0;
var defaultHeight = 0;

	var schedulingMode;
	var speakerForScheduling;
	var ss;

  var day = new Date().getDay();

  var myChart_timing; 

  var canvas_timing = document.getElementById("timingChart");
  var chart_width_scalar_timing = .8;
  var chart_height_scalar_timing = .52;
 	

 	function saveDefault() {
 		if(defaultWidth == 0) {
 			defaultWidth = document.getElementById("card-content").offsetWidth;
 		}
 		if(defaultHeight == 0) {
 			defaultHeight = document.getElementById("card-content").offsetHeight;
 		}
 	}

  function drawBasic() {
  	ss.schedulingSelector.selected = "basic";

  	var localWidth = $(window).width() * .20;
  	var localHeight = $(window).height() * .5;
  	
  	document.getElementById("card-content").style.width = localWidth + "px";
  	document.getElementById("card-content").style.height = localHeight + "px";
  	//console.log(speaker);

    loadDayOutside();
	}

  function drawAdvanced() {
  	ss.schedulingSelector.selected = "advanced";

    var localWidth = $(window).width() * .82;
  	var localHeight = ($(window).height() + 36) * .62;

  	document.getElementById("card-content").style.width = localWidth + "px";
  	document.getElementById("card-content").style.height = localHeight + "px";

    canvas_timing = document.getElementById("timingChart");
    canvas_timing.width = $(window).width() * chart_width_scalar_timing;
    canvas_timing.height = $(window).height() * chart_height_scalar_timing;

    timingDataArray = speaker.Schedule[ss.dayListbox.selected].slice(4, 28);

    $(document.getElementsByClassName("times")).each(function(i) {
      this.value = timingDataArray[i];
    });

    var timingData = {
	    labels: timingLabels,
	    datasets: [{
	      fillColor: "rgba(63, 81, 181, 0.25)",
	      strokeColor: "#3f51b5",
	      pointColor: "#fff",
	      pointStrokeColor: "#3f51b5",  
	      data: timingDataArray
	    }]
	  };

    //console.log(timingDataArray);

    loadDayOutside();
	}

  $(window).resize(function() {
  	if(schedulingMode.selected == "advanced") {
      redrawTimingOnInput();
   	}
   	adjustSize();

  });

  function adjustSize() {	
  	var localWidth = $(window).width() * .2;
  	var localHeight = $(window).height() * .5;

  	if(ss.schedulingSelector.selected == "basic") {  
  		document.getElementById("card-content").style.height = localHeight + "px";
  		saveDefault();
	   	var basic_container = document.getElementById("basic_container");
	   	var card = document.getElementById("card");
	   	basic_container.style.marginLeft = (Math.abs(card.offsetWidth - basic_container.offsetWidth) / 2) - 28 + "px";
  	}
  }

  var timingLabels = ["00:00", "01:00", "02:00", "03:00", "04:00", "05:00", "06:00", "07:00", "08:00", "09:00", "10:00", "11:00", "12:00", "13:00", "14:00", "15:00", "16:00", "17:00", "18:00", "19:00", "20:00", "21:00", "22:00", "23:00", "24:00"];

  var timingData = {
    labels: timingLabels,
    datasets: [{
      fillColor: "rgba(63, 81, 181, 0.25)",
      strokeColor: "#3f51b5",
      pointColor: "#fff",
      pointStrokeColor: "#3f51b5",  
      data: timingDataArray
    }]
  };

  var timingLineAttributes = {
    scaleOverride: true,
    scaleSteps: 10,   
  	scaleLabel: "<%= value + \"%\" %>",
    scaleStartValue: 0,
    scaleStepWidth: 10,
    scaleShowLabels: true,
    scaleFontSize: 12,
  };


  var timingResizeLineAttributes = {
    animation: false,
    scaleOverride: true,
    scaleSteps: 10,
  	scaleLabel: "<%= value + \"%\" %>",
    scaleStartValue: 0,
    scaleStepWidth: 10,
    scaleShowLabels: true,
    scaleFontSize: 12,
  };

  var timingRedrawLineAttributes = {
    showTooltips: true,
    scaleOverride: true,
    scaleSteps: 10,
  	scaleLabel: "<%= value + \"%\" %>",
    scaleStartValue: 0,
    scaleStepWidth: 10,
    scaleShowLabels: true,
    scaleFontSize: 12,
  };
	
 Polymer({
      is: 'scheduling-controller',

      behaviors: [
        Polymer.NeonAnimatableBehavior
      ],

      properties: {

        speaker: {
          type: Object
        },

      	animationConfig: {
          value: function() {
            return {
              'entry': {
                  name: 'fade-in-animation',
                  node: this
                },
              'exit': {
                name: 'slide-left-animation',
                node: this
              }
            }
          }
        }
      },

      loadDay: loadDayOutside,

      scheduleTime: function() {

      	//console.log(this.$.intervalListbox.selected);

	      var scheduleTimeAjax = document.querySelector("#scheduleTimeAjax");

	      if(schedulingMode.selected == "advanced") {
			    var times = document.getElementsByClassName("times");
					var timesArray = [];
					for (var i = 0; i < times.length; i++) {
						var timesParse = parseInt(times[i].value);

						if(isNaN(timesParse)) {
							return;
						}
						timesArray.push(timesParse);
					}

					scheduleTimeAjax.body = {
						"speaker" 	: speakerid,
						"mode"			: 1,
						"start"			: -1,
						"end"				: -1,
						"day" 			: parseInt(this.$.dayListbox.selected),
						"interval"	: 60,
						"times" 		:	timesArray
					};

					scheduleTimeAjax.generateRequest();
	      } else if(schedulingMode.selected == "basic") {
	      	var upTime = this.$.upTimeListBox.selected;
	      	var downTime = this.$.downTimeListBox.selected;

	      	//console.log("this is the interval", this.$.interval);

	      	var timesArray = [];
	      	rampUpLevel = this.$.rampUpLevelListbox.selected
	      	rampDownLevel = this.$.rampDownLevelListbox.selected
	      	var i = 0;
	      	for(i = 0; i < upTime; i++) {
	      		timesArray.push(rampDownLevel * 10);
	      	}	
	      	for(i; i < downTime; i++) {
	      		timesArray.push(rampUpLevel * 10);
	      	}
	      	for(i; i < 24; i++) {
	      		timesArray.push(rampDownLevel * 10);
	      	}

					scheduleTimeAjax.body = {
						"speaker" 	: speakerid, 
						"mode"			: 0,
						"start"			: parseInt(this.$.upTimeListBox.selected),
						"end"				: parseInt(this.$.downTimeListBox.selected),
						"day" 			: parseInt(this.$.dayListbox.selected),
						"interval"	: parseInt(this.$.intervalListbox.selected),
						"times" 		:	timesArray
					};

					scheduleTimeAjax.generateRequest();

	      	//console.log(timesArray);
	      	this.$.scheduleTimeAjax.body = {

	      	};
	      }
		  },

    	redrawTimingOnInput: redrawTimingOnInput,

    	_logResponse: function(event, data) {
    		//console.log(data.response);
    	},

      ready: function() {
      	//console.log("THIS IS THE DAY", day);
      	this.$.dayListbox.selected = day;
      	//console.log(this.$.dayListbox.selected);
      	schedulingMode = this.$.schedulingSelector; 
      	//console.log("i have read the selected value", schedulingMode); 	
      	//console.log(schedulingMode);
      	//console.log("im the ready selected  ", this.$);

  			//console.log(this.speaker);
  			ss = this.$;

      }

    });

	function loadDayOutside() {
		//console.log(ss);

		var selectedDay = ss.dayListbox.selected;
		//console.log("day" + selectedDay);
		//console.log(speaker);
		//console.log(ss.schedulingSelector.selected);
		if(ss.schedulingSelector.selected == "basic") {
			if(speaker.Schedule[selectedDay][2] != -1 || speaker.Schedule[selectedDay][3] != -1) {
				ss.upTimeListBox.selected = speaker.Schedule[selectedDay][2];
				ss.downTimeListBox.selected = speaker.Schedule[selectedDay][3];

				ss.rampUpLevelListbox.selected = speaker.Schedule[selectedDay][speaker.Schedule[selectedDay][2] + 4] / 10;
				ss.rampDownLevelListbox.selected = speaker.Schedule[selectedDay][speaker.Schedule[selectedDay][3] + 4] / 10;
				ss.intervalListbox.selected = speaker.Schedule[selectedDay][1];
			} else {
				ss.upTimeListBox.selected = -1;
				ss.downTimeListBox.selected = -1;
				ss.rampUpLevelListbox.selected = -1;
				ss.rampDownLevelListbox.selected = -1;
				ss.intervalListbox.selected = -1;
			}
		} else if(ss.schedulingSelector.selected == "advanced") {
			timingDataArray = speaker.Schedule[selectedDay].slice(4, 29);
			var times = document.getElementsByClassName("times");

			for(var x = 0; x < timingDataArray.length; x++) {
				times[x].value = timingDataArray[x];
				times[x + 1].value = timingDataArray[0];
			}
			redrawTimingOnInput()

			//console.log(timingDataArray);
		} else {
			alert("i hava make a break");
		}
	}


  function redrawTiming(drawingData, drawingLineAttributes) {
  	//console.log("I got here");
		$($(canvas_timing).parent()).prepend("<canvas id=\"timingChart\" style=\"margin-left: calc(.02 * 100vh); margin-right: calc(.02 * 100vh); margin-top: calc(.02 * 100vh);\" class=\"chartClass\"></canvas>");

		//console.log("now im here");

    canvas_timing.remove();		

    var localWidth = $(window).width() * .82;
  	var localHeight = ($(window).height() + 36) * .62;

  	document.getElementById("card-content").style.width = localWidth + "px";
  	document.getElementById("card-content").style.height = localHeight + "px";

    canvas_timing = document.getElementById("timingChart");
    canvas_timing.width = $(window).width() * chart_width_scalar_timing;
    canvas_timing.height = $(window).height() * chart_height_scalar_timing;

    //console.log("i think im here");

    //console.log(timingDataArray);

    myChart_timing = new Chart(canvas_timing.getContext("2d")).Line(drawingData, drawingLineAttributes);

    //console.log("and now im here...");
  }

  function redrawTimingOnInput() {
    var errorAdjustToast = document.getElementById("errorAdjustToast");
    var strokeColor = "#3f51b5";
    var pointStrokeColor = "#3f51b5";
    var fillColor = "rgba(63, 81, 181, .25)";

    //console.log("redrawTimingOnInput");

    $(document.getElementsByClassName("times")).each(function(i) {
    	//console.log(i);
      if(this.value > -1 && this.value < 101 && this.value !== "") {
        timingDataArray[i] = parseInt(this.value);
      } else {
        timingDataArray[i] = null;
        
        // Set the error colors
        strokeColor = "#F33F31";
        pointStrokeColor = "#F33F31";
        fillColor = "rgba(243, 63, 49, .25)";

        if(isNaN(this.value)) {
          errorAdjustToast.text = "Invalid entry for " + timingLabels[i] + "Hz: Only numerical values are accepted";
          errorAdjustToast.show();
        } else if (this.value === "") {

        } else {
          errorAdjustToast.text = "Invalid entry for " + timingLabels[i] + "Hz: Accepted range of values is -20 to 80";
          errorAdjustToast.show();
        }
      }
    });

    var timingRedrawData = {
      labels: timingLabels,
      datasets: [{
        fillColor: fillColor,
        strokeColor: strokeColor,
        pointColor: "#fff",
        pointStrokeColor: pointStrokeColor,  
        data: timingDataArray,
        }]
    };

    redrawTiming(timingRedrawData, timingRedrawLineAttributes);
  }

