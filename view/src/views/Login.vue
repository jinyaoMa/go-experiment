<template>
  <div class="login">
    <Copyright />
    <div class="form">
      <div class="form-item">
        <label for="account">{{ $locale.login.account }}</label>
        <div class="custom-input">
          <input
            id="account"
            type="text"
            :placeholder="$locale.login.accountPlaceholder"
            v-model="account"
            @keyup.enter="handleLogin"
          />
          <div v-if="account" class="clean-input">
            <i class="fas fa-times" @click="handleCleanAccount"></i>
          </div>
        </div>
      </div>
      <div class="form-item">
        <label for="password">{{ $locale.login.password }}</label>
        <div class="custom-input">
          <input
            id="password"
            type="password"
            :placeholder="$locale.login.passwordPlaceholder"
            v-model="password"
            @keyup.enter="handleLogin"
          />
          <div v-if="password" class="clean-input">
            <i class="fas fa-times" @click="handleCleanPassword"></i>
          </div>
        </div>
      </div>
      <div class="form-item">
        <button class="btn-login" @click="handleLogin">
          {{ $locale.login.login }}
        </button>
      </div>
      <div class="form-item">
        <router-link class="btn-signup" to="/signup">
          {{ $locale.login.signup }}
        </router-link>
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
      account: "",
      password: "",
    };
  },
  methods: {
    handleLogin() {
      this.$startLoading();
      this.$http
        .post("/api/login", {
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
              usedSpace: data.usedSpace,
              allSpace: data.allSpace,
              token: data.token,
            });
            this.$router.push("/");
            this.$stopLoading();
          } else {
            this.$stopLoading();
          }
        })
        .catch((err) => {
          this.$stopLoading();
        });
    },
    handleCleanAccount() {
      this.account = "";
    },
    handleCleanPassword() {
      this.password = "";
    },
  },
};
</script>

<style lang="scss" scoped>
.login {
  position: relative;
  min-height: 100vh;
  display: flex;
  justify-content: center;
  align-items: center;
  letter-spacing: 0.1em;
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
  .btn-login,
  .btn-signup {
    border: none;
    padding: 0.5em;
    font-family: inherit;
    font-size: 0.9em;
    letter-spacing: 0.1em;
    outline: none;
    transition: 0.2s;
    border-radius: 4px;
    user-select: none;
  }
  .btn-login {
    font-weight: 500;
    color: #ffffff;
    background-color: #333333;
    &:hover {
      background-color: #555555;
    }
  }
  .btn-signup {
    cursor: default;
    text-align: center;
    text-decoration: none;
    color: #2196f3;
    background-color: #f1f2f3;
    &:hover {
      background-color: #eaebec;
    }
  }
}
.float-panel {
  position: absolute;
  bottom: 1em;
  right: 1em;
}
</style>