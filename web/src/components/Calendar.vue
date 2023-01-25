<script setup lang="ts">
// Documentation: https://antoniandre.github.io/vue-cal

import VueCal from "vue-cal";
import { useCharacterStore } from "@/stores";
const store = useCharacterStore();

const flag = ref(false);
store.listCalendar().then(() => {
  flag.value = true;
});
const print = console.log;

const events = computed(() =>
  store.calendars.map((c) => {
    const start = new Date(c.dateTime);
    const end = new Date(c.dateTime);
    end.setHours(end.getHours() + 1);
    return {
      start,
      end,
      title: c.title,
    };
  })
);
</script>

<template>
  <VueCal
    v-if="flag"
    class="vuecal--blue-theme cal"
    locale="zh-cn"
    :disable-views="['years', 'year']"
    :time-from="8 * 60"
    :time-to="24 * 60"
    :time-step="120"
    timeFormat=""
    :events="events"
  >
    <template #cell-content="{ cell, view, events, goNarrower }">
      <span
        style="backgroud: #abacdb"
        class="vuecal__cell-date"
        :class="view.id"
        v-if="view.id === 'day'"
        @click="goNarrower"
      >
        {{ cell.date.getDate() }}1111111
      </span>
      <span
        class="vuecal__cell-events-count"
        v-if="view.id === 'month' && events.length"
        >{{ events.length }}</span
      >
      <span
        class="vuecal__no-event"
        v-if="['week', 'day'].includes(view.id) && !events.length"
        >Nothing here 👌</span
      >
    </template>
  </VueCal>
</template>

<style scoped>
.cal {
  max-height: 500px;
}
</style>
