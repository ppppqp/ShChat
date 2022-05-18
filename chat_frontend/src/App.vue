

<script>
import VueTerminal from './components/Terminal.vue'
export default {
  data: ()=>({
      text: "",
      response: "abcac",
      messages: ['a', 'b', 'c', 'd'],
      connections: null
    }),
  components:{
    VueTerminal
  },
  methods:{
    sendMessage: function(message) {
      this.connection.send(message);
    },
    onCliCommand(data, resolve, reject){
      // typed command is available in data.text
      // don't forget to resolve or reject the Promise
      console.log(data)
      setTimeout(()=> {
        resolve('')
      }, 300)
    }
  },
  created: function() {
    console.log("Starting connection to WebSocket Server")
    this.connection = new WebSocket("ws://127.0.0.1:8080/connect")

    this.connection.onmessage = (event) => {
      console.log(event.data);
      this.response = event.data;
      console.log("response is", this.response);
    }

    this.connection.onopen = (event) => {
      console.log("Successfully connected to the echo websocket server...")
    }

  }
}

</script>

<template>
  <VueTerminal :intro="intro"
            console-sign="$"
            allow-arbitrary
            height="500px"
            @command="onCliCommand"
            executeCommand="sendMessage"
            ></VueTerminal>
  <input v-model="text">{{text}}
  <h1>{{ response }}</h1>
  <button @click="sendMessage(text)">Send</button>


</template>