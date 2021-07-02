<template>
  <div
    class="btn-lang-swap"
    @click="$swapLang"
    :title="$locale.common.btnLangTips"
  >
    <span class="text">{{ $locale.common.langTag }}</span>
  </div>
</template>

<script>
export default {
  name: "btn-lang-swap",
};
</script>

<style lang="scss" scoped>
.btn-lang-swap {
  display: flex;
  width: 3em;
  height: 3em;
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
.text {
  display: inline-block;
  line-height: 1;
}
</style>