<template>
  <div class="reset-password">
    <div class="title">
      {{ $locale.settings.resetPassword.title }}
    </div>
    <div class="form">
      <div class="form-item">
        <label for="oldPassword">
          {{ $locale.settings.resetPassword.oldPassword }}
        </label>
        <div class="custom-input">
          <input
            id="oldPassword"
            type="password"
            :placeholder="$locale.settings.resetPassword.oldPasswordPlaceholder"
            v-model="oldPassword"
            @keyup.enter="handleSubmit"
          />
          <div v-if="oldPassword" class="clean-input">
            <i class="fas fa-times" @click="oldPassword = ''"></i>
          </div>
        </div>
      </div>
      <div class="form-item">
        <label for="newPassword">
          {{ $locale.settings.resetPassword.newPassword }}
        </label>
        <div class="custom-input">
          <input
            id="newPassword"
            type="password"
            :placeholder="$locale.settings.resetPassword.newPasswordPlaceholder"
            v-model="newPassword"
            @keyup.enter="handleSubmit"
          />
          <div v-if="newPassword" class="clean-input">
            <i class="fas fa-times" @click="newPassword = ''"></i>
          </div>
        </div>
      </div>
      <div class="form-item">
        <label for="confirmNewPassword">
          {{ $locale.settings.resetPassword.confirmNewPassword }}
        </label>
        <div class="custom-input">
          <input
            id="confirmNewPassword"
            type="password"
            :placeholder="
              $locale.settings.resetPassword.confirmNewPasswordPlaceholder
            "
            v-model="confirmNewPassword"
            @keyup.enter="handleSubmit"
          />
          <div v-if="confirmNewPassword" class="clean-input">
            <i class="fas fa-times" @click="confirmNewPassword = ''"></i>
          </div>
        </div>
      </div>
      <div class="form-item">
        <button class="btn-submit" @click="handleSubmit">
          {{ $locale.settings.resetPassword.submit }}
        </button>
      </div>
    </div>
  </div>
</template>

<script>
export default {
  methods: {
    clearForm() {
      this.oldPassword = "";
      this.newPassword = "";
      this.confirmNewPassword = "";
    },
    handleSubmit() {
      this.$startLoading();
      this.$http
        .post("/api/admin/resetPassword", {
          id: this.$user.id,
          token: this.$user.token,
          oldPassword: this.oldPassword,
          newPassword: this.newPassword,
        })
        .then((res) => {
          if (res.data.success) {
            this.clearForm();
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
      oldPassword: "",
      newPassword: "",
      confirmNewPassword: "",
    };
  },
};
</script>

<style lang="scss" scoped>
.reset-password {
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