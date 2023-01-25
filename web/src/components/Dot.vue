<script setup lang="ts">
const attr = useAttrs();
const props = defineProps<{ color: string }>();
</script>

<template>
  <svg
    t="1674634629675"
    class="dot"
    viewBox="0 0 1024 1024"
    version="1.1"
    xmlns="http://www.w3.org/2000/svg"
    p-id="2709"
    height="30"
    width="30"
    v-bind="attr"
    :fill="props.color"
  >
    <path
      d="M512 298.666667c117.333333 0 213.333333 96 213.333333 213.333333s-96 213.333333-213.333333 213.333333-213.333333-96-213.333333-213.333333S394.666667 298.666667 512 298.666667z"
      p-id="2710"
    ></path>
  </svg>
</template>

<style scoped>
.dot {
  top: 3px;
  border-radius: 100%;
}
</style>
