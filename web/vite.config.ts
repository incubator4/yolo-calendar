import { fileURLToPath, URL } from "node:url";

import { defineConfig } from "vite";
import vue from "@vitejs/plugin-vue";
import vueJsx from "@vitejs/plugin-vue-jsx";
import importToCDN, { autoComplete } from "vite-plugin-cdn-import";
import AutoImport from "unplugin-auto-import/vite";
import viteCompression from "vite-plugin-compression";
import { chunkSplitPlugin } from "vite-plugin-chunk-split";

// import postcssImport from "postcss-import";
// import tailwindcss from "tailwindcss";

// https://vitejs.dev/config/
export default defineConfig({
  plugins: [
    vue(),
    vueJsx(),
    AutoImport({
      imports: ["vue"],
      dts: "src/auto-import.d.ts",
    }),
    viteCompression(),
    importToCDN({
      modules: [
        autoComplete("vue"),
        autoComplete("axios"),
        autoComplete("lodash"),
        autoComplete("@vueuse/core"),
        {
          name: "vue-router",
          var: "VueRouter",
          path: "dist/vue-router.global.prod.js",
        },
        {
          name: "element-plus",
          var: "ElementPlus",
          path: "dist/index.full.min.js",
        },
        {
          name: "moment",
          var: "Moment",
          path: "min/moment.min.js",
        },
      ],
      prodUrl: "//unpkg.com/{name}@{version}/{path}",
      // prodUrl: "//cdn.bootcdn.net/ajax/libs/{name}/{version}/{path}",
    }),
    chunkSplitPlugin(),
  ],
  resolve: {
    alias: {
      "@": fileURLToPath(new URL("./src", import.meta.url)),
    },
  },
  server: {
    proxy: {
      "/api": {
        target: "https://yolo.incubator4.com",
        changeOrigin: true,
      },
    },
  },
  css: {
    postcss: {
      plugins: [],
    },
  },
});
