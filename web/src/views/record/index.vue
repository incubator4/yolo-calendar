<script setup lang="ts">
import { useCharacterStore } from "@/stores";
import { groupBy } from "lodash";
import moment from "moment";
const store = useCharacterStore();
store.listCalendar({
  end: moment().format("yyyy-MM-DD"),
});

const dailyCalendars = computed(() =>
  groupBy(
    store.calendars.sort((a, b) => (a.start_time > b.start_time ? 1 : -1)),
    (event) => moment(event.start_time).format("yyyy-MM-DD")
  )
);

const date = ref(new Date());
</script>
<template>
  <el-calendar v-model="date">
    <template #date-cell="{ data, date }">
      <p :class="data.isSelected ? 'is-selected' : ''">
        {{ data.day.split("-").slice(1).join("-") }}
        {{ data.isSelected ? "✔️" : "" }}
      </p>

      <div style="witdh: 160px; height: 160px">
        <el-row>
          <el-col :span="4" v-for="cal in dailyCalendars[data.day]">
            <div
              :style="{
                width: '16px',
                height: '16px',

                backgroundColor: cal.main_color,
              }"
            ></div>
          </el-col>
        </el-row>
      </div>
    </template>
  </el-calendar>
</template>

<style scoped>
.el-calendar-day {
  height: 2400px;
}
</style>
