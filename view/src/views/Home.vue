<template>
  <Layout>
    <div class="home">
      <TaskBar
        ref="taskBar"
        :currentFilesCount="currentFiles.length"
        :filesToDownload="filesToDownload"
        :cutOptions="cutOptions"
        :canNewFolder="searchKeyword.length === 0"
        @search="handleSearch"
        @newFolder="handleNewFolderClick"
      />
      <div class="home-frame">
        <Breadcrumb
          :checkedCount="checkedItemsAmount"
          :allChecked="checkedItemsAmount === currentFiles.length"
        />
        <table>
          <thead>
            <tr>
              <td class="col-select">
                <CheckSquare
                  :prepared="
                    checkedItemsAmount > 0 &&
                    checkedItemsAmount < currentFiles.length
                  "
                  :checked="checkedItemsAmount === currentFiles.length"
                  @click="handleAllCheck"
                />
              </td>
              <td @click="handleSortName">
                {{ $locale.home.filename }}
                <template v-if="sortBy === 'Name'">
                  <i v-if="sortOrder === 'DESC'" class="fas fa-sort-alpha-up" />
                  <i v-else class="fas fa-sort-alpha-down" />
                </template>
              </td>
              <td class="col-size" @click="handleSortSize">
                {{ $locale.home.size }}
                <template v-if="sortBy === 'Size'">
                  <i
                    v-if="sortOrder === 'DESC'"
                    class="fas fa-sort-numeric-up-alt"
                  />
                  <i v-else class="fas fa-sort-numeric-down" />
                </template>
              </td>
              <td class="col-modified-at" @click="handleSortUpdatedAt">
                {{ $locale.home.modifiedAt }}
                <template v-if="sortBy === 'UpdatedAt'">
                  <i
                    v-if="sortOrder === 'DESC'"
                    class="fas fa-sort-amount-up"
                  />
                  <i v-else class="fas fa-sort-amount-down" />
                </template>
              </td>
            </tr>
          </thead>
          <tbody>
            <tr v-if="parentPath != null || searchKeyword">
              <td class="col-select"></td>
              <td colspan="3">
                <div class="col-filename">
                  <span
                    v-if="searchKeyword"
                    class="col-filename-start back"
                    @click="handleClearOpts"
                  >
                    <i class="fas fa-arrow-left fa-fw" />
                    {{ $locale.home.endSearch }}
                  </span>
                  <router-link
                    v-else
                    class="col-filename-start back"
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
              </td>
            </tr>
            <tr v-if="showNewFolderForm">
              <td class="col-select"></td>
              <td colspan="3">
                <div class="col-filename">
                  <div class="col-filename-start">
                    <div class="form-rename">
                      <input
                        type="text"
                        v-model="newFolderName"
                        @keyup.enter="handleNewFolder"
                      />
                      <i class="fas fa-check fa-fw" @click="handleNewFolder" />
                      <i
                        class="fas fa-times fa-fw"
                        @click="showNewFolderForm = false"
                      />
                    </div>
                  </div>
                </div>
              </td>
            </tr>
            <tr v-for="file in currentFiles" :key="file.ID">
              <td class="col-select">
                <CheckSquare
                  :fileId="file.ID"
                  :checked="file.isChecked"
                  @click="handleSingleCheck"
                />
              </td>
              <td>
                <div class="col-filename">
                  <div
                    v-if="file.Type === 'directory'"
                    class="col-filename-start"
                  >
                    <div v-if="file.canRename" class="form-rename">
                      <input
                        type="text"
                        v-model="renameFileName"
                        @keyup.enter="handleRenameFile(file)"
                      />
                      <i
                        class="fas fa-check fa-fw"
                        @click="handleRenameFile(file)"
                      />
                      <i
                        class="fas fa-times fa-fw"
                        @click="file.canRename = false"
                      />
                    </div>
                    <div v-else>
                      <router-link
                        :to="{
                          query: {
                            currentPath: file.Path,
                          },
                        }"
                      >
                        <i class="fas fa-folder fa-fw" />
                        {{ file.Name }}
                        <span class="hint" v-if="searchKeyword">
                          {{ file.Path }}
                        </span>
                      </router-link>
                    </div>
                  </div>
                  <div v-else class="col-filename-start">
                    <div v-if="file.canRename" class="form-rename">
                      <input
                        type="text"
                        v-model="renameFileName"
                        @keyup.enter="handleRenameFile(file)"
                      />
                      <i
                        class="fas fa-check fa-fw"
                        @click="handleRenameFile(file)"
                      />
                      <i
                        class="fas fa-times fa-fw"
                        @click="file.canRename = false"
                      />
                    </div>
                    <div v-else>
                      <i class="fas fa-file fa-fw" />
                      {{ file.Name }}
                      <span class="hint" v-if="searchKeyword">
                        {{ file.Path }}
                      </span>
                    </div>
                  </div>
                  <div class="col-filename-end">
                    <BtnDownloadFile
                      v-if="file.Type === 'file'"
                      :url="getFileDownloadLink(file)"
                    />
                    <BtnSharing
                      v-if="file.Type === 'file'"
                      @click="handleSharingClick(file)"
                    />
                    <BtnCut
                      v-if="file.Type === 'file'"
                      @click="handleCutClick(file)"
                    />
                    <BtnRename @click="handleRenameClick(file)" />
                    <BtnRecycling @click="handleRecyclingClick(file)" />
                  </div>
                </div>
              </td>
              <td v-if="file.Type === 'directory'">-</td>
              <td v-else>{{ $convertSpace2String(file.Size) }}</td>
              <td>{{ $formatDate(file.UpdatedAt, "YYYY-MM-DD HH:mm") }}</td>
            </tr>
          </tbody>
        </table>
      </div>
    </div>
  </Layout>
