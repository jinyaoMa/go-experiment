<template>
  <div class="btn-settings" :title="$locale.common.btnSettingsTips">
    <i class="fas fa-cog" />
  </div>
</template>

<style lang="scss" scoped>
.btn-settings {
  display: flex;
  width: 3em;
  height: 3em;
  line-height: 3em;
  justify-content: center;
  align-items: center;
  cursor: pointer;
  user-select: none;
  transition: 0.2s;
  font-weight: bold;
  background-color: #ffffffaa;
  letter-spacing: 0;
  border-radius: 4px;
  &:hover {
    text-shadow: 0.1em 0.1em #dcbdcb;
  }
  &:active {
    background-color: #f1f2f3aa;
  }
}
</style>