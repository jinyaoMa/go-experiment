import Vue from "vue";
import VueRouter from "vue-router";
import { loadUser } from "../store";

Vue.use(VueRouter);

const routes = [
  {
    path: "/login",
    name: "Login",
    component() {
      return import(/* webpackChunkName: "login" */ "../views/Login.vue");
    },
    meta: {
      spinAtCorner: true,
      requireAuth: false
    }
  },
  {
    path: "/signup",
    name: "Signup",
    component() {
      return import(/* webpackChunkName: "signup" */ "../views/Signup.vue");
    },
    meta: {
      spinAtCorner: true,
      requireAuth: false
    }
  },
  {
    path: "/settings",
    name: "Settings",
    component() {
      return import(/* webpackChunkName: "settings" */ "../views/Settings.vue");
    },
    meta: {
      spinAtCorner: false,
      requireAuth: true
    }
  },
  {
    path: "/recycle",
    name: "Recycle",
    component() {
      return import(/* webpackChunkName: "recycle" */ "../views/Recycle.vue");
    },
    meta: {
      spinAtCorner: false,
      requireAuth: true
    }
  },
  {
    path: "/share",
    name: "Share",
    component() {
      return import(/* webpackChunkName: "share" */ "../views/Share.vue");
    },
    meta: {
      spinAtCorner: false,
      requireAuth: true
    }
  },
  {
    path: "/minip",
    name: "Minip",
    component() {
      return import(/* webpackChunkName: "minip" */ "../views/Minip.vue");
    },
    meta: {
      spinAtCorner: false,
      requireAuth: false
    }
  },
  {
    path: "/miniv",
    name: "Miniv",
    component() {
      return import(/* webpackChunkName: "miniv" */ "../views/Miniv.vue");
    },
    meta: {
      spinAtCorner: false,
      requireAuth: false
    }
  },
  {
    path: "/:cat",
    name: "Cat",
    component() {
      return import(/* webpackChunkName: "cat" */ "../views/Cat.vue");
    },
    meta: {
      spinAtCorner: false,
      requireAuth: true
    }
  },
  {
    path: "*",
    name: "Home",
    component() {
      return import(/* webpackChunkName: "home" */ "../views/Home.vue");
    },
    meta: {
      spinAtCorner: false,
      requireAuth: true
    }
  }
];

const router = new VueRouter({
  routes
});

const isLogin = (o) => {
  const data = loadUser();
  if (data) {
    return (
      typeof data.user.id === "number" &&
      data.user.id > 0 &&
      typeof data.user.token === "string" &&
      data.user.token.length === 128 &&
      typeof data.user.permission === "string" &&
      data.user.permission.includes("CORE:1")
    );
  }
  return false;
};

router.beforeEach((to, from, next) => {
  if (
    /(phone|pad|pod|iPhone|iPod|ios|iPad|Android|Mobile|BlackBerry|IEMobile|MQQBrowser|JUC|Fennec|wOSBrowser|BrowserNG|WebOS|Symbian|Windows Phone)/i.test(
      window.navigator.userAgent
    ) &&
    to.path != "/minip" &&
    to.path != "/miniv"
  ) {
    next({
      path: "/minip",
      query: { redirect: to.fullPath }
    });
  } else if (to.matched.some((record) => record.meta.requireAuth)) {
    if (!isLogin()) {
      next({
        path: "/login",
        query: { redirect: to.fullPath }
      });
    } else {
      next();
    }
  } else {
    next();
  }
});

export default router;
