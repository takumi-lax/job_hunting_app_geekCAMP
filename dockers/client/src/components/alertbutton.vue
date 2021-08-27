<template>
<div class="hello">
    <!-- `greet` は、あらかじめ定義したメソッドの名前 -->
    <button v-on:click="greet" class="btn-flat-border">{{ button_name }}</button>
</div>
</template>

<script>
export default {
  name: 'alertButton',
  props: ['button_name'],
  methods: {
  greet: function (event) {
    // メソッド内の `this` は、 Vue インスタンスを参照します
    alert('Hello ' + this.button_name + '!')
    // `event` は、ネイティブ DOM イベントです
    if (event) {
      alert(event.target.tagName)
    }
  }
  }
}
</script>

<style>
#alertBotton {
  font-family: Avenir, Helvetica, Arial, sans-serif;
  -webkit-font-smoothing: antialiased;
  -moz-osx-font-smoothing: grayscale;
  text-align: center;
  color: #2c3e50;
  margin-top: 60px;
  @import "../css/styles.css";
}
</style>