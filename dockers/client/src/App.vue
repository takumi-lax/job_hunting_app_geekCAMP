<template>
  <div>
  <div class="bg"></div>
  <div class="bg bg2"></div>
  <div class="bg bg3"></div>

    <div class='center'>
      <h1 class="ribbon1">CircleRecommender</h1>
    </div>
    
    <div class='searchbox'>
      <displaybutton button_name='検索'/>
    </div>

    {{ search_data }}


    <body  class='center'>
      <select name="category">
        <option value="restaurants">飲食店</option>
      </select>

      <select name="radius">
        <option value=1.5>1.5km</option>
        <option value=3>3km</option>
        <option value=5>5km</option>
      </select>
    </body>

    <div class='center'>
      <GoogleMap/> 
    </div>

  </div>
</template>

<script>
import displaybutton from './components/input_str_display.vue'
// import coolsearchbar from './components/coolsearchbar.vue'
import GoogleMap from './components/GoogleMap.vue'
import store from './components/store.vue'

export default {
  name: 'App',
data() {
    return {
      number: 1,
      // 受け取ったデータを格納する初期値
      search_data: {Lat: 35.863, Lng: 139.607},
      flag:true,
    }},
  
  components: {
    store,
    displaybutton,
    // coolsearchbar,
    GoogleMap
  },
  }
</script>

<style>
#app {
  font-family: Avenir, Helvetica, Arial, sans-serif;
  -webkit-font-smoothing: antialiased;
  -moz-osx-font-smoothing: grayscale;
  text-align: center;
  color: #0f80f1;
  margin-top: 60px;
  @import "./css/styles.css";
}


.searchbox {
  font-family: Avenir, Helvetica, Arial, sans-serif;
  -webkit-font-smoothing: antialiased;
  -moz-osx-font-smoothing: grayscale;
  text-align: center;
  color: #2c3e50;
  margin-top: 80px;
  transform: scale(2.0);
  @import "./css/styles.css";
}

.center {
  text-align:center;
  margin-top: 20px;
}


select {
-webkit-appearance: none;/* ベンダープレフィックス(Google Chrome、Safari用) */
-moz-appearance: none; /* ベンダープレフィックス(Firefox用) */
appearance: none; 標準のスタイルを無効にする
display: block; /* ブロックレベル要素化する */
width: 150px; /* ボックスの高さを指定する */
height: 40px;
padding: 10px 15px; /* ボックス内側の余白を指定する */
/* color: #b20000; フォントの色を指定 */
/* border: none; 境界線をなしに指定する */
border-radius: 5px; /* 角丸を指定する */
background: #FFF;/* 背景を透過指定する */
}
::-ms-expand { /* select要素のデザインを無効にする（IE用） */
display: none;
}

.ribbon1 {
  display: inline-block;
  position: relative;
  height: 50px;/*リボンの高さ*/
  line-height: 50px;/*リボンの高さ*/
  text-align: center;
  padding: 0 30px;/*横の大きさ*/
  font-size: 24px;/*文字の大きさ*/
  background: #f57a78;/*塗りつぶし色*/
  color: #FFF;/*文字色*/
  box-sizing: border-box;
  transform: scale(2.0);
  @import url('https://fonts.googleapis.com/css?family=Noto+Serif+JP');
  font-family: 'Noto Serif JP', sans-serif;
}

.ribbon1:before, .ribbon1:after {
  position: absolute;
  content: '';
  width: 0px;
  height: 0px;
  z-index: 1;
  
}

.ribbon1:before {
  top: 0;
  left: 0;
  border-width: 25px 0px 25px 15px;
  border-color: transparent transparent transparent #fff;
  border-style: solid;
  
}

.ribbon1:after {
  top: 0;
  right: 0;
  border-width: 25px 15px 25px 0px;
  border-color: transparent #fff transparent transparent;
  border-style: solid;
}

</style>