import { fileURLToPath, URL } from "node:url";

import { defineConfig } from "vite";
import vue from "@vitejs/plugin-vue";
import vueJsx from "@vitejs/plugin-vue-jsx";
import importToCDN, { autoComplete } from "vite-plugin-cdn-import";
import AutoImport from "unplugin-auto-import/vite";
import viteCompression from "vite-plugin-compression";
import { chunkSplitPlugin } from "vite-plugin-chunk-split";
import commonjs from "rollup-plugin-commonjs";
import { VitePluginFonts } from "vite-plugin-fonts";

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
        autoComplete("axios"),
        autoComplete("lodash"),
        // {
        //   name: "element-plus",
        //   var: "ElementPlus",
        //   path: "dist/index.full.min.js",
        // },
        // {
        //   name: "moment",
        //   var: "Moment",
        //   path: "min/moment.min.js",
        // },
      ],
      prodUrl: "//unpkg.com/{name}@{version}/{path}",
      // prodUrl: "//cdn.bootcdn.net/ajax/libs/{name}/{version}/{path}",
    }),
    VitePluginFonts({
      custom: {
        families: [
          // {
          //   name: "灵动指书",
          //   local: "灵动指书",
          //   src: "https://yolo-1256553639.cos.ap-shanghai.myqcloud.com/fonts/灵动指书手机字体.ttf",
          // },
        ],
      },
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
        target: "https://calendar.incubator4.com",
        changeOrigin: true,
      },
    },
  },
  css: {
    postcss: {
      plugins: [],
    },
  },
  build: {
    rollupOptions: {
      // external: ["element-plus"],
      plugins: [commonjs()],
    },
  },
});
