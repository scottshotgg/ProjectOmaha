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

      adjustSizeZone: function() {
        adjustSizeZone();
      },

      _initializeStatus: function(speaker) {
        this.speakerId = speaker.id;
        speakerid = speaker.id;
        this.volume = speaker.VolumeLevel[0];
        this.equalizer = speaker.equalizerone;
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
