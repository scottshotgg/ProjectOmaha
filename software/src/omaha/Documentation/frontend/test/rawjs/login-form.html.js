    Polymer({
      is: 'login-form',

      _submit: function() {
        var loginInfo = {

        }
        this.$.loginAjax.body = {
          username: this.$.username.value,
          password: this.$.password.value
        }
        this.$.loginAjax.generateRequest();
      },

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
