    window.addEventListener('WebComponentsReady', function(e) {
      //console.log("WebComponentsReady");
      //console.log();
    });

    var dataArray = [];
    //console.log(dataArray);

    var labels = ["80", "100", "125", "160", "200", "250", "315", "400", "500", "630", "800", "1k", "1.25k", "1.6k", "2k", "2.5k", "3.15k", "4k", "5k", "6.3k", "8k"];


    var data = {
      labels: labels,
      datasets: [{
        fillColor: "rgba(63, 81, 181, 0.25)",
        strokeColor: "#3f51b5",
        pointColor: "#fff",
        pointStrokeColor: "#3f51b5",  
        data: dataArray,
      }]
    };

    var lineAttributes = {
      scaleOverride: true,
      scaleSteps: 10,
      scaleStartValue: -20,
      scaleStepWidth: 10,
      scaleShowLabels: true,
      scaleFontSize: 12,
    };

    var resizeLineAttributes = {
      animation: false,
      scaleOverride: true,
      scaleSteps: 10,
      scaleStartValue: -20,
      scaleStepWidth: 10,
      scaleShowLabels: true,
      scaleFontSize: 12,
    };

    var redrawLineAttributes = {
      showTooltips: true,
      scaleOverride: true,
      scaleSteps: 10,
      scaleStartValue: -20,
      scaleStepWidth: 10,
      scaleShowLabels: true,
      scaleFontSize: 12,
    };

    Polymer({
      is: 'target-selector',
      
      properties: {
        internalTarget: String,

        target: {
          type: String,
          notify: true,
          observer: '_setInternalTarget'
        },

        speaker: {
          type: Object,
        },
        speakerId: {
          type: Number
        }
      },

      redrawOnInput: redrawOnInput,

      /**
       * Used for setting the internal target value that will trigger updates.
       */
      _setInternalTarget: function(dataArrayString) {
        this.internalTarget = dataArrayString;
        //console.log("_setInternalTarget" + dataArrayString);
      },
      
      /**
       * Used when the 'apply target' button is clicked. It validates the entered values and sets the internal value.
       */
      _setTargetClicked: function() {
        //console.log("i got pressed _setTargetClicked");
        var checkInt = 0;
        var checkBands = [];
        var targetText = "Invalid input at ";

        $(document.getElementsByClassName("inputs")).each(function(i) {
          if(isNaN(this.value) || this.value < -20 || this.value > 80) {
            //console.log(typeof this.value, parseInt(this.value));
            //console.log(this.name);
            checkInt++;
            checkBands.push(labels[i]); // may have to put new string here
            //console.log(checkBands);
          }
        });

        //console.log(checkBands);

        if(checkInt > 0) {
          if(checkBands.length == 1) {
            targetText += checkBands[0] + "Hz";
          } else {
            for (var i = 0; i < checkBands.length - 1; i++) {
              targetText += checkBands[i] + "Hz, "
            }
            targetText += "and " + checkBands[checkBands.length - 1] + "Hz";
          }
          this.$.toaster.duration = 2000 + (500 * checkBands.length);
          this.$.toaster.text = targetText;
          this.$.toaster.show();
        } else {
          //console.log("can you just fucking work", this.target, dataArray);
          this.dataArrayString = dataArray.join(", ");
          //console.log("current target length", this.speaker.CurrentTarget.length);
          for (var i = 0; i < this.speaker.CurrentTarget.length; i++) { 
            this.speaker.CurrentTarget[i] = dataArray[i];
          }
          //console.log("can you just fucking work", this.target, dataArray);
          this.target = dataArray.join(" ");
          //console.log(this.target, dataArray);
          //console.log(dataArray);
          this.$.toaster.text = "Target set to: " + dataArray.join(", ") + ", switch to the equalizer page for further adjustment";
          this.$.toaster.show();
          //console.log("_setTargetClicked");
        }
      },

      _logResponse: function(event, data) {
        //console.log(data.response);
      },

      /**
       * This is used when the user presses the 'add' button on the dialog for adding a preset.
       */
      _acceptPressed: function(event) {
        //console.log("i got pressed _acceptPressed");
        var name = this.$.popupInput.value.valueOf().replace("\n", "").replace("\t", "");
        var ifvar = 0;
        if(this.speaker.TargetNames.length > 0) {
          for(var i = 0; i < this.speaker.TargetNames.length; i++) {
            if(name.toLowerCase() == this.speaker.TargetNames[i].valueOf().toLowerCase()) {
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
          this.$.addTargetAjax.body = {
            "speaker": this.speakerId,
            "name": name,
            "update": false,
            "constants": dataArray.join(" ")
          };

          //console.log("addTargetAjax:", addTargetAjax.body);
          this.$.addTargetAjax.generateRequest();

          this.$.dialog.toggle();

          this.$.errorAdjustToast.text = "Added target: " + name + " " + dataArray.join(", ");
          //console.log(this.$.errorAdjustToast.text);

          this.$.errorAdjustToast.show();

          temp = document.createElement("paper-item");
          temp.setAttribute("value", name);
          temp.setAttribute("id", name);
          temp.innerHTML = name;
          Polymer.dom(this.$.targetListbox).appendChild(temp);

          var length = this.speaker.Target.length;

          //console.log("dataArray length", dataArray);
          this.speaker.Target.push(dataArray.slice());

          this.speaker.TargetNames[this.speaker.TargetNames.length] = name;

          //console.log(this.speaker.Target);

          this.$.targetListbox.selected = name;
          this.$.popupInput.value = "";
        }
      },

      /**
       * This is used to dismiss the dialog and reset the selected target.
       */
      _cancelPressed: function(event) {
        this.$.targetListbox.selected = -1;
        //console.log(lastSelected);
      },

      listeners: {
        'targetListbox.iron-select': '_presetSelectionChanged'
      },
      
      /**
       * This function is used to when the user selected a different preset from the dropdown box. It loads the preset into the graph and then fires the redraw function to redraw the graph.
       */
      _presetSelectionChanged: function(event) {
        var selected = this.$.targetListbox.selected.valueOf().toLowerCase().replace("\n", "");

        if(selected == "currently applied") {
          this.$.errorAdjustToast.text = "Currently applied constants loaded! " + dataArray.join(", ");
          this.$.errorAdjustToast.show();

          lastSelected = this.$.targetListbox.selected;

          for (var i = 0; i < this.speaker.CurrentTarget.length; i++) { 
            dataArray[i] = this.speaker.CurrentTarget[i];
          }

          redraw(data, lineAttributes);

          $(document.getElementsByClassName("inputs")).each(function(j) {
            this.value = dataArray[j];
          });
        } else if(selected == "add new target") {
            var addTargetAjax = document.getElementById("addTargetAjax");
            //console.log(addTargetAjax);
            this.$.dialog.toggle();
        } else if(this.speaker.TargetNames.length > 0) {
            lastSelected = this.$.targetListbox.selected;
            var associationLink = -1;
            for (k in this.speaker.TargetNames) {
              if(selected == this.speaker.TargetNames[k].valueOf().toLowerCase()) {
                associationLink = k;
                //console.log("associationLink", associationLink);
                break;
              }
            }
            if(associationLink != -1) {
              //console.log(this.speaker.Target.length);

              for (var j = 0; j < this.speaker.Target[0].length; j++) {
               dataArray[j] = this.speaker.Target[associationLink][j]; 
              }

              $(document.getElementsByClassName("inputs")).each(function(k) {
                this.value = dataArray[k];
              });

              redraw(data, lineAttributes);
              this.$.errorAdjustToast.text = this.$.targetListbox.selected + " target loaded! " + dataArray.join(", ");
              this.$.errorAdjustToast.show();
              //console.log("Speaker " + this.speakerId + " selected target: " + selected);
            }
          }
        },


      ready: function() {
        //console.log("target stuff", this.internalTarget, this.target);
        this.internalTarget = this.target;
        //console.log("I am ready " + this.internalTarget);
      }
    });

    var myChart;

    var canvas = document.getElementById("targetChart");
    var chart_width_scalar = .8;
    var chart_height_scalar = .55;


  $(document).ready(function() {
    canvas = document.getElementById("targetChart");
    canvas.width = $(window).width() * chart_width_scalar;
    canvas.height = $(window).height() * chart_height_scalar;

    $(document.getElementsByClassName("inputs")).each(function(i) {
      this.value = dataArray[i];
    });

    //console.log("dater");
    //console.log(data, lineAttributes);

    myChart = new Chart(canvas.getContext("2d")).Line(data, lineAttributes);
  });


  $(window).resize(function() {
    $($(canvas).parent()).prepend("<canvas id=\"targetChart\" style=\"margin-left: calc(.02 * 100vh); margin-right: calc(.02 * 100vh); margin-top: calc(.02 * 100vh);\" class=\"chartClass\"></canvas>");
    canvas.remove();

    //console.log(canvas);

    canvas = document.getElementById("targetChart");
    canvas.width = $(window).width() * chart_width_scalar;
    canvas.height = $(window).height() * chart_height_scalar;
    myChart = new Chart(canvas.getContext("2d")).Line(data, resizeLineAttributes);
  });

  /**
   * This function removes the graph, draws a new one and then puts the graph in the same position the old one was in to update it.
   */
  function redraw(drawingData, drawingLineAttributes) {
    $($(canvas).parent()).prepend("<canvas id=\"targetChart\" style=\"margin-left: calc(.02 * 100vh); margin-right: calc(.02 * 100vh); margin-top: calc(.02 * 100vh);\" class=\"chartClass\"></canvas>");
    canvas.remove();

    canvas = document.getElementById("targetChart");
    canvas.width = $(window).width() * chart_width_scalar;
    canvas.height = $(window).height() * chart_height_scalar;
    myChart = new Chart(canvas.getContext("2d")).Line(drawingData, drawingLineAttributes);

  }

  /**
   * This function fires when the user inputs their coefficients into the text boxes.
   */
  function redrawOnInput() {
    var errorAdjustToast = document.getElementById("errorAdjustToast");
    var strokeColor = "#3f51b5";
    var pointStrokeColor = "#3f51b5";
    var fillColor = "rgba(63, 81, 181, .25)";

    //console.log(dataArray);

    $(document.getElementsByClassName("inputs")).each(function(i) {
      //console.log("this is the i thing that we printed", i);
      //console.log(dataArray.length);
      if(this.value > -21 && this.value < 81 && this.value !== "") {
        dataArray[i] = parseInt(this.value);
      } else {
        dataArray[i] = null;
        
        // Set the error colors
        strokeColor = "#F33F31";
        pointStrokeColor = "#F33F31";
        fillColor = "rgba(243, 63, 49, .25)";

        if(isNaN(this.value)) {
          errorAdjustToast.text = "Invalid entry for " + labels[i] + "Hz: Only numerical values are accepted";
          errorAdjustToast.show();
        } else if (this.value === "") {

        } else {
          errorAdjustToast.text = "Invalid entry for " + labels[i] + "Hz: Accepted range of values is -20 to 80";
          errorAdjustToast.show();
        }
      }
    });

    var redrawData = {
      labels,
      datasets: [{
        fillColor: fillColor,
        strokeColor: strokeColor,
        pointColor: "#fff",
        pointStrokeColor: pointStrokeColor,  
        data: dataArray,
        }]
    };

    redraw(redrawData, redrawLineAttributes);
  }
