<template>
  <div class="breadcrumb">
    <div class="breadcrumb-start">
      <router-link
        :to="{
          query: {
            currentPath: '',
          },
        }"
        exact
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
            exact
          >
            {{ path.name }}
          </router-link>
        </div>
      </template>
    </div>
    <div class="breadcrumb-end"></div>
  </div>
</template>

<script>
export default {
  computed: {
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
  justify-content: flex-start;
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
    &.router-link-active {
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