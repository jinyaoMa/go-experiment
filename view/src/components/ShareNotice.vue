<template>
  <div class="share-notice">
    <div class="share-notice-inner">
      <div class="title">
        {{ $locale.common.sharing }}
      </div>
      <div class="msg">
        <a download="share" :href="$shareNoticeLink">
          {{ $shareNoticeLink }}
        </a>
      </div>
      <button @click="$hideShareNotice">
        {{ $locale.common.confirm }}
      </button>
    </div>
  </div>
</template>

<style lang="scss" scoped>
.share-notice {
  box-sizing: border-box;
  position: fixed;
  top: calc(50% - 10em);
  left: 0;
  display: flex;
  flex-direction: column;
  align-items: center;
  width: 100%;
  letter-spacing: 0.1em;
}
.share-notice-inner {
  margin: 1.8em;
  background-color: #ffffff;
  box-sizing: border-box;
  border: 2px solid;
  border-radius: 1.8em;
  padding: 1.8em;
  min-width: 280px;
  display: flex;
  flex-direction: column;
  align-items: center;
  .title {
    font-size: 2em;
    margin-bottom: 0.5em;
    font-weight: bold;
  }
  .msg {
    margin-top: 0.6em;
    margin-bottom: 2em;
    position: relative;
    padding: 1em 2em;
    &::before {
      content: "「";
      position: absolute;
      left: 0;
      top: 0;
    }
    &::after {
      content: "」";
      position: absolute;
      right: 0;
      bottom: 0;
    }
  }
  button {
    border: none;
    padding: 0.5em 1em;
    font-family: inherit;
    font-size: 0.9em;
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