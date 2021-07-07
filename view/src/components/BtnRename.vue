<template>
  <span
    class="btn-rename"
    :title="$locale.common.rename"
    @click="$emit('click')"
  >
    <i class="fas fa-pen fa-fw" />
  </span>
</template>

<style lang="scss" scoped>
.btn-rename {
  color: #333333;
  text-decoration: none;
  transition: 0.2s;
  opacity: 0.5;
  &:hover {
    opacity: 1;
  }
}
</style>