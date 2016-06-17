    Polymer({
      is: 'averaging-controller',

      properties: {
        averaging: Number,

        speakerId: {
            type: Number,
          },

        speaker: {
          type: Object
        }
      },
      
      _selectionChanged: function(event) {
        //console.log(this.speakerId);

        var menu = document.getElementById('menu');

        var attributeValues;

        var updatedAttributes = ["averaging"];

        attributeValues = {
          effectiveness: this.$.effectiveness.value,
          pleasantness: this.$.pleasantness.value
        };
        if(zoneFlag == 0) {
          this.$.selectModeAjax.body = {
            "updatedAttributes": updatedAttributes, 
            "attributeValues": attributeValues,                        
            "speaker": speaker.id
          };

          //console.log(this.$.selectModeAjax.body);
          this.$.selectModeAjax.generateRequest();
        } else {
          this.$.selectModeAjaxZone.body = {                      
            "zone": speaker.zone.zoneID,
            "effectiveness": this.$.effectiveness.value,
            "pleasantness": this.$.pleasantness.value
          };

          //console.log(this.$.selectModeAjaxZone.body);
          this.$.selectModeAjaxZone.generateRequest();
          //console.log(zones);
        }
        //console.log("this is the speaker", speaker);
        var selectModeAjax = document.getElementById("selectModeAjax");
        //console.log(selectModeAjax);
        this.$.toast.text = "Effectiveness set to " + this.$.effectiveness.value + ". Pleasantness set to " + this.$.pleasantness.value + "."; 
        this.$.toast.show();
      },

      _logResponse: function(event, data) {
        //console.log(data.response);
      },

      decrementE: function() {
        this.$.effectiveness.decrement();
        this.$.effectiveness._expandKnob();
        if(this.prevTimeout != 0) {
          window.clearTimeout(this.prevTimeout)
        }
        var elem = this;
        var func = this.resetKnob;
        this.prevTimeout = window.setTimeout(function(){
          func(elem);
        }, 500);
      },
      
      incrementE: function() {
        this.$.effectiveness.increment();
        this.$.effectiveness._expandKnob();
        if(this.prevTimeout != 0) {
          window.clearTimeout(this.prevTimeout)
        }
        var elem = this;
        var func = this.resetKnob;
        this.prevTimeout = window.setTimeout(function(){
          func(elem);
        }, 500);
      },

      decrementC: function() {
        this.$.pleasantness.decrement();
        this.$.pleasantness._expandKnob();
        if(this.prevTimeout != 0) {
          window.clearTimeout(this.prevTimeout)
        }
        var elem = this;
        var func = this.resetKnob;
        this.prevTimeout = window.setTimeout(function(){
          func(elem);
        }, 500);
      },
      
      incrementC: function() {
        this.$.pleasantness.increment();
        this.$.pleasantness._expandKnob();
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
        elem.$.effectiveness._resetKnob();
        elem.$.pleasantness._resetKnob();
      },
      
    });
