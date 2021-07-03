import Vue from "vue";
import Vuex from "vuex";
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

const loadUser = (o) => {
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
      $loading: "loading",
      $locale: "locale",
      $user: "user"
    })
  },
  methods: {
    ...Vuex.mapActions({
      $startLoading: "startLoading",
      $stopLoading: "stopLoading",
      $swapLang: "swapLang",
      $setUser: "setUser"
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
    }
  }
});

export default new Vuex.Store({
  state: {
    loading: true,
    lang: isChinese() ? langZH : langEN,
    ...loadLang(),
    user: {
      id: 0,
      name: "",
      role: "",
      usedSpace: 0,
      allSpace: 0,
      token: ""
    },
    ...loadUser()
  },
  getters: {
    loading(state) {
      return state.loading;
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
    }
  },
  mutations: {
    startLoading(state) {
      state.loading = true;
    },
    stopLoading(state) {
      state.loading = false;
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
    }
  },
  actions: {
    startLoading(context) {
      context.commit("startLoading");
    },
    stopLoading(context) {
      context.commit("stopLoading");
    },
    swapLang(context) {
      context.commit("swapLang");
    },
    setUser(context, user) {
      context.commit("setUser", user);
    }
  }
});
