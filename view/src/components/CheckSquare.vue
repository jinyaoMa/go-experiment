<template>
  <i
    class="check-square far"
    :class="{
      'fa-check-square': this.isChecked,
      'fa-square': !this.isChecked,
      'fa-minus-square': this.isPrepared,
    }"
    @click="handleClick"
  />
</template>

<script>
export default {
  props: {
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
  data() {
    return {
      isChecked: this.checked,
      isPrepared: this.prepared,
    };
  },
  methods: {
    handleClick() {
      this.isChecked = !this.isChecked;
      this.$emit("click", this.isChecked);
    },
  },
};
</script>

<style lang="scss" scoped>
.check-square {
  font-size: 1em;
}
</style>