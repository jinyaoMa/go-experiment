<template>
  <div class="user-management">
    <div class="title">
      {{ $locale.settings.userManagement.title }}
    </div>
    <div class="form-item" v-for="user in options.users" :key="user.ID">
      <div class="custom-input">
        <input
          type="text"
          v-model="user.Account"
          :placeholder="$locale.settings.userManagement.account"
          disabled
        />
        <input
          type="password"
          v-model="user.Password"
          :placeholder="$locale.settings.userManagement.password"
          autocomplete="off"
        />
        <div class="custom-select">
          <select
            v-model="user.RoleID"
            :placeholder="$locale.settings.userManagement.role"
          >
            <option
              v-for="role in options.roles"
              :key="role.ID"
              :value="role.ID"
            >
              {{ role.Name }}
            </option>
          </select>
          <i class="fas fa-sort-down" />
        </div>
      </div>
      <div class="action">
        <button class="btn-update" @click="handleSave(user)">
          <i class="fas fa-save" />
        </button>
      </div>
    </div>
  </div>
</template>

<script>
export default {
  props: {
    options: {
      type: Object,
      default() {
        return {
          users: [],
          roles: [],
        };
      },
    },
  },
  methods: {
    handleSave(user) {
      this.$startLoading();
      this.$http
        .post("/api/admin/user/save", {
          id: this.$user.id,
          token: this.$user.token,
          userId: user.ID,
          roleId: user.RoleID,
          password: user.Password,
        })
        .then((res) => {
          if (res.data.success) {
            let data = res.data.data;
            this.options.users = data.users;
          } else {
            this.$showError(this.$locale.common.errorMsg);
          }
          this.$stopLoading();
        })
        .catch((err) => {
          this.$showError(this.$locale.common.errorServer);
          this.$stopLoading();
        });
    },
  },
};
</script>

<style lang="scss" scoped>
.user-management,
.form {
  display: flex;
  flex-direction: column;
  align-items: center;
}
.title {
  font-size: 1.5em;
  font-weight: bold;
  margin-bottom: 1em;
}
.line {
  border-top: 2px solid #dddedf;
  margin: 1em 0;
  width: 100%;
}
.form-item,
.custom-input,
.action,
.custom-select {
  display: flex;
  flex-direction: row;
  position: relative;
}
.custom-select > i {
  position: absolute;
  right: 0.6em;
  top: 0.4em;
}
.form-item:not(:first-child):not(:last-child) {
  margin-bottom: 1em;
}
input,
select,
button {
  border-radius: 0;
  border-width: 1px;
  border-style: solid;
  outline: none;
  font-size: 0.9em;
  line-height: 2.33;
}
input,
select {
  border-right-width: 0;
}
input {
  padding: 0 1em;
  &:first-child {
    border-radius: 4px 0 0 4px;
  }
  &:nth-child(2) {
    min-width: 200px;
  }
  &:nth-child(3) {
    max-width: 100px;
  }
  &:disabled {
    border-color: #333333;
  }
}
select {
  padding: 0 2em 0 1em;
  appearance: none;
  position: relative;
}
button {
  border-radius: 0 4px 4px 0;
  width: 2.6em;
}
.btn-add {
  background-color: #333333;
  color: #ffffff;
  &:hover {
    background-color: #555555;
  }
}
.btn-update {
  background-color: #f1f2f3;
  color: #333333;
  &:hover {
    background-color: #eaebec;
  }
}
</style>