<template>

  <div id="app">
    <h1>{{ message }}</h1>
    <a href="https://www.google.co.jp/maps/?hl=ja">自分で調べてください</a>
    <a href="#" class="btn-flat-border">BUTTON</a>
  </div>

</template>

<script>
export default {
  data(){
    return {
      message :'google map'
          }
  }
  
}

alert('アラートです')
</script>

<style>
#app {
  font-family: Avenir, Helvetica, Arial, sans-serif;
  -webkit-font-smoothing: antialiased;
  -moz-osx-font-smoothing: grayscale;
  text-align: center;
  color: #2c3e50;
  margin-top: 60px;
  @import "./css/styles.css";
}
</style>