<template>
  <div class='displayButton'>
    <!-- `greet` は、あらかじめ定義したメソッドの名前 -->
    <input type="text" v-model="name">
    <button v-on:click="display" class="btn-flat-border">{{ button_name }}</button>
    {{ display_name }}
  </div>
</template>

<script>
export default {
  name: 'displaybutton',
  props: ['button_name'],
  serch_name: ''
  data () {
    return {
      name: '',
      display_name :'',
    }
  },
  methods: {
    display: function() {
      
      this.name

    },
  }
}
</script>

<style>

</style>