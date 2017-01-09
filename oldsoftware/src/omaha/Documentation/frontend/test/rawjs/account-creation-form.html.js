    Polymer({
      is: 'account-creation-form',

      _submit: function() {
        if(!this.$.form.validate())
            return;
        var password = this.$.password.value;
        var confirm = this.$.confirm.value
        this.err = "";
        this.successMsg = "";
        this.$.progress.customStyle['--paper-progress-active-color'] = "#4285f4";
        this.$.progress.updateStyles();
        this.$.progress.indeterminate = true;
        if(password != confirm) {
          this.async(function() {
            this.$.progress.customStyle['--paper-progress-active-color'] = "#F44336";
            this.$.progress.updateStyles();
            this.$.progress.value = 100;
            this.$.progress.indeterminate = false;
            this.err = "Passwords do not match";
            this.$.password.value = "";
            this.$.confirm.value = "";
            return;
          }, 1000);
        } else {
          hash = document.cookie.split('; ');
          this.$.loginAjax.body = {
            level:      parseInt(localStorage.getItem("level")) - 1,
            username:   this.$.username.value,
            password:   this.$.password.value,
            name:       this.$.name.value,
            email:      this.$.email.value,
            phone:      this.$.phone.value,
            speakerid:  parseInt(this.$.speakerid.value), 
            zoneid:     parseInt(this.$.zoneid.value)
          }
          this.$.loginAjax.generateRequest();
        }
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
          this.successMsg = "Account created";
          this.$.username.value = "";
          this.$.password.value = "";
          this.$.confirm.value = "";
        }
    }
    });
