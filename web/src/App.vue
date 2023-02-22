<script setup lang="ts">
import { RouterLink, RouterView } from "vue-router";
import { useToggle, useDark } from "@vueuse/core";
import { useRouter } from "vue-router";
import { Sunny, Moon } from "@element-plus/icons-vue";
import { useScreen } from "vue-screen";

const screen = useScreen();

const isDark = useDark();
const toggleDark = useToggle(isDark);
const router = useRouter();

const backgroundHeight = computed(() => {
  return screen.height > screen.width ? "60vh" : "100vh";
});

const handleSelect = (key: string, keyPath: string[]) => {
  router.push(key);
};
const routes = [
  {
    text: "Home",
    path: "/",
  },
  {
    text: "Vtuber",
    path: "/vtuber",
  },
  {
    text: "Record",
    path: "/record",
  },
  {
    text: "About",
    path: "/about",
  },
];

// todo 添加 COS API 自动检测 banners 文件名
</script>

<template>
  <div class="wrapper">
    <header class="header">
      <el-row>
        <el-col style="flex: auto" class="logo"></el-col>
        <el-col
          style="
            flex: 1 1 auto;
            text-align: center;
            max-width: calc(100vw - 120px);
          "
        >
          <el-menu
            :default-active="$route.path"
            class="el-menu-demo"
            background-color="var(--el-bg-color-overlay)"
            mode="horizontal"
            @select="handleSelect"
          >
            <el-menu-item
              v-for="({ text, path }, index) in routes"
              :index="path"
              >{{ text }}</el-menu-item
            >
          </el-menu>
        </el-col>
        <el-col style="flex: 0 0 100px">
          <div class="right-menubar">
            <el-switch
              @change="toggleDark"
              :active-icon="Sunny"
              :inactive-icon="Moon"
              v-model="isDark"
              class="ml-2"
            />
          </div>
        </el-col>
      </el-row>
    </header>

    <div class="background-container" :style="{ height: backgroundHeight }">
      <el-carousel height="100vh" :interval="5000" arrow="always">
        <el-carousel-item v-for="item in 4" :key="item">
          <img
            class="background-image"
            :src="`https://yolo-1256553639.cos.ap-shanghai.myqcloud.com/banners/${item}.webp`"
          />
        </el-carousel-item>
      </el-carousel>

      <div class="mask" />
      <div class="title">
        <p class="main">YOLO</p>
        <p class="secondary">YOU ONLY LIVE ONCE</p>
      </div>
    </div>

    <div class="route-backgroud"><RouterView /></div>
  </div>
</template>

<style scoped lang="less">
nav a.router-link-exact-active {
  color: var(--color-text);
  color: yellow;
}

// nav a.router-link-exact-active:hover {
//   background-color: transparent;
// }

nav a {
  display: inline-block;
  padding: 0 1rem;
  border-left: 1px solid var(--color-border);
}

nav a:first-of-type {
  border: 0;
}

// @media (min-width: 1024px) {
//   header {
//     display: flex;
//     place-items: center;
//     padding-right: calc(var(--section-gap) / 2);
//   }

//   .logo {
//     margin: 0 2rem 0 0;
//   }

//   header .wrapper {
//     display: flex;
//     place-items: flex-start;
//     flex-wrap: wrap;
//   }

//   nav {
//     text-align: left;
//     margin-left: -1rem;
//     font-size: 1rem;

//     padding: 1rem 0;
//     margin-top: 1rem;
//   }
// }

.wrapper {
  display: flex;
  flex-direction: column;
  min-height: 100vh;
  .header {
    z-index: 100;
    position: fixed;
    top: 0;
    left: 0;
    right: 0;
    height: 60px;
    background-color: var(--el-bg-color-overlay);
    .header-nav nav {
      color: white;
      text-align: center;
    }
  }
  .background-container {
    position: relative;
    width: 100vw;
    height: 100vh;
    max-width: 100%;
    overflow: hidden;
    // overflow: hidden;
    .title {
      font-family: "Times New Roman";
      position: absolute;
      width: 40%;
      height: 40%;
      top: 50%;
      left: 50%;
      transform: translate(-50%, -50%);
      text-align: center;
      color: white;
      .main {
        font-size: 96px;
        text-align: center;
        margin-left: -40px;
      }
      .secondary {
        text-align: center;
        color: #fdd911;
        margin-left: -10px;
        margin-top: -30px;
        font-size: 18px;
      }
      @media (min-width: 640px) {
        .main {
          font-size: 144px;
        }
        .secondary {
          margin-top: -30px;
          font-size: 24px;
        }
      }
      @media (min-width: 1920px) {
        .main {
          font-size: 256px;
        }
        .secondary {
          margin-top: -30px;
          font-size: 48px;
        }
      }
    }
    .background-image {
      margin-left: -10px;
      margin-right: -10px;
      width: 100%;
      height: 100%;
      object-fit: cover;
    }
    .mask {
      position: absolute;
      top: 0;
      left: 0;
      right: 0;
      bottom: 0;
      background-color: rgba(0, 0, 0, 0.5);
      background-image: radial-gradient(#000000 20%, transparent 0);
      background-size: 18px 18px;
    }
  }
}

.right-menubar {
  margin-top: 12px;
}
.route-backgroud {
  background-color: var(--el-bg-color-overlay);
}
</style>
