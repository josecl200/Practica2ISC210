var socket;

socket = new WebSocket("ws://127.0.0.1:8080/ss");

socket.onmessage = function(evt) {
  var newData = JSON.parse(evt.data);
  console.log(evt.data); 
  tictactoe.gameState = newData;
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
