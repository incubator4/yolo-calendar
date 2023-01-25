<script setup lang="ts">
import { Calendar, PopoverRow } from "v-calendar";
import { Vue3StatusIndicator } from "vue3-status-indicator";
import "vue3-status-indicator/dist/style.css";
import { useCharacterStore } from "@/stores";
import colorMatrix from "@/tools/color";
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
    dates.setHours(dates.getHours() - 8);
    return {
      key: `${c.dateTime} - ${c.title}`,
      dates,
      popover: {
        label: c.title,
        visibility: "focus",
      },
      customData: {
        title: c.title,
        time: dates.getHours(),
        color: colorMatrix(c.cid),
      },
      dot: {
        style: {
          marginTop: "5px",
          backgroundColor: colorMatrix(c.cid),
        },
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
      <div>
        <div
          v-for="attr in attributes"
          :hideIndicator="true"
          :key="attr.key"
          :attribute="attr"
        >
          <Vue3StatusIndicator
            style="width: 8px; height: 8px"
            :bgColor="attr.customData.color"
            pause
          />
          {{ zeroPad(attr.customData.time, 2) }} - {{ attr.customData.title }}
        </div>
      </div>
    </template>
  </Calendar>
</template>
