import Vue from "vue";
import Vuex from "vuex";
import * as moment from "moment";
import ZH from "../strings/zh";
import EN from "../strings/en";

const loadLang = (o) => {
  if (typeof window === "undefined") return {};
  try {
    return JSON.parse(
      window.localStorage.getItem(window.btoa(`${window.location.host}:lang`))
    );
  } catch (error) {
    return {};
  }
};

const saveLang = (lang) => {
  window.localStorage.setItem(
    window.btoa(`${window.location.host}:lang`),
    JSON.stringify({ lang })
  );
};

export const loadUser = (o) => {
  if (typeof window === "undefined") return {};
  try {
    return JSON.parse(
      window.localStorage.getItem(window.btoa(`${window.location.host}:user`))
    );
  } catch (error) {
    return {};
  }
};

const saveUser = (user) => {
  window.localStorage.setItem(
    window.btoa(`${window.location.host}:user`),
    JSON.stringify({ user })
  );
};

const isChinese = (o) => {
  if (typeof window === "undefined") return;
  return /^(zh)/i.test(
    window.navigator.browserLanguage || window.navigator.language || "zh"
  );
};

const langZH = "zh";
const langEN = "en";

const B = 1;
const KB = 1024 * B;
const MB = 1024 * KB;
const GB = 1024 * MB;

Vue.use(Vuex);

Vue.mixin({
  computed: {
    ...Vuex.mapGetters({
      $progress: "progress",
      $progressing: "progressing",
      $loading: "loading",
      $hasError: "hasError",
      $errorMsg: "errorMsg",
      $hasShareNotice: "hasShareNotice",
      $shareNoticeLink: "shareNoticeLink",
      $locale: "locale",
      $user: "user",
      $files: "files"
    })
  },
  methods: {
    ...Vuex.mapActions({
      $setProgress: "setProgress",
      $startProgressing: "startProgressing",
      $stopProgressing: "stopProgressing",
      $startLoading: "startLoading",
      $stopLoading: "stopLoading",
      $showError: "showError",
      $hideError: "hideError",
      $showShareNotice: "showShareNotice",
      $hideShareNotice: "hideShareNotice",
      $swapLang: "swapLang",
      $setUser: "setUser",
      $setFiles: "setFiles",
      $setFileCheckById: "setFileCheckById",
      $clearChecklist: "clearChecklist"
    }),
    $getUnitBySpace(space) {
      if (space < KB) {
        return "B";
      } else if (space < MB) {
        return "KB";
      } else if (space < GB) {
        return "MB";
      } else {
        return "GB";
      }
    },
    $convertSpaceByUnit(space, unit) {
      switch (unit) {
        case "KB":
          return Math.round((space / KB) * 100) / 100;
        case "MB":
          return Math.round((space / MB) * 100) / 100;
        case "GB":
          return Math.round((space / GB) * 100) / 100;
        default:
          return space;
      }
    },
    $revertSpaceByUnit(space, unit) {
      switch (unit) {
        case "KB":
          return space * KB;
        case "MB":
          return space * MB;
        case "GB":
          return space * GB;
        default:
          return space;
      }
    },
    $convertSpace2String(space) {
      const unit = this.$getUnitBySpace(space);
      return `${this.$convertSpaceByUnit(space, unit)} ${unit}`;
    },
    $formatDate(date, fmtStr) {
      return moment(date).format(fmtStr);
    },
    $isDateBefore(a, b) {
      return moment(a).isBefore(b);
    },
    $getCurrentPathId() {
      let currentPath = this.$route.query.currentPath;
      if (typeof currentPath === "string" && currentPath.length > 0) {
        let file = this.$files.find((f) => f.Path === currentPath);
        if (file) {
          return file.ID;
        }
      } else {
        let file = this.$files.find((f) => f.Path === ".");
        if (file) {
          return file.ID;
        }
      }
      return 0;
    }
  }
});

export default new Vuex.Store({
  state: {
    progress: 0,
    progressing: false,
    loading: true,
    hasError: false,
    errorMsg: "",
    hasShareNotice: false,
    shareNoticeLink: "",
    lang: isChinese() ? langZH : langEN,
    ...loadLang(),
    user: {
      id: 0,
      name: "",
      role: "",
      permission: "",
      usedSpace: 0,
      allSpace: 0,
      token: ""
    },
    ...loadUser(),
    files: []
  },
  getters: {
    progress(state) {
      return state.progress;
    },
    progressing(state) {
      return state.progressing;
    },
    loading(state) {
      return state.loading;
    },
    hasError(state) {
      return state.hasError;
    },
    errorMsg(state) {
      return state.errorMsg;
    },
    hasShareNotice(state) {
      return state.hasShareNotice;
    },
    shareNoticeLink(state) {
      return state.shareNoticeLink;
    },
    locale(state) {
      switch (state.lang) {
        case langZH:
          return ZH;
        case langEN:
          return EN;
      }
      return ZH;
    },
    user(state) {
      return state.user;
    },
    files(state) {
      return state.files;
    }
  },
  mutations: {
    setProgress(state, percent) {
      state.progress = percent;
    },
    startProgressing(state) {
      state.progressing = true;
    },
    stopProgressing(state) {
      state.progressing = false;
    },
    startLoading(state) {
      state.loading = true;
    },
    stopLoading(state) {
      state.loading = false;
    },
    showError(state, text) {
      state.errorMsg = text;
      state.hasError = true;
    },
    hideError(state) {
      state.hasError = false;
    },
    showShareNotice(state, text) {
      state.shareNoticeLink = text;
      state.hasShareNotice = true;
    },
    hideShareNotice(state) {
      state.hasShareNotice = false;
    },
    swapLang(state) {
      switch (state.lang) {
        case langZH:
          state.lang = langEN;
          break;
        case langEN:
          state.lang = langZH;
          break;
      }
      saveLang(state.lang);
    },
    setUser(state, user) {
      Object.assign(state.user, user);
      saveUser(state.user);
    },
    setFiles(state, files) {
      for (let i = 0; i < files.length; i++) {
        files[i].isChecked = false; // add for checklist
        files[i].canRename = false; // add for rename
      }
      state.files = files;
    },
    setFileCheckById(state, { fileId, isChecked }) {
      for (let i = 0; i < state.files.length; i++) {
        if (state.files[i].ID === fileId) {
          state.files[i].isChecked = isChecked;
          break;
        }
      }
    },
    clearChecklist(state) {
      for (let i = 0; i < state.files.length; i++) {
        state.files[i].isChecked = false;
      }
    }
  },
  actions: {
    setProgress(context, percent) {
      context.commit("setProgress", percent);
    },
    startProgressing(context) {
      context.commit("startProgressing");
    },
    stopProgressing(context) {
      context.commit("stopProgressing");
    },
    startLoading(context) {
      context.commit("startLoading");
    },
    stopLoading(context) {
      context.commit("stopLoading");
    },
    showError(context, text) {
      context.commit("showError", text);
    },
    hideError(context) {
      context.commit("hideError");
    },
    showShareNotice(context, text) {
      context.commit("showShareNotice", text);
    },
    hideShareNotice(context) {
      context.commit("hideShareNotice");
    },
    swapLang(context) {
      context.commit("swapLang");
    },
    setUser(context, user) {
      context.commit("setUser", user);
    },
    setFiles(context, files) {
      context.commit("setFiles", files);
    },
    setFileCheckById(context, options) {
      context.commit("setFileCheckById", options);
    },
    clearChecklist(context) {
      context.commit("clearChecklist");
    }
  }
});
