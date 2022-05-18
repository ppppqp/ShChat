

<script>
import VueTerminal from './components/Terminal.vue'
export default {
  data: ()=>({
      messages: [
        {text: "hello", id: 1},
        {text: "world", id: 2}
      ],
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
      this.sendMessage(data.text)
      setTimeout(()=> {
        resolve('')
      }, 1000)
    }
  },
  created: function() {
    console.log("Starting connection to WebSocket Server")
    this.connection = new WebSocket("ws://127.0.0.1:8080/connect")

    this.connection.onmessage = (event) => {
      const currentDate = new Date(); 
      const timestamp = currentDate. getTime()
      this.messages.push({text: event.data, id: timestamp})
      this.$refs.terminal.push(event.data);
    }

    this.connection.onopen = (event) => {
      console.log("Successfully connected to the echo websocket server...")
    }

  }
}

</script>

<template>
    <VueTerminal 
        ref="terminal"
        intro="Welcome to ShChat!"
        console-sign="$"
        allow-arbitrary
        height="90vh"
        @command="onCliCommand"
        ></VueTerminal>
</template>