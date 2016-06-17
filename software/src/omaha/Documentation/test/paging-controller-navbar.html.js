
  var map;
  var features;
  var mapController;
  var idIterator = 0;
  var currentSpeakers = [];


  function pulsateCircle(id, coordinates) {
    var elem = document.createElement('div');
    elem.setAttribute('class', 'gps_ring');
    return new ol.Overlay({
      id: id,
      element: elem,
      position: coordinates,
      positioning: 'center-center'
    });
  }

  function findZoneIndex(zoneName) {
    //console.log(zoneName);
    for(var i = 0; i < zones.length; i++) {
      if(zones[i].name === zoneName) {        
        for(var k = 0; k < zones[i].speakers.length; k++) {
          currentSpeakers[k] = zones[i].speakers[k].id;
        }
        //console.log("found a name");
        return i;
      }
    }

    for(var i = 0; i < pagingZones.length; i++) {
      if(pagingZones[i].name === zoneName) {
        //console.log(pagingZones[i]);
        for(var k = 0; k < pagingZones[i].speakers.length; k++) {
          currentSpeakers[k] = pagingZones[i].speakers[k].id;
        }
        return i;
      }
    }
    
    var zoneNameInt = parseInt(zoneName);
    for(var i = 0; i < zones[0].speakers.length; i++) {
      if(zones[0].speakers[i].id == zoneNameInt) {
        currentSpeakers = [];
        currentSpeakers[0] = zones[0].speakers[i].id;
        return i;
      }
    }
  }

  function updateViewForSelectedSpeakers(elem) {
    //console.log(elem.id);
    //console.log(features);
    //console.log(map.a);
    var overlays = mapController.map.getOverlays();
    //console.log(overlays);

    //console.log(elem);
    var zoneIndex = findZoneIndex(elem.id);
    //console.log(typeof currentSpeakers);
    //console.log(currentSpeakers, currentSpeakers.length);
    //console.log(zoneIndex);
    var zoneOverlayArray = [];
    //console.log(elem.id);
    if(displayedZones[elem.id] === undefined) {
      elem.style.background = "rgba(63, 81, 181, .25)";

      for(var i = 0; i < currentSpeakers.length; i++) {
        var tempOverlay = pulsateCircle(i, features[currentSpeakers[i] - 1].getGeometry().getCoordinates());
        tempOverlay.id = idIterator;
        mapController.map.addOverlay(tempOverlay);
        zoneOverlayArray.push(tempOverlay.id);
        idIterator++;
      }
      displayedZones[elem.id] = zoneOverlayArray;
      mapController.map.updateSize();
    }
    else {
      elem.style.background = "rgba(63, 81, 181, .1)";

      var overlays = mapController.map.getOverlays();
      //console.log(overlays);
      //console.log(overlays.a.length);

      //console.log(displayedZones[elem.id]);
      for(var j = displayedZones[elem.id].length - 1; j >= 0; j--) {
        var elementPos = overlays.a.map(function(x) { return x.id; }).indexOf(displayedZones[elem.id][j]);

        var objectFound = overlays.a[elementPos];
        //console.log(objectFound);
        mapController.map.removeOverlay(objectFound);
      }
      mapController.map.updateSize();
      displayedZones[elem.id] = undefined;
      currentSpeakers = [];
    }
    //console.log(displayedZones);

    //console.log(displayedZones);
  }

  (function() {
    $(".pagingZoneStyle").each(function(i) {
      //console.log(i);
    });
  })();

    var totalInfo = [];
    
    Polymer({
      is: 'paging-controller-navbar',

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

      _acceptPressed: function(e) {
        this.$.zoneDialog.toggle();
        this.$.toast.show();
        //console.log("_acceptPressed", totalInfo);
        var name = this.$.popupInput.value.valueOf().replace("\n", "").replace("\t", "");

        this.$.createPagingZoneAjax.body = {
          "name":     name,
          "speakers": totalInfo
        }

        totalInfo = [];
        document.getElementById("pagingSpeakerHolder").innerHTML = "";

        //console.log("createPagingZoneAjax:", createPagingZoneAjax.body);
        this.$.createPagingZoneAjax.generateRequest();
      },

      _cancelPressed: function(e) {
      },

      _dial: function(e) {
        this.$.zoneDialog.toggle();
      },
      
      _refreshSpeaker: function(e) {
        var speakerId = e.detail.speakerId;
        this.featureMap[speakerId].changed();
      },
      
      _assignZone: function() {
        this._hideOverlay();
        this.$.dialog.open();
      },
      
      _hideOverlay: function() {
        this.overlay.setPosition(undefined);
      },

      renderMap: function() {
        mapController = this;
        mapController.featureMap = {};
        var container = mapController.$.popup;
        var overlay = new ol.Overlay(({
          element: container,
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
        features = new Array(locations.length);
        for(var i = 0; i < features.length; i++) {
            var coordinates = [locations[i].x, 1700-locations[i].y];
            features[i] = new ol.Feature({
              geometry: new ol.geom.Point(coordinates),
              speaker: locations[i]
            });
            mapController.featureMap[locations[i].id] = features[i]
        }

        var vectorSource = new ol.source.Vector({
            features: features
        });

        var style = {}

       var allColors = [
          '#f44336'
        ];
        
        var unusedColors = [
          '#f44336'
        ];
        var colorMap = {};
        
        var vectorLayer = new ol.layer.Vector({
            source: vectorSource,
            style: function(feature, resolution) {
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
                }
                if(!(zone in style) || style[zone].resolution !== resolution) {
                  style[zone] = {};
                  style[zone].style = 
                  [new ol.style.Style({
                        image: new ol.style.Circle({
                            radius: 15/Math.max(resolution,1),
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
                return style[zone].style;
            }
        });

        map = new ol.Map({
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
          overlays: [overlay],
          target: 'pagingMapContainer',
          view: new ol.View({
            projection: projection,
            center: ol.extent.getCenter(extent),
            zoom: 2
          })
        });
        mapController.map = map;

        var box = new ol.interaction.Select();

        var selectedFeatures = box.getFeatures();

        var dragBox = new ol.interaction.DragBox({
          condition: ol.events.condition.shiftKeyOnly,
          style: new ol.style.Style({
            stroke: new ol.style.Stroke({
              color: [63, 81, 181, 1]
            })
          })
        });

        dragBox.on('boxend', function(e) {
          var info = [];
          var extent = dragBox.getGeometry().getExtent();

          //console.log(extent);

          vectorSource.forEachFeatureIntersectingExtent(extent, function(feature) {
            //console.log(feature);
            selectedFeatures.push(feature);

            info.push(feature.B.speaker.id);
          });
          //console.log("selectedFeatures: ", selectedFeatures);

          if (info.length > 0) {
            //console.log("Selected speakers: ", info);

            for (var i = 0; i < info.length; i++) {
              var areYouInThere = totalInfo.indexOf(info[i]);
              if(areYouInThere == -1) {
                totalInfo.push(info[i]); 
                //console.log(totalInfo);
              } else {
                //console.log("Speaker", info[i], "is already in the list of selected speakers.");
                delete totalInfo[areYouInThere];

                //console.log(totalInfo);
              }
            }

            //console.log("Total selected speakers: ", totalInfo);
            
            //console.log("Starting sort: ");
            totalInfo = mergeSort(totalInfo);

            pagingSpeakerHolder = document.getElementById("pagingSpeakerHolder");
            var innerHTML = "";
            pagingSpeakerHolder.innerHTML = "";
            for (var k = 0; k < totalInfo.length; k++) {
              if (totalInfo[k] !== undefined) {
               innerHTML += "<paper-item center-justified flex style=\"border-style: solid; border-radius: 5px; border-width: 1px; border-color: #3f51b5; margin: 5px;\"> Speaker " + totalInfo[k] + "</paper-item>";
              }
            }
            pagingSpeakerHolder.innerHTML = innerHTML;
            //console.log(totalInfo);
          }
        });

        dragBox.on('boxstart', function(e) {

        });
        map.on('click', function() {

          //console.log(selectedFeatures);
          selectedFeatures.clear();
        });
      }
    });

    function mergeSort(arr) {
        if (arr.length < 2)
            return arr;

        var middle = parseInt(arr.length / 2);
        var left   = arr.slice(0, middle);
        var right  = arr.slice(middle, arr.length);

        return merge(mergeSort(left), mergeSort(right));
    }

    function merge(left, right) {
        var result = [];

        while (left.length && right.length) {
            if (left[0] <= right[0]) {
                result.push(left.shift());
            } else {
                result.push(right.shift());
            }
        }

        while (left.length)
            result.push(left.shift());

        while (right.length)
            result.push(right.shift());

        return result;
    }
