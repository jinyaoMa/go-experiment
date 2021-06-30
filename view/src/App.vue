<template>
  <div id="app">
    <router-view />
    <BtnLangSwap />
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
  font-family: Avenir, Helvetica, Arial, sans-serif;
  -webkit-font-smoothing: antialiased;
  -moz-osx-font-smoothing: grayscale;
  color: #2c3e50;
}
</style>
