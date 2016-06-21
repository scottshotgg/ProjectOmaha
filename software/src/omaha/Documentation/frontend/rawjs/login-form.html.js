    Polymer({
      is: 'login-form',

      /**
       * Submit takes the text fields in the login screen and sends them off to the server for validation.
       */
      _submit: function() {
        var loginInfo = {

        }
        this.$.loginAjax.body = {
          username: this.$.username.value,
          password: this.$.password.value
        }
        this.$.loginAjax.generateRequest();
      },

      /**
       * The handlelogin function is a callback for the submit function that is fired on response. The function stores the level, speakerid, and zoneid in localStorage so that it can be retrieved by other classes. The function recieves as session hash and assigns it as a session cookie.
       */
      _handleLogin: function(event, data, x) {
        console.log(data.response);
        localStorage.setItem('level', data.response.level);
        localStorage.setItem('speakerid', data.response.speakerid);
        localStorage.setItem('zoneid', data.response.zoneid);
        if(data.response.hash !== undefined) {
          var sessionCookie = data.response.hash;
          document.cookie = "session=" + sessionCookie + ";";
          document.location.href = "/app/"
        } else if (data.response.success){
          document.location.href = "/app/"
        } else {
          console.log(data.response);
          this.err = data.response.err;
          this.$.toast.show();
        }
      }
    });
