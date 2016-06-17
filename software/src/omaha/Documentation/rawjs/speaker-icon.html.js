    var colors = [
      '--paper-brown-500',
      '--paper-deep-orange-500',
      '--paper-orange-500',
      '--paper-amber-500',
      '--paper-lime-500',
      '--paper-light-green-500',
      '--paper-green-500',
      '--paper-teal-500',
      '--paper-cyan-500',
      '--paper-light-blue-500',
      '--paper-blue-500',
      '--paper-indigo-500',
      '--paper-deep-purple-500',
      '--paper-purple-500',
      '--paper-pink-500',
      '--paper-red-500'
    ];

    Polymer({
      is: 'speaker-icon',

      properties: {
        x: Number,
        y: Number,
        speakerId: Number
      },

      ready: function () {
          this.$.icon.style.left = this.x.toString() + "px";
          this.$.icon.style.top = this.y.toString() + "px";
          var color = colors[Math.floor(Math.random()*colors.length)];
          this.customStyle['--speaker-background-color'] = 'var(' + color + ')';
          this.updateStyles();
      }
    });
