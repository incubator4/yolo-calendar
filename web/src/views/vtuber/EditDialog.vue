<script setup lang="ts">
import { useCharacterStore } from "@/stores";
import moment from "moment";
import { cloneDeep, pick } from "lodash";

const store = useCharacterStore();

const props = defineProps<{
  modelValue: boolean;
}>();

const emit = defineEmits(["update:modelValue"]);

const startOfWeek = moment().startOf("week").toDate();
const endofWeek = moment().endOf("week").toDate();

const minusHour = (start: Date, end: Date) => {
  const diffTime = end.getTime() - start.getTime();
  return Math.floor(diffTime / (1000 * 60 * 60));
};

const calcValue = computed<boolean>({
  get() {
    return props.modelValue;
  },
  set(val) {
    emit("update:modelValue", val);
  },
});

const currentWeekCal = computed(() =>
  store.calendars
    .filter((event) => {
      const date = new Date(event.start_time);
      return date > startOfWeek && date < endofWeek;
    })
    .map((event) =>
      pick(event, [
        "id",
        "title",
        "start_time",
        "end_time",
        "title",
        "cid",
        "tag_id",
      ])
    )
);

const editWeekCal = ref<ICalendar[]>([]);

watch(
  () => props.modelValue,
  (val) => {
    if (val) {
      store.listTags();
      editWeekCal.value = cloneDeep(currentWeekCal.value);
    }
  }
);

const append = () => {
  const c = store.curentCharacter;
  if (c) {
    editWeekCal.value = [
      ...editWeekCal.value,
      {
        title: "",
        start_time: new Date(),
        end_time: moment().add(1, "h").toDate(),
        tag_id: 0,
        cid: c.id,
      },
    ];
  }
};

const handleClose = () => {
  emit("update:modelValue", false);
};
</script>
<template>
  <el-dialog v-model="calcValue" title="Tips" width="70%">
    <el-card class="event" v-for="event in editWeekCal">
      <el-row>
        <el-col :span="12">
          <el-form-item label="标题">
            <el-input style="width: 90%" v-model="event.title"></el-input>
          </el-form-item>
        </el-col>
        <el-col :span="12">
          <el-form-item label="类型">
            <el-select v-model="event.tag_id" placeholder="">
              <el-option label="暂无" :value="0"></el-option>
              <el-option
                v-for="tag in store.tags"
                :label="tag.name"
                :value="tag.id"
              ></el-option>
            </el-select>
          </el-form-item>
        </el-col>
        <el-col :span="12">
          <el-form-item label="时间">
            <el-date-picker
              v-model="event.start_time"
              type="datetime"
              placeholder="Pick a date"
              style="width: 90%"
            />
          </el-form-item>
        </el-col>
        <el-col :span="12">
          <el-form-item label="时长">
            <el-select
              :model-value="
                minusHour(new Date(event.start_time), new Date(event.end_time))
              "
              @change="(val: number) => {
                event.end_time = moment(event.start_time).add(val, 'hour').toDate()
              }"
              placeholder=""
            >
              <el-option
                v-for="item in 8"
                :label="item + '小时'"
                :value="item"
              ></el-option>
            </el-select>
          </el-form-item>
        </el-col>
      </el-row>
    </el-card>
    <template #footer>
      <span class="dialog-footer">
        <el-button @click="append">新增</el-button>
        <el-button type="danger" @click="handleClose">取消</el-button>
        <el-button type="primary" @click="handleClose"> 提交 </el-button>
      </span>
    </template>
  </el-dialog>
</template>
<style scoped>
.event {
  border: 2px;
  border-radius: 2px;
  margin-top: 5px;
}
</style>
