<template>
  <div class="task-bar">
    <div class="start">
      <button class="btn-upload">{{ $locale.common.upload }}</button>
      <button v-if="canNewFolder" class="btn-new-folder">
        {{ $locale.common.newFolder }}
      </button>
      <button v-if="filesToDownload.length > 0" class="btn-new-folder">
        {{ $locale.common.download }}
      </button>
      <button
        v-if="cutOptions.canPaste"
        class="btn-new-folder"
        @click="handlePaste"
      >
        {{ $locale.common.paste }}
      </button>
      <span class="src-path-hint">{{ cutOptions.srcPath }}</span>
    </div>
    <div class="end">
      <div class="loaded">
        <template v-if="searchComplete">
          {{ searchCompleteText }}
        </template>
        <template v-else>
          {{ loadedText }}
        </template>
      </div>
      <div class="search-input">
        <input
          type="text"
          :placeholder="$locale.common.searchPlaceholder"
          v-model="keyword"
        />
        <button
          class="btn-search"
          @click="$emit('search', keyword, searchCompleteCallback)"
        >
          <i class="fas fa-search" />
        </button>
      </div>
    </div>
  </div>
</template>

<script>
export default {
  props: {
    currentFilesCount: {
      type: Number,
      default() {
        return 0;
      },
    },
    filesToDownload: {
      type: Array,
      default() {
        return [];
      },
    },
    cutOptions: {
      type: Object,
      default() {
        return {
          canPaste: false,
          srcId: 0,
          srcPath: "",
        };
      },
    },
    canNewFolder: {
      type: Boolean,
      default() {
        return false;
      },
    },
  },
  computed: {
    searchCompleteText() {
      if (this.currentFilesCount === 1) {
        return this.$locale.common.searchCompleted.one.replace(
          "?",
          this.currentFilesCount
        );
      }
      return this.$locale.common.searchCompleted.more.replace(
        "?",
        this.currentFilesCount
      );
    },
    loadedText() {
      if (this.currentFilesCount === 1) {
        return this.$locale.common.loaded.one.replace(
          "?",
          this.currentFilesCount
        );
      }
      return this.$locale.common.loaded.more.replace(
        "?",
        this.currentFilesCount
      );
    },
  },
  methods: {
    handlePaste() {
      this.$startLoading();
      this.$http
        .post("/api/files/moveFile", {
          id: this.$user.id,
          token: this.$user.token,
          srcId: this.cutOptions.srcId,
          desId: this.$getCurrentPathId(),
        })
        .then((res) => {
          if (res.data.success) {
            this.$setFiles(res.data.data.files);
            this.clearCutOptions();
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
    clearCutOptions() {
      this.cutOptions.canPaste = false;
      this.cutOptions.srcId = 0;
      this.cutOptions.srcPath = "";
    },
    searchCompleteCallback() {
      this.searchComplete = true;
    },
    clearSearchKeyword() {
      this.keyword = "";
      this.searchComplete = false;
    },
  },
  data() {
    return {
      keyword: "",
      searchComplete: false,
    };
  },
};
</script>

<style lang="scss" scoped>
.task-bar {
  display: flex;
  flex-direction: row;
  align-items: center;
  justify-content: space-between;
  box-sizing: border-box;
  font-size: 0.9em;
  height: 3.33em;
}
[class^="btn-"] {
  border: none;
  padding: 0 1em;
  margin-right: 1em;
  line-height: 2.33;
  font-family: inherit;
  outline: none;
  transition: 0.2s;
  border-radius: 4px;
  user-select: none;
}
.start,
.end {
  display: flex;
  flex-direction: row;
  align-items: center;
}
.btn-upload {
  font-weight: 500;
  color: #ffffff;
  background-color: #333333;
  &:hover {
    background-color: #555555;
  }
}
.btn-new-folder {
  font-weight: 500;
  background-color: #f1f2f3;
  &:hover {
    background-color: #eaebec;
  }
}
.btn-search {
  margin-right: 0;
  color: #ffffff;
  background-color: #333333;
  &:hover {
    background-color: #555555;
  }
  border-radius: 0 4px 4px 0;
}
.search-input {
  position: relative;
  display: flex;
  flex-direction: row;
  input {
    box-sizing: border-box;
    border-top-width: 2px;
    border-left-width: 2px;
    border-right-width: 0;
    border-bottom-width: 2px;
    border-style: solid;
    border-color: #333333;
    border-radius: 4px 0 0 4px;
    font-family: inherit;
    height: 2.33em;
    padding-left: 0.66em;
    padding-right: 2.33em;
    letter-spacing: 0.1em;
    background-color: transparent;
    outline: none;
    min-width: 240px;
    font-size: 0.9em;
    color: #333333;
    transition: 0.2s;
  }
}
.loaded {
  margin-right: 1em;
  margin-left: 1em;
  font-size: 0.9em;
}
.src-path-hint {
  font-size: 0.8em;
  color: #989796;
  align-self: flex-end;
  margin-bottom: 0.6em;
  letter-spacing: 0;
}
</style>