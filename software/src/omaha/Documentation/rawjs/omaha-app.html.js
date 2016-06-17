
    var zoneFlag = 0;
    var globalZone = 0;

    var localLevel = parseInt(localStorage.getItem("level"));

    var localSpeakerID = parseInt(localStorage.getItem("speakerid"));
    var localZoneID = parseInt(localStorage.getItem("zoneid"));


    var lastSelected = 0;
    var zones;
    var pagingZones;
    var displayedZones = {};
    
    Polymer({
      is: 'omaha-app',

      properties: {        
        selected: {
          type: String,
          value: "map"
        },
        zones: Object,
        controllers: Object
      },

      listeners: {
        'nav.backpressed': '_showMap',
        'drawer.showmap': '_showMap',
        'drawer.iron-select': '_closeDrawer',
        'nav.toggledrawer': '_toggleDrawer',
        'mapcontroller.speakerselected': '_speakerSelected',
        'mapcontroller.zoneselected': '_zoneSelected',
        'controller.equalizer.equalizerMenu.iron-select': '_presetSelectionChanged'
        // put a mapcontroller.logresponse thing here
      },

      /**
        * This function displays the map.
      */
      _showMap: function() {
        this.selected = "map";
      },

      _presetSelectionChanged: function(event) {    // Why does this have event in the thing? It is not like that in volume selector
        //console.log("hi");    // This says the speakerId is undefined
        // dont think we need this
        },

        /**
          * This function is fired when the mapcontroller detects that a speaker was selected by the user; when the mapcontroller.speakerselected fires.
        */
      _speakerSelected: function(e) {
        //console.log("this is where i am");
        //console.log(this.$);
        var speaker = e.detail.speaker;
        //console.log(e);
        //console.log(speaker);
        //console.log(localStorage.getItem("speakerid"));

        if(speaker === undefined) {
          this.$.dialog.open();
        } else if(speaker.id == localSpeakerID || speaker.zone.zoneID == localZoneID || localZoneID == 1 || localLevel > 0) {
          this.$.controller.speaker = speaker;
          //console.log(this.$.speaker);
          this.$.nav.showBackArrow();

          this.selected = "speakerController";

          // Get controller object
          var controller = document.getElementById('controller');

          // Set volume levels
          document.getElementById('volume2').value = this.$.controller.speaker.VolumeLevel[0];
          document.getElementById('volume3').value = this.$.controller.speaker.VolumeLevel[1];
          document.getElementById('volume4').value = this.$.controller.speaker.VolumeLevel[2];
          document.getElementById('volume5').value = this.$.controller.speaker.VolumeLevel[3];

          document.getElementById('musicMin').value = this.$.controller.speaker.LowerThreshold[0];
          document.getElementById('musicMax').value = this.$.controller.speaker.UpperThreshold[0];
          document.getElementById('pagingMin').value = this.$.controller.speaker.LowerThreshold[1];
          document.getElementById('pagingMax').value = this.$.controller.speaker.UpperThreshold[1];
          document.getElementById('maskingMin').value = this.$.controller.speaker.LowerThreshold[2];
          document.getElementById('maskingMax').value = this.$.controller.speaker.UpperThreshold[2];

          //console.log(document.getElementById('equalizerMenu'));

          // Set averaging mode
          document.getElementById('effectiveness').value = this.$.controller.speaker.Effectiveness;
          document.getElementById('pleasantness').value = this.$.controller.speaker.Pleasantness;

          // Set paging levels
          document.getElementById('pagingvolume').value = this.$.controller.speaker.VolumeLevel[2];
          document.getElementById("fadetime").value = this.$.controller.speaker.FadeTime;
          document.getElementById("fadelevel").value = this.$.controller.speaker.FadeLevel;

          // Set equalizer bands
          var sliders = document.getElementsByClassName("ui-slider-range ui-widget-header ui-corner-all ui-slider-range-min");
          var handles = document.getElementsByClassName("ui-slider-handle ui-state-default ui-corner-all");

          for (i in controller.speaker.CurrentPreset) {   // this was "current" for some reason
            sliders[i].style.height = (controller.speaker.CurrentPreset[i] + 40) * 2 + "%";    // why does this have to be lower case
            handles[i].style.bottom = (controller.speaker.CurrentPreset[i] + 40) * 2 + "%";
            $("#text" + (parseInt(i) + 1)).val(controller.speaker.CurrentPreset[i]);
            iirConstants[i] = controller.speaker.CurrentPreset[i];
          }
          for (var i = 0; i < controller.speaker.CurrentMusicPreset.length; i++) {   // this was "current" for some reason
            //console.log("I am here");
            sliders[i + 21].style.height = (controller.speaker.CurrentMusicPreset[i] + 40) * 2 + "%";
            handles[i + 21].style.bottom = (controller.speaker.CurrentMusicPreset[i] + 40) * 2 + "%";
            $("#text" + (parseInt(i + 21) + 1)).val(controller.speaker.CurrentMusicPreset[i]);
            iirConstants[i + 21] = controller.speaker.CurrentMusicPreset[i];
          }
          for (var i = 0; i < controller.speaker.CurrentPagingPreset.length; i++) {   // this was "current" for some reason
            //console.log("I am here");
            sliders[i + 42].style.height = (controller.speaker.CurrentPagingPreset[i] + 40) * 2 + "%";
            handles[i + 42].style.bottom = (controller.speaker.CurrentPagingPreset[i] + 40) * 2 + "%";
            $("#text" + (parseInt(i + 42) + 1)).val(controller.speaker.CurrentPagingPreset[i]);
            iirConstants[i + 42] = controller.speaker.CurrentPagingPreset[i];
          }

          for (var i = 0; i < this.$.controller.speaker.CurrentTarget.length; i++) {
            dataArray[i] = this.$.controller.speaker.CurrentTarget[i];

            dataArrayZone[i] = 0;
          }

          var equalizerListbox = document.getElementById("equalizerListbox");
          //console.log(Polymer.dom(equalizerListbox).children);

          var motherPolymer = Polymer.dom(equalizerListbox).children; 

          // Remove the paper-items from the last time and the other speakers, basically throwing away the stuff that we dont need
          if(motherPolymer.length > 1) {
            for (var i = motherPolymer.length - 1; i > 1; i--)
              Polymer.dom(equalizerListbox).removeChild(motherPolymer[i]);
          }

          // Injecting the paper-items for the presets

          if(controller.speaker.PresetNames.length > 0) {
            for (i in controller.speaker.PresetNames) {
              temp = document.createElement("paper-item");
              temp.setAttribute("value", controller.speaker.PresetNames[i].valueOf().replace("\n", ""));
              temp.setAttribute("id", controller.speaker.PresetNames[i].valueOf().replace("\n", ""));
              // temp.setAttribute("speakerId", controller.speaker.speakerId); try doing some optimization stuff with the speakerId
              temp.innerHTML = controller.speaker.PresetNames[i].valueOf().replace("\n", "");
              Polymer.dom(equalizerListbox).appendChild(temp);
            }
          }


          var musicListbox = document.getElementById("musicListbox");
          //console.log(Polymer.dom(musicListbox).children);
          var motherPolymer = Polymer.dom(musicListbox).children; 


          // Remove the paper-items from the last time and the other speakers, basically throwing away the stuff that we dont need
          if(motherPolymer.length > 1) {
            for (var i = motherPolymer.length - 1; i > 1; i--)
              Polymer.dom(musicListbox).removeChild(motherPolymer[i]);
          }

          if(controller.speaker.MusicPresetNames.length > 0) {
            for (i in controller.speaker.MusicPresetNames) {
              temp = document.createElement("paper-item");
              temp.setAttribute("value", controller.speaker.MusicPresetNames[i].valueOf().replace("\n", ""));
              temp.setAttribute("id", controller.speaker.MusicPresetNames[i].valueOf().replace("\n", ""));
              temp.innerHTML = controller.speaker.MusicPresetNames[i].valueOf().replace("\n", "");
              Polymer.dom(musicListbox).appendChild(temp);
            }
          }


          
          var targetListbox = document.getElementById("targetListbox");
          //console.log(Polymer.dom(targetListbox).children);
          motherPolymer = Polymer.dom(targetListbox).children; 

          // Remove the paper-items from the last time and the other speakers, basically throwing away the stuff that we dont need
          if(motherPolymer.length > 1) {
            for (var i = motherPolymer.length - 1; i > 1; i--)
              Polymer.dom(targetListbox).removeChild(motherPolymer[i]);
          }

          if(controller.speaker.TargetNames.length > 0) {
            for (i in controller.speaker.TargetNames) {
              temp = document.createElement("paper-item");
              temp.setAttribute("value", controller.speaker.TargetNames[i].valueOf().replace("\n", ""));
              temp.setAttribute("id", controller.speaker.TargetNames[i].valueOf().replace("\n", ""));
              temp.innerHTML = controller.speaker.TargetNames[i].valueOf().replace("\n", "");
              Polymer.dom(targetListbox).appendChild(temp);
            }
          }
        } else {
          this.$.errorNoAccessToast.innerHTML = "You only have access to";
          if(localSpeakerID > 0) {
            this.$.errorNoAccessToast.innerHTML += " speaker " + localSpeakerID;
            if(localZoneID > 0)
              this.$.errorNoAccessToast.innerHTML += " and zone " + localZoneID;
          }
          else if(localZoneID > 0)
            this.$.errorNoAccessToast.innerHTML += " zone " + localZoneID;
          else
            this.$.errorNoAccessToast.innerHTML = "You don't have access to anything!";

          this.$.errorNoAccessToast.duration = 2000;
          this.$.errorNoAccessToast.show();
          //console.log("{username} does not have access to this speaker");

        }
    var basic_container = document.getElementById("basic_container");
    var card = document.getElementById("card");
    basic_container.style.marginLeft = (Math.abs(card.offsetWidth - basic_container.offsetWidth) / 2) - 28 + "px";

      },
      
      /**
       * This function is fired when a zone is selected. It will go through and inject all the information pertaining to the zone into the controller.
       */
      _zoneSelected: function(e) {
        var zone = e.detail.zone;
        this.$.zoneController.zone = zone;
        this.selected = "zoneController"
        this.$.nav.showBackArrow();

        var zoneVolumePage = document.getElementById("volumeZone"); 
        //console.log(volumeZone, this.volumeZone);

        // Set volume levels
        volumeZone.$.volume2.value = zone.VolumeLevel[0];
        volumeZone.$.volume3.value = zone.VolumeLevel[1];
        volumeZone.$.volume4.value = zone.VolumeLevel[2];
        volumeZone.$.volume5.value = zone.VolumeLevel[3];

        volumeZone.$.musicMin.value = zone.LowerThreshold[0];
        volumeZone.$.pagingMin.value = zone.LowerThreshold[1];
        volumeZone.$.maskingMin.value = zone.LowerThreshold[2];
        volumeZone.$.musicMax.value = zone.UpperThreshold[0];
        volumeZone.$.pagingMax.value = zone.UpperThreshold[1];
        volumeZone.$.maskingMax.value = zone.UpperThreshold[2];

        averagingZone.$.effectiveness.value = zone.effectiveness;
        averagingZone.$.pleasantness.value = zone.pleasantness;

        var targetZoneChildren = targetZone.$.adjustment.children;
        
        for(var i = 0; i < targetZoneChildren.length; i++) {
          targetZoneChildren[i].value = zone.currentTarget[i];
        }
        redrawOnInputZone();

        var levelZones = document.getElementsByClassName("levelZone");
        var i = 0;

        for(; i < levelZones.length / 3; i++) {
          levelZones[i].value = zone.currentPreset[i];
          sliders[i + 63].style.height = (zone.currentPreset[i] + 40) * 2 + "%";
          handles[i + 63].style.bottom = (zone.currentPreset[i] + 40) * 2 + "%";
        }

        for(; i < levelZones.length / 3; i++) {
          levelZones[i].value = zone.currentMusicPreset[i];
          sliders[i + 84].style.height = (zone.currentPreset[i] + 40) * 2 + "%";
          handles[i + 84].style.bottom = (zone.currentPreset[i] + 40) * 2 + "%";
        }

        for(; i < levelZones.length / 3; i++) {
          levelZones[i].value = zone.currentPagingPreset[i];
          sliders[i + 105].style.height = (zone.currentPreset[i] + 40) * 2 + "%";
          handles[i + 105].style.bottom = (zone.currentPreset[i] + 40) * 2 + "%";
        }

        pagingZone.$.fadetime.value = zone.PagingLevel[0];
        pagingZone.$.fadelevel.value = zone.PagingLevel[1];
        pagingZone.$.pagingvolume.value = zone.VolumeLevel[2];


        // Test for null / undefined values and set them to a value
        if(zone.Equalizer == undefined) {
          zone.Equalizer = [];
        }

        if(zone.MusicEqualizer == undefined) {
          zone.MusicEqualizer = [];
        }

        if(zone.MusicPresetNames == undefined) {
          zone.MusicPresetNames = [];
        }

        if(zone.PagingEqualizer == undefined) {
          zone.PagingEqualizer = [];
        }

        if(zone.PagingPresetNames == undefined) {
          zone.PagingPresetNames = [];
        }

        if(zone.PresetNames == undefined) {
          zone.PresetNames = [];
        }

        if(zone.Target == undefined) {
          zone.Target = [];
        }

        if(zone.TargetNames == undefined) {
          zone.TargetNames = [];
        }
        

        //console.log("its zonerino time", zone);
      },

      /**
       * This function sets the basic controller up, sets up the zone options in the scrollbox, as well as calling for the rendering of the maps.
       */
      _initializeStatus: function(event, data) {

        //console.log("data", data.response);


        var speakers = [];
        var speakersJustID = [];

        var callingInnerHTMLString = "";
        var innerHTMLSpeakersString = "";

        var innerHTMLString = "";
        zones = data.response.zones;
        //console.log(zones);
        for (var i = 0; i < zones.length; i++) {
          displayedZones[zones[i].name] = undefined;
        }
        var zoneListbox = document.querySelector("#zoneListbox");
        zoneListbox.innerHTML = "";

        //console.log("zones", zones);

        zones.forEach(function(zone){
          zone.speakers.forEach(function(speaker){
            speaker.zone = zone;
          });
        })

        callingInnerHTMLString += "<paper-item center-justified flex class=\"pagingZoneStyle\"><font style=\"text-align: center; width: 100%; font-weight: bold;\">Masking Zones:</font></paper-item>";
        //console.log(data.response);
        for(var i = 0; i < zones.length; i++) {
          ////console.log(zones);
          innerHTMLString += "<paper-item center-justified flex ><font style=\"width: 100%; text-align: center;\">" + zones[i].name + "</font></paper-item>";
          callingInnerHTMLString += "<paper-item center-justified flex id=\"" + zones[i].name + "\" class=\"pagingZoneStyle\" onclick=\"updateViewForSelectedSpeakers(this)\"><font style=\"text-align: center; width: 100%;\">" + zones[i].name + "</font></paper-item>";

          for (var k = 0; k < zones[i].speakers.length; k++) {
            var isItInYetThatsWhatSheSaid = speakersJustID.indexOf(zones[i].speakers[k].id);
            if(isItInYetThatsWhatSheSaid == -1) {
              speakers.push(zones[i].speakers[k]);
              speakersJustID.push(zones[i].speakers[k].id);
            } else {
              speakers[isItInYetThatsWhatSheSaid].zone = zones[i].speakers[k].zone;
            }
          }
        }
        zoneListbox.innerHTML = innerHTMLString;


        innerHTMLString = "";
        pagingZones = data.response.pagingZones;
        pagingZoneListbox.innerHTML = "";

        pagingZones.forEach(function(pagingZone){
          pagingZone.speakers.forEach(function(speaker){
            speaker.pagingZone = pagingZone;
          });
        })

        callingInnerHTMLString += "<paper-item center-justified flex><font style=\"text-align: center; width: 100%;\"></font></paper-item>";
        callingInnerHTMLString += "<paper-item center-justified flex class=\"pagingZoneStyle\"><font style=\"text-align: center; width: 100%; font-weight: bold;\">Paging Zones:</font></paper-item>";
        
        for(var i = 0; i < pagingZones.length; i++) {
          displayedZones[pagingZones[i].name] = undefined;
          innerHTMLString += "<paper-item id=\"" + pagingZones[i].name + "\" value=\"" + pagingZones[i].name + "\"center-justified flex ><font style=\"width: 100%; text-align: center;\">" + pagingZones[i].name + "</font></paper-item>";

          callingInnerHTMLString += "<paper-item id=\"" + pagingZones[i].name + "\" value=\"" + pagingZones[i].name + "\" center-justified flex class=\"pagingZoneStyle\" onclick=\"updateViewForSelectedSpeakers(this)\"><font style=\"text-align: center; width: 100%;\">" + pagingZones[i].name + "</font></paper-item>";

          for (var k = 0; k < pagingZones[i].speakers.length; k++) {
            var isItInYetThatsWhatSheSaid = speakersJustID.indexOf(pagingZones[i].speakers[k].id);
            if(isItInYetThatsWhatSheSaid == -1) {
              speakers.push(pagingZones[i].speakers[k]);
              speakersJustID.push(pagingZones[i].speakers[k].id);
            } else {
              speakers[isItInYetThatsWhatSheSaid].pagingZone = pagingZones[i].speakers[k].pagingZone;
            }
          }
        } 
        pagingZoneListbox.innerHTML = innerHTMLString;

        var targetListbox = document.getElementById("targetListbox");

        callingInnerHTMLString += "<paper-item center-justified flex><font style=\"text-align: center; width: 100%;\"></font></paper-item>";
        callingInnerHTMLString += "<paper-item center-justified flex class=\"pagingZoneStyle\"><font style=\"text-align: center; width: 100%; font-weight: bold;\">Named Speakers:</font></paper-item>";
        callingInnerHTMLString += "<paper-item center-justified flex class=\"pagingZoneStyle\" onclick=\"updateViewForSelectedSpeakers(this)\"><font style=\"text-align: center; width: 100%;\">Dummy name</font></paper-item>";



        callingInnerHTMLString += "<paper-item center-justified flex><font style=\"text-align: center; width: 100%;\"></font></paper-item>";
        callingInnerHTMLString += "<paper-item center-justified flex class=\"pagingZoneStyle\"><font style=\"text-align: center; width: 100%; font-weight: bold;\">Speakers:</font></paper-item>";

        //console.log(zones[0].speakers[0]);

        for(var k = 0; k < zones[0].speakers.length; k++) {
          displayedZones[zones[0].speakers[k].id] = undefined;
          callingInnerHTMLString += "<paper-item id=\"" + zones[0].speakers[k].id + "\" value=\"" + zones[0].speakers[k].id + "\"  center-justified flex class=\"pagingZoneStyle\" onclick=\"updateViewForSelectedSpeakers(this)\"><font style=\"text-align: center; width: 100%;\">Speaker " + zones[0].speakers[k].id + "</font></paper-item>";
        }

        var pagingSpeakerHolder = document.querySelector("#pagingSpeakerHolder");    
        pagingSpeakerHolder.innerHTML = callingInnerHTMLString;

        //console.log(this);
        this.controllers = speakers;
        this.zones = zones;
        //console.log(this.$);
        this.$.information.renderMap();
        this.$.pagingMap.renderMap();
        this.$.mapcontroller.renderMap();
        this.$.zoneMap.renderMap();
        this.$.zoneMap.renderPagingMap();
      },

      _showAccountCreation: function() {
        this.selected = "accountCreation";
        this.$.nav.showAccountCreation = false;
      },

      /**
       * This function toggles the drawer on the left side of the website.
       */
      _toggleDrawer: function() {
        this.$.panel.togglePanel();
      },
      
      /**
       * This function closes the drawer on the left side.
       */
      _closeDrawer: function() {
        this.$.panel.closeDrawer();
      },
      
      /**
       * This function will update the map size, as well as determine the selected map.
       */
      refreshMap: function() {
        switch(this.$.pages.selected) {
          case "map":
          case "pagingMap":
          case "information":
          case "zoneMap":
            var map = this.$.pages.selectedItem.map;
            //console.log("THIS IS THE MAP THING THAT YOU ARE LOOKING FOR " + map);
            if(map !== undefined)
              map.updateSize();
            break;

          default:
            //console.log("tell someone of higher authority if you see this, like me, lol");
        }
      },

      /**
       * The keep alive response is a callback for a function that schedules the speaker update querying to check whether they are still functioning or not.
       */
      _keepAliveResponse: function(event, data) {
        //console.log(data.response);
        if(data.response.status !== this.controllers[data.response.id - 1].status) { 
          updateMapForError(data.response.id, (this.controllers[data.response.id - 1].status) === 0 ? true : false);
          this.controllers[data.response.id - 1].status = data.response.status;
        }
      },

      /**
       * This function loads after the webcomponents load and sets some variables up and scheduling a self-executing function that will continuously execute itself every second.
       */
      ready: function() {       
        this.$.systemStatusAjax.generateRequest();
        this.$.nav.hideBackArrow();
        this.$.panel.forceNarrow = true;

        controllerUpdateAjax = this.$.controllerUpdateAjax;
        setUI();

        // this is asynchonous and a better usage that using setInterval
        (function keepAlive() {
          controllerUpdateAjax.generateRequest();
          setTimeout(keepAlive, 1000);
        })();
      }

    });

  /**
   * This function is used to set the UI based on the permission received when the user logs in.
   */
  function setUI() {
    switch(localLevel) {
      case 0:
        Polymer.dom(mode).removeChild(pagingButtonRC);
        Polymer.dom(mode).removeChild(schedulingButton);
        Polymer.dom(menu).removeChild(averagingButton);
        Polymer.dom(menu).removeChild(accountCreation);
        Polymer.dom(menu).removeChild(mapForZone);
        Polymer.dom(menu).removeChild(mapForPaging);
        // Just let it fall through

      case 1:
        Polymer.dom(mode).removeChild(targetButton);
        Polymer.dom(mode).removeChild(equalizerButton);
        break;

      default:
        //console.log("lel");
    }
  }
  
