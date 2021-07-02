<template>
  <div id="app">
    <div class="copyright">© 2021 jinyaoMa</div>
    <router-view />
    <div class="float-panel">
      <BtnLangSwap />
    </div>
  </div>
</template>

<script>
import BtnLangSwap from "./components/BtnLangSwap.vue";

export default {
  components: {
    BtnLangSwap,
  },
};
</script>

<style lang="scss">
body {
  margin: 0;
}
#app {
  font-family: -apple-system, BlinkMacSystemFont, Segoe UI, Helvetica, Arial,
    sans-serif;
  -webkit-font-smoothing: antialiased;
  -moz-osx-font-smoothing: grayscale;
  color: #333333;
  background-color: #ffffff;
  font-weight: 500;
  position: relative;
  font-size: 16px;
}
.copyright {
  position: absolute;
  bottom: 2em;
  left: 1.8em;
  line-height: 1;
  color: #999999;
  font-weight: 400;
  transform-origin: left;
  transform: scale(0.8);
}
.float-panel {
  position: absolute;
  bottom: 1em;
  right: 1em;
}
</style>
