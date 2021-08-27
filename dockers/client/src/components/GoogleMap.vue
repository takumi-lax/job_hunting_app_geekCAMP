
<template>
<div>
<h1>Google Map</h1>



<div ref="map" style="height:500px;width:1000px;"></div>

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
  data(){
    return {
      map:'',
      myLatLng :{lat: 35.863, lng: 139.607},
      myLatLng2:{lat: 35.862, lng: 139.608},
      myLatLng3:{lat: 35.862, lng: 139.606},
    }
  },
  mounted(){
    let script = document.createElement('script');
    script.src = process.env.Google_API_KEY;
    script.async = true;
    document.head.appendChild(script);
    let timer = setInterval(() => {
      if(window.google){
        clearInterval(timer);
        this.map = new window.google.maps.Map(this.$refs.map, {
          center: this.myLatLng,
          zoom: 18,

          mapTypeControl: true, //マップタイプ コントロール
          fullscreenControl: true, //全画面表示コントロール
          streetViewControl: true, //ストリートビュー コントロール
          zoomControl: true, //ズーム コントロール

        });
        new window.google.maps.Marker({position:this.myLatLng,map:this.map})
        new window.google.maps.Marker({position:this.myLatLng2,map:this.map})
        new window.google.maps.Marker({position:this.myLatLng3,map:this.map})
      }
    }, 500)

  }
}

</script>