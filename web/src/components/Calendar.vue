<script setup lang="tsx">
import { useCharacterStore } from "@/stores";
import moment from "moment";
import { groupBy } from "lodash";
import Avatar from "./icons/Avatar.vue";

const store = useCharacterStore();

const getUID = (cid: number) =>
  (store.characters.find((c) => c.id === cid) || { uid: 0 }).uid;

const startOfWeek = moment().startOf("week").toDate();
const endofWeek = moment().endOf("week").toDate();

const currentWeekRange = computed(() => {
  const dates = [];
  for (let i = 0; i <= 6; i++) {
    dates.push(moment(startOfWeek).add(i, "days"));
  }
  return dates;
});

const currentDay = ref(moment().days() === 0 ? 7 : moment().days());

store.listCalendar();

const currentWeekCal = computed(() =>
  groupBy(
    store.calendars
      .filter(({ start_time }) => {
        const date = new Date(start_time);
        return date > startOfWeek && date < endofWeek;
      })
      .sort((a, b) => (a.start_time > b.start_time ? 1 : -1)),
    (event) => {
      const index = new Date(event.start_time).getDay();
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

const matchWeek = (index: number) => {
  switch (index) {
    case 1:
      return "星期一";
    case 2:
      return "星期二";
    case 3:
      return "星期三";
    case 4:
      return "星期四";
    case 5:
      return "星期五";
    case 6:
      return "星期六";
    case 7:
      return "星期日";
  }
};

const Icon = (uid: number) => {
  <Avatar uid={uid}></Avatar>;
};
</script>

<template>
  <el-skeleton v-if="false" :rows="5" />
  <el-carousel
    class="carousel"
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
            <span
              :style="{ fontWeight: currentDay === item ? 'bolder' : 'normal' }"
              >{{ currentWeekRange[item - 1].format("yyyy-MM-DD") }}
              {{ matchWeek(item) }}
            </span>
          </div>
        </template>
        <el-timeline>
          <el-timeline-item
            v-for="(event, index) in currentWeekCal[item]"
            :key="index"
            :color="event.main_color"
            :icon="Icon(event.uid)"
            :timestamp="new Date(event.start_time).getHours() + '点'"
          >
            <el-link
              :style="{ marginTop: '-8px' }"
              target="_blank"
              :href="`https://live.bilibili.com/${event.live_id}`"
              >{{ event.title }}</el-link
            >
          </el-timeline-item>
        </el-timeline>
      </el-card>
    </el-carousel-item>
  </el-carousel>
</template>

<style scoped>
.carousel {
  margin-top: 40px;
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
