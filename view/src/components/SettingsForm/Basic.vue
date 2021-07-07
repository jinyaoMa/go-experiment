<template>
  <div class="basic">
    <div class="title">
      {{ $locale.settings.basic.title }}
    </div>
    <div class="form">
      <div v-if="$user.permission.includes('ADMIN:1')" class="form-item">
        <label for="userLimit">
          {{ $locale.settings.basic.userLimit }}
        </label>
        <div class="custom-input">
          <input
            id="userLimit"
            type="number"
            min="1"
            step="1"
            v-model="options.userLimit"
            @keyup.enter="handleSubmit"
          />
        </div>
      </div>
      <div class="form-item">
        <button class="btn-submit" @click="handleSubmit">
          {{ $locale.settings.basic.submit }}
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
          userLimit: 1,
        };
      },
    },
  },
  methods: {
    handleSubmit() {
      this.$startLoading();
      this.$http
        .post("/api/admin/basic", {
          id: this.$user.id,
          token: this.$user.token,
          userLimit: this.options.userLimit,
        })
        .then((res) => {
          if (res.data.success) {
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
};
</script>

<style lang="scss" scoped>
.basic {
  display: flex;
  flex-direction: column;
  align-items: center;
}
.title {
  font-size: 1.5em;
  font-weight: bold;
  margin-bottom: 1em;
}
.form {
  background-color: #ffffff;
  padding: 0 0.5em;
  max-width: 280px;
}
.form-item {
  display: flex;
  flex-direction: column;
  justify-content: space-between;
  padding: 0.5em 0;
  label {
    display: block;
    font-size: 0.8em;
    user-select: none;
  }
  .custom-input {
    position: relative;
  }
  input {
    border-top-width: 0;
    border-left-width: 0;
    border-right-width: 0;
    border-bottom-width: 1px;
    border-style: solid;
    border-color: #ccccccaa;
    font-family: inherit;
    line-height: 1;
    padding: 0.75em;
    padding-right: 1.5em;
    letter-spacing: 0.1em;
    background-color: transparent;
    outline: none;
    min-width: 280px;
    width: 100%;
    box-sizing: border-box;
    font-size: 0.9em;
    color: #333333;
    transition: 0.2s;
    opacity: 0.9;
    &:focus {
      border-color: #333333;
      font-size: 1.1em;
      opacity: 1;
    }
  }
  .clean-input {
    position: absolute;
    top: 0;
    right: 0.5em;
    height: 100%;
    box-sizing: border-box;
    display: flex;
    flex-direction: column;
    justify-content: center;
    > i {
      cursor: pointer;
    }
  }
  .btn-submit {
    border: none;
    padding: 0.5em;
    font-family: inherit;
    font-size: 0.9em;
    letter-spacing: 0.1em;
    outline: none;
    transition: 0.2s;
    border-radius: 4px;
    user-select: none;
    font-weight: 500;
    color: #ffffff;
    background-color: #333333;
    &:hover {
      background-color: #555555;
    }
  }
}
</style>