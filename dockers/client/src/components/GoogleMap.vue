
<template>
<div>
      <div ref="map" style="height:600px;width:1300px;margin-left:auto;margin-right:auto;" ></div>
      <button v-on:click="dis" class="btn-flat-border">{{ button_name }}</button>
</div>
</template>

<script>
{/* export default {
  data(){
    return {
      map:'',
    }
  },
  mounted(){
    let timer = setInterval(() => {
      if(window.google){
        clearInterval(timer);
        this.map = new window.google.maps.Map(this.$refs.map, {
          center: {lat: 35.863, lng: 139.608},
          zoom: 18
        });       
      }
    },500)
  }
} */}

export default {
  name:'GoogleMap',
  props: ['s_data'],
  data(){
    return {
      map:'',
      myLatLng :{lat: 35.863, lng: 139.607},
      myLatLng2:{lat: 35.862, lng: 139.608},
      myLatLng3:{lat: 35.862, lng: 139.606},
    }
  },
  mounted() {
    let script = document.createElement('script');
    script.src = process.env.Google_API_KEY;
    script.async = true;
    document.head.appendChild(script);
    let timer = setInterval(() => {
      if(window.google){
        clearInterval(timer);
        this.map = new window.google.maps.Map(this.$refs.map, {
          // center: {lat: this.$parent.search_data.Lat, lng:this.$parent.search_data.Lng},
          center: this.myLatLng,
          zoom: 18,

          mapTypeControl: true, //マップタイプ コントロール
          fullscreenControl: true, //全画面表示コントロール
          streetViewControl: true, //ストリートビュー コントロール
          zoomControl: true, //ズーム コントロール

        });
        new window.google.maps.Marker({position:this.myLatLng,map:this.map})
        // new window.google.maps.Marker({position:this.myLatLng2,map:this.map})
        // new window.google.maps.Marker({position:this.myLatLng3,map:this.map})
      }
    }, 500)},

  methods: {
  //   isNavigatorAvailable: () => {
  //   return navigator.geolocation;
  // },
      
      dis : function() {

        // console.log('update center: ', this.lat, this.lng, this);
          console.log(this.$parent.search_data)
          alert(this.$parent.search_data.Lat)
          this.map.setCenter({lat: this.$parent.search_data.Lat, lng:this.$parent.search_data.Lng});
          this.map.Maker.setPosition({lat: this.$parent.search_data.Lat, lng:this.$parent.search_data.Lng});
    
  }
  }
}

</script>

<style>
.center {
  text-align:center
}
</style>