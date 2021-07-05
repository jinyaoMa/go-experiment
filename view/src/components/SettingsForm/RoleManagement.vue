<template>
  <div class="role-management">
    <div class="title">
      {{ $locale.settings.roleManagement.title }}
    </div>
  </div>
</template>

<style lang="scss" scoped>
.role-management {
  display: flex;
  flex-direction: column;
  align-items: center;
}
.title {
  font-size: 1.5em;
  font-weight: bold;
  margin-bottom: 1em;
}
</style>