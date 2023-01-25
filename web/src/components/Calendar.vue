<script setup lang="ts">
import { Calendar } from "v-calendar";
import Dot from "./Dot.vue";

import { useCharacterStore } from "@/stores";
import colorMatrix from "@/tools/color";
import useClipboard from "vue-clipboard3";
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

const ics = ref("webcal://yolo.incubator4.com/api/ics");
const { toClipboard } = useClipboard();
const onClipboard = () => {
  toClipboard(ics.value);
  alert("复制成功");
};
</script>

<template>
  <div style="margin: 5px">
    <a :href="ics" style="margin: 10px">
      <button class="button">订阅到日历</button>
    </a>
    <a @click="onClipboard" style="margin: 10px">
      <button class="button">复制到剪贴板</button>
    </a>
  </div>
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
