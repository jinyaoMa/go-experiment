<template>
  <div class="head">
    <div class="options">
      <BtnLogout />
      <BtnSettings />
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
      <div class="user-space">
        <div
          class="used"
          :style="{
            width: usedSpaceWidth,
          }"
        >
          <span>
            {{ usedSpace }} {{ usedUnit }} / {{ allSpace }} {{ allUnit }}
          </span>
        </div>
      </div>
    </div>
  </div>
</template>

<script>
import BtnLogout from "../components/BtnLogout.vue";
import BtnSettings from "../components/BtnSettings.vue";
import BtnLangSwap from "../components/BtnLangSwap.vue";

export default {
  components: {
    BtnLogout,
    BtnSettings,
    BtnLangSwap,
  },
  computed: {
    usedSpaceWidth() {
      return `${Math.round(this.$user.usedSpace / this.$user.allSpace)}%`;
    },
    usedSpace() {
      return this.$convertSpaceByUnit(this.$user.usedSpace, this.usedUnit);
    },
    usedUnit() {
      return this.$getUnitBySpace(this.$user.usedSpace);
    },
    allSpace() {
      return this.$convertSpaceByUnit(this.$user.allSpace, this.allUnit);
    },
    allUnit() {
      return this.$getUnitBySpace(this.$user.allSpace);
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
  &.empty {
    width: 2.33em;
  }
}
.user-space {
  margin-left: 1em;
  width: 280px;
  border: 1px solid;
  box-sizing: border-box;
  border-radius: 4px;
  overflow: hidden;
  background-color: #f1f2f3;
  .used {
    box-sizing: border-box;
    padding: 0 0.4em;
    height: 100%;
    line-height: 100%;
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