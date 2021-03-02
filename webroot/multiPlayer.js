var socket;
var firstMessage = false;
socket = new WebSocket("ws://127.0.0.1:8080/ws");


socket.onmessage = function(evt) {
  var newData = JSON.parse(evt.data);
  console.log(evt.data); 
  tictactoe.gameState = newData;
  if (!firstMessage){
    firstMessage = true;
    tictactoe.letter = newData.started ? 'O' : 'X'
  }
};

Vue.config.debug = true; 

Vue.transition('board',
               {enterClass : 'bounceInDown', leaveClass : 'bounceOutDown'});

var tictactoe = new Vue({
  el : '#tictactoe',
  data : {

    gameState : {
      started : false,
      fields : [],
    },
    // Special Move coding scheme
    RESTART : 10,
    letter: 'x'
  },
  computed : {
    row1 : function() { return this.gameState.fields.slice(0, 3); },
    row2 : function() { return this.gameState.fields.slice(3, 6); },
    row3 : function() { return this.gameState.fields.slice(6, 9); },
  },
  methods : {
    makeMove : function(fieldNum) { socket.send(fieldNum); },
  }
});
