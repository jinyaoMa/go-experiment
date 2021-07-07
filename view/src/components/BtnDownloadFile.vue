<template>
  <a class="btn-download-file" :href="url" download="file">
    <i class="fas fa-download fa-fw" />
  </a>
</template>

<script>
export default {
  props: {
    url: {
      type: String,
    },
  },
};
</script>

<style lang="scss" scoped>
.btn-download-file {
  color: #333333;
  text-decoration: none;
  transition: 0.2s;
  opacity: 0.5;
  &:hover {
    opacity: 1;
  }
}
</style>