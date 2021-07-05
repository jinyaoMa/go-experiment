<template>
  <div
    class="btn-home"
    :title="$locale.common.btnHomeTips"
    @click="handleClick"
  >
    <i class="fas fa-hdd" />
  </div>
</template>

<script>
export default {
  methods: {
    handleClick() {
      this.$router.push("/");
    },
  },
};
</script>

<style lang="scss" scoped>
.btn-home {
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