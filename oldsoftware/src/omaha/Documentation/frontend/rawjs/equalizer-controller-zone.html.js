
  var slidersZone;
  var handlesZone;
  var backgroundZone;

  var equalizerToastZone = 0;
  var equalizerModeZone;
  var typeFoundZone = 0;

  /**
   * This is the zone version of getValues.
   * GetValues is used to get the type of equalizer that is currently selected.
   * @returns {Number} typeFoundZone is an integer that represents the type of equalizer for the zone.
   */
  function getValuesZone() {
    switch (equalizerModeZone.selected) {
      case "masking": 
        typeFoundZone = 0;
        //console.log(equalizerModeZone.selected);
        break;
      case "music":
        typeFoundZone = 1;
        //console.log(equalizerModeZone.selected);
        break;
      case "paging":
        typeFoundZone = 2;
        //console.log(equalizerModeZone.selected);
        break;

      default:
        //console.log("something is messed up with the equalizerModeZone = this.$.equalizerSelectorZone.selected;");
      }

      return typeFoundZone;
  }


  $(document).ready(function() {
    //console.log("equalizer zoneFlag", zoneFlag);
    slidersZone = document.getElementsByClassName("ui-slider-range ui-widget-header ui-corner-all ui-slider-range-min");
    handlesZone = document.getElementsByClassName("ui-slider-handle ui-state-default ui-corner-all");
    backgroundZone = document.getElementsByClassName("custom_slider_zone style-scope equalizer-controller-zone ui-slider ui-slider-vertical ui-widget ui-widget-content ui-corner-all");
    //console.log(slidersZone, "\n", handlesZone, "\n", background);
  });
  
  var lastSelectedZone = 0;
  var isThereAnErrorZone = false;

  var lastEQModeZone = "masking";
  var iirConstantsZone = [];

  for(loopVar = 0; loopVar < 63; loopVar++) {
    iirConstantsZone[loopVar] = 0;
  }
  //console.log(iirConstantsZone);

  /**
   * This is the zone version of the showEqualizerSwitchDialog function.
   * This function is used to show the switching dialog that notifies the user of the implications of their changes when switching equalizers.
   */
  function showEqualizerSwitchDialogZone(mode) {
    //console.log(mode);
    switch(mode) {
      case "masking": {
        document.getElementById("eqmodeZone").innerHTML = "You are switching the speaker to masking mode<br>In this mode, the speaker will produce equalized white noise and unequalized music.";
        break;
      }
      case "music": {
        document.getElementById("eqmodeZone").innerHTML = "You are switching the speaker to music mode<br>In this mode the speaker will produce equalized music only.";
        break;
      }
      case "paging": {
        document.getElementById("eqmodeZone").innerHTML = "You are switching the speaker to paging mode;<br>Something needs to go here, ask Matt"
        break;
      }
    }
    document.querySelector("#switchingEQModeDialogZone").toggle(); 
    //console.log("i got here zone");
  }

  /**
   * This is the zone version of the modeChanged function.
   * This function is used for changing the equalizer mode. It also takes care of some UI switching as well, such as hiding and showing the appropriate menus for the equalizers.
   */
  function modeChangedZone(mode) {

    switch(mode) {
      case "masking":
        musicMenuZone.style.visibility = "hidden";
        equalizerMenuZone.style.visibility = "visible";
        pagingMenuZone.style.visibility = "hidden";
        showEqualizerSwitchDialogZone("masking");
        break;

      case "music":
        equalizerMenuZone.style.visibility = "hidden";
        musicMenuZone.style.visibility = "visible";
        pagingMenuZone.style.visibility = "hidden";
        showEqualizerSwitchDialogZone("music");
        break;
      
      case "paging":
        musicMenuZone.style.visibility = "hidden";
        equalizerMenuZone.style.visibility = "hidden";
        pagingMenuZone.style.visibility = "visible";
        showEqualizerSwitchDialogZone("paging");
        break;

      default: 
        //console.log("SOMETHING WAS WRONG WITH WHAT WAS PASSED TO modeChangedZone()");
    }
  }


    Polymer({
      is: 'equalizer-controller-zone',
      
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
          observer: '_setInternalEqualizerZone'
        }
      },

      _setInternalEqualizerZone: function(iirConstantsStringZone) {
        this.internalEqualizer = iirConstantsStringZone;
        //console.log("_setInternalEqualizerZone " + iirConstantsStringZone);
      },
      
      /**
       * This is the zone version of the _setEqualizerClicked function.
       * This function sets the equalizer value and is the function that is fired when the user presses the 'set equalizer' button. This function validates the bands to make sure that correct values are being submitted.
       */
      _setEqualizerClickedZone: function() {

        //console.log("equalizer clicked zoneFlag", zoneFlag);

        var checkIntZone = 0;
        var checkBandsZone = [];
        var equalizerTextZone = "Invalid input at ";

        $(document.getElementsByClassName("levelZone")).each(function(i) {
          typeof this.value;
          if(isNaN(this.value) || this.value == "" || this.value < -40 || this.value > 10) {
            //console.log(this.name);
            checkIntZone++;
            checkBandsZone.push(this.name);
          }
        });

        if(checkIntZone > 0) {
          if(checkBandsZone.length == 1) {
            equalizerTextZone += checkBandsZone[0] + "Hz";
          } else {
            for (var i = 0; i < checkBandsZone.length - 1; i++) {
              equalizerTextZone += checkBandsZone[i] + "Hz, "
            }
            equalizerTextZone += "and " + checkBandsZone[checkBandsZone.length - 1] + "Hz";
          }
          this.$.equalizerToastZone.duration = 2000 + (500 * checkBandsZone.length);
          this.$.equalizerToastZone.text = equalizerTextZone;
          this.$.equalizerToastZone.show();
        } else {
          //console.log(this.equalizer);
          this.iirConstantsStringZone = iirConstantsZone.slice(0 + (21 * typeFoundZone), 21 + (21 * typeFoundZone)).join(", ");


          this.$.equalizerToastZone.text = "Equalizer set to " + this.iirConstantsStringZone;
          this.$.equalizerToastZone.show();
          //console.log("_setEqualizerClickedZone"); 
          //console.log(this.iirConstantsStringZone);
          
          // make this zone wide
          for (var i = 0; i < speaker.CurrentPreset.length; i++) { 
            //speaker.CurrentPreset[i] = iirConstantsZone[i];
          }

          //console.log(iirConstantsZone.slice(0 + (21 * typeFoundZone), 21 + (21 * typeFoundZone)).length);

          this.$.updateEqualizerZone.body = {
            "zone": speaker.zone.zoneID,
            "constants": iirConstantsZone.slice(0 + (21 * typeFoundZone), 21 + (21 * typeFoundZone)).join(" "),
            "mode": getValuesZone()
          };
          this.$.updateEqualizerZone.generateRequest();
        }
      },

      _logResponse: function(event, data) {
        //console.log("shuddup");
        //console.log(event, data);
      },

      /**
       * This function is used when the the user adds a preset and clicks the 'add' button on the assisting dialog.
       */
      _acceptPressedZone: function(event) {
        var name = this.$.popupInputZone.value.valueOf().replace("\n", "").replace("\t", "");
        var ifvar = 0;
        if(speaker.zone.PresetNames.length > 0) {
          for(var i = 0; i < speaker.zone.PresetNames.length; i++) {
            if(name.toLowerCase() == speaker.zone.PresetNames[i].valueOf().toLowerCase()) {
              //console.log("same name dude");
              ifvar++;    // make more checks dummy
            }
          }
        }

        if(ifvar > 0) {
          //console.log("we are in here right now");
          this.err = "Invalid name: Name already taken";
          this.$.toastyZone.text = this.err;
          this.$.toastyZone.show();
        } else {
          this.$.addPresetAjaxZone.body = {
            "zone": speaker.zone.zoneID,
            "name": name,
            "type": getValuesZone(),
            "constants": iirConstantsZone.slice(0 + (21 * typeFoundZone), 21 + (21 * typeFoundZone)).join(" "),
          };

          //console.log(iirConstantsZone.join(" "));

          //console.log("addPresetAjaxZone:", addPresetAjaxZone.body);
          this.$.addPresetAjaxZone.generateRequest();

          this.$.dialogZone.toggle();

          this.$.toastyZone.text = "Added preset: " + name + " " + iirConstantsZone.slice(0 + (21 * typeFoundZone), 21 + (21 * typeFoundZone)).join(" ");
          //console.log(this.$.toastyZone.text);

          this.$.toastyZone.show();

          temp = document.createElement("paper-item");
          temp.setAttribute("value", name);
          temp.setAttribute("id", name);
          temp.innerHTML = name;
          Polymer.dom(this.$.equalizerListboxZone).appendChild(temp);

          var length = speaker.zone.Equalizer.length;
          //console.log(length);

          //console.log(iirConstantsZone);

          var iirConstantsCopy = iirConstantsZone.slice();

          //console.log("iirConstantsCopy: ", iirConstantsCopy);
          speaker.zone.Equalizer.push(iirConstantsZone.slice(0 + (21 * typeFoundZone), 21 + (21 * typeFoundZone)));
          //}

          speaker.zone.PresetNames[speaker.zone.PresetNames.length] = name;

          //console.log(speaker.zone.Equalizer);

          this.$.equalizerListboxZone.selected = name;
          this.$.popupInputZone.value = "";
        }
      },

      /**
       * This function dismisses the dialog and resets the selected preset.
       */
      _cancelPressedZone: function(event) {
        this.$.equalizerListboxZone.selected = -1;
        //console.log(lastSelectedZone);
      },

      /**
       * This function is used when the the user accepts the implications of their selection using the 'switch' button on the dialog.
       */
      _switchPressedZone: function(event) {
        this.$.changeEQModeAjaxZone.body = {
          "zone":  speaker.zone.zoneID,
          "mode":  getValuesZone()
        };

        this.$.changeEQModeAjaxZone.generateRequest();

        this.$.switchingEQModeDialogZone.toggle();
        lastEQModeZone = this.$.equalizerSelectorZone.selected;
        //console.log(lastEQModeZone);
      },

     /**
       * This function is used when the user decided against switching equalizers and clicks the cancel button on the switch dialog.
       */
      _cancelSwitchPressedZone: function(event) {
        switch(lastEQModeZone) {
          case "masking":
            musicMenuZone.style.visibility = "hidden";
            pagingMenuZone.style.visibility = "hidden";
            equalizerMenuZone.style.visibility = "visible";
            break;

          case "music":
            musicMenuZone.style.visibility = "visible";
            equalizerMenuZone.style.visibility = "hidden";
            pagingMenuZone.style.visibility = "hidden";
            break;

          case "paging":
            pagingMenuZone.style.visibility = "visible";
            equalizerMenuZone.style.visibility = "hidden";
            musicMenuZone.style.visibility = "hidden";


          default:
        }
        this.$.equalizerSelectorZone.selected = lastEQModeZone;
      },

      listeners: {
        'equalizerListboxZone.iron-select': '_presetSelectionChangedZone'
      },

      /**
       * PresetSelectionChanged is fired when the user selects a new preset that is then used to set the equalizer to the corresponding values.
       */
      _presetSelectionChangedZone: function(event) {
        //console.log(this.speakerId);
        //console.log(this.speaker);

        var selected = this.$.equalizerListboxZone.selected.valueOf().toLowerCase().replace("\n", "");

        if(selected == "currently applied") {
          this.$.toastyZone.text = "Currently applied constants loaded! " + iirConstantsZone.join(", ");
          this.$.toastyZone.show();

          lastSelectedZone = this.$.equalizerListboxZone.selected;

          for (var i = 0; i < speaker.zone.currentPreset.length; i++) { 
            slidersZone[i].style.height = (speaker.zone.currentPreset[i] + 40) * 2 + "%";
            //console.log(this);
            handlesZone[i].style.bottom = (speaker.zone.currentPreset[i] + 40) * 2 + "%";
            $("#text" + (parseInt(i) + 1)).val(speaker.zone.currentPreset[i]);
            iirConstantsZone[i] = speaker.zone.currentPreset[i];       
          }
        } else if(selected == "add new preset") {
            var addPresetAjaxZone = document.getElementById("addPresetAjaxZone");
            //console.log(addPresetAjaxZone);
            this.$.dialogZone.toggle();
        } else if(speaker.zone.PresetNames.length > 0) {
            lastSelectedZone = this.$.equalizerListboxZone.selected;
            var associationLink = -1;
            for (k in speaker.zone.PresetNames) {
              if(selected == speaker.zone.PresetNames[k].valueOf().toLowerCase()) {
                this.$.toastyZone.text = this.$.equalizerListboxZone.selected + " preset loaded! " + iirConstantsZone.slice(0 + (21 * typeFoundZone), 21 + (21 * typeFoundZone)).join(", ");
                this.$.toastyZone.show();
                //console.log("Zone " + speaker.zone.zoneID + " selected preset: " + selected);
                associationLink = k;
                //console.log("associationLink", associationLink);
                break;
              }
            }
            if(associationLink != -1) {
              //console.log(speaker.zone.Equalizer.length);
              for (var j = 0; j < speaker.zone.Equalizer[0].length; j++) {
                slidersZone[j + 63].style.height = (speaker.zone.Equalizer[associationLink][j] + 40) * 2 + "%";
                handlesZone[j + 63].style.bottom = (speaker.zone.Equalizer[associationLink][j] + 40) * 2 + "%";
                $("#text" + (parseInt(j) + 1)).val(speaker.zone.Equalizer[associationLink][j]);
                iirConstantsZone[j] = speaker.zone.Equalizer[associationLink][j];
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
        this.internalEqualizer = this.iirConstantsStringZone;
        equalizerToastZone = this.$.equalizerToastZone;
        equalizerModeZone = this.$.equalizerSelectorZone;
      }
    });

    function showToast() {
      //console.log("gi");
    }

    /**
     * InputChanged is used when the user types a value into the text box for the slider. This function will validate that the value entered is completely numeric (ie, integer or float) and is between 10 and -40.
     */
    function inputChangedZone() {
        var error = false;
      $(document.getElementsByClassName("levelZone")).each(function(i) {
        //console.log(this.value);
        if(isNaN(this.value) || this.value == "" || this.value < -40 || this.value > 10) {
          error = true;
          slidersZone[i].style.background = "#F33F31";
          handlesZone[i].style.background = "#F33F31";
          backgroundZone[i].style.background = "rgba(243, 63, 49, .2)";
          
        } else {        
          iirConstantsZone[i] = parseInt(this.value);
          slidersZone[i].style.background = "#3f51b5";
          handlesZone[i].style.background = "#3f51b5";
          backgroundZone[i].style.background = "rgba(63, 81, 181, .2)";

          slidersZone[i].style.height = (iirConstantsZone[i] + 40) * 2 + "%";
          handlesZone[i].style.bottom = (iirConstantsZone[i] + 40) * 2 + "%";
        }
      });
        //console.log(error);

        if(error == false) {
          isThereAnErrorZone = false;
        } else {
          isThereAnErrorZone = true;
        }

        //console.log(isThereAnErrorZone);
    }

  var MIN_WIDTH_SUPPORTED = 0;
  var MIN_HEIGHT_SUPPORTED = 0;

  if($(window).width() > MIN_WIDTH_SUPPORTED || $(window).height() > MIN_HEIGHT_SUPPORTED) {
    $(function() {
      $(".custom_slider_zone").each(function(i) {
        j = i - 63;
        //console.log("this is me, the eq: ", i);
        var texts = $("#text" + (i + 1) +":nth-child(1)");
        $(this).slider({
          orientation: "vertical",
          range: "min",
          step: .01,
          min: -40,
          max: 10,
          value: 0,
          slide: function( event, ui) {
            isThereAnErrorZone = false;
            slidersZone[i].style.background = "#3f51b5";
            handlesZone[i].style.background = "#3f51b5";
            //console.log(i);
            //console.log(texts);
            backgroundZone[i].style.background = "rgba(63, 81, 181, .2)";
            $(texts[1]).val(ui.value);

            iirConstantsZone[i] = ui.value;
          }
        });
        $(texts[1]).val($(".custom > .custom_slider_zone").slider("value"));
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

