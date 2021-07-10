<template>
  <span class="btn-recycling" :title="customTitle" @click="$emit('click')">
    <i v-if="forCancelShare" class="fas fa-times-circle fa-fw" />
    <i v-else class="fas fa-trash fa-fw" />
  </span>
</template>

<script>
export default {
  props: {
    forCancelShare: {
      type: Boolean,
      default() {
        return false;
      },
    },
  },
  computed: {
    customTitle() {
      return this.forCancelShare
        ? this.$locale.common.cancelShare
        : this.$locale.common.recycling;
    },
  },
};
</script>

<style lang="scss" scoped>
.btn-recycling {
  color: #333333;
  text-decoration: none;
  transition: 0.2s;
  opacity: 0.5;
  &:hover {
    color: #dd3333;
    opacity: 1;
  }
}
</style>