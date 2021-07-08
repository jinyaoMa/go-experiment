<template>
  <span
    class="btn-recycling"
    :title="$locale.common.recycling"
    @click="$emit('click')"
  >
    <i class="fas fa-trash fa-fw" />
  </span>
</template>

<style lang="scss" scoped>
.btn-recycling {
  color: #333333;
  text-decoration: none;
  transition: 0.2s;
  opacity: 0.5;
  &:hover {
    opacity: 1;
  }
}
</style>