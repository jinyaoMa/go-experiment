<template>
  <div class="mini">
    <input
      class="secret"
      type="password"
      autocomplete="off"
      v-model="secret"
      placeholder="© 2021 jinyaoMa"
      @keypress.enter="handleLoad"
    />
    <div class="tab">
      <div class="tab-item active">
        <i class="fas fa-images" />
      </div>
      <div class="tab-item" @click="handleTab">
        <i class="fas fa-video" />
      </div>
    </div>
    <div class="list">
      <div v-if="parentPath != null" class="list-item">
        <router-link
          class="list-item-link"
          :to="{
            query: {
              currentPath: parentPath,
            },
          }"
        >
          <i class="fas fa-reply fa-fw" />
          {{ $locale.home.back2Parent }}
        </router-link>
      </div>
      <div class="list-item" v-for="file in currentFiles" :key="file.ID">
        <router-link
          v-if="file.Type === 'directory'"
          class="list-item-link"
          :to="{
            query: {
              currentPath: file.Path,
            },
          }"
        >
          <i class="fas fa-folder fa-fw" />
          {{ file.Name }}
        </router-link>
        <div v-else class="list-item-text">
          <img
            v-if="isImage(file)"
            :src="getImgLink(file)"
            preview="0"
            @click="$previewRefresh()"
          />
          <div v-else>
            <i class="fas fa-file fa-fw" />
            {{ file.Name }}
          </div>
        </div>
      </div>
    </div>
  </div>
</template>

<script>
export default {
  methods: {
    getImgLink(file) {
      return `${this.$http.defaults.baseURL}/api/secret/${file.Apath}?secret=${this.secret}`;
    },
    isImage(file) {
      return /^\.(jpg|jpeg|git|png|bmp|webp|tif|svg)$/i.test(file.Extension);
    },
    handleLoad() {
      this.$startLoading();
      this.$http
        .get(`/api/secret?secret=${this.secret}`)
        .then((res) => {
          if (res.data.success) {
            window.localStorage.setItem(
              "sEcReTGo-EXpiRemENT55699",
              this.secret
            );
            this.$setFiles(res.data.data.files);
          }
          this.$stopLoading();
        })
        .catch((err) => {
          this.$stopLoading();
        });
    },
    handleTab() {
      this.$router.replace("miniv");
    },
  },
  computed: {
    parentPath() {
      let currentPath = this.$route.query.currentPath;
      if (typeof currentPath === "string" && currentPath.length > 0) {
        let lastSlashIndex = currentPath.lastIndexOf("\\");
        if (lastSlashIndex !== -1) {
          return currentPath.substr(0, lastSlashIndex); // parent
        }
        return ""; // parent => root
      }
      return null; // already root
    },
    currentFiles() {
      return [
        ...this.currentSortedFiles.filter((file) => file.Type === "directory"),
        ...this.currentSortedFiles.filter((file) => file.Type === "file"),
      ];
    },
    currentImageCount() {
      return this.currentSortedFiles.filter((f) => f.Type === "file").length;
    },
    currentSortedFiles() {
      let currentPath = this.$route.query.currentPath;
      if (typeof currentPath === "string" && currentPath.length > 0) {
        let matches = currentPath.match(/\\/g);
        let depth = matches ? matches.length : 0;
        // current folder
        return this.$files.filter((file) => {
          let find = file.Path.match(/\\/g);
          let temp = find ? find.length : 0;
          if (file.Path.startsWith(currentPath) && temp === depth + 1) {
            if (file.Type === "directory") {
              return true;
            }
            return /^\.(jpg|jpeg|git|png|bmp|webp|tif|svg)$/i.test(
              file.Extension
            );
          }
          return false;
        });
      }

      // root
      return this.$files.filter((file) => {
        if (file.Path === "." || file.Path.includes("\\")) {
          return false;
        }
        if (file.Type === "directory") {
          return true;
        }
        return /^\.(jpg|jpeg|gif|png|bmp|webp|tif|svg)$/i.test(file.Extension);
      });
    },
  },
  data() {
    return {
      secret: "",
    };
  },
  mounted() {
    this.secret = window.localStorage.getItem("sEcReTGo-EXpiRemENT55699") || "";
    if (this.secret != "") {
      this.handleLoad();
    }
  },
};
</script>

<style lang="scss" scoped>
.mini {
  display: flex;
  flex-direction: column;
  box-sizing: border-box;
  width: 100vw;
  font-size: 14px;
}
.secret {
  border: 0;
  font-size: 1.2em;
  padding: 1em;
  line-height: 1;
  outline: none;
}
.tab {
  display: flex;
  flex-direction: row;
  background-color: #f1f2f3;
}
.tab-item {
  flex-grow: 1;
  box-sizing: border-box;
  padding: 1em;
  text-align: center;
  transition: 0.2s;
  &.active {
    background-color: #333333;
    color: #ffffff;
  }
}
.list {
  padding: 0.5em 0;
  display: flex;
  flex-direction: column;
}
.list-item {
  box-sizing: border-box;
  line-height: 2;
}
.list-item-link,
.list-item-text {
  white-space: nowrap;
  overflow: hidden;
  text-overflow: ellipsis;
  &:active {
    background-color: #fafbfc;
  }
  > i {
    margin: 0 10px;
  }
}
.list-item-link {
  display: block;
  text-decoration: none;
  color: #2196f3;
  padding: 10px;
}
.list-item-text {
  > img {
    display: block;
    margin-right: 0.5em;
    width: 100%;
    object-fit: contain;
    object-position: center;
  }
}
</style>