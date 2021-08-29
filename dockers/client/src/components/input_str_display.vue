<template>
  <div>
  <div class='displayButton'>
    <!-- `greet` は、あらかじめ定義したメソッドの名前 -->
    <input type="text" v-model="name" value="yahoo" placeholder="What are you looking for?">
    <button v-on:click="display" class="btn-flat-border">{{ button_name }}</button>
  </div>
  </div>
</template>

<script>
import axios from 'axios'
import GoogleMap from './GoogleMap.vue'
export default {
  name: 'displaybutton',
  props: ['button_name'],
  serch_name: '',
  data () {
    return {
      name: '',
      display_name :'',
      info:null,
      errored:null,
      s_data: 'init',
    }
  },
  methods: {
    display: function() {
      // GoogleMap.map.setCenter({lat: this.$parent.search_data.Lat, lng:this.$parent.search_data.Lng});
    
      // test = GoogleMap.data()
    // alert(test.myLatLng.lat)
    // new window.google.maps.Marker({position:this.myLatLng2,map:this.map})
    // this.$parent.search_data = 1
    axios.post('http://localhost:8888', {
    Name: this.name
  }, {
        headers: {
          "Content-Type": "application/json",
        }
      })
  .then(response => this.$parent.search_data = response.data)
  .catch(error => console.log(error))

  this.$parent.flag = true
  },

  }
}

</script>

<style>

</style>