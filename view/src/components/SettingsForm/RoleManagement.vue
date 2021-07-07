<template>
  <div class="role-management">
    <div class="title">
      {{ $locale.settings.roleManagement.title }}
    </div>
    <div class="form">
      <div class="form-item">
        <div class="custom-input">
          <input
            type="text"
            :placeholder="$locale.settings.roleManagement.name"
            v-model="newRole.name"
          />
          <input
            type="text"
            :placeholder="$locale.settings.roleManagement.permission"
            v-model="newRole.permission"
          />
          <input
            type="number"
            min="0"
            step="1"
            :placeholder="$locale.settings.roleManagement.space"
            v-model="newRole.space"
          />
          <div class="custom-select">
            <select v-model="newRole.unit">
              <option
                v-for="(u, i) in ['GB', 'MB', 'KB', 'B']"
                :key="i"
                :value="u"
              >
                {{ u }}
              </option>
            </select>
            <i class="fas fa-sort-down" />
          </div>
        </div>
        <div class="action">
          <button class="btn-add" @click="handleAdd">
            <i class="fas fa-plus" />
          </button>
        </div>
      </div>
      <div class="line"></div>
      <div class="form-item" v-for="role in options.roles" :key="role.ID">
        <div class="custom-input">
          <input type="text" v-model="role.Name" />
          <input type="text" v-model="role.Permission" />
          <input type="number" min="0" step="1" v-model="role.convertedSpace" />
          <div class="custom-select">
            <select v-model="role.unit">
              <option
                v-for="(u, i) in ['GB', 'MB', 'KB', 'B']"
                :key="i"
                :value="u"
              >
                {{ u }}
              </option>
            </select>
            <i class="fas fa-sort-down" />
          </div>
        </div>
        <div class="action">
          <button class="btn-update" @click="handleSave(role)">
            <i class="fas fa-save" />
          </button>
        </div>
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
          roles: [],
        };
      },
    },
  },
  methods: {
    clearAdd() {
      this.newRole = {
        name: "",
        permission: "",
        space: "",
        unit: "GB",
      };
    },
    handleAdd() {
      this.$startLoading();
      this.$http
        .post("/api/admin/role/add", {
          id: this.$user.id,
          token: this.$user.token,
          name: this.newRole.name,
          permission: this.newRole.permission,
          space: this.$revertSpaceByUnit(this.newRole.space, this.newRole.unit),
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
            this.options.roles = data.roles;
            this.clearAdd();
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
    handleSave(role) {
      this.$startLoading();
      this.$http
        .post("/api/admin/role/save", {
          id: this.$user.id,
          token: this.$user.token,
          roleId: role.ID,
          name: role.Name,
          permission: role.Permission,
          space: this.$revertSpaceByUnit(role.convertedSpace, role.unit),
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
            this.options.roles = data.roles;
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
  },
  data() {
    return {
      newRole: {
        name: "",
        permission: "",
        space: "",
        unit: "GB",
      },
    };
  },
};
</script>

<style lang="scss" scoped>
.role-management,
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