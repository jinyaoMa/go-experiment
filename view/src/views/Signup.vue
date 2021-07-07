<template>
  <div class="signup">
    <Copyright />
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
            @keyup.enter="handleSignup"
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
            @keyup.enter="handleSignup"
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
            @keyup.enter="handleSignup"
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
            @keyup.enter="handleSignup"
          />
          <div v-if="confirmPassword" class="clean-input">
            <i class="fas fa-times" @click="handleCleanConfirmPassword"></i>
          </div>
        </div>
      </div>
      <div class="form-item">
        <button class="btn-submit" @click="handleSignup">
          {{ $locale.signup.submit }}
        </button>
      </div>
      <div class="form-item back-option">
        {{ $locale.signup.back }}
        <router-link to="/login">{{ $locale.signup.login }}</router-link>
      </div>
    </div>
    <div class="float-panel">
      <BtnLangSwap />
    </div>
  </div>
</template>

<script>
import Copyright from "../components/Copyright.vue";
import BtnLangSwap from "../components/BtnLangSwap.vue";

export default {
  components: {
    Copyright,
    BtnLangSwap,
  },
  data() {
    return {
      username: "",
      account: "",
      password: "",
      confirmPassword: "",
    };
  },
  methods: {
    handleSignup() {
      if (this.password != this.confirmPassword) {
        this.$showError(this.$locale.common.errorConfirmPassword);
        return;
      }
      this.$startLoading();
      this.$http
        .post("/api/signup", {
          username: this.username,
          account: this.account,
          password: this.password,
        })
        .then((res) => {
          if (res.data.success) {
            let data = res.data.data;
            this.$setUser({
              id: data.userid,
              name: data.username,
              role: data.role,
              permission: data.permission,
              usedSpace: data.usedSpace,
              allSpace: data.allSpace,
              token: data.token,
            });
            this.$router.push("/");
            this.$stopLoading();
          } else {
            this.$showError(this.$locale.common.errorMsg);
            this.$stopLoading();
          }
        })
        .catch((err) => {
          this.$showError(this.$locale.common.errorServer);
          this.$stopLoading();
        });
    },
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
    user-select: none;
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
  font-size: 0.9em;
  display: block;
  text-align: center;
  > a {
    font-weight: bold;
    color: #2196f3;
  }
}
.float-panel {
  position: absolute;
  bottom: 1em;
  right: 1em;
}
</style>