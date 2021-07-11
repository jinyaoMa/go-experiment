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
      <div class="tab-item" @click="handleTab">
        <i class="fas fa-images" />
      </div>
      <div class="tab-item active">
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
      <div class="list-item" v-for="(file, i) in currentFiles" :key="file.ID">
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
        <div v-else class="list-item-text" @click="handlePlay(file, i)">
          <i class="fas fa-file fa-fw" />
          {{ file.Name }}
        </div>
      </div>
    </div>
    <div v-if="canPlay" class="video-section">
      <i class="fas fa-times video-close" @click="handleClose" />
      <video-player
        ref="videoPlayer"
        :options="playerOptions"
        @ended="onPlayerEnded($event)"
      >
      </video-player>
    </div>
  </div>
</template>

<script>
export default {
  methods: {
    onPlayerEnded(event) {
      let next = this.currentSource + 1;
      if (next >= this.currentFiles.length) {
        this.currentSource = 0;
      } else {
        this.currentSource = next;
      }
    },
    handleClose() {
      this.canPlay = false;
    },
    handlePlay(file, index) {
      this.currentSource = index;
      this.canPlay = true;
    },
    getVideoLink(file) {
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
      this.$router.replace("minip");
    },
    getMime(ext) {
      switch (ext) {
        case ".mp4":
          return "video/mp4";
        case ".webm":
          return "video/webm";
        case ".flv":
          return "flv";
        case ".mkv":
          return "video/webm";
        case ".mov":
          return "video/quicktime";
        case ".swf":
          return "application/x-shockwave-flash";
        default:
          return "";
      }
    },
  },
  computed: {
    playerOptions() {
      if (
        this.currentSource < 0 ||
        this.currentSource >= this.currentFiles.length
      ) {
        return {};
      }
      let targetFile = this.currentFiles[this.currentSource];
      let link = this.getVideoLink(targetFile);
      let mime = this.getMime(targetFile.Extension.toLowerCase());
      return {
        language: this.$locale.LANG,
        sources: [
          {
            type: mime,
            src: link,
          },
        ],
      };
    },
    player() {
      return this.$refs.videoPlayer.player;
    },
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
            return /^\.(wmv|mp4|3gp|mov|m4v|avi|mkv|flv)$/i.test(
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
        return /^\.(wmv|mp4|3gp|mov|m4v|avi|mkv|flv)$/i.test(file.Extension);
      });
    },
  },
  data() {
    return {
      secret: "",
      canPlay: false,
      currentSource: -1,
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
  position: relative;
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
  padding: 10px;
  &:active {
    background-color: #fafbfc;
  }
  i {
    margin: 0 10px;
  }
}
.list-item-link {
  display: block;
  text-decoration: none;
  color: #2196f3;
}
.list-item-text {
  > img {
    display: block;
    margin-right: 0.5em;
    width: 100%;
    height: 6em;
    object-fit: contain;
    object-position: center;
  }
}
.video-section {
  position: absolute;
  top: 0;
  left: 0;
  width: 100%;
  min-height: 100vh;
  background-color: #000000;
  display: flex;
  flex-direction: column;
}
.video-close {
  color: #ffffff;
  text-shadow: 0 0 2px #000000;
  font-size: 2em;
  line-height: 3;
  height: 3em;
  width: 100%;
  box-sizing: border-box;
  text-align: right;
  padding-right: 1em;
}
</style>