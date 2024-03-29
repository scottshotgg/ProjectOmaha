
  // the name needs to be limited to 50 characters

    var totalInfo = [];
    
    Polymer({
      is: 'information-viewer',

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
                // provided by neon-animation-animations/fade-out-animation.html
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

        // this is where we strip out the undefineds

        this.$.createZoneAjax.body = {
          // make sure that this isn't empty, do the disable button thing
          // send the name with this too dummy
          "name":     name,
          "speakers": totalInfo   // stripem out, maybe we can just takem out when we show them in list
        }

        totalInfo = [];
        document.getElementById("informationSpeakerHolder").innerHTML = "";

        //console.log("createZoneAjax:", createZoneAjax.body);
        this.$.createZoneAjax.generateRequest();
      },

      _cancelPressed: function(e) {
      },

      _createZone: function(e) {
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
        //console.log("gi");
        var mapController = this;
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
        
        // Map views always need a projection.  Here we just want to map image
        // coordinates directly to map coordinates, so we create a projection that uses
        // the image 4 in pixels.
        var extent = [0, 0, 2200, 1700];
        var projection = new ol.proj.Projection({
          code: 'pixel',
          units: 'pixels',
          extent: extent
        });

        var locations = this.controllers;
        var features = new Array(locations.length);
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
          /*'#795548',
          '#ff5722',
          '#ff9800',
          '#ffc107',
          //'--paper-yellow-500', this is kinda hard to see
          '#cddc39',
          '#8bc34a',
          '#009688',
          '#00bcd4',
          '#03a9f4',
          '#2196f3',
          '#3f51b5',
          '#673ab7',
          '#9c27b0',
          '#e91e63',*/
          '#f44336'
        ];
        
        var unusedColors = [
          '#795548',
          '#ff5722',
          '#ff9800',
          '#ffc107',
          //'--paper-yellow-500', this is kinda hard to see
          '#cddc39',
          '#8bc34a',
          '#009688',
          '#00bcd4',
          '#03a9f4',
          '#2196f3',
          '#3f51b5',
          '#673ab7',
          '#9c27b0',
          '#e91e63',
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
                  style[zone].style = [new ol.style.Style({
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
          overlays: [overlay],
          target: 'informationContainer',
          view: new ol.View({
            projection: projection,
            center: ol.extent.getCenter(extent),
            zoom: 2
          })
        });
        
        mapController.map = map;

        // this is the selection handler for the box dragging

        // we need to add a button if it is a touch screen
        var box = new ol.interaction.Select();
        map.addInteraction(box);

        var selectedFeatures = box.getFeatures();

        // a DragBox interaction used to select features by drawing boxes
        var dragBox = new ol.interaction.DragBox({
          condition: ol.events.condition.shiftKeyOnly,
          style: new ol.style.Style({
            stroke: new ol.style.Stroke({
              color: [63, 81, 181, 1]
            })
          })
        });

        map.addInteraction(dragBox);

        dragBox.on('boxend', function(e) {
          // features that intersect the box are added to the collection of
          // selected features, and their names are displayed in the "info"
          // div
          var info = [];
          var extent = dragBox.getGeometry().getExtent();

          //console.log(extent);

          vectorSource.forEachFeatureIntersectingExtent(extent, function(feature) {
            //console.log(feature);
            selectedFeatures.push(feature);

            // just make it grab every other one
            info.push(feature.B.speaker.id);
          });
          //console.log("selectedFeatures: ", selectedFeatures);

          // To make sure that we are no running this on times when nothing is selected
          if (info.length > 0) {
             //infoBox.innerHTML = info.join(', ');
            //console.log("Selected speakers: ", info);

            // i have a feeling that this is not very optimized

            // Here we figure out what is not already in the total selected and add the ones that aren't and delete it if they are
            for (var i = 0; i < info.length; i++) {
              var areYouInThere = totalInfo.indexOf(info[i]);
              if(areYouInThere == -1) {
                totalInfo.push(info[i]); 
                //console.log(totalInfo);
              } else {
                //console.log("Speaker", info[i], "is already in the list of selected speakers.");
                delete totalInfo[areYouInThere];    // make something for delete

                //console.log(totalInfo);
              }
            }

            //console.log("Total selected speakers: ", totalInfo);
            // make this a checkbox thing that will fire when checked, which means we need to keep track of two arrays?
            
            //console.log("Starting sort: ");
            // Sort the array so that when we display it, the items are sorted
            totalInfo = mergeSort(totalInfo);   // This mergeSort isn't any faster wtf

            // Clear the speaker holder element and then reset the entities to everything that has been selected, updating the view if you will
            informationSpeakerHolder = document.getElementById("informationSpeakerHolder");
            var innerHTML = "";
            informationSpeakerHolder.innerHTML = "";
            for (var k = 0; k < totalInfo.length; k++) {
              // The selections in the total selection array take a value of 'undefined' when deleted, and we don't want those in our view, so we parse them out
              if (totalInfo[k] !== undefined) {
               innerHTML += "<paper-item center-justified flex style=\"border-style: solid; border-radius: 5px; border-width: 1px; border-color: #3f51b5; margin: 5px;\"> Speaker " + totalInfo[k] + "</paper-item>";
                //Polymer.dom(speakerHolder).innerHTML = Polymer.dom(speakerHolder).innerHTML + "<paper-item center-justified flex style=\"border-style: solid; border-radius: 5px; border-width: 1px; border-color: #3f51b5; margin: 5px;\"> Speaker " + totalInfo[k] + "</paper-item>";/* + document.getElementById("speakerHolder").innerHTML + "<paper-item> Speaker " + totalInfo[k] + "</paper-item>" +document.getElementById("speakerHolder").innerHTML + "<paper-item> Speaker " + totalInfo[k] + "</paper-item>" +document.getElementById("speakerHolder").innerHTML + "<paper-item> Speaker " + totalInfo[k] + "</paper-item>" +document.getElementById("speakerHolder").innerHTML + "<paper-item> Speaker " + totalInfo[k] + "</paper-item>" +document.getElementById("speakerHolder").innerHTML + "<paper-item> Speaker " + totalInfo[k] + "</paper-item>" +document.getElementById("speakerHolder").innerHTML + "<paper-item> Speaker " + totalInfo[k] + "</paper-item>" +document.getElementById("speakerHolder").innerHTML + "<paper-item> Speaker " + totalInfo[k] + "</paper-item>" +document.getElementById("speakerHolder").innerHTML + "<paper-item> Speaker " + totalInfo[k];*/

                // when we send it off make a new array and splice out the undefineds 
              }
            }
            informationSpeakerHolder.innerHTML = innerHTML;
            //console.log(totalInfo);
          }
        });

        // clear selection when drawing a new box and when clicking on the map
        dragBox.on('boxstart', function(e) {

          // make something that removes them from the animation here when you feel like it

          //selectedFeatures.clear();
          //infoBox.innerHTML = '&nbsp;';
        });
        map.on('click', function() {

          //console.log(selectedFeatures);
          selectedFeatures.clear();
          //infoBox.innerHTML = '&nbsp;';
        });

        // this is the selection for the tooltip shit

       /* var select = new ol.interaction.Select({
          style: function(feature, resolution) {
                var zone = feature.getProperties().speaker.zone.name;
                return style[zone].style;
            }
        });
        map.addInteraction(select);
        select.on('select', function(e) {
            if(e.selected.length > 0) {
              var speaker = e.selected[0].getProperties().speaker

              mapController.selected = speaker;
              // get the point's coordinate
              var coordinate = e.selected[0].getProperties().geometry.getCoordinates();
              overlay.setPosition(coordinate);
            } else {
              // make sure the overlay is hidden
              mapController._hideOverlay();
            }
            
        }); */
      }
    });

//console.log("this is me");

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
