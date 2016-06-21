    Polymer({
      is: 'zone-assigner',

      properties: {
        speaker: {
          type: Object,
          observer: "_speakerUpdated"
        }
      },
      
      _acceptClicked: function() {
        this.updateZoneId();
        return true;
      },
      
      updateZoneId: function() {
        var zone = this.$.zoneMenu.selectedItem.zone;
        var updatedAttributes = ["zoneId"];
          var attributeValues = {
            "zoneId": zone.zoneID,
          };
          this.$.speakerUpdateAjax.body = {
            "updatedAttributes": updatedAttributes, 
            "attributeValues": attributeValues,
            "speaker": this.speaker.id
          };
          this.$.speakerUpdateAjax.generateRequest();
       var oldZone = this.speaker.zone;
       zone.speakers.push(this.speaker);
       oldZone.speakers.splice(oldZone.speakers.indexOf(this.speaker), 1);
       this.speaker.zone = zone;
       this.fire('zoneupdated', {speakerId: this.speaker.id});
      },
      
      _speakerUpdated: function() {
        var speaker = this.speaker;
        var zoneName = speaker.zone.name;
        this.$.zoneMenu.selected = zoneName;
        console.log(this.speaker);
      },
      
       _logResponse: function(event, data) {
        console.log(data.response);
      }
    });
