<template>
  <el-avatar v-if="src !== ''" class="avatar" v-bind="attrs" :src="src" />
  <svg
    v-else
    t="1674584839181"
    class="icon avatar"
    viewBox="0 0 1024 1024"
    version="1.1"
    xmlns="http://www.w3.org/2000/svg"
    p-id="2690"
    width="200"
    height="200"
    v-bind="attrs"
  >
    <path
      d="M270.4 214.4C336 160 420 128 512 128c212 0 384 172 384 384h64c0-247.2-200.8-448-448-448-107.2 0-205.6 37.6-282.4 100l40.8 50.4z"
      p-id="2691"
    ></path>
  </svg>
</template>

<script setup lang="ts">
// import axios from "axios";
const props = defineProps({
  uid: Number,
});
const attrs = useAttrs();
const src = ref(
  `	https://yolo-1256553639.cos.ap-shanghai.myqcloud.com/assets/avatar/${props.uid}.webp`
);

// axios
//   .get<{ face: string }>(`https://api.vtbs.moe/v1/detail/${props.uid}/`)
//   .then((res) => {
//     src.value = res.data.face;
//   });
</script>

<style scoped>
.avatar {
  width: 100px;
  height: 100px;
  border-radius: 10px;
}
@keyframes loading {
  from {
    transform: rotate(0deg);
  }
  to {
    transform: rotate(360deg);
  }
}
.icon {
  animation-name: loading;
  animation-duration: 2s;
}
</style>
