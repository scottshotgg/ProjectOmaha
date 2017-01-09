
    var speakerid = 0;

    Polymer({
      is: 'speaker-controller',

      behaviors: [
        Polymer.NeonAnimatableBehavior
      ],

      properties: {
        speaker: {
          type: Object,
          observer: "_speakerUpdated"
        },

        isReady: {
          type: Boolean,
          value: false
        },

        mode: {
          type: String,
          value: "volume"
        },

        speakerId: {
          type: Number,
          value: 0
        },

        volumeLevel: {    // this shit and all the other shitters were volumeone for some reason ??? changed to volumeLevel
          type: String,
          value: "0",
          observer: '_volumeUpdate'
        },
        
        musicLevel: {
          type: Number,
          value: 0
        },

        pagingLevel: {
          type: String,
          value: "0",
          observer: '_pagingUpdate'
        },

        equalizerone: {
          type: String,
          value: "0",
          observer: '_equalizerUpdate'
        },

        musicone: {
          type: String,
          value: "0",
          observer: '_equalizerUpdate'
        },

        target: {
          type: String,
          value: "0",
          observer: '_targetUpdate'
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

      adjustSize: function() {
        adjustSize();
      },

      _volumeUpdate: function() {
        if(this.isReady) {
          var updatedAttributes = ["volume"];
          var attributeValues = {
            "volume": this.volumeLevel,
          };
          this.$.speakerUpdateAjax.body = {
            "updatedAttributes": updatedAttributes, 
            "attributeValues": attributeValues,
            "speaker": this.speakerId
          };
          //console.log(this.$, this, "this is the stuff yo")
          this.$.speakerUpdateAjax.generateRequest();
        }
        //console.log("_volumeUpdate " + this.volumeLevel);
      },

      _pagingUpdate: function() {
        if(this.isReady) {
          var updatedAttributes = ["paging"];
          var attributeValues = {
            "paging": this.pagingLevel,
          };
          this.$.speakerUpdateAjax.body = {
            "updatedAttributes": updatedAttributes, 
            "attributeValues": attributeValues,
            "speaker": this.speakerId
          };
          this.$.speakerUpdateAjax.generateRequest();
        }
        //console.log("_pagingUpdate " + this.pagingLevel);
      },

      _equalizerUpdate: function() {
        if(this.isReady) {
          var updatedAttributes = ["equalizer"];
          var attributeValues = {
            "equalizer": this.equalizerone,
          };
          this.$.speakerUpdateAjax.body = {
            "updatedAttributes": updatedAttributes, 
            "attributeValues": attributeValues,
            "speaker": this.speakerId
          };
          this.$.speakerUpdateAjax.generateRequest();

          //console.log(this.$.speakerUpdateAjax.body);

        }
      },

      _musicEqualizerUpdate: function() {
        if(this.isReady) {
          var updatedAttributes = ["music"];
          var attributeValues = {
            "music": this.musicone,
          };
          this.$.speakerUpdateAjax.body = {
            "updatedAttributes": updatedAttributes, 
            "attributeValues": attributeValues,
            "speaker": this.speakerId
          };
          this.$.speakerUpdateAjax.generateRequest();

          //console.log(this.$.speakerUpdateAjax.body);

        }
      },

      _targetUpdate: function() {
        if(this.isReady) {
          var updatedAttributes = ["target"];
          var attributeValues = {
            "target": this.target,
          };
          this.$.speakerUpdateAjax.body = {
            "updatedAttributes": updatedAttributes, 
            "attributeValues": attributeValues,
            "speaker": this.speakerId
          };
          this.$.speakerUpdateAjax.generateRequest();
        }
      },

      _targetUpdateZone: function() {
        if(this.isReady) {
          var updatedAttributes = ["target"];
          var attributeValues = {
            "target": this.target,
          };
          this.$.speakerUpdateAjax.body = {
            "updatedAttributes": updatedAttributes, 
            "attributeValues": attributeValues,
            "speaker": this.speakerId
          };
          this.$.speakerUpdateAjax.generateRequest();
        }
      },

      _logResponse: function(event, data) {
        //console.log(data.response);
      },

      _initializeStatus: function(speaker) {
        this.speakerId = speaker.id;
        speakerid = speaker.id;
        this.volume = speaker.VolumeLevel[0];
        this.equalizer = speaker.equalizerone;
      },

      _speakerUpdated: function() {
        this.isReady = false;
        this._initializeStatus(this.speaker);
        this.isReady = true;
      },

      ready: function() {
        //console.log(document.getElementById('target'));
        if(this.speaker !== undefined)
          this._initializeStatus(this.speaker);
        this.isReady = true;
        //console.log("this is the zone flag", zoneFlag);
      }
    });
  
  (function hideSomething() {
    
  })();

