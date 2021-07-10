<template>
  <div class="settings">
    <Head isSettings />
    <div class="main">
      <Basic :options="options" />
      <div v-if="$user.permission.includes('ADMIN:1')" class="divider" />
      <UserManagement
        v-if="$user.permission.includes('ADMIN:1')"
        :options="options"
      />
      <div v-if="$user.permission.includes('ADMIN:1')" class="divider" />
      <RoleManagement
        v-if="$user.permission.includes('ADMIN:1')"
        :options="options"
      />
      <div class="divider" />
      <ResetPassword />
    </div>
  </div>
</template>

<script>
import Head from "../components/Head.vue";
import Basic from "../components/SettingsForm/Basic.vue";
import RoleManagement from "../components/SettingsForm/RoleManagement.vue";
import ResetPassword from "../components/SettingsForm/ResetPassword.vue";
import UserManagement from "../components/SettingsForm/UserManagement.vue";

export default {
  components: {
    Head,
    Basic,
    RoleManagement,
    ResetPassword,
    UserManagement,
  },
  data() {
    return {
      options: {},
    };
  },
  mounted() {
    this.$startLoading();
    this.$http
      .post("/api/admin/settings", {
        id: this.$user.id,
        token: this.$user.token,
      })
      .then((res) => {
        if (res.data.success) {
          let data = res.data.data;
          data.roles.forEach((role) => {
            role.unit = this.$getUnitBySpace(role.Space);
            role.convertedSpace = this.$convertSpaceByUnit(
              role.Space,
              role.unit
            );
          });
          this.options = data;
        } else {
          this.$router.push("/login");
        }
        this.$stopLoading();
      })
      .catch((err) => {
        this.$router.push("/login");
        this.$stopLoading();
      });
  },
};
</script>

<style lang="scss" scoped>
.settings {
  display: flex;
  flex-direction: column;
  height: 100vh;
}
.main {
  flex-grow: 1;
  display: flex;
  flex-direction: column;
  letter-spacing: 0.1em;
  padding: 1.8em;
}
.divider {
  border-top: 2px solid;
  margin: 2.8em 0 1.8em;
}
</style>