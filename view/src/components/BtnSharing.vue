<template>
  <span class="btn-sharing" :title="customTitle" @click="$emit('click')">
    <i v-if="forUpdate" class="fas fa-sync-alt fa-fw" />
    <i v-else class="fas fa-share-alt fa-fw" />
  </span>
</template>

<script>
export default {
  props: {
    forUpdate: {
      type: Boolean,
      default() {
        return false;
      },
    },
  },
  computed: {
    customTitle() {
      return this.forUpdate
        ? this.$locale.common.updateShare
        : this.$locale.common.sharing;
    },
  },
};
</script>

<style lang="scss" scoped>
.btn-sharing {
  color: #333333;
  text-decoration: none;
  transition: 0.2s;
  opacity: 0.5;
  &:hover {
    opacity: 1;
  }
}
</style>