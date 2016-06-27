
  var sliders;
  var handles;
  var background;

  var equalizerToast = 0;
  var equalizerMode;
  var typeFound = 0;

  /**
   * GetValues is used to get the type of equalizer that is currently selected.
   * @returns {Number} typeFound is an integer that represents the type of equalizer for the speaker.
   */

  function getValues() {
    switch (equalizerMode.selected) {
      case "masking": 
        typeFound = 0;
        //console.log(equalizerMode.selected);
        break;
      case "music":
        typeFound = 1;
        //console.log(equalizerMode.selected);
        break;
      case "paging":
        typeFound = 2;
        //console.log(equalizerMode.selected);
        break;

      default:
        //console.log("something is messed up with the equalizerMode = this.$.equalizerSelector.selected;");
      }

      return typeFound;
  }


  $(document).ready(function() {
    //console.log("equalizer zoneFlag", zoneFlag);
    sliders = document.getElementsByClassName("ui-slider-range ui-widget-header ui-corner-all ui-slider-range-min");
    handles = document.getElementsByClassName("ui-slider-handle ui-state-default ui-corner-all");
    background = document.getElementsByClassName("custom_slider style-scope equalizer-controller ui-slider ui-slider-vertical ui-widget ui-widget-content ui-corner-all");

    //console.log(sliders, "\n", handles, "\n", background);

  });
  
  var lastSelected = 0;
  var isThereAnError = false;



  var lastEQMode = "masking";
  var iirConstants = [];

  for(loopVar = 0; loopVar < 63; loopVar++) {
    iirConstants[loopVar] = 0;
  }
  //console.log(iirConstants);

  /**
   * This function is used to show the switching dialog that notifies the user of the implications of their changes when switching equalizers.
   */
  function showEqualizerSwitchDialog(mode) {
    //console.log(mode);
    switch(mode) {
      case "masking": {
        document.getElementById("eqmode").innerHTML = "You are switching the speaker to masking mode<br>In this mode, the speaker will produce equalized white noise and unequalized music.";
        break;
      }
      case "music": {
        document.getElementById("eqmode").innerHTML = "You are switching the speaker to music mode<br>In this mode the speaker will produce equalized music only.";
        break;
      }
      case "paging": {
        document.getElementById("eqmode").innerHTML = "You are switching the speaker to paging mode;<br>Something needs to go here, ask Matt"
        break;
      }
    }
    document.querySelector("#switchingEQModeDialog").toggle();
    //console.log("i got here");
  }

  /**
   * This function is used for changing the equalizer mode. It also takes care of some UI switching as well, such as hiding and showing the appropriate menus for the equalizers.
   */
  function modeChanged(mode) {

    switch(mode) {
      case "masking":
        equalizerMenu.style.visibility = "visible";
        musicMenu.style.visibility = "hidden";
        pagingMenu.style.visibility = "hidden";
        showEqualizerSwitchDialog("masking");
        break;

      case "music":
        equalizerMenu.style.visibility = "hidden";
        musicMenu.style.visibility = "visible";
        pagingMenu.style.visibility = "hidden";
        showEqualizerSwitchDialog("music");
        break;
      
      case "paging":
        equalizerMenu.style.visibility = "hidden";
        musicMenu.style.visibility = "hidden";
        pagingMenu.style.visibility = "visible";
        showEqualizerSwitchDialog("paging");
        break;

      default: 
        //console.log("SOMETHING WAS WRONG WITH WHAT WAS PASSED TO modeChanged()");
    }
  }


    Polymer({
      is: 'equalizer-controller',
      
      properties: {
        internalEqualizer: String,

        mode: {
        type:  String,
        value: "masking"
        },

        speakerId: {
          type: Number
        },

        speaker: {
          type: Object
        },

        equalizer: {
          type: String,
          notify: true,
          observer: '_setInternalEqualizer'
        }
      },

      /**
       * Setting the internal equalizer has internal effects so that the system can keep track of itself. It sets a variable that when set fires a function to update the server of its value.
       */
      _setInternalEqualizer: function(iirConstantsString) {
        this.internalEqualizer = iirConstantsString;
        //console.log("_setInternalEqualizer " + iirConstantsString);
      },
      
      /**
       * This function sets the equalizer value and is the function that is fired when the user presses the 'set equalizer' button. This function validates the bands to make sure that correct values are being submitted.
       */
      _setEqualizerClicked: function() {

        //console.log("equalizer clicked zoneFlag", zoneFlag);

        var checkInt = 0;
        var checkBands = [];
        var equalizerText = "Invalid input at ";

        $(document.getElementsByClassName("level")).each(function(i) {
          typeof this.value;
          if(isNaN(this.value) || this.value == "" || this.value < -40 || this.value > 10) {
            //console.log(this, this.name, this.value);
            checkInt++;
            checkBands.push(this.name);
          }
        });

        if(checkInt > 0) {
          if(checkBands.length == 1) {
            equalizerText += checkBands[0] + "Hz";
          } else {
            for (var i = 0; i < checkBands.length - 1; i++) {
              equalizerText += checkBands[i] + "Hz, "
            }
            equalizerText += "and " + checkBands[checkBands.length - 1] + "Hz";
          }
          this.$.equalizerToast.duration = 2000 + (500 * checkBands.length);
          this.$.equalizerToast.text = equalizerText;
          this.$.equalizerToast.show();
        } else {


          this.equalizer = iirConstants.slice(0 + (21 * typeFound), 21 + (21 * typeFound)).join(" ");
          //console.log(this.equalizer);
          this.iirConstantsString = iirConstants.slice(0 + (21 * typeFound), 21 + (21 * typeFound)).join(", ");


          this.$.equalizerToast.text = "Equalizer set to " + this.iirConstantsString;
          this.$.equalizerToast.show();
          //console.log("_setEqualizerClicked");
          //console.log(iirConstants);
          
          for (var i = 0; i < this.speaker.CurrentPreset.length; i++) { 
            this.speaker.CurrentPreset[i] = iirConstants[i];
          }
        }
      },

      _logResponse: function(event, data) {
        //console.log("shuddup");
        //console.log(event, data);
      },

      /**
       * This function is used when the the user adds a preset and clicks the 'add' button on the assisting dialog.
       */
      _acceptPressed: function(event) {
        var name = this.$.popupInput.value.valueOf().replace("\n", "").replace("\t", "");
        var ifvar = 0;
        if(this.speaker.PresetNames.length > 0) {
          for(var i = 0; i < this.speaker.PresetNames.length; i++) {
            if(name.toLowerCase() == this.speaker.PresetNames[i].valueOf().toLowerCase()) {
              //console.log("same name dude");
              ifvar++;
            }
          }
        }

        if(ifvar > 0) {
          //console.log("we are in here right now");
          this.err = "Invalid name: Name already taken";
          this.$.toasty.text = this.err;
          this.$.toasty.show();
        } else {
          this.$.addPresetAjax.body = {
            "speaker": this.speakerId,
            "name": name,
            "type": getValues(),
            "constants": iirConstants.join(" "),
          };

          //console.log(iirConstants.join(" "));

          //console.log("addPresetAjax:", addPresetAjax.body);
          this.$.addPresetAjax.generateRequest();

          this.$.dialog.toggle();

          this.$.toasty.text = "Added preset: " + name + " " + iirConstants.join(", ");
          //console.log(this.$.toasty.text);

          this.$.toasty.show();

          temp = document.createElement("paper-item");
          temp.setAttribute("value", name);
          temp.setAttribute("id", name);
          temp.innerHTML = name;
          Polymer.dom(this.$.equalizerListbox).appendChild(temp);

          var length = this.speaker.Equalizer.length;
          //console.log(length);

          //console.log(iirConstants);

          var iirConstantsCopy = iirConstants.slice();

          //console.log("iirConstantsCopy: ", iirConstantsCopy);
          this.speaker.Equalizer.push(iirConstants.slice());
          //}

          this.speaker.PresetNames[this.speaker.PresetNames.length] = name;

          //console.log(this.speaker.Equalizer);

          this.$.equalizerListbox.selected = name;
          this.$.popupInput.value = "";
        }
      },

      /**
       * This function dismisses the dialog and resets the selected preset.
       */
      _cancelPressed: function(event) {
        this.$.equalizerListbox.selected = -1;
        //console.log(lastSelected);
      },

      /**
       * This function is used when the the user accepts the implications of their selection using the 'switch' button on the dialog.
       */
      _switchPressed: function(event) {
        this.$.changeEQModeAjax.body = {
          "speaker":  this.speakerId,
          "mode":     getValues(),
        };

        this.$.changeEQModeAjax.generateRequest();

        this.$.switchingEQModeDialog.toggle();
        lastEQMode = this.$.equalizerSelector.selected;
        //console.log(lastEQMode);
      },

      /**
       * This function is used when the user decided against switching equalizers and clicks the cancel button on the switch dialog.
       */
      _cancelSwitchPressed: function(event) {
        switch(lastEQMode) {
          case "masking":
            equalizerMenu.style.visibility = "visible";
            musicMenu.style.visibility = "hidden";
            pagingMenu.style.visibility = "hidden";
            break;

          case "music":
            equalizerMenu.style.visibility = "hidden";
            musicMenu.style.visibility = "visible";
            pagingMenu.style.visibility = "hidden";
            break;

          case "paging":
            equalizerMenu.style.visibility = "hidden";
            musicMenu.style.visibility = "hidden";
            pagingMenu.style.visibility = "visible";
            break;

          default:
        }
        this.$.equalizerSelector.selected = lastEQMode;
      },

      listeners: {
        'equalizerListbox.iron-select': '_presetSelectionChanged'
      },

      /**
       * PresetSelectionChanged is fired when the user selects a new preset that is then used to set the equalizer to the corresponding values.
       */
      _presetSelectionChanged: function(event) {
        //console.log(this.speakerId);
        //console.log(this.speaker);

        var selected = this.$.equalizerListbox.selected.valueOf().toLowerCase().replace("\n", "");

        if(selected == "currently applied") {
          this.$.toasty.text = "Currently applied constants loaded! " + iirConstants.join(", ");
          this.$.toasty.show();

          lastSelected = this.$.equalizerListbox.selected;

          for (var i = 0; i < this.speaker.CurrentPreset.length; i++) { 
            sliders[i].style.height = (this.speaker.CurrentPreset[i] + 40) * 2 + "%";
            //console.log(this);
            handles[i].style.bottom = (this.speaker.CurrentPreset[i] + 40) * 2 + "%";
            $("#text" + (parseInt(i) + 1)).val(this.speaker.CurrentPreset[i]);
            iirConstants[i] = this.speaker.CurrentPreset[i];       
          }
        } else if(selected == "add new preset") {
            var addPresetAjax = document.getElementById("addPresetAjax");
            //console.log(addPresetAjax);
            this.$.dialog.toggle();
        } else if(this.speaker.PresetNames.length > 0) {
            lastSelected = this.$.equalizerListbox.selected;
            var associationLink = -1;
            for (k in this.speaker.PresetNames) {
              if(selected == this.speaker.PresetNames[k].valueOf().toLowerCase()) {
                this.$.toasty.text = this.$.equalizerListbox.selected + " preset loaded! " + iirConstants.join(", ");
                this.$.toasty.show();
                //console.log("Speaker " + this.speakerId + " selected preset: " + selected);
                associationLink = k;
                //console.log("associationLink", associationLink);
                break;
              }
            }
            if(associationLink != -1) {
              //console.log(this.speaker.Equalizer.length);
              for (var j = 0; j < this.speaker.Equalizer[0].length; j++) {
                sliders[j].style.height = (this.speaker.Equalizer[associationLink][j] + 40) * 2 + "%";
                handles[j].style.bottom = (this.speaker.Equalizer[associationLink][j] + 40) * 2 + "%";
                $("#text" + (parseInt(j) + 1)).val(this.speaker.Equalizer[associationLink][j]);
                iirConstants[j] = this.speaker.Equalizer[associationLink][j];
              }
            }
          }
      },
      
      resetKnob: function(elem) {
        elem.prevTimeout = 0;
        elem.$.volume2._resetKnob();
      },

      showToast: function() {
        //console.log("gi");
      },

      /**
       * On ready, the function takes the equalizer constants and sets the internal equalizer and assigns the toasts to global values so that they are accessible outside the selectors scope.
       */
      ready: function() {
        //console.log("this.$", this.$);
        this.internalEqualizer = this.iirConstantsString;
        equalizerToast = this.$.equalizerToast;
        equalizerMode = this.$.equalizerSelector;
      }
    });

    function showToast() {
      //console.log("gi");
    }

    /**
     * InputChanged is used when the user types a value into the text box for the slider. This function will validate that the value entered is completely numeric (ie, integer or float) and is between 10 and -40.
     */
    function inputChanged() {
        var error = false;
      $(document.getElementsByClassName("level")).each(function(i) {
        //console.log(this.value);
        if(isNaN(this.value) || this.value == "" || this.value < -40 || this.value > 10) {
          error = true;
          sliders[i].style.background = "#F33F31";
          handles[i].style.background = "#F33F31";
          background[i].style.background = "rgba(243, 63, 49, .2)";
          
        } else {        
          iirConstants[i] = parseInt(this.value);
          sliders[i].style.background = "#3f51b5";
          handles[i].style.background = "#3f51b5";
          background[i].style.background = "rgba(63, 81, 181, .2)";

          sliders[i].style.height = (iirConstants[i] + 40) * 2 + "%";
          handles[i].style.bottom = (iirConstants[i] + 40) * 2 + "%";
        }
      });
        //console.log(error);

        if(error == false) {
          isThereAnError = false;
        } else {
          isThereAnError = true;
        }

        //console.log(isThereAnError);
    }

  var MIN_WIDTH_SUPPORTED = 0;
  var MIN_HEIGHT_SUPPORTED = 0;

  if($(window).width() > MIN_WIDTH_SUPPORTED || $(window).height() > MIN_HEIGHT_SUPPORTED) {
    $(function() {
      $(".custom_slider").each(function(i) {
        //console.log("this is me, the eq: ", i);
        var texts = $("#text" + (i + 1) + ":nth-child(1)");
        $(this).slider({
          orientation: "vertical",
          range: "min",
          step: .01,
          min: -40,
          max: 10,
          value: 0,
          slide: function( event, ui) {
            isThereAnError = false;
            sliders[i].style.background = "#3f51b5";
            handles[i].style.background = "#3f51b5";
            background[i].style.background = "rgba(63, 81, 181, .2)";
            //console.log(texts);
            $(texts[0]).val(ui.value);

            iirConstants[i] = ui.value;
          }
        });
        $(texts[0]).val($(".custom > .custom_slider").slider("value"));
        //console.log(texts);
      });
    });
  } else {
      //console.log("there is not enough space to display the equalizer");     
    }
  $(window).resize(function() {
    if($(window).width() < MIN_WIDTH_SUPPORTED || $(window).height() < MIN_HEIGHT_SUPPORTED) {
      //console.log($(window).width());

      //console.log(document.getElementById("eighty").offsetWidth);
    } else {
      var custom_container = document.getElementById("custom_container");
    }
  });

