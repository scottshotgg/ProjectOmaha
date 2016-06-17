    Polymer({
      is: 'account-creation-form',

      /**
       * Submit is an on-click that is fired when the user clicks the submit button. The submit function takes care of collecting the data from the form,
       * packaging it up, and then sending it off. It also does some validation and fills in some dots, such as the level; this variable is inferred to be the [level of user making the account] - 1. 
       */
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

      /**
       * This is the logResponse for the submit funtion, this changes some styles to reflect that an account was created.
       */
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
