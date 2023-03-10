import { createApp } from "vue";
import { createPinia } from "pinia";

import App from "./App.vue";
import VueKonva from "vue-konva";
import router from "./router";

import ElementPlus from "element-plus";
import zhCn from "element-plus/lib/locale/lang/zh-cn";
import "element-plus/dist/index.css";
import "element-plus/theme-chalk/dark/css-vars.css";

import "./assets/main.css";
// import "./assets/font.css";
import "virtual:fonts.css";
import "v-calendar/dist/style.css";

import moment from "moment";

moment.locale("zh-cn", {
  week: {
    dow: 1,
  },
});

const app = createApp(App);

app.use(ElementPlus, {
  locale: zhCn,
});
app.use(VueKonva);
app.use(createPinia());
app.use(router);

app.mount("#app");
