    Polymer({
      is: 'omaha-navbar',

      properties: {
        selected: {
          type: Number,
          notify: true
        }
      },

      _toggleDrawerPressed: function() {
        this.fire('toggledrawer', {});
      },

      _createAccountPressed: function() {
        this.fire('showaccountcreation', {})
        this.showBackArrow();
      },

      _backPressed: function() {
        this.fire('backpressed', {});
        this.hideBackArrow();
      },

      _logout: function() {
        if(document.cookie.includes("expires")) {
          this.$.logoutAjax.body = {
            //"hash": document.cookie.slice(8, document.cookie.length)
            "hash": document.cookie.split('; ')[2].slice(8, 38)
          };
        } else {
          this.$.logoutAjax.body = {
            //"hash": document.cookie.slice(8, document.cookie.length)
            "hash": document.cookie.split('; ')[1].slice(8, 38)
          };
        }

        this.$.logoutAjax.generateRequest();

        /*var d = new Date();
        d.setYear(d.getYear() - 20);
        document.cookie = "name=" + sessionCookie + "; expires=" + d.toUTCString() + ";";
        */
        document.cookie = "expires=Thu, 01 Jan 1970 00:00:00 UTC";
        //document.cookie = -1;
        console.log(document.cookie);
        window.location = "http://" + window.location.hostname + ":8080";
      },

      showBackArrow: function() {
        this.$.back.icon = "arrow-back";
        this.$.back.disabled = false;
      },

      hideBackArrow: function() {
        this.$.back.icon = "none";
        this.$.back.disabled = true;
      }
    });
