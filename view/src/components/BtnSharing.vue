<template>
  <span
    class="btn-sharing"
    :title="$locale.common.sharing"
    @click="$emit('click')"
  >
    <i class="fas fa-share-alt fa-fw" />
  </span>
</template>

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