import Vue from "vue";
import Vuex from "vuex";
import ZH from "../strings/zh";
import EN from "../strings/en";

const isChinese = (o) => {
  if (typeof window === "undefined") return;
  return /^(zh)/i.test(
    window.navigator.browserLanguage || window.navigator.language || "zh"
  );
};

const langZH = "zh";
const langEN = "en";

Vue.use(Vuex);

Vue.mixin({
  computed: {
    ...Vuex.mapGetters({
      $locale: "locale"
    })
  },
  methods: {
    ...Vuex.mapActions({
      $swapLang: "swapLang"
    })
  }
});

export default new Vuex.Store({
  state: {
    lang: isChinese() ? langZH : langEN,
    user: {
      username: "",
      role: "",
      usedSpace: 0,
      allSpace: 0,
      unitSpace: "B"
    }
  },
  getters: {
    locale(state) {
      switch (state.lang) {
        case langZH:
          return ZH;
        case langEN:
          return EN;
      }
      return ZH;
    }
  },
  mutations: {
    swapLang(state) {
      switch (state.lang) {
        case langZH:
          state.lang = langEN;
          break;
        case langEN:
          state.lang = langZH;
          break;
      }
    }
  },
  actions: {
    swapLang(context) {
      context.commit("swapLang");
    }
  }
});
