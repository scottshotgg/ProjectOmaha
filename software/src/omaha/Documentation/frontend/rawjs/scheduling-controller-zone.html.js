
	var timingDataArrayZone = [];
	var schedulingModeZone;
	var speakerForScheduling;
	var ssZone;

	var day = new Date().getDay();

	var myChart_timingZone; 

	var canvas_timingZone = document.getElementById("timingChartZone");
	var chart_width_scalar_timing = .8;
	var chart_height_scalar_timing = .52;

	/**
	 * This function is the zone version of drawBasic.
 	 * This function is used to draw the basic scheduling card.
	 */
	function drawBasicZone() {
		ssZone.schedulingSelectorZone.selected = "basic";

		var localWidthZone = $(window).width() * .20;
		var localHeightZone = $(window).height() * .5;
		
		document.getElementById("card-contentZone").style.width = localWidthZone + "px";
		document.getElementById("card-contentZone").style.height = localHeightZone + "px";

		//console.log(speaker);


	  loadDayOutsideZone();
	}

	/**
	 * This function is the zone version of drawAdvanced.
	 * This function is used to draw the advanced scheduling card.
	 */
	function drawAdvancedZone() {
		canvas_timingZone = document.getElementById("timingChartZone");
		//console.log(canvas_timingZone);
		ssZone.schedulingSelectorZone.selected = "advanced";

	  var localWidthZone = $(window).width() * .82;
		var localHeightZone = ($(window).height() + 36) * .62;

		document.getElementById("card-contentZone").style.width = localWidthZone + "px";
		document.getElementById("card-contentZone").style.height = localHeightZone + "px";

	  canvas_timingZone = document.getElementById("timingChartZone");
	  canvas_timingZone.width = $(window).width() * chart_width_scalar_timing;
	  canvas_timingZone.height = $(window).height() * chart_height_scalar_timing;

	  timingDataArrayZone = zones[0].Schedule[ssZone.dayListboxZone.selected].slice(4, 28);

	  $(document.getElementsByClassName("timesZone")).each(function(i) {
	    this.value = timingDataArrayZone[i];
	  });

	  var timingData = {
	    labels: timingLabels,
	    datasets: [{
	      fillColor: "rgba(63, 81, 181, 0.25)",
	      strokeColor: "#3f51b5",
	      pointColor: "#fff",
	      pointStrokeColor: "#3f51b5",  
	      data: timingDataArrayZone
	    }]
	  };

	  //console.log(timingDataArrayZone);


	  loadDayOutsideZone();
	}

	$(window).resize(function() {
		if(schedulingModeZone.selected == "advanced") {
	    redrawTimingOnInputZone()
	 	}

	});

	var timingLabels = ["00:00", "01:00", "02:00", "03:00", "04:00", "05:00", "06:00", "07:00", "08:00", "09:00", "10:00", "11:00", "12:00", "13:00", "14:00", "15:00", "16:00", "17:00", "18:00", "19:00", "20:00", "21:00", "22:00", "23:00", "24:00"];

	var timingData = {
	  labels: timingLabels,
	  datasets: [{
	    fillColor: "rgba(63, 81, 181, 0.25)",
	    strokeColor: "#3f51b5",
	    pointColor: "#fff",
	    pointStrokeColor: "#3f51b5",  
	    data: timingDataArrayZone
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
	    is: 'scheduling-controller-zone',

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

	    loadDayZone: loadDayOutsideZone,

	    /**
	     * This function is the zone version of scheduleTime.
       * ScheduleTime is used to apply a new timing schedule to a speaker. 
	     */
	    scheduleTimeZone: function() {

	    	//console.log(this.$.intervalListboxZone.selected);

	      var scheduleTimeAjaxZone = document.querySelector("#scheduleTimeAjaxZone");

	      if(schedulingModeZone.selected == "advanced") {
			    var timesZone = document.getElementsByClassName("timesZone");
					var timesArrayZone = [];
					for (var i = 0; i < timesZone.length; i++) {
						var timesParse = parseInt(times[i].value);

						if(isNaN(timesParse)) {
							return;
						}
						timesArrayZone.push(timesParse);
					}

					scheduleTimeAjaxZone.body = {
						"zone" 			: speaker.zone.zoneID,
						"mode"			: 1,
						"start"			: -1,
						"end"				: -1,
						"day" 			: parseInt(this.$.dayListboxZone.selected),
						"interval"	: 60,
						"times" 		:	timesArrayZone
					};

					scheduleTimeAjaxZone.generateRequest();
	      } else if(schedulingModeZone.selected == "basic") {
	      	var upTime = this.$.upTimeListBoxZone.selected;
	      	var downTime = this.$.downTimeListBoxZone.selected;
	      	//console.log("this is the interval", this.$.interval);

	      	var timesArrayZone = [];
	      	rampUpLevel = this.$.rampUpLevelListboxZone.selected
	      	rampDownLevel = this.$.rampDownLevelListboxZone.selected
	      	var i = 0;
	      	for(i = 0; i < upTime; i++) {
	      		timesArrayZone.push(rampDownLevel * 10);
	      	}	
	      	for(i; i < downTime; i++) {
	      		timesArrayZone.push(rampUpLevel * 10);
	      	}
	      	for(i; i < 24; i++) {
	      		timesArrayZone.push(rampDownLevel * 10);
	      	}

					scheduleTimeAjaxZone.body = {
						"zone" 			: speaker.zone.zoneID, 
						"mode"			: 0,
						"start"			: parseInt(this.$.upTimeListBoxZone.selected),
						"end"				: parseInt(this.$.downTimeListBoxZone.selected),
						"day" 			: parseInt(this.$.dayListboxZone.selected),
						"interval"	: parseInt(this.$.intervalListboxZone.selected),
						"times" 		:	timesArrayZone
					};

					scheduleTimeAjaxZone.generateRequest();

	      	//console.log(timesArrayZone);
	      }
		  },

	  	redrawTimingOnInputZone: redrawTimingOnInputZone,

	  	_logResponse: function(event, data) {
	  		//console.log(data.response);
	  	},

	    ready: function() {
	    	//console.log("THIS IS THE DAY", day);
	    	this.$.dayListboxZone.selected = day;
	    	//console.log(this.$.dayListboxZone.selected);
	    	schedulingModeZone = this.$.schedulingSelectorZone; 
	    	//console.log("i have read the selected value", schedulingModeZone); 	
	    	//console.log(schedulingModeZone);
	    	//console.log("im the ready selected  ", this.$);

				//console.log(this.zone);
				ssZone = this.$;

	    }

	  });

	/**
	 * This is the zone version of the loadDayOutside function.
 	 * This function is used to load the selectors with the correct values for that day.
	 */
	function loadDayOutsideZone() {
		//console.log(ss);

		var selectedDay = ssZone.dayListboxZone.selected;
		//console.log("day" + selectedDay);
		//console.log(ssZone.schedulingSelectorZone.selected);
		if(ssZone.schedulingSelectorZone.selected == "basic") {
			if(zones[speaker.zone.zoneID - 1].Schedule[selectedDay][2] != -1 || zones[speaker.zone.zoneID - 1].Schedule[selectedDay][3] != -1) {
				ssZone.upTimeListBoxZone.selected = zones[speaker.zone.zoneID - 1].Schedule[selectedDay][2];
				ssZone.downTimeListBoxZone.selected = zones[speaker.zone.zoneID - 1].Schedule[selectedDay][3];

				ssZone.rampUpLevelListboxZone.selected = zones[speaker.zone.zoneID - 1].Schedule[selectedDay][zones[speaker.zone.zoneID - 1].Schedule[selectedDay][2] + 4] / 10;
				ssZone.rampDownLevelListboxZone.selected = zones[speaker.zone.zoneID - 1].Schedule[selectedDay][zones[speaker.zone.zoneID - 1].Schedule[selectedDay][3] + 4] / 10;
				ssZone.intervalListboxZone.selected = zones[speaker.zone.zoneID - 1].Schedule[selectedDay][1];
			} else {
				ssZone.upTimeListBoxZone.selected = -1;
				ssZone.downTimeListBoxZone.selected = -1;
				ssZone.rampUpLevelListboxZone.selected = -1;
				ssZone.rampDownLevelListboxZone.selected = -1;
				ssZone.intervalListboxZone.selected = -1;
			}
		} else if(ssZone.schedulingSelectorZone.selected == "advanced") {
			timingDataArrayZone = zones[speaker.zone.zoneID - 1].Schedule[selectedDay].slice(4, 29);
			var times = document.getElementsByClassName("timesZone");

			for(var x = 0; x < timingDataArrayZone.length; x++) {
				times[x].value = timingDataArrayZone[x];
				times[x + 1].value = timingDataArrayZone[0];
			}
			redrawTimingOnInputZone()

			//console.log(timingDataArrayZone);
		} else {
			alert("i hava make a break");
		}
	}

	/**
	 * This is the zone version of redrawTiming.
	 * This function is used to redraw the graph.
	 */
	function redrawTimingZone(drawingData, drawingLineAttributes) {
		//console.log("I got here");
		$($(canvas_timingZone).parent()).prepend("<canvas id=\"timingChartZone\" style=\"margin-left: calc(.02 * 100vh); margin-right: calc(.02 * 100vh); margin-top: calc(.02 * 100vh);\" class=\"chartClass\"></canvas>");

		//console.log("now im here");

	  canvas_timingZone.remove();		

	  var localWidthZone = $(window).width() * .82;
		var localHeightZone = ($(window).height() + 36) * .62;

		document.getElementById("card-contentZone").style.width = localWidthZone + "px";
		document.getElementById("card-contentZone").style.height = localHeightZone + "px";

	  canvas_timingZone = document.getElementById("timingChartZone");
	  canvas_timingZone.width = $(window).width() * chart_width_scalar_timing;
	  canvas_timingZone.height = $(window).height() * chart_height_scalar_timing;

	  //console.log("i think im here");

	  //console.log(timingDataArrayZone);

	  myChart_timingZone = new Chart(canvas_timingZone.getContext("2d")).Line(drawingData, drawingLineAttributes);

	  //console.log("and now im here...");
	}

	/**
	 * This is the zone version of redrawTimingOnInput.
   * This function is fired on input into the text box. It validates the data that is input and changes the graph to red if something is detected as invalid.
	 */
	function redrawTimingOnInputZone() {
	  var errorAdjustToastZone = document.getElementById("errorAdjustToastZone");
	  var strokeColor = "#3f51b5";
	  var pointStrokeColor = "#3f51b5";
	  var fillColor = "rgba(63, 81, 181, .25)";

	  //console.log("redrawTimingOnInputZone");

	  $(document.getElementsByClassName("timesZone")).each(function(i) {
	  	//console.log(i, this.value);

	    if(this.value > -1 && this.value < 101 && this.value !== "") {
	    	//console.log("in the bigger One");
	      timingDataArrayZone[i] = parseInt(this.value);
	    } else {
	    	//console.log("not in the bigger one");
	      timingDataArrayZone[i] = null;
	      
	      // Set the error colors
	      strokeColor = "#F33F31";
	      pointStrokeColor = "#F33F31";
	      fillColor = "rgba(243, 63, 49, .25)";

	      if(isNaN(this.value)) {
	        errorAdjustToastZone.text = "Invalid entry for " + timingLabels[i] + "Hz: Only numerical values are accepted";
	        errorAdjustToastZone.show();
	      } else if (this.value === "") {

	      } else {
	        errorAdjustToastZone.text = "Invalid entry for " + timingLabels[i] + "Hz: Accepted range of values is -20 to 80";
	        errorAdjustToastZone.show();
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
	      data: timingDataArrayZone,
	      }]
	  };

	  redrawTimingZone(timingRedrawData, timingRedrawLineAttributes);
	}

