<template>
  <img class="avatar" v-bind="attrs" :src="src" />
</template>

<script setup lang="ts">
const props = defineProps({
  uid: Number,
});
const attrs = useAttrs();
const src = ref(`/src/assets/avatar/${props.uid}.jpg`);
</script>

<style scoped>
.avatar {
  width: 100px;
  height: 100px;
  border-radius: 10px;
}
</style>
