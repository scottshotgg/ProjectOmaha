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

      showBackArrow: function() {
        this.$.back.icon = "arrow-back";
        this.$.back.disabled = false;
      },

      hideBackArrow: function() {
        this.$.back.icon = "none";
        this.$.back.disabled = true;
      }
    });
