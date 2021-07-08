<template>
  <span class="btn-cut" :title="$locale.common.cut" @click="$emit('click')">
    <i class="fas fa-cut fa-fw" />
  </span>
</template>

<style lang="scss" scoped>
.btn-cut {
  color: #333333;
  text-decoration: none;
  transition: 0.2s;
  opacity: 0.5;
  &:hover {
    opacity: 1;
  }
}
</style>