<template>
  <div class="progressing">
    <div class="progressing-inner">
      <div class="progressing-loaded" :style="loadedStyle"></div>
    </div>
  </div>
</template>

<script>
export default {
  computed: {
    loadedStyle() {
      return {
        width: this.$progress + "%",
      };
    },
  },
};
</script>

<style lang="scss" scoped>
.progressing {
  box-sizing: border-box;
  position: fixed;
  top: calc(50% - 2em);
  left: 0;
  display: flex;
  flex-direction: column;
  align-items: center;
  width: 100%;
}
.progressing-inner {
  box-sizing: border-box;
  width: 280px;
  height: 1em;
  border: 2px solid;
  border-radius: 1em;
  background-color: #ffffff;
  overflow: hidden;
}
.progressing-loaded {
  min-width: 0;
  max-width: 100%;
  width: 0;
  height: 100%;
  background-color: #333333;
  transition: 0.2s;
}
</style>