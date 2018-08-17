function Melody(opts) {
  opts = opts || {};
  this.conn = null;
  this.host = opts.host;
  this.path = opts.path
  this.secure = opts.secure;
  this.id = opts.id;
  this.WebSocket = opts.WebSocket;

  var self = this;

  this.onopen = function() {
    console.log('connected!');
  };

  this.onmessage = function() {
    console.log('message!');
  };

  this.onclose = function() {
    console.log('disconnected!');
  };

  this.genId = function() {
    return Math.random()
      .toString(36)
      .replace(/[^a-z]+/g, '')
      .substr(0, 5);
  };

  this.initId = function() {
    self.id = self.id || self.genId();
  };

  this.initWebSocket = function() {
    var proto = self.secure ? 'wss' : 'ws';
    var id = '?id=' + self.id;
    var wsURL = proto + '://' + self.host + '/ws/' + self.path + id;
    var conn = new self.WebSocket(wsURL);
    conn.onopen = self.onopen;
    conn.onmessage = self.onmessage;
    conn.onclose = self.onclose;

    self.conn = conn;
  };

  this.init = function() {
    self.initId()
    self.initWebSocket();
  };

  this.init();
};

if (typeof window === 'undefined') {
  module.exports = Melody;
}

