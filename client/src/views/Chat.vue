<template>
  <div class="container">
    <div class="row">
      <div class="frame">
        <ul id="msgs">
          <li v-for="(msg, index) in msgs" :key="index" id="li">
            <div class="msj macro" v-if="msg.senderId !== userId">
              <div class="avatar">
                <img class="img-circle" src="images/akshatm.png " />
              </div>
              <div class="text text-l">
                <p>{{ msg.body }}</p>
                <p>
                  <small>{{ msg.time | agoCalculation }}</small>
                </p>
              </div>
            </div>
            <div class="msj-rta macro" v-else>
              <div class="text text-r">
                <p>{{ msg.body }}</p>
                <p>
                  <small>{{ msg.time | agoCalculation }}</small>
                </p>
              </div>
              <div class="avatar" style="padding:0px 0px 0px 10px !important">
                <img class="img-circle" src="images/bot.png " />
              </div>
            </div>
          </li>
        </ul>
      </div>
    </div>
    <div class="row type-msg">
      <div class="col-9 row-custom">
        <div class="type-r" style="background:whitesmoke !important">
          <input class="mytext" v-model="msg" placeholder="Type a message" />
        </div>
      </div>

      <div class="col-3 mt-1">
        <b-button variant="primary" @click="sendMessage" class="send-btn">Send</b-button>
      </div>
    </div>
  </div>
</template>

<script>
export default {
  name: 'Chat',
  data() {
    return {
      connection: null,
      msg: '',
      msgList: [],
      userId: ''
    }
  },
  mounted() {
    this.openSocketConnection()
  },
  computed: {
    msgs() {
      return this.msgList.filter(m => m.messageType === 'message')
    }
  },
  methods: {
    openSocketConnection() {
      this.connection = new WebSocket("ws://localhost:3600/chat");
      console.log("Attempting web socket connection");
      this.connection.onopen = () => {
        console.log("connected successfully");
      };
      this.connection.onerror = err => {
        console.log(err);
      };
      this.connection.onmessage = msg => {
        let newMsg = JSON.parse(msg.data);
        console.log(newMsg.userId);
        if (newMsg.userId) {
          this.userId = newMsg.userId;
        } else {
          this.msgList.push(JSON.parse(msg.data));
        }
        var objDiv = document.getElementById("msgs");
        objDiv.scrollTop = objDiv.scrollHeight;
      };
    },
    sendMessage() {
      if (this.msg !== '') {
        this.connection.send(this.msg)
        this.msg = ''
      }
    }
  },
  filters: {
    agoCalculation(time) {
      var d = new Date(time * 1000)
      // const monthNames = [
      //   'Jan',
      //   'Feb',
      //   'Mar',
      //   'Apr',
      //   'May',
      //   'Jun',
      //   'Jul',
      //   'Aug',
      //   'Sep',
      //   'Oct',
      //   'Nov',
      //   'Dec'
      // ]
      return d.toLocaleString()
      // +
      // '' +
      // d.getDate() +
      // '-' +
      // monthNames[d.getMonth() + 1] +
      // '-' +
      // d.getFullYear()
    }
  },
  destroyed() {
    this.connection.onclose = () => {
      console.log('Disconnected Successfully')
    }
  }
}
</script>

<style scoped>
/* Chat containers */
</style>
