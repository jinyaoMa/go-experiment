<template>
  <div class="settings">
    <Head isSettings />
    <div class="main">
      <Basic />
      <div class="divider" />
      <RoleManagement />
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

export default {
  components: {
    Head,
    Basic,
    RoleManagement,
    ResetPassword,
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