    var ss;
    Polymer({
      is: 'login-form',

      listeners: {
        'controllerListBox.iron-select': '_changeController'
      },

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

      _changeController: function(event) {
        if(this.$.controllerListBox.selected == "addController") {
          this.$.dialog.toggle();
        } else {
          window.location = "http://" + this.$.controllerListBox.selected + ":8080";
          var timerID = setInterval(
            function() {
              document.getElementById("message").innerHTML = "Controller from that address is not responding";
              cancelInterval(timerID);
          }, 5000)
        }
      },

      _acceptPressed: function(event) {
        console.log("added new controller");

        this.$.addController.body = {
          ip: this.$.ipa.value,
          name: this.$.name.value
        }
        console.log(this.$.addController.body);
        this.$.addController.generateRequest();

        this.$.dialog.toggle();

        temp = document.createElement("paper-item");
        temp.setAttribute("value", this.$.ipa.value);
        temp.setAttribute("id", this.$.name.value);
        temp.innerHTML = this.$.name.value + " (" + this.$.ipa.value + ")";
        Polymer.dom(this.$.controllerListBox).appendChild(temp);

        this.$.controllerListBox.selected = -1;
      },

      _addControllerResponse: function(event, data) {
        console.log(data.response);
      },
      /**
       * The handlelogin function is a callback for the submit function that is fired on response. The function stores the level, speakerid, and zoneid in localStorage so that it can be retrieved by other classes. The function recieves as session hash and assigns it as a session cookie. During this time it also makes the expiration date of the cookie 36 hours from the time it was created.
       */
      _handleLogin: function(event, data, x) {
        console.log(data.response);
        localStorage.setItem('level', data.response.level);
        localStorage.setItem('speakerid', data.response.speakerid);
        localStorage.setItem('zoneid', data.response.zoneid);
        if(data.response.hash !== undefined) {
          var sessionCookie = data.response.hash;
          console.log(sessionCookie)
          /*var d = new Date();
          d.setHours(d.getHours() + 36);
          document.cookie = "name=" + sessionCookie + "; session=" + sessionCookie + "; expires=" + d.toUTCString() + ";";*/

          document.cookie = "path=/";
          document.cookie = "session=" + sessionCookie;
          console.log(window.location);
          console.log(document.cookie);
          //document.location.href = "/app/"
          window.location = "http://" + window.location.hostname + ":8080/app/";
        } else if (data.response.success){
          console.log(data.response.hash);
          //document.location.href = "/app/"
          //window.location = "http://" + window.location.hostname + ":8080/app/";
        } else {
          console.log(data.response);
          console.log(this);
          this.err = data.response.err;
          console.log(typeof data.response.err);
          this.$.toast.show();
        }
      },

      _loadControllerResponse: function(event, data) {
        console.log(data.response);

        var controllerids = data.response.Controllerids;
        var ips = data.response.Ips;
        var names = data.response.Names;

        for(var x = 0; x < controllerids.length; x++) {
          temp = document.createElement("paper-item");
          temp.setAttribute("value", ips[x]);
          temp.setAttribute("id", names[x]);
          temp.innerHTML = names[x] + " (" + ips[x] + ")";
          Polymer.dom(this.$.controllerListBox).appendChild(temp);
        }

      },

      ready: function() {
        this.$.loadControllers.generateRequest();
      }
    });
