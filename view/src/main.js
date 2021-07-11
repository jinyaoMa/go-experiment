import Vue from "vue";
import App from "./App.vue";
import router from "./router";
import store from "./store";
import "@fortawesome/fontawesome-free/css/all.min.css";
import axios from "axios";
import VueAxios from "vue-axios";
import qs from "qs";
import preview from "vue-photo-preview";
import "vue-photo-preview/dist/skin.css";
import VueVideoPlayer from "vue-video-player";
import "video.js/dist/video-js.css";
// import 'vue-video-player/src/custom-theme.css'

Vue.use(preview);
Vue.use(VueVideoPlayer, {
  options: {
    autoplay: "muted",
    playbackRates: [0.7, 1.0, 1.5, 2.0],
    fluid: true,
    controls: true
  }
});

axios.defaults.baseURL = "http://192.168.1.29:55699";
axios.defaults.transformRequest = [
  function(data) {
    return qs.stringify(data);
  }
];
Vue.use(VueAxios, axios);

Vue.config.productionTip = false;

new Vue({
  router,
  store,
  render: function(h) {
    return h(App);
  }
}).$mount("#app");
