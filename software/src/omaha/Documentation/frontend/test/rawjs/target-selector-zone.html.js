    window.addEventListener('WebComponentsReady', function(e) {
      //console.log("WebComponentsReady");
      //console.log();
    });

    //console.log(document.getElementById("targetZone"));

    var dataArrayZone = [];
    //console.log(dataArrayZone);

    var labelsZone = ["80", "100", "125", "160", "200", "250", "315", "400", "500", "630", "800", "1k", "1.25k", "1.6k", "2k", "2.5k", "3.15k", "4k", "5k", "6.3k", "8k"];


    var dataZone = {
      labels: labelsZone,
      datasets: [{
        fillColor: "rgba(63, 81, 181, 0.25)",
        strokeColor: "#3f51b5",
        pointColor: "#fff",
        pointStrokeColor: "#3f51b5",  
        data: dataArrayZone,
      }]
    };

    var lineAttributesZone = {
      scaleOverride: true,
      scaleSteps: 10,
      scaleStartValue: -20,
      scaleStepWidth: 10,
      scaleShowLabels: true,
      scaleFontSize: 12,
    };

    var resizeLineAttributesZone = {
      animation: false,
      scaleOverride: true,
      scaleSteps: 10,
      scaleStartValue: -20,
      scaleStepWidth: 10,
      scaleShowLabels: true,
      scaleFontSize: 12,
    };

    var redrawLineAttributesZone = {
      showTooltips: true,
      scaleOverride: true,
      scaleSteps: 10,
      scaleStartValue: -20,
      scaleStepWidth: 10,
      scaleShowLabels: true,
      scaleFontSize: 12,
    };

    Polymer({
      is: 'target-selector-zone',
      
      properties: {
        internalTargetZone: String,

        targetZone: {
          type: String,
          notify: true,
          observer: '_setInternalTargetZone'
        },

        speaker: {
          type: Object,
        },

        speakerId: {
          type: Number
        }
      },

    redrawOnInputZone: redrawOnInputZone,

    properties: {
      internalVolume2: Number,

       musicVolume: {
        type: Number,
        notify: true,
        observer: '_setInternalMusicVolume'
      },

      speakerId: {
        type: Number
      }
    },
      

      _setInternalTargetZone: function(dataArrayZoneString) {
        this.internalTargetZone = dataArrayZoneString;
        //console.log("_setInternalTargetZone" + dataArrayZoneString);
      },
      
      _setTargetClickedZone: function() {
        //console.log("i was just clicked");
        var checkIntZone = 0;
        var checkBandsZone = [];
        var targetTextZone = "Invalid input at ";



        $(document.getElementsByClassName("inputs-zone")).each(function(i) {
          if(isNaN(this.value) || this.value < -20 || this.value > 80) {
            //console.log(typeof this.value, parseInt(this.value));
            //console.log(this.name);
            checkIntZone++;
            //console.log(checkBandsZone);
          }
        });

        //console.log(checkBandsZone);

        if(checkIntZone > 0) {
          if(checkBandsZone.length == 1) {
            targetTextZone += checkBandsZone[0] + "Hz";
          } else {
            for (var i = 0; i < checkBandsZone.length - 1; i++) {
              targetTextZone += checkBandsZone[i] + "Hz, "
            }
            targetTextZone += "and " + checkBandsZone[checkBandsZone.length - 1] + "Hz";
          }
          this.$.toaster.duration = 2000 + (500 * checkBandsZone.length);
          this.$.toaster.text = targetTextZone;
          this.$.toaster.show();
        } else {
          var inputsZoneArray = document.getElementsByClassName("inputs-zone");
          //console.log(inputsZoneArray);
          for(var x = 0; x < inputsZoneArray.length; x++) {
            dataArrayZone[x] = inputsZoneArray[x].value;
          }

          this.dataArrayZoneString = dataArrayZone.join(", ");
          for (var i = 0; i < speaker.CurrentTarget.length; i++) { 
            speaker.CurrentTarget[i] = dataArrayZone[i];
          }
          this.target = dataArrayZone.join(" ");
          //console.log(dataArrayZone);
          this.$.toaster.text = "Target set to: " + dataArrayZone.join(", ") + ", switch to the equalizer page for further adjustment";
          this.$.toaster.show();
          //console.log("_setTargetClicked");


          this.$.updateTargetAjaxZone.body = {
            "zone": speaker.zone.zoneID,
            "constants": this.target,
          };

          this.$.updateTargetAjaxZone.generateRequest();
        }
      },

      _logResponse: function(event, data) {
        //console.log(data.response);
      },

      _acceptPressed: function(event) {
        var name = this.$.popupInput.value.valueOf().replace("\n", "").replace("\t", "");
        var ifvar = 0;
        if(speaker.zone.TargetNames.length > 0) {
          for(var i = 0; i < speaker.zone.TargetNames.length; i++) {
            if(name.toLowerCase() == speaker.zone.TargetNames[i].valueOf().toLowerCase()) {
              //console.log("same name dude");
              ifvar++;
            }
          }
        }

        if(ifvar > 0) {
          //console.log("we are in here right now");
          this.err = "Invalid name: Name already taken";
          this.$.errorAdjustToast.text = this.err;
          this.$.errorAdjustToast.show();
        } else {
          this.$.addTargetAjaxZone.body = {
            "zone": speaker.zone.zoneID,
            "name": name,
            "update": false,
            "constants": dataArrayZone.join(" ")
          };

          //console.log("addTargetAjax:", addTargetAjaxZone.body);
          this.$.addTargetAjaxZone.generateRequest();
          //console.log(this.$.addTargetAjaxZone.body);
          this.$.dialog.toggle();

          this.$.errorAdjustToast.text = "Added target: " + name + " " + dataArrayZone.join(", ");
          //console.log(this.$.errorAdjustToast.text);

          this.$.errorAdjustToast.show();

          temp = document.createElement("paper-item");
          temp.setAttribute("value", name);
          temp.setAttribute("id", name);
          temp.innerHTML = name;
          Polymer.dom(this.$.targetListbox).appendChild(temp);

          var length = speaker.zone.Target.length;

          //console.log(dataArrayZone);

          var dataArrayZoneCopy = dataArrayZone.slice();

          //console.log("dataArrayZoneCopy: ", dataArrayZoneCopy);

          speaker.zone.TargetNames[speaker.zone.TargetNames.length] = name;

          //console.log(speaker.zone.Target);

          this.$.targetListbox.selected = name;
          this.$.popupInput.value = "";
        }
      },

      _cancelPressed: function(event) {
        this.$.targetListbox.selected = -1;
        //console.log(lastSelected);
      },

      listeners: {
        'targetListbox.iron-select': '_presetSelectionChanged'
      },
      
      _presetSelectionChanged: function(event) {
        var selected = this.$.targetListbox.selected.valueOf().toLowerCase().replace("\n", "");

        if(selected == "currently applied") {
          this.$.errorAdjustToast.text = "Currently applied constants loaded! " + dataArrayZone.join(", ");
          this.$.errorAdjustToast.show();

          lastSelected = this.$.targetListbox.selected;

          for (var i = 0; i < speaker.zone.CurrentTarget.length; i++) { 
            dataArrayZone[i] = speaker.zone.CurrentTarget[i];
          }

          redrawZone(dataZone, lineAttributesZone);

          $(document.getElementsByClassName("inputs-zone")).each(function(j) {
            this.value = dataArrayZone[j];
          });
        } else if(selected == "add new target") {
            var addTargetAjaxZone = document.getElementById("addTargetAjaxZone");
            //console.log(addTargetAjaxZone);
            this.$.dialog.toggle();
        } else if(speaker.zone.TargetNames.length > 0) {
            lastSelected = this.$.targetListbox.selected;
            var associationLink = -1;
            for (k in speaker.zone.TargetNames) {
              if(selected == speaker.zone.TargetNames[k].valueOf().toLowerCase()) {
                associationLink = k;
                //console.log("associationLink", associationLink);
                break;
              }
            }
            if(associationLink != -1) {
              //console.log(speaker.zone.Target.length);

              for (var j = 0; j < speaker.zone.Target[0].length; j++) {
                dataArrayZone[j] = speaker.zone.Target[associationLink - 1][j]; 
              }

              $(document.getElementsByClassName("inputs-zone")).each(function(k) {
                this.value = dataArrayZone[k];
              });

              redrawOnInputZone(dataZone, lineAttributesZone);
              this.$.errorAdjustToast.text = this.$.targetListbox.selected + " target loaded! " + dataArrayZone.join(", ");
              this.$.errorAdjustToast.show();
              //console.log("Speaker " + speaker.zone.zoneID + " selected target: " + selected);
            }
          }
        },

      ready: function() {
        //console.log("target stuff zone", this.internalTarget, this.target);
        this.internalTarget = this.target;
        //console.log("I am ready " + this.internalTarget);
        //console.log("ready shit", this.$.targetChartZone);
        
      }
    });

    var myChartZone;

    var canvasZone = document.getElementById("targetChartZone");
    var chart_width_scalar = .8;
    var chart_height_scalar = .55;


  $(document).ready(function() {
    canvasZone = document.getElementById("targetChartZone");
    canvasZone.width = $(window).width() * chart_width_scalar;
    canvasZone.height = $(window).height() * chart_height_scalar;

    $(document.getElementsByClassName("inputs-zone")).each(function(i) {
      this.value = dataArrayZone[i];
    });

    //console.log("dater zone");
    //console.log(dataZone, lineAttributesZone);

    myChartZone = new Chart(canvasZone.getContext("2d")).Line(dataZone, lineAttributesZone);
  });


  $(window).resize(function() {
    $($(canvasZone).parent()).prepend("<canvas id=\"targetChartZone\" style=\"margin-left: calc(.02 * 100vh); margin-right: calc(.02 * 100vh); margin-top: calc(.02 * 100vh);\" class=\"chartClass\"></canvas>");
    canvasZone.remove();

    //console.log(canvasZone);

    canvasZone = document.getElementById("targetChartZone");
    canvasZone.width = $(window).width() * chart_width_scalar;
    canvasZone.height = $(window).height() * chart_height_scalar;
    myChartZone = new Chart(canvasZone.getContext("2d")).Line(dataZone, resizeLineAttributesZone);
  });

  function redrawZone(drawingData, drawingLineAttributes) {
    $($(canvasZone).parent()).prepend("<canvas id=\"targetChartZone\" style=\"margin-left: calc(.02 * 100vh); margin-right: calc(.02 * 100vh); margin-top: calc(.02 * 100vh);\" class=\"chartClass\"></canvas>");
    canvasZone.remove();

    canvasZone = document.getElementById("targetChartZone");
    //console.log("this is the canvasZone", canvasZone);
    canvasZone.width = $(window).width() * chart_width_scalar;
    canvasZone.height = $(window).height() * chart_height_scalar;
    myChartZone = new Chart(canvasZone.getContext("2d")).Line(drawingData, drawingLineAttributes);

  }

  function redrawOnInputZone() {
    //console.log("I am being called: redrawOnInputZone");
    var errorAdjustToast = document.getElementById("errorAdjustToast");
    var strokeColor = "#3f51b5";
    var pointStrokeColor = "#3f51b5";
    var fillColor = "rgba(63, 81, 181, .25)";

    //console.log(dataArrayZone);

    $(document.getElementsByClassName("inputs-zone")).each(function(i) {
      //console.log(dataArrayZone.length);
      if(this.value > -21 && this.value < 81 && this.value !== "") {
        dataArrayZone[i] = parseInt(this.value);
      } else {
        dataArrayZone[i] = null;
        
        // Set the error colors
        strokeColor = "#F33F31";
        pointStrokeColor = "#F33F31";
        fillColor = "rgba(243, 63, 49, .25)";

        if(isNaN(this.value)) {
          errorAdjustToast.text = "Invalid entry for " + labelsZone[i] + "Hz: Only numerical values are accepted";
          errorAdjustToast.show();
        } else if (this.value === "") {

        } else {
          errorAdjustToast.text = "Invalid entry for " + labelsZone[i] + "Hz: Accepted range of values is -20 to 80";
          errorAdjustToast.show();
        }
      }
    });

    var redrawDataZone = {
      labels: labelsZone,
      datasets: [{
        fillColor: fillColor,
        strokeColor: strokeColor,
        pointColor: "#fff",
        pointStrokeColor: pointStrokeColor,  
        data: dataArrayZone,
        }]
    };

    redrawZone(redrawDataZone, redrawLineAttributesZone);
  }
