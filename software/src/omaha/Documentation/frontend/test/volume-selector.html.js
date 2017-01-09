
    var localLevelForVolumeSelector = parseInt(localStorage.getItem("level"));

    Polymer({
      is: 'volume-selector',
      
      properties: {
        internalVolume: String,

        volume: {
          type: String,
          notify: true,
          observer: '_setInternalVolume'
        },

        speakerId: {
          type: Number
        }
      },

      _setInternalVolume: function(volume) {
        this.internalVolume = volume;
        //console.log("_setInternalVolume");
      },
      
      _setVolumeClicked: function() {
        if(zoneFlag == 0) {
          if(localLevelForVolumeSelector > 0) {
            this.$.thresholdValue.body = {
              "speaker": speakerid, 

              "musicMin":   ((!isNaN(parseInt(this.$.musicMin.value)))   ? parseInt(this.$.musicMin.value)   : 0),
              "musicMax":   ((!isNaN(parseInt(this.$.musicMax.value)))   ? parseInt(this.$.musicMax.value)   : 100),

              "pagingMin":  ((!isNaN(parseInt(this.$.pagingMin.value)))  ? parseInt(this.$.pagingMin.value)  : 0),
              "pagingMax":  ((!isNaN(parseInt(this.$.pagingMax.value)))  ? parseInt(this.$.pagingMax.value)  : 100),

              "maskingMin": ((!isNaN(parseInt(this.$.maskingMin.value))) ? parseInt(this.$.maskingMin.value) : 0),
              "maskingMax": ((!isNaN(parseInt(this.$.maskingMax.value))) ? parseInt(this.$.maskingMax.value) : 100)
            };
            this.$.thresholdValue.generateRequest();
          }

          if(localLevelForVolumeSelector < 1 && 
            (this.$.volume3.value < this.$.musicMin.value || this.$.volume3.value > this.$.musicMax.value || 
            this.$.volume4.value < this.$.pagingMin.value || this.$.volume4.value > this.$.pagingMax.value || 
            this.$.volume5.value < this.$.maskingMin.value || this.$.volume5.value > this.$.maskingMax.value)) {

            this.$.invalidVolumeToast.show();
            return;
          }

          this.volume = this.internalVolume + " " + this.internalMusicVolume + " " + this.internalPagingVolume + " " + this.internalSoundMaskingVolume;
          this.$.toast.show();
          //console.log("_setVolumeClicked");
        } else {
          if(localLevelForVolumeSelector > 0) {
            this.$.thresholdValueZone.body = {
              "zone": speaker.zone.zoneID, 

              "musicMin":   ((!isNaN(parseInt(this.$.musicMin.value)))   ? parseInt(this.$.musicMin.value)   : 0),
              "musicMax":   ((!isNaN(parseInt(this.$.musicMax.value)))   ? parseInt(this.$.musicMax.value)   : 100),

              "pagingMin":  ((!isNaN(parseInt(this.$.pagingMin.value)))  ? parseInt(this.$.pagingMin.value)  : 0),
              "pagingMax":  ((!isNaN(parseInt(this.$.pagingMax.value)))  ? parseInt(this.$.pagingMax.value)  : 100),

              "maskingMin": ((!isNaN(parseInt(this.$.maskingMin.value))) ? parseInt(this.$.maskingMin.value) : 0),
              "maskingMax": ((!isNaN(parseInt(this.$.maskingMax.value))) ? parseInt(this.$.maskingMax.value) : 100)
            };
            //console.log(this.$.thresholdValueZone.body);
            this.$.thresholdValueZone.generateRequest();
          }

          if(localLevelForVolumeSelector < 1 && 
            (this.$.volume3.value < this.$.musicMin.value || this.$.volume3.value > this.$.musicMax.value || 
            this.$.volume4.value < this.$.pagingMin.value || this.$.volume4.value > this.$.pagingMax.value || 
            this.$.volume5.value < this.$.maskingMin.value || this.$.volume5.value > this.$.maskingMax.value)) {

            this.$.invalidVolumeToast.show();
            return;
          } else {
            this.$.updateVolumesZone.body = {
              "zone": speaker.zone.zoneID,
              "volume": this.$.volume2.value,
              "music": this.$.volume3.value,
              "paging": this.$.volume4.value,
              "masking": this.$.volume5.value
            };
            this.$.updateVolumesZone.generateRequest();
          }
          this.$.toast.show();
          //console.log("_setVolumeClicked");
        }
      },
      
      decrementVolume: function() {
        if(localLevelForVolumeSelector < 1) {
          return;
        }
        this.$.volume2.decrement();
        this.$.volume2._expandKnob();
        if(this.prevTimeout != 0) {
          window.clearTimeout(this.prevTimeout)
        }
        var elem = this;
        var func = this.resetKnob;
        this.prevTimeout = window.setTimeout(function(){
          func(elem);
        }, 500);
      },
      
      incrementVolume: function() {
        if(localLevelForVolumeSelector < 1) {
          return;
        }
        this.$.volume2.increment();
        this.$.volume2._expandKnob();
        if(this.prevTimeout != 0) {
          window.clearTimeout(this.prevTimeout)
        }
        var elem = this;
        var func = this.resetKnob;
        this.prevTimeout = window.setTimeout(function(){
          func(elem);
        }, 500);
      },
      
      decrementMusicVolume: function() {
        if(this.internalMusicVolume - 1 < LowerThreshold[0] && localLevelForVolumeSelector < 1) {
          return;
        }
        this.$.volume3.decrement();
        this.$.volume3._expandKnob();
        if(this.prevTimeout != 0) {
          window.clearTimeout(this.prevTimeout)
        }
        var elem = this;
        var func = this.resetKnob;
        this.prevTimeout = window.setTimeout(function(){
          func(elem);
        }, 500);
      },
      
      incrementMusicVolume: function() {
        if(this.internalMusicVolume + 1 > UpperThreshold[0] && localLevelForVolumeSelector < 1) {
          return;
        }
        this.$.volume3.increment();
        this.$.volume3._expandKnob();
        if(this.prevTimeout != 0) {
          window.clearTimeout(this.prevTimeout)
        }
        var elem = this;
        var func = this.resetKnob;
        this.prevTimeout = window.setTimeout(function(){
          func(elem);
        }, 500);
      },
      
      decrementPagingVolume: function() {
        if(this.internalPagingVolume - 1 < LowerThreshold[1] && localLevelForVolumeSelector < 1) {
          return;
        }
        this.$.volume4.decrement();
        this.$.volume4._expandKnob();
        if(this.prevTimeout != 0) {
          window.clearTimeout(this.prevTimeout)
        }
        var elem = this;
        var func = this.resetKnob;
        this.prevTimeout = window.setTimeout(function(){
          func(elem);
        }, 500);
      },
      
      incrementPagingVolume: function() {
        if(this.internalPagingVolume + 1 > UpperThreshold[1] && localLevelForVolumeSelector < 1) {
          return;
        }
        this.$.volume4.increment();
        this.$.volume4._expandKnob();
        if(this.prevTimeout != 0) {
          window.clearTimeout(this.prevTimeout)
        }
        var elem = this;
        var func = this.resetKnob;
        this.prevTimeout = window.setTimeout(function(){
          func(elem);
        }, 500);
      },
      
      decrementMaskingVolume: function() {
        if(this.internalMaskingVolume - 1 < LowerThreshold[2] && localLevelForVolumeSelector < 1) {
          return;
        }
        this.$.volume5.decrement();
        this.$.volume5._expandKnob();
        if(this.prevTimeout != 0) {
          window.clearTimeout(this.prevTimeout)
        }
        var elem = this;
        var func = this.resetKnob;
        this.prevTimeout = window.setTimeout(function(){
          func(elem);
        }, 500);
      },
      
      incrementMaskingVolume: function() {
        if(this.internalMaskingVolume + 1 > UpperThreshold[2] && localLevelForVolumeSelector < 1) {
          return;
        }
        this.$.volume5.increment();
        this.$.volume5._expandKnob();
        if(this.prevTimeout != 0) {
          window.clearTimeout(this.prevTimeout)
        }
        var elem = this;
        var func = this.resetKnob;
        this.prevTimeout = window.setTimeout(function(){
          func(elem);
        }, 500);
      },
      
      resetKnob: function(elem) {
        elem.prevTimeout = 0;
        elem.$.volume2._resetKnob();
        elem.$.volume3._resetKnob();
        elem.$.volume4._resetKnob();
        elem.$.volume5._resetKnob();
      },

      _logResponse: function(event, data) {
        //console.log(data.response);
      },

      ready: function() {
        //console.log(speakerid);
        //console.log(this, "\n", this.$, "\n", this.$.speakerId);

        switch(localLevelForVolumeSelector) {
          case 0:
            this.$.musicMin.disabled = true;
            this.$.musicMax.disabled = true;

            this.$.pagingMin.disabled = true;
            this.$.pagingMax.disabled = true;
            
            this.$.maskingMin.disabled = true;
            this.$.maskingMax.disabled = true;

            this.$.volume2.disabled = true;
            break;
        }

        this.internalVolume = this.volume;
        //console.log("I am ready " + this.internalVolume + " " + this.internalMusicVolume + " " + this.internalPagingVolume) + " " + this.internalSoundMaskingVolume;
      }
    });
