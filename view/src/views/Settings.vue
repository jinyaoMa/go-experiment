<template>
  <div class="settings">
    <Head />
    <div class="main"></div>
  </div>
</template>

<script>
import Head from "../components/Head.vue";
export default {
  components: {
    Head,
  },
};
</script>

<style lang="scss" scoped>
.settings {
  display: flex;
  flex-direction: column;
  height: 100vh;
}
.main {
  flex-grow: 1;
  display: flex;
  flex-direction: row;
  letter-spacing: 0.1em;
}
</style>