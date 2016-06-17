    var on = 0;

    Polymer({
      is: 'paging-controller',
      
      properties: {
        internalFadeTime: String,

        speaker: {
          type: Object
        },

        speakerId: {
          type: Number
        },

        paging: {
          type: String,
          notify: true,
          observer: '_setInternalPaging'
        },

        value: {
          type: String
        }
      },

      _setInternalPaging: function(paging) {
        this.internalFadeTime = paging;
        //console.log("_setInternalPaging");
      },
      
      _setPagingClicked: function() {
        if(zoneFlag == 0) {
          this.paging = this.internalFadeTime + " " + this.internalFadeLevel + " " + this.internalPagingVolume;
          this.$.toast.show();
          //console.log("_setPagingClicked");
        } else {
          this.$.updatePagingValuesZone.body = {
            "zone": speaker.zone.zoneID,
            "fadeTime": this.internalFadeTime,
            "fadeLevel": this.internalFadeLevel,
            "pagingVolume": this.internalPagingVolume
          };

          this.$.updatePagingValuesZone.generateRequest();
          //console.log(this.$.updatePagingValuesZone.body);
          //console.log("hi");
        }
      },

      _sendPageClicked: function() {
        //console.log(on);
        if(on == 1) {
          //console.log("on");
          on--;
          Polymer.dom(sendPageButton).innerHTML = "stop page";
        } else {
          //console.log("off");
          on++;
          Polymer.dom(sendPageButton).innerHTML = "send page";
        }

        this.$.pagingRequest.body = {
            "speaker": this.speakerId
        };
        this.$.pagingRequest.generateRequest();

        //console.log("Making a paging request ", this.speakerId, this.$.pagingRequest.body);
        this.$.toasty.show();
      },

      _logResponse: function(event, data) {
        //console.log(data.response);
      },

      ready: function() {
        this.internalFadeTime = this.paging;
        //console.log("I am ready from paging " + this.internalFadeTime);
      }
    });
