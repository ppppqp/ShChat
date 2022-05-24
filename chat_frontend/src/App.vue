

<script>
import VueTerminal from './components/Terminal.vue'
import { ref } from 'vue'
export default {
  data: ()=>({
      messages: [
        {text: "hello", id: 1},
        {text: "world", id: 2}
      ],
      connections: null,
      helpMessage: ref(false),
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
      this.sendMessage(JSON.stringify(data))
      setTimeout(()=> {
        resolve('')
      }, 1000)
    },
    turnOnHelpMessage: function(){this.helpMessage = true}
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
    // window.onbeforeunload = function() {
    //   this.connection.onclose = function () {}; // disable onclose handler first
    //   this.connection.close();
    // };
  }
}

</script>
<style scoped>
  a{
    text-decoration: none;
    color: white;
  }
  code{
    color: rgb(224, 108, 117)
  }
</style>
<template>
    <q-bar class="bg-black text-white">
      <div class="cursor-pointer gt-xs">View</div>
      <div class="cursor-pointer" @click="turnOnHelpMessage">Help</div>
      <div class="cursor-pointer gt-xs"><a href=https://github.com/ppppqp/ShChat>GitHub</a></div>
      <q-space />
      <q-btn dense flat icon="minimize" />
      <q-btn dense flat icon="crop_square" />
      <q-btn dense flat icon="close" />
    </q-bar>
    <VueTerminal 
        ref="terminal"
        intro="Welcome to ShChat!
  ____  _          _           _   
 / ___|| |__   ___| |__   __ _| |_ 
 \___ \| '_ \ / __| '_ \ / _` | __|
  ___) | | | | (__| | | | (_| | |_ 
 |____/|_| |_|\___|_| |_|\__,_|\__|
                                   
"
        console-sign="$"
        allow-arbitrary
        height="100vh-32px"
        @command="onCliCommand"
        ></VueTerminal>


    <q-dialog v-model="helpMessage">
      <q-card>
        <q-card-section>
          <div class="text-h6">Manual</div>
        </q-card-section>

        <q-card-section class="q-pt-none">
          <strong></strong><br>

          A brief overview of the commands:
          <ul>
            <li><code>help</code> : If you are geek enough, use this command to find out the rest for yourself!</li>
            <li><code>history</code> : History list of all the commands.</li>
            <li><code>clear</code> : Clear everything on the terminal</li>
            <li><code>/nick &lt;nick_name&gt;</code> : Create a &lt;nick_name&gt; for yourself.</li>
            <li><code>/rooms</code> : See the list of all existing rooms.
              <ul>
                <li><code>/rooms -v </code> See a verbose list.</li>
              </ul>
            </li>
            <li><code>/join &lt;room_name&gt;</code> : Join the &lt;room_name&gt;. If &lt;room_name&gt; doesn't exists, create it.</li>
            <li><code>/leave </code> : Leave the current room.</li>
            <li><code>/msg &lt;message&gt;</code> : Send the &lt;message&gt; to everyone else in the room.</li>
            <li><code>/quit</code> : Disconnect.</li>

          </ul>
        </q-card-section>

        <q-card-actions align="right">
          <q-btn flat label="OK" color="primary" v-close-popup />
        </q-card-actions>
      </q-card>
    </q-dialog>
</template>