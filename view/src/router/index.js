import Vue from "vue";
import VueRouter from "vue-router";

Vue.use(VueRouter);

const routes = [
  {
    path: "/login",
    name: "Login",
    component() {
      return import(/* webpackChunkName: "login" */ "../views/Login.vue");
    }
  },
  {
    path: "/signup",
    name: "Signup",
    component() {
      return import(/* webpackChunkName: "signup" */ "../views/Signup.vue");
    }
  },
  {
    path: "*",
    name: "Home",
    component() {
      return import(/* webpackChunkName: "home" */ "../views/Home.vue");
    }
  }
];

const router = new VueRouter({
  routes
});

export default router;
