<template>
  <div>
    <!-- `greet` は、あらかじめ定義したメソッドの名前 -->
    <input type="text" v-model="name">
    <button v-on:click="display" class="btn-flat-border">{{ button_name }}</button>
    {{ display_name }}
  </div>
</template>

<script>
export default {
  name: 'displaybotton',
  props: ['button_name'],
  data () {
    return {
      name: '',
      display_name :'',
    }
  },
  methods: {
    display: function() {
      this.display_name = this.name
    },
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