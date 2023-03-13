import { createRouter, createWebHistory } from "vue-router";

const router = createRouter({
  history: createWebHistory(import.meta.env.BASE_URL),
  routes: [
    {
      path: "/",
      name: "home",
      component: () => import("../views/HomeView.vue"),
    },
    {
      path: "/about",
      name: "about",
      // route level code-splitting
      // this generates a separate chunk (About.[hash].js) for this route
      // which is lazy-loaded when the route is visited.
      component: () => import("../views/AboutView.vue"),
    },
    {
      path: "/vtuber",
      name: "Vtuber",
      component: () => import("../views/vtuber/List.vue"),
    },
    {
      path: "/record",
      name: "Record",
      component: () => import("../views/record/index.vue"),
    },
    {
      path: "/vtuber/:uid",
      name: "VtuberDetail",
      props: true,
      component: () => import("../views/vtuber/Detail.vue"),
    },
    {
      path: "/status",
      name: "Status",
      component: () => import("../views/StatusView.vue"),
    },
  ],
});

export default router;
