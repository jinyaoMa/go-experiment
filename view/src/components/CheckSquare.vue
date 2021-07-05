<template>
  <i
    class="check-square"
    :class="{
      'fas fa-check-square': !this.prepared && this.checked,
      'far fa-square': !this.prepared && !this.checked,
      'fas fa-minus-square': this.prepared,
    }"
    @click="handleClick"
  />
</template>

<script>
export default {
  props: {
    fileId: {
      type: Number,
      default() {
        return 0;
      },
    },
    checked: {
      type: Boolean,
      default() {
        return false;
      },
    },
    prepared: {
      type: Boolean,
      default() {
        return false;
      },
    },
  },
  methods: {
    handleClick() {
      this.$emit("click", this.fileId, !this.checked);
    },
  },
};
</script>

<style lang="scss" scoped>
.check-square {
  font-size: 1em;
}
</style>