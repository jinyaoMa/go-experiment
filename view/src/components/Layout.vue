<template>
  <div class="layout">
    <Head />
    <div class="main">
      <div class="cat">
        <Copyright />
        <router-link to="/">
          <i class="fas fa-file fa-fw" />
          {{ $locale.common.all }}
        </router-link>
        <router-link to="/picture">
          <i class="fas fa-fw" />
          {{ $locale.common.pictures }}
        </router-link>
        <router-link to="/documents">
          <i class="fas fa-fw" />
          {{ $locale.common.documents }}
        </router-link>
        <router-link to="/videos">
          <i class="fas fa-fw" />
          {{ $locale.common.videos }}
        </router-link>
        <router-link to="/music">
          <i class="fas fa-fw" />
          {{ $locale.common.music }}
        </router-link>
        <router-link to="/other">
          <i class="fas fa-fw" />
          {{ $locale.common.other }}
        </router-link>
        <router-link to="/share">
          <i class="fas fa-share-alt fa-fw" />
          {{ $locale.common.share }}
        </router-link>
        <router-link to="/recycle">
          <i class="fas fa-trash-alt fa-fw" />
          {{ $locale.common.recycle }}
        </router-link>
      </div>
      <div class="frame">
        <slot></slot>
      </div>
    </div>
  </div>
</template>

<script>
import Head from "./Head.vue";
import Copyright from "./Copyright.vue";

export default {
  components: {
    Head,
    Copyright,
  },
};
</script>

<style lang="scss" scoped>
.layout {
  display: flex;
  flex-direction: column;
  height: 100vh;
}
.main {
  flex-grow: 1;
  display: flex;
  flex-direction: row;
  letter-spacing: 0.1em;
}
.cat {
  position: relative;
  min-width: 240px;
  display: flex;
  flex-direction: column;
  padding: 4px 0;
  > a {
    padding: 0 1.8em;
    line-height: 3em;
    text-decoration: none;
    color: #333333;
    transition: 0.2s;
    &:hover {
      background-color: #f1f2f3;
    }
    &.router-link-exact-active {
      color: #ffffff;
      background-color: #333333;
    }
    > i {
      margin-right: 0.6em;
    }
  }
}
</style>