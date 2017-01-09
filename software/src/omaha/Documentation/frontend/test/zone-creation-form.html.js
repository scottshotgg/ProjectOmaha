    Polymer({
      is: 'zone-creation-form',

      _submit: function() {
        if(!this.$.form.validate())
            return;
        this.err = "";
        this.successMsg = "";
        this.$.progress.customStyle['--paper-progress-active-color'] = "#4285f4";
        this.$.progress.updateStyles();
        this.$.progress.indeterminate = true;
		this.$.zoneCreationAjax.body = {
			name: this.$.name.value,
		}
		this.$.zoneCreationAjax.generateRequest();
      },

      _logResponse: function(event, data) {
        if(!data.response.success) {
          this.$.progress.customStyle['--paper-progress-active-color'] = "#F44336";
          this.$.progress.updateStyles();
          this.$.progress.indeterminate = false;
          this.$.progress.value = 100;
          this.err = data.response.err
        } else {
          this.$.progress.value = 0;
          this.$.progress.indeterminate = false;
          this.successMsg = "Zone created";
          this.$.name.value = "";
        }
    }
    });
