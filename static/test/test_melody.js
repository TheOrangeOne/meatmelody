var assert = require('assert');
var Melody = require('../js/meatmelody.js');
var { WebSocket, Server } = require('./mock-socket.min.js');

describe('Melody', () => {
  describe('initialization', () => {
    describe('Melody()', () => {
      it('should not err out', () => {
        var melody = new Melody({
          WebSocket: WebSocket
        });
        assert(melody);
      });
    });
  });
});
