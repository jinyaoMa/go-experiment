<template>
  <div class="breadcrumb">
    <div class="breadcrumb-start">
      <router-link
        to="/"
        :class="{
          'router-link-exact-active': $route.fullPath === '/',
        }"
      >
        {{ $locale.common.all }}
      </router-link>
      <template v-for="(path, i) in paths">
        <div :key="i">
          <i class="fas fa-angle-right"></i>
          <router-link
            :to="{
              query: {
                currentPath: path.path,
              },
            }"
          >
            {{ path.name }}
          </router-link>
        </div>
      </template>
    </div>
    <div class="breadcrumb-end">
      <span v-if="checkedCount > 0">{{ selectedText }}</span>
    </div>
  </div>
</template>

<script>
export default {
  props: {
    checkedCount: {
      type: Number,
      default() {
        return 0;
      },
    },
    allChecked: {
      type: Boolean,
      default() {
        return false;
      },
    },
  },
  computed: {
    selectedText() {
      if (this.allChecked) {
        return this.$locale.common.selected.all.replace("?", this.checkedCount);
      }
      if (this.checkedCount === 1) {
        return this.$locale.common.selected.one.replace("?", this.checkedCount);
      }
      return this.$locale.common.selected.more.replace("?", this.checkedCount);
    },
    paths() {
      let currentPath = this.$route.query.currentPath;
      let arr = [];
      if (typeof currentPath === "string" && currentPath.length > 0) {
        let lastSlashIndex = currentPath.lastIndexOf("\\");
        while (lastSlashIndex !== -1) {
          let pathname = currentPath.substr(lastSlashIndex + 1);
          arr.unshift({
            name: pathname,
            path: currentPath,
          });
          currentPath = currentPath.substr(0, lastSlashIndex);
          lastSlashIndex = currentPath.lastIndexOf("\\");
        }
        arr.unshift({
          name: currentPath,
          path: currentPath,
        });
      }
      return arr;
    },
  },
};
</script>

<style lang="scss" scoped>
.breadcrumb {
  display: flex;
  flex-direction: row;
  justify-content: space-between;
  align-items: center;
  font-size: 0.8em;
  line-height: 3.75;
  margin-top: 1em;
  background-color: #fafbfc;
  padding: 0 1.12em;
  div {
    display: flex;
    flex-direction: row;
    align-items: center;
  }
  a {
    text-decoration: none;
    color: #2196f3;
    &:hover {
      text-decoration: underline;
    }
    &.router-link-exact-active {
      cursor: default;
      text-decoration: none;
      color: #333333;
    }
  }
  i {
    margin: 0.28em 0.5em 0;
    color: #bbbcbd;
  }
}
.slash {
  margin: 0 0.5em;
  &::before {
    content: ">";
    font-weight: bold;
  }
}
</style>