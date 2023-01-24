<script setup lang="ts">
import { useCharacterStore } from "@/stores";
const props = defineProps({ uid: String });

const store = useCharacterStore();

watch(
  () => props.uid,
  (newUID, oldUID) => {
    if (newUID !== "" && newUID !== undefined && newUID !== oldUID) {
      const uid = +(newUID as string);
      store.clearCalendar();
      store.getCalendar(uid);
    }
  }
);
</script>

<template>
  <p>{{ uid }}</p>
  <div v-for="c in store.calendars">{{ c }}</div>
</template>
