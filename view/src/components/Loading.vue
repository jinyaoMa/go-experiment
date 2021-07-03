<template>
  <div
    class="loading"
    :class="{
      normal: !$route.meta.spinAtCorner,
      corner: $route.meta.spinAtCorner,
    }"
  ></div>
</template>

<style lang="scss" scoped>
.loading {
  box-sizing: border-box;
  position: fixed;
  border-bottom-color: transparent !important;
  border-radius: 50%;
  animation: spin 1s infinite linear;
}
.normal {
  top: calc(50% - 2em);
  right: calc(50% - 2em);
  width: 4em;
  height: 4em;
  border: 0.5em solid;
}
.corner {
  top: 2em;
  right: 2em;
  width: 1em;
  height: 1em;
  border: 0.2em solid;
}
@keyframes spin {
  0% {
    transform: rotate(0deg);
  }
  100% {
    transform: rotate(360deg);
  }
}
</style>