<template>
  <div class="head">
    <div class="options">
      <BtnLogout />
      <BtnHome v-if="isSettings" />
      <BtnSettings v-else />
      <BtnLangSwap />
    </div>
    <div class="user">
      <div
        class="user-name"
        :class="{
          empty: $user.name.length === 0,
        }"
      >
        {{ $user.name }}
      </div>
      <div
        class="user-role"
        :class="{
          empty: $user.role.length === 0,
        }"
      >
        {{ $user.role }}
      </div>
      <div class="user-space" :title="userSpaceTitle">
        <div
          class="used"
          :style="{
            width: usedSpaceWidth,
          }"
        >
          <span> {{ usedSpace }} / {{ allSpace }} </span>
        </div>
      </div>
    </div>
  </div>
</template>

<script>
import BtnLogout from "../components/BtnLogout.vue";
import BtnHome from "../components/BtnHome.vue";
import BtnSettings from "../components/BtnSettings.vue";
import BtnLangSwap from "../components/BtnLangSwap.vue";

export default {
  props: {
    isSettings: {
      type: Boolean,
      default() {
        return false;
      },
    },
  },
  components: {
    BtnLogout,
    BtnHome,
    BtnSettings,
    BtnLangSwap,
  },
  computed: {
    userSpaceTitle() {
      return `${this.$locale.common.usedSpace}: ${this.usedSpace}\n${this.$locale.common.allSpace}: ${this.allSpace}`;
    },
    usedSpaceWidth() {
      return `${Math.round(this.$user.usedSpace / this.$user.allSpace)}%`;
    },
    usedSpace() {
      return this.$convertSpace2String(this.$user.usedSpace);
    },
    allSpace() {
      return this.$convertSpace2String(this.$user.allSpace);
    },
  },
};
</script>

<style lang="scss" scoped>
.head {
  display: flex;
  flex-direction: row-reverse;
  justify-content: space-between;
  padding: 1em;
}
.options {
  display: flex;
  flex-direction: row-reverse;
}
.user {
  display: flex;
  flex-direction: row;
  align-items: center;
  > div {
    height: 1.3em;
  }
}
.user-name {
  font-size: 1.5em;
  font-weight: bold;
  margin-left: 0.6em;
  &.empty {
    width: 6.66em;
    background-color: #f1f2f3;
  }
}
.user-role {
  margin-left: 0.6em;
  background-color: #f1f2f3;
  padding: 0.3em 0.6em;
  border-radius: 4px;
  cursor: default;
  user-select: none;
  &.empty {
    width: 2.33em;
  }
}
.user-space {
  margin-left: 1.9em;
  width: 280px;
  border: 2px solid;
  box-sizing: border-box;
  border-radius: 4px;
  overflow: hidden;
  background-color: #f1f2f3;
  user-select: none;
  .used {
    box-sizing: border-box;
    padding: 0 0.4em;
    height: 100%;
    line-height: calc(100% - 2px);
    min-width: min-content;
    max-width: 100%;
    width: 0%;
    background-color: #333333;
    text-align: right;
    > span {
      color: #ffffff;
      font-size: 0.8em;
      white-space: nowrap;
    }
  }
}
</style>