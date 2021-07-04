<template>
  <div
    class="btn-logout"
    @click="handleLogout"
    :title="$locale.common.btnLogout"
  >
    <i class="fas fa-sign-out-alt" />
  </div>
</template>

<script>
export default {
  methods: {
    handleLogout() {
      this.$startLoading();
      this.$http
        .post("/api/logout", {
          id: this.$user.id,
          token: this.$user.token,
        })
        .then((res) => {
          if (res.data.success) {
            this.$setUser({
              id: 0,
            });
            this.$router.push("/login");
            this.$stopLoading();
          } else {
            this.$stopLoading();
          }
        });
    },
  },
};
</script>

<style lang="scss" scoped>
.btn-logout {
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