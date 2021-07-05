<template>
  <Layout>
    <div class="home">
      <TaskBar />
      <div class="home-frame">
        <table>
          <thead>
            <tr>
              <td class="col-select">
                <CheckBox />
                <input type="checkbox" v-model="isSelectAll" />
              </td>
              <td>{{ $locale.home.filename }}</td>
              <td class="col-size">{{ $locale.home.size }}</td>
              <td class="col-modified-at">{{ $locale.home.modifiedAt }}</td>
            </tr>
          </thead>
          <tbody>
            <tr v-for="file in $files" :key="file.ID">
              <td>
                <input type="checkbox" v-model="file.isSelected" />
              </td>
              <td>{{ file.Path }}</td>
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
import CheckBox from "../components/CheckBox.vue";

export default {
  components: {
    Layout,
    TaskBar,
    CheckBox,
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
  td {
    font-size: 0.9em;
    line-height: 3.2;
  }
  .col-select {
    width: 2em;
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