</template>

<script>
import Layout from "../components/Layout.vue";
import TaskBar from "../components/TaskBar.vue";
import CheckSquare from "../components/CheckSquare.vue";
import Breadcrumb from "../components/Breadcrumb.vue";
import BtnDownloadFile from "../components/BtnDownloadFile.vue";
import BtnRename from "../components/BtnRename.vue";
import BtnCut from "../components/BtnCut.vue";
import BtnSharing from "../components/BtnSharing.vue";
import BtnRecycling from "../components/BtnRecycling.vue";

export default {
  components: {
    Layout,
    TaskBar,
    CheckSquare,
    Breadcrumb,
    BtnDownloadFile,
    BtnRename,
    BtnCut,
    BtnSharing,
    BtnRecycling,
  },
  watch: {
    $route: {
      handler() {
        this.handleClearOpts();
      },
    },
  },
  methods: {
    handleRecyclingClick(file) {
      this.$startLoading();
      this.$http
        .post("/api/files/recycle", {
          id: this.$user.id,
          token: this.$user.token,
          fileId: file.ID,
        })
        .then((res) => {
          if (res.data.success) {
            this.$setFiles(res.data.data.files);
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
    handleNewFolder() {
      this.$startLoading();
      this.$http
        .post("/api/files/newFolder", {
          id: this.$user.id,
          token: this.$user.token,
          folderName: this.newFolderName,
          desId: this.$getCurrentPathId(),
        })
        .then((res) => {
          if (res.data.success) {
            this.$setFiles(res.data.data.files);
            this.showNewFolderForm = false;
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
    handleNewFolderClick() {
      this.newFolderName = this.$locale.common.newFolder;
      this.showNewFolderForm = true;
    },
    handleSharingClick(file) {
      this.$startLoading();
      this.$http
        .post("/api/files/shareFile", {
          id: this.$user.id,
          token: this.$user.token,
          fileId: file.ID,
        })
        .then((res) => {
          if (res.data.success) {
            let data = res.data.data;
            let shareLink = `${this.$http.defaults.baseURL}/api/service/download?c=${data.c}&d=${data.d}&e=${data.e}`;
            this.$setFiles(data.files);
            this.$showShareNotice(shareLink);
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
    handleCutClick(file) {
      this.cutOptions.srcId = file.ID;
      this.cutOptions.srcPath = file.Path;
      this.cutOptions.canPaste = true;
    },
    handleRenameFile(file) {
      this.$startLoading();
      this.$http
        .post("/api/files/renameFile", {
          id: this.$user.id,
          token: this.$user.token,
          fileId: this.renameFileId,
          filename: this.renameFileName,
        })
        .then((res) => {
          if (res.data.success) {
            file.Name = this.renameFileName;
            file.canRename = false;
            this.renameFileId = 0;
            this.renameFileName = "";
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
    handleRenameClick(file) {
      this.currentFiles.forEach((f) => {
        f.canRename = false;
      });
      file.canRename = true;
      this.renameFileId = file.ID;
      this.renameFileName = file.Name;
    },
    getFileDownloadLink(file) {
      let i = this.$user.usedSpace % 64;
      return `${this.$http.defaults.baseURL}/api/service/download?a=${
        i + 1
      }&b=${this.$user.token.substr(i)}&c=${file.ShareCode}&d=${file.ID}`;
    },
    handleAllCheck(_, checked) {
      if (!checked) {
        this.$clearChecklist();
      } else {
        this.currentFiles.forEach((file) => {
          file.isChecked = true;
        });
      }
    },
    handleSingleCheck(fileId, checked) {
      this.$setFileCheckById({ fileId, isChecked: checked });
    },
    handleSortName() {
      if (this.sortBy === "Name") {
        this.sortOrder = this.sortOrder === "ASC" ? "DESC" : "ASC";
      } else {
        this.sortBy = "Name";
        this.sortOrder = "ASC";
      }
    },
    handleSortSize() {
      if (this.sortBy === "Size") {
        this.sortOrder = this.sortOrder === "DESC" ? "ASC" : "DESC";
      } else {
        this.sortBy = "Size";
        this.sortOrder = "DESC";
      }
    },
    handleSortUpdatedAt() {
      if (this.sortBy === "UpdatedAt") {
        this.sortOrder = this.sortOrder === "DESC" ? "ASC" : "DESC";
      } else {
        this.sortBy = "UpdatedAt";
        this.sortOrder = "DESC";
      }
    },
    handleClearOpts() {
      this.searchKeyword = "";
      this.$refs.taskBar.clearSearchKeyword();
      this.$clearChecklist();
    },
    handleSearch(keyword, complete) {
      this.searchKeyword = keyword;
      complete();
    },
  },
  computed: {
    filesToDownload() {
      return this.currentUnsortedFiles.filter(
        (file) => file.Type === "file" && file.isChecked
      );
    },
    checkedItemsAmount() {
      return this.currentUnsortedFiles.filter((item) => item.isChecked).length;
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
      if (this.sortOrder === "DESC") {
        return this.currentUnsortedFiles.sort((a, b) => {
          if (a[this.sortBy] > b[this.sortBy]) {
            return -1;
          } else if (a[this.sortBy] < b[this.sortBy]) {
            return 1;
          }
          return 0;
        });
      }
      return this.currentUnsortedFiles.sort((a, b) => {
        if (a[this.sortBy] > b[this.sortBy]) {
          return 1;
        } else if (a[this.sortBy] < b[this.sortBy]) {
          return -1;
        }
        return 0;
      });
    },
    currentUnsortedFiles() {
      let currentPath = this.$route.query.currentPath;
      if (typeof currentPath === "string" && currentPath.length > 0) {
        let matches = currentPath.match(/\\/g);
        let depth = matches ? matches.length : 0;
        // search current folder
        if (this.searchKeyword.length > 0) {
          return this.$files.filter((file) => {
            let find = file.Path.match(/\\/g);
            let temp = find ? find.length : 0;
            if (
              file.Path.startsWith(currentPath) &&
              file.Path.replace(currentPath, "").includes(this.searchKeyword) &&
              temp > depth
            ) {
              return true;
            }
            return false;
          });
        }
        // current folder
        return this.$files.filter((file) => {
          let find = file.Path.match(/\\/g);
          let temp = find ? find.length : 0;
          if (file.Path.startsWith(currentPath) && temp === depth + 1) {
            return true;
          }
          return false;
        });
      }

      // search root folder
      if (this.searchKeyword.length > 0) {
        return this.$files.filter((file) => {
          if (file.Path === ".") {
            return false;
          }
          if (file.Path.includes(this.searchKeyword)) {
            return true;
          }
          return false;
        });
      }
      // root
      return this.$files.filter((file) => {
        if (file.Path === "." || file.Path.includes("\\")) {
          return false;
        }
        return true;
      });
    },
  },
  data() {
    return {
      isSelectAll: false,
      searchKeyword: "",
      sortBy: "UpdatedAt",
      sortOrder: "DESC",
      renameFileId: 0,
      renameFileName: "",
      newFolderName: "",
      cutOptions: {
        canPaste: false,
        srcId: 0,
        srcPath: "",
      },
      showNewFolderForm: false,
    };
  },
  mounted() {
    this.$startLoading();
    this.$http
      .post("/api/files", {
        id: this.$user.id,
        token: this.$user.token,
      })
      .then((res) => {
        if (res.data.success) {
          this.$setFiles(res.data.data.files);
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
.home {
  display: flex;
  flex-direction: column;
  max-height: 100%;
  height: 100%;
}
.home-frame {
  flex-grow: 1;
  box-sizing: border-box;
}
table {
  border-collapse: collapse;
  width: 100%;
  table-layout: fixed;
  margin-bottom: 1em;
  thead {
    font-weight: bold;
    td {
      cursor: pointer;
      transition: 0.2s;
      &:not(.col-select) {
        &:hover {
          background-color: #2196f309;
        }
        i {
          color: #2196f3;
          transform-origin: left center;
          transform: scale(1.12);
        }
      }
    }
  }
  td {
    font-size: 0.9em;
    line-height: 3.6;
    white-space: nowrap;
    overflow: hidden;
    text-overflow: ellipsis;
    padding: 0 1em;
  }
  tr {
    transition: 0.2s;
    border-bottom: 1px solid #f1f2f3;
    &:hover {
      background-color: #fafbfc;
    }
  }
  .col-select {
    width: 1em;
  }
  .col-filename {
    display: flex;
    flex-direction: row;
    justify-content: space-between;
  }
  .col-filename-start {
    flex-grow: 1;
    cursor: default;
    max-width: 100%;
    overflow: hidden;
    text-overflow: ellipsis;
    > i {
      margin-right: 0.5em;
      transform: scale(1.3);
      transform-origin: center center;
    }
    .hint {
      color: #aaabac;
      font-size: 0.8em;
      margin-left: 0.5em;
    }
    a,
    span {
      color: #333333;
      text-decoration: none;
      &.back {
        color: #aaabac;
        user-select: none;
      }
      &:hover {
        color: #2196f3;
      }
    }
  }
  a,
  span {
    &.col-filename-start {
      color: #333333;
      text-decoration: none;
      &.back {
        color: #aaabac;
        user-select: none;
      }
    }
  }
  a.col-filename-start:hover {
    color: #2196f3;
  }
  span.col-filename-start:hover {
    color: #882222;
  }
  .col-filename-end {
    margin-right: 1em;
    > * {
      margin-left: 1em;
    }
  }
  .col-size {
    width: 100px;
  }
  .col-modified-at {
    width: 160px;
  }
  input[type="checkbox"] {
    margin: 0;
  }
}
.form-rename {
  > input {
    margin-right: 0.5em;
  }
  > i {
    color: #2196f3;
  }
  .fa-check:hover {
    color: #4caf50;
  }
  .fa-times:hover {
    color: #dd3333;
  }
}
</style>