<template>
  <div id="app">
    <router-view
      :class="{
        covered: $loading || $hasError,
      }"
    />
    <Loading v-if="$loading" />
    <Error v-if="$hasError" />
  </div>
</template>

<script>
import Loading from "./components/Loading.vue";
import Error from "./components/Error.vue";
export default {
  components: {
    Loading,
    Error,
  },
  mounted() {
    console.log(this);
    this.$stopLoading();
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
  font-size: 16px;
  line-height: 1.3;
}
.covered {
  filter: opacity(0.3);
  pointer-events: none;
}
</style>
