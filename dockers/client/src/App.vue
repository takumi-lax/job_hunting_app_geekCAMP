<template>
  
  <div>

    <div class='center'>
      <h1 class="ribbon1">就活が楽しくなるアプリ</h1>
    </div>
    
    <div class='searchbox'>
      <body class="align">
        <div class="grid">
          <form action="" method="get" class="search">
            <div class="form__field">
              <input type="search" name="search" placeholder="What are you looking for?" class="form__input">
              <input type="submit" value="Search" class="button">
            </div>
          </form>
        </div>
      </body>
      <displaybutton button_name='検索'/>
    </div>

    <body  class='center'>
      <select name="category">
        <option value="restaurants">飲食店</option>
        <!-- <option value="feb">2月</option>
        <option value="mar">3月</option> -->
      </select>

      <select name="radius">
        <option value=1>1km</option>
        <option value=3>3km</option>
        <option value=5>5km</option>
      </select>
    </body>

  <!-- <div>
    <coolsearchbar/>
  </div> -->

    <div class='center'>
      <GoogleMap/>
    </div>
  </div>
</template>

<script>
import displaybutton from './components/input_str_display.vue'
// import coolsearchbar from './components/coolsearchbar.vue'
import GoogleMap from './components/GoogleMap.vue'

export default {
  name: 'App',
  components: {
    displaybutton,
    // coolsearchbar,
    GoogleMap
  }
}
</script>

<style>

</style>