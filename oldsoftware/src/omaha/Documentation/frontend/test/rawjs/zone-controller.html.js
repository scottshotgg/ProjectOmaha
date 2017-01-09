    Polymer({
      is: 'zone-controller',

      behaviors: [
        Polymer.NeonAnimatableBehavior
      ],

      properties: {
        zone: {
          type: Object,
          observer: "_zoneUpdated"
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

        volumeLevel: {
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
          this._initializeStatus(this.speaker);
        this.isReady = true;
        //console.log("this is the zone flag", zoneFlag);
      },

      _initializeStatus: function(speaker) {
        this.volume = 0;
      },

      _zoneUpdated: function() {
        this.isReady = false;
        this._initializeStatus();
        this.isReady = true;
      }
    });
