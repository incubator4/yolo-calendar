<script setup lang="tsx">
import { useCharacterStore } from "@/stores";
import moment from "moment";
import { groupBy } from "lodash";
import Avatar from "./icons/Avatar.vue";
import Calendar from "./Calendar/index.vue";
import { useScreen } from "vue-screen";

const store = useCharacterStore();
const screen = useScreen();

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

store.fetchAll().then(() => {
  checkAll.value = true;
  isIndeterminate.value = false;
  handleCheckAllChange(true);
});

store.listCalendar({
  // start: moment(startOfWeek).format("yyyy-MM-DD"),
  end: moment(endofWeek).add(1, "days").format("yyyy-MM-DD"),
});

const currentWeekCal = computed(() =>
  groupBy(
    store.calendars
      // .filter(({ start_time }) => {
      //   const date = new Date(start_time);
      //   return date > startOfWeek && date < endofWeek;
      // })
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

const checkboxVtubers = ref<Array<number>>([]);

const checkAll = ref(false);
const isIndeterminate = ref(true);
const handleCheckAllChange = (val: boolean) => {
  checkboxVtubers.value = val ? store.characters.map((c) => c.id) : [];
  isIndeterminate.value = false;
};

const handleCheckedVtubersChange = (value: string[]) => {
  const checkedCount = value.length;
  checkAll.value = checkedCount === store.characters.length;
  isIndeterminate.value =
    checkedCount > 0 && checkedCount < store.characters.length;
};
</script>

<template>
  <div class="wrapper">
    <div class="check-wrapper">
      <!--
      <el-checkbox
        v-model="checkAll"
        :indeterminate="isIndeterminate"
        @change="handleCheckAllChange"
        >Check all</el-checkbox
      >
      <el-checkbox-group
        v-model="checkboxVtubers"
        @change="handleCheckedVtubersChange"
        size="large"
      >
        <template v-if="screen.width > 440">
          <el-checkbox-button
            v-for="c in store.characters"
            :key="c.id"
            :label="c.id"
          >
            <Avatar style="width: 30px; height: 30px" :uid="c.uid" />
          </el-checkbox-button>
        </template>
        <template v-else>
          <el-checkbox v-for="c in store.characters" :key="c.id" :label="c.id">
            {{ c.name }}
          </el-checkbox>
        </template>
      </el-checkbox-group>

    -->
      <el-select
        style="width: 100%"
        v-model="checkboxVtubers"
        multiple
        placeholder="Select"
      >
        <el-option
          v-for="c in store.characters"
          :key="c.id"
          :label="c.name"
          :value="c.id"
        >
          <span style="float: left">{{ c.name }}</span>
          <span
            style="
              float: right;
              color: var(--el-text-color-secondary);
              font-size: 13px;
            "
            >{{ c.uid }}</span
          >
        </el-option>
      </el-select>
    </div>
    <Calendar
      :calendars="
        store.calendars.filter((cal) => checkboxVtubers.includes(cal.cid))
      "
    />
  </div>
</template>

<style scoped>
.wrapper {
  margin: 10px;
}

.check-wrapper {
  /* margin: 5px; */
  margin-bottom: 5px;
}
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
