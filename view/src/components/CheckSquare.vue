<template>
  <i
    class="check-square far"
    :class="{
      'fa-check-square': this.isChecked,
      'fa-square': !this.isChecked,
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
  },
  data() {
    return {
      isChecked: this.checked,
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