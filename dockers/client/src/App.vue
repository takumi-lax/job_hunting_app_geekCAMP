<template>
  <div>
  <alertBotton button_name='Hello'/>
  <displaybotton button_name='abc'/>
  </div>
</template>

<script>
import Textinput from './components/Textinput.vue'
import alertBotton from './components/alertbutton.vue'
import displaybotton from './components/input_str_display.vue'
import HelloWorld from './components/HelloWorld.vue'


export default {
  name: 'App',
  components: {
    Textinput,
    alertBotton,
    displaybotton,
    HelloWorld,
  }
}

</script>

<style>
#app {
  font-family: Avenir, Helvetica, Arial, sans-serif;
  -webkit-font-smoothing: antialiased;
  -moz-osx-font-smoothing: grayscale;
  /* text-align: center; */
  color: #2c3e50;
  margin-top: 60px;
  @import "./css/styles.css";
}
</style>