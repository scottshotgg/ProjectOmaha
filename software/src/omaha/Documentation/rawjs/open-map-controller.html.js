
  var LowerThreshold = [];
  var UpperThreshold = [];

  $(document).ready(function() {
    //console.log($(document.getElementById("goodOrBadFont")).innerWidth(), $(document.getElementById("selectedSpeakers")).innerWidth());
  });


  $(window).resize(function() {
  });

  var savedResolution;
  var savedStyle = [];
  var savedStyle_i = 1;

  var controllers = []; 
  var renderMapCopy;
  var vectorLayerCopy;
  var goodCircle = document.getElementById("goodCircle");
  var goodFont = document.getElementById("goodFont"); 
  var badCircle = document.getElementById("badCircle");
  var badFont = document.getElementById("badFont"); 

    Polymer({
      is: 'open-map-controller',

      behaviors: [
        Polymer.NeonAnimatableBehavior
      ],

      properties: {
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
      
      _selectSpeaker: function() {
        this._hideOverlay();

        //console.log("i am fired");
      },
      
      _selectZone: function() {
        this._hideOverlay();
      },
      
      _hideOverlay: function() {
        //console.log(this.overlay);
      },

      _logResponse: function(event, data) {
        //console.log("this is the speaker response", data.response);
        var a = document.getElementById('speaker-tooltip');
  
        a.speaker.name = data.response.name;
        a.speaker.VolumeLevel[0] = data.response.volume;
        a.speaker.VolumeLevel[1] = data.response.music;
        a.speaker.VolumeLevel[2] = data.response.paging;
        a.speaker.VolumeLevel[3] = data.response.masking;

        a.speaker.LowerThreshold = [];
        a.speaker.UpperThreshold = [];

        a.speaker.LowerThreshold[0] = data.response.lowerMusicThreshold
        a.speaker.UpperThreshold[0] = data.response.upperMusicThreshold
        a.speaker.LowerThreshold[1] = data.response.lowerPagingThreshold
        a.speaker.UpperThreshold[1] = data.response.upperPagingThreshold
        a.speaker.LowerThreshold[2] = data.response.lowerMaskingThreshold
        a.speaker.UpperThreshold[2] = data.response.upperMaskingThreshold

        LowerThreshold[0] = data.response.lowerMusicThreshold
        UpperThreshold[0] = data.response.upperMusicThreshold
        LowerThreshold[1] = data.response.lowerPagingThreshold
        UpperThreshold[1] = data.response.upperPagingThreshold
        LowerThreshold[2] = data.response.lowerMaskingThreshold
        UpperThreshold[2] = data.response.upperMaskingThreshold

        a.speaker.Effectiveness = data.response.effectiveness;
        a.speaker.Pleasantness = data.response.pleasantness;
        a.speaker.FadeTime = data.response.fadetime;
        a.speaker.FadeLevel = data.response.fadelevel;
        a.speaker.EqualizerMode = data.response.equalizerMode;
        a.speaker.SchedulingMode = data.response.schedulingMode;

        a.speaker.Schedule = data.response.schedule;

        if(a.speaker.EqualizerMode == 1) {
          equalizerSelector.selected = "music";
          lastEQMode = "music";
          getValues();
          musicMenu.style.visibility = "visible";
          equalizerMenu.style.visibility = "hidden";
        } else if(a.speaker.EqualizerMode == 2) {
          equalizerSelector.selected = "paging"; 
          lastEQMode = "paging";
          getValues();
        }
        //console.log(schedulingSelector);
        speaker = a.speaker;

        if(a.speaker.SchedulingMode == 0) {
          schedulingSelector.selected = "basic";
          schedulingSelectorZone.selected = "basic";
          drawBasic();
          drawBasicZone();
        } else {
          schedulingSelector.selected = "advanced";
          schedulingSelectorZone.selected = "advanced";
          drawAdvanced();
          drawAdvancedZone();
        }

        //console.log(document.querySelector("#scheduling-controller"));

        //console.log("i have applied the selected value", schedulingSelector.selected, schedulingSelector);

        //console.log(this.$$);
        if(a.speaker.name !== "") {
          var tooltipInformation = "<b id=\"name\">" + a.speaker.name + "</b><br><b id=\"speaker\">Speaker ID: " + a.speaker.id + "</b><br><b id=\"zone\">Zone ID: " + a.speaker.zoneID + "</b><br><b id=\"volume\">Volume: " + a.speaker.VolumeLevel[0] + "</b>";
          //console.log("i am here");
        } else {
          //console.log("i am not here");
          var tooltipInformation = "<b id=\"speaker\">Speaker ID: " + a.speaker.id + "</b><br><b id=\"zone\">Zone ID: " + a.speaker.zone.zoneID + "</b><br><b id=\"volume\">Volume: " + a.speaker.VolumeLevel[0] + "</b>";
        }
        document.getElementById("tooltipInformation").innerHTML = tooltipInformation;

        //console.log("Response:", data.response);
        //console.log("Speaker:", a.speaker);

        a.speaker.TargetNames   = [];
        a.speaker.CurrentTarget = [];

        a.speaker.PresetNames   = [];
        a.speaker.CurrentPreset = [];

        a.speaker.MusicPresetNames   = [];
        a.speaker.CurrentMusicPreset = [];

        a.speaker.PagingPresetNames   = [];
        a.speaker.CurrentPagingPreset = [];

        a.speaker.Target        = [];
        a.speaker.Equalizer     = [];

        a.speaker.Music         = [];

        for (i in data.response.TargetNames) {
          a.speaker.TargetNames[i] = data.response.TargetNames[i].replace("\n", "");
          //console.log(i, data.response.TargetNames[i]);
        }

        for (i in data.response.PresetNames) {
          a.speaker.PresetNames[i] = data.response.PresetNames[i].replace("\n", "");
          //console.log(i, data.response.PresetNames[i]);
        }

        for (i in data.response.MusicPresetNames) {
          a.speaker.MusicPresetNames[i] = data.response.MusicPresetNames[i].replace("\n", "");
          //console.log(i, data.response.MusicPresetNames[i]);
        }

        for (i in data.response.PagingPresetNames) {
          a.speaker.PagingPresetNames[i] = data.response.PagingPresetNames[i].replace("\n", "");
          //console.log(i, data.response.PagingPresetNames[i]);
        }

        for (i in data.response.CurrentPreset) {
          a.speaker.CurrentPreset[i] = data.response.CurrentPreset[i];
          //console.log(i, data.response.CurrentPreset[i]);
        }

        for (i in data.response.CurrentMusicPreset) {
          a.speaker.CurrentMusicPreset[i] = data.response.CurrentMusicPreset[i];
          //console.log(i, data.response.CurrentMusicPreset[i]);
        } 

        for (i in data.response.CurrentPagingPreset) {
          a.speaker.CurrentPagingPreset[i] = data.response.CurrentPagingPreset[i];
          //console.log(i, data.response.CurrentPagingPreset[i]);
        } 

        for (i in data.response.CurrentTarget) {
          a.speaker.CurrentTarget[i] = data.response.CurrentTarget[i];
          //console.log(i, data.response.CurrentTarget[i]);
          //dataArray.push(data.response.CurrentTarget[i]);
          dataArray[i] = a.speaker.CurrentTarget[i];
        }

        $(document.getElementsByClassName("inputs")).each(function(j) {
          this.value = dataArray[j];
        });
        redrawOnInput();

        dataArray = [];

        for(i in data.response.Equalizer) {
          a.speaker.Equalizer[i] = data.response.Equalizer[i];
          //console.log(data.response.PresetNames[i], data.response.Equalizer[i]);
        }

        for(i in data.response.MusicEqualizer) {
          a.speaker.MusicEqualizer[i] = data.response.MusicEqualizer[i];
          //console.log(data.response.MusicPresetNames[i], data.response.MusicEqualizer[i]);
        }

        for(i in data.response.PagingEqualizer) {
          a.speaker.PagingEqualizer[i] = data.response.PagingEqualizer[i];
          //console.log(data.response.PagingPresetNames[i], data.response.PagingEqualizer[i]);
        }

        for(i in data.response.Target) {
          a.speaker.Target[i] = data.response.Target[i];
          //console.log(data.response.TargetNames[i], data.response.Target[i]);
        }

        a.speaker.Response = "true";

        //console.log("Speaker", a.speaker);
      },
    
      renderMap: function() {
        var mapController = this;
        var container = document.getElementById('popup');
        var overlay = new ol.Overlay(({
          element: container,
          autoPan: true,
          autoPanAnimation: {
            duration: 250
          }
        }));
        
        var tt = document.getElementById('test');
        var tooltip = new ol.Overlay(({
          element: tt,
          autoPan: true,
          autoPanAnimation: {
            duration: 250
          }
        }));
        
        mapController.overlay = overlay;

        var extent = [0, 0, 2200, 1700];
        var projection = new ol.proj.Projection({
          code: 'pixel',
          units: 'pixels',
          extent: extent
        });

        var locations = this.controllers;
        controllers = this.controllers;

        var features = new Array(locations.length);
        for(var i = 0; i < features.length; i++) {
            var coordinates = [locations[i].x, 1700-locations[i].y];
            var point = new ol.geom.Point(coordinates);
            features[i] = new ol.Feature({
              geometry: point,
              speaker: locations[i]
            });
        }
        
        var vectorSource = new ol.source.Vector({
            features: features
        });

        var style = {}
        var allColors = [
          '#3f51b5'
        ];
        
        var unusedColors = [
          '#3f51b5'
        ];
        var colorMap = {};
        var vectorLayer = new ol.layer.Vector({
            source: vectorSource,
            style: function(feature, resolution) {
              //console.log(feature.getProperties().speaker);
              var zone = feature.getProperties().speaker.zone.name;
              if(!(zone in colorMap)) {
                var color;
                if(unusedColors.length > 0) {
                  var index = Math.floor(Math.random() * unusedColors.length)
                  color = unusedColors[index]
                  unusedColors.splice(index, 1);
                } else {
                  color = colors[Math.floor(Math.random() * colors.length)]
                }

                colorMap[zone] = color;
                savedStyle[zone] = color;
                savedStyle_i++;
                //console.log(savedStyle);
              }
              if(!(zone in style) || style[zone].resolution !== resolution) {
                style[zone] = {};
                style[zone].style = [new ol.style.Style({
                    image: new ol.style.Circle({
                      radius: 15 / Math.max(resolution, 1),
                      stroke: new ol.style.Stroke({
                          color: '#000000'
                      }),
                      fill: new ol.style.Fill({
                          color: colorMap[zone]
                      })
                    })
                })];
                style[zone].resolution = resolution;
              }

              savedResolution = resolution;
              savedFeature = feature;

              var errorStyle = new ol.style.Style({
                image: new ol.style.Icon({
                  src: "/css/brokenlink.png"
                })
              });

              if(zone == "default") {
                var speakers = feature.getProperties().speaker;
                  if(speakers.status !== 0) {
                    //console.log("updateMapForError", speakers.id);

                    feature.setStyle(errorStyle);
                  }
              }
              return style[zone].style;
            }
        });

        vectorLayerCopy = vectorLayer;

        var map = new ol.Map({
          layers: [
            new ol.layer.Image({
              source: new ol.source.ImageStatic({
                url: '/css/map.png',
                projection: projection,
                imageExtent: extent
              })
            }),
              vectorLayer
          ],
          overlays: [overlay, tooltip],
          target: 'map',
          view: new ol.View({
            projection: projection,
            center: ol.extent.getCenter(extent),
            zoom: 6,
            zoomFactor: 1.35
          })
        });
        mapController.map = map;

        var select = new ol.interaction.Select({
          style: function(feature, resolution) {
                var zone = feature.getProperties().speaker.zone.name;
                return style[zone].style;
            }
        });


        map.addInteraction(select);

        select.on('select', function(e) {
            if(e.selected.length > 0) {
              //console.log("Selected");
              var speaker = e.selected[0].getProperties().speaker
              //console.log(e)
              //console.log(speaker);

              var frontendSpeakerUpdateAjax = document.getElementById("frontendSpeakerUpdateAjax");
              //console.log("If I do say myself sir, this is a pretty snazzy way to go about doing this!");
              frontendSpeakerUpdateAjax.body = {
                "speaker": speaker.id
              };
              //console.log("frontendSpeakerUpdateAjax ", frontendSpeakerUpdateAjax.body);
              frontendSpeakerUpdateAjax.generateRequest(); 

              mapController.selected = speaker;
              var coordinate = e.selected[0].getProperties().geometry.getCoordinates();

              //console.log(coordinate)

              //console.log(speaker);

              overlay.setPosition(coordinate);

              //console.log(colorMap);
            } else {
              mapController._hideOverlay();
            }
        });
        
        var hoverInteraction = new ol.interaction.Select({
            condition: ol.events.condition.pointerMove,
            style: function(feature, resolution) {
                return style.style;
            }
        });
        hoverInteraction.on('select', function(e){
          if(e.selected.length > 0) {
              //console.log("Selected");
              var speaker = e.selected[0].getProperties().speaker;
              //console.log(speaker);
              var coordinate = e.selected[0].getProperties().geometry.getCoordinates();
              tooltip.setPosition(coordinate);
            }
        })
      }
    });

  function updateMapForError(id, change) {
    var features = vectorLayerCopy.getSource().getFeatures();
    var i = 0;
    for(i = 0; i < features.length; i++) {
      if(features[i].getProperties().speaker.id == id) {
        //console.log(features[i].getProperties().speaker);
        break;
      }
    }

    //console.log("i: ", i);
    //console.log("features of i", features[i].getProperties().speaker);

    var old_style = features[i].getStyle(style);

    //console.log(features);
    var style;

    if(change) {
      style = new ol.style.Style({
        image: new ol.style.Icon({
          src: "/css/brokenlink.png"
        })
      });
      //console.log(statusCircle);

      // update the system status icon
      goodCircle.style.visibility = "hidden";
      badCircle.style.visibility = "visible";
      goodFont.style.visibility = "hidden";
      badFont.style.visibility = "visible";

      var badCircleText = document.getElementById("badCircleText");
      var badFontText = document.getElementById("badFontText");

      if(badCircleText.innerHTML.length != 0) {
        badCircleText.innerHTML = "There was an error trying to retrieve the status for speaker " + features[i].getProperties().speaker.id;
      } else {
        badCircleText.innerHTML += ", speaker " + features[i].getProperties().speaker.id; 
      }
      if(badFontText.innerHTML.length != 0) {
        badFontText.innerHTML = "There was an error trying to retrieve the status for speaker " + features[i].getProperties().speaker.id;

      } else {
        badFontText.innerHTML += ", speaker " + features[i].getProperties().speaker.id; 
      }
    } else {
       style = new ol.style.Style({
                    image: new ol.style.Circle({
                      radius: 15 / Math.max(savedResolution, 1),
                      stroke: new ol.style.Stroke({
                          color: '#000000'
                      }),
                      fill: new ol.style.Fill({
                          color: savedStyle[features[i].getProperties().speaker.zone.name]
                      })
                    })
                });

      var badCircleText = document.getElementById("badCircleText");
      var badFontText = document.getElementById("badFontText");
      badCircleText.innerHTML = "";
      badFontText.innerHTML = "";
       
      goodCircle.style.visibility = "visible";
      badCircle.style.visibility = "hidden";
      goodFont.style.visibility = "visible";
      badFont.style.visibility = "hidden";
    }

    features[i].setStyle(style);
    vectorLayerCopy.changed();
  }

