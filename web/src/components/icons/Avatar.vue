<template>
  <img :src="src" />
</template>

<script setup lang="ts">
import request from "@/tools/axios";
const props = defineProps({
  uid: Number,
});

const src = ref("");

request
  .get<{ face: string }>(`https://api.vtbs.moe/v1/detail/${props.uid}/`)
  .then((res) => {
    console.log(res.data);
    src.value = res.data.face;
  });
</script>
