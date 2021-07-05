<template>
  <Layout>
    <div class="home">
      <TaskBar />
      <div class="home-frame">
        <table>
          <thead>
            <tr>
              <td class="col-select">
                <CheckSquare />
              </td>
              <td>{{ $locale.home.filename }}</td>
              <td class="col-size">{{ $locale.home.size }}</td>
              <td class="col-modified-at">{{ $locale.home.modifiedAt }}</td>
            </tr>
          </thead>
          <tbody>
            <tr v-for="file in $files" :key="file.ID">
              <td class="col-select">
                <CheckSquare :checked="file.isSelected" />
              </td>
              <td>
                <div class="col-filename">
                  <div class="col-filename-start">
                    <i
                      class="fas fa-fw"
                      :class="{
                        'fa-folder': file.Type === 'directory',
                        'fa-file': file.Type === 'file',
                      }"
                    />
                    {{ file.Path }}
                  </div>
                  <div class="col-filename-end"></div>
                </div>
              </td>
              <td>{{ $convertSpace2String(file.Size) }}</td>
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

export default {
  components: {
    Layout,
    TaskBar,
    CheckSquare,
  },
  data() {
    return {
      isSelectAll: false,
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
  margin-top: 1em;
  td {
    font-size: 0.9em;
    line-height: 3.2;
    white-space: nowrap;
    overflow: hidden;
    text-overflow: ellipsis;
  }
  tr {
    transition: 0.2s;
    &:hover {
      background-color: #fafbfc;
    }
  }
  .col-select {
    width: 2em;
    padding-left: 1em;
  }
  .col-filename {
    display: flex;
    flex-direction: row;
    justify-content: space-between;
  }
  .col-filename-start {
    flex-grow: 1;
    > i {
      margin-right: 0.5em;
      transform: scale(1.3);
      transform-origin: center center;
    }
  }
  .col-filename-end {
    margin-right: 1em;
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
</style>