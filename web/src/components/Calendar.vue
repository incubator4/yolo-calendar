<script setup lang="ts">
import { Calendar } from "v-calendar";
import Dot from "./Dot.vue";

import { useCharacterStore } from "@/stores";
import colorMatrix from "@/tools/color";
import moment from "moment";
import { groupBy, cloneDeep } from "lodash";

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

const startOfWeek = moment().startOf("week").add(1, "days").toDate();
const endofWeek = moment().endOf("week").add(1, "days").toDate();

const currentWeekRange = computed(() => {
  const dates = [];
  for (let i = 0; i <= 6; i++) {
    dates.push(moment(startOfWeek).add(i, "days"));
  }
  return dates;
});

const currentDay = ref(today[0].dates.getDay());

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

const currentWeekCal = computed(() =>
  groupBy(
    store.calendars
      .filter(({ dateTime }) => {
        const date = new Date(dateTime);
        return date > startOfWeek && date < endofWeek;
      })
      .sort((a, b) => (a.dateTime > b.dateTime ? 1 : -1)),
    (event) => {
      const index = new Date(event.dateTime).getDay();
      return index === 0 ? 7 : index;
    }
  )
);

interface attr {
  customData: {
    time: number;
  };
}
const sortAttr = (a: attr, b: attr) =>
  a.customData.time > b.customData.time ? 1 : -1;
</script>

<template>
  <el-carousel
    :interval="4000"
    :initial-index="currentDay === 0 ? 6 : currentDay - 1"
    indicator-position="none"
    :autoplay="false"
    :loop="false"
    type="card"
    height="600px"
  >
    <el-carousel-item v-for="item in 7" :key="item">
      <el-card class="box-card">
        <template #header>
          <div class="card-header">
            <span>{{ currentWeekRange[item - 1].format("yyyy-MM-DD") }}</span>
          </div>
        </template>
        <el-timeline>
          <el-timeline-item
            v-for="(event, index) in currentWeekCal[item]"
            :key="index"
            :color="colorMatrix(event.cid)"
            :timestamp="new Date(event.dateTime).getHours() + '点'"
          >
            {{ event.title }}
          </el-timeline-item>
        </el-timeline>
      </el-card>
    </el-carousel-item>
  </el-carousel>
</template>

<style scoped>
.button {
  padding: 10px 20px;
  border: 1px solid #ddd;
  /* color: #333; */
  background-color: #fff;
  border-radius: 4px;
  font-size: 14px;
  font-family: "微软雅黑", arail;
  cursor: pointer;
}
.el-carousel__item h3 {
  /* color: #475669; */
  opacity: 0.75;
  line-height: 200px;
  margin: 0;
  text-align: center;
}

.el-carousel__item:nth-child(2n) {
  /* background-color: #99a9bf; */
}

.el-carousel__item:nth-child(2n + 1) {
  /* background-color: #d3dce6; */
}
</style>
