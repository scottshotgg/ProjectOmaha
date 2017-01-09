
    Polymer({
      is: 'speaker-tooltip',

      attributes: {
        speaker: Object
      },

      _selectSpeaker: function() {
        zoneFlag = 0;
        this.fire('speakerselected', {speaker: this.speaker});
        //console.log(this.speaker)
      },

      _selectZone: function() {
        zoneFlag = 1;
        this.fire('zoneselected', {zone: this.speaker.zone});
      },

      _selectAllZone: function() {
        zoneFlag = -1;
        this.fire('zoneselected', {zone: 1})
      }
    });

    $(document).ready(function() {
      /*if(parseInt(localStorage.getItem("level")) == 0) {
          document.getElementById("zoneSelectionButton").innerHTML = "";
          document.getElementById("allZoneSelectionButton").innerHTML = "";
      }*/
  });
