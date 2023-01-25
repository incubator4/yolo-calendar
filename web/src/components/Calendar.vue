<script setup lang="ts">
import { Calendar } from "v-calendar";

import { useCharacterStore } from "@/stores";
const zeroPad = (num: number, places: number) =>
  String(num).padStart(places, "0");
const store = useCharacterStore();

const today = [
  {
    key: "today",
    highlight: "red",
    dates: new Date(),
  },
];

store.listCalendar();
const attrs = computed(() => [
  ...store.calendars.map((c) => {
    const dates = new Date(c.dateTime);
    console.log(c.dateTime, "   ", dates);
    return {
      key: `${c.dateTime} - ${c.title}`,
      dates,
      popover: {
        label: c.title,
      },
      customData: {
        title: c.title,
        time: dates.getHours(),
      },
      highlight: {
        color: "purple",
        fillMode: "light",
      },
    };
  }),
  ...today,
]);
</script>

<template>
  <Calendar is-expanded :attributes="attrs" :rows="2">
    <template #day-popover="{ day, format, masks, attributes }">
      <div>{{ format(day.date, masks.dayPopover) }}</div>
      <div v-for="attr in attributes">
        <popover-row :key="attr.key" :attribute="attr">
          {{ zeroPad(attr.customData.time, 2) }} - {{ attr.customData.title }}
        </popover-row>
      </div>
    </template>
  </Calendar>
</template>
