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
      
      /**
       * This function is fired when the 'set ratio' button is clicked by the user to set the effectiveness and pleasantness. 
       */
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

      /**
       * This function is used to decrement the effectiveness.
       */
      decrementEffectiveness: function() {
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
      
      /**
       * This function is used to increment the effectiveness.
       */
      incrementEffectiveness: function() {
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

      /**
       * This function is used to decrement the pleasantness.
       */
      decrementPleasantnesss: function() {
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

      /**
       * This function is used to increment the pleasantness.
       */
      incrementPleasantnesss: function() {
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

      /**
       * Resetting the knob resets the knob to remove the pin that is hoisted up when the knob is moved.
       */
      resetKnob: function(elem) {
        elem.prevTimeout = 0;
        elem.$.effectiveness._resetKnob();
        elem.$.pleasantness._resetKnob();
      },
      
    });
