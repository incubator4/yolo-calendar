import { createApp } from "vue";
import { createPinia } from "pinia";

import App from "./App.vue";
import router from "./router";

import ElementPlus from "element-plus";
import "element-plus/dist/index.css";
import "element-plus/theme-chalk/dark/css-vars.css";

import "./assets/main.css";
import "v-calendar/dist/style.css";

import moment from "moment";

moment.locale("zh-cn", {
  week: {
    dow: 1,
  },
});

const app = createApp(App);

app.use(ElementPlus);
app.use(createPinia());
app.use(router);

app.mount("#app");
