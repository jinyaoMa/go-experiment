<template>
  <div id="app">
    <router-view :class="{ covered }" />
    <Loading v-if="$loading" />
    <ShareNotice v-if="$hasShareNotice" />
    <Error v-if="$hasError" />
  </div>
</template>

<script>
import Loading from "./components/Loading.vue";
import Error from "./components/Error.vue";
import ShareNotice from "./components/ShareNotice.vue";
export default {
  components: {
    Loading,
    Error,
    ShareNotice,
  },
  computed: {
    covered() {
      return this.$loading || this.$hasShareNotice || this.$hasError;
    },
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
  min-width: 100vw;
  width: fit-content;
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
