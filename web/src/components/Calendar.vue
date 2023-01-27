<script setup lang="ts">
import { Calendar } from "v-calendar";
import Dot from "./Dot.vue";

import { useCharacterStore } from "@/stores";
import colorMatrix from "@/tools/color";
const zeroPad = (num: number, places: number) =>
  String(num).padStart(places, "0");
const store = useCharacterStore();

const today = [
  {
    key: "today",
    highlight: {
      fillMode: "light",
      style: {
        backgroundColor: colorMatrix(0),
      },
    },
    dates: new Date(),
  },
];

store.listCalendar();
const attrs = computed(() => [
  ...store.calendars.map((c) => {
    const dates = new Date(c.dateTime);
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

interface attr {
  customData: {
    time: number;
  };
}
const sortAttr = (a: attr, b: attr) =>
  a.customData.time > b.customData.time ? 1 : -1;
</script>

<template>
  <Calendar is-expanded :attributes="attrs" :rows="2">
    <template #day-popover="{ day, format, masks, attributes }">
      <div>{{ format(day.date, masks.dayPopover) }}</div>
      <div>
        <div
          v-for="attr in attributes.sort(sortAttr)"
          :hideIndicator="true"
          :key="attr.key"
          :attribute="attr"
        >
          <Dot
            style="width: 16px; height: 16px"
            :color="attr.customData.color"
          />
          {{ zeroPad(attr.customData.time, 2) }} - {{ attr.customData.title }}
        </div>
      </div>
    </template>
  </Calendar>
</template>

<style scoped>
.button {
  padding: 10px 20px;
  border: 1px solid #ddd;
  color: #333;
  background-color: #fff;
  border-radius: 4px;
  font-size: 14px;
  font-family: "微软雅黑", arail;
  cursor: pointer;
}
</style>
