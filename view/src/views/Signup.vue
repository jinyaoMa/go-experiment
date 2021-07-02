<template>
  <div class="signup">
    <div class="title">{{ $locale.signup.title }}</div>
    <div class="form">
      <div class="form-item">
        <label for="username">{{ $locale.signup.username }}</label>
        <div class="custom-input">
          <input
            id="username"
            type="text"
            :placeholder="$locale.signup.usernamePlaceholder"
            v-model="username"
          />
          <div v-if="username" class="clean-input">
            <i class="fas fa-times" @click="handleCleanUsername"></i>
          </div>
        </div>
      </div>
      <div class="form-item">
        <label for="account">{{ $locale.signup.account }}</label>
        <div class="custom-input">
          <input
            id="account"
            type="text"
            :placeholder="$locale.signup.accountPlaceholder"
            v-model="account"
          />
          <div v-if="account" class="clean-input">
            <i class="fas fa-times" @click="handleCleanAccount"></i>
          </div>
        </div>
      </div>
      <div class="form-item">
        <label for="password">{{ $locale.signup.password }}</label>
        <div class="custom-input">
          <input
            id="password"
            type="password"
            :placeholder="$locale.signup.passwordPlaceholder"
            v-model="password"
          />
          <div v-if="password" class="clean-input">
            <i class="fas fa-times" @click="handleCleanPassword"></i>
          </div>
        </div>
      </div>
      <div class="form-item">
        <label for="confirmPassword">{{
          $locale.signup.confirmPassword
        }}</label>
        <div class="custom-input">
          <input
            id="confirmPassword"
            type="password"
            :placeholder="$locale.signup.confirmPasswordPlaceholder"
            v-model="confirmPassword"
          />
          <div v-if="confirmPassword" class="clean-input">
            <i class="fas fa-times" @click="handleCleanConfirmPassword"></i>
          </div>
        </div>
      </div>
      <div class="form-item">
        <button class="btn-submit">{{ $locale.signup.submit }}</button>
      </div>
      <div class="form-item back-option">
        {{ $locale.signup.back }}
        <router-link to="/login">{{ $locale.signup.login }}</router-link>
      </div>
    </div>
  </div>
</template>

<script>
export default {
  data() {
    return {
      username: "",
      account: "",
      password: "",
      confirmPassword: "",
    };
  },
  methods: {
    handleCleanUsername() {
      this.username = "";
    },
    handleCleanAccount() {
      this.account = "";
    },
    handleCleanPassword() {
      this.password = "";
    },
    handleCleanConfirmPassword() {
      this.confirmPassword = "";
    },
  },
};
</script>

<style lang="scss" scoped>
.signup {
  position: relative;
  min-height: 560px;
  height: 100vh;
  display: flex;
  flex-direction: column;
  justify-content: center;
  align-items: center;
  letter-spacing: 0.1em;
}
.title {
  font-size: 1.5em;
  font-weight: bold;
  margin-bottom: 0.25em;
}
.form {
  background-color: #ffffff;
  padding: 0 0.5em;
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
  button {
    border: none;
    padding: 0.5em;
    font-family: inherit;
    font-size: 0.9em;
    letter-spacing: 0.1em;
    outline: none;
    transition: 0.2s;
    font-weight: 500;
    border-radius: 4px;
  }
  .btn-submit {
    color: #ffffff;
    background-color: #333333;
    &:hover {
      background-color: #555555;
    }
  }
}
.back-option {
  font-size: 0.8em;
  display: block;
  text-align: center;
}
</style>