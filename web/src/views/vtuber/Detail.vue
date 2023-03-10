<script setup lang="ts">
import {
  useVtuberStore,
  useCalendarStore,
  useImageRenderConfig,
} from "@/stores";
import { useDark } from "@vueuse/core";
import { useScreen } from "vue-screen";
import moment from "moment";
import Avatar from "@/components/icons/Avatar.vue";
import EditDialog from "./EditDialog.vue";
import { Hide, View } from "@element-plus/icons-vue";
import { ElMessage } from "element-plus";
import ImageRender from "@/components/imageRender/index.vue";

const props = defineProps({ uid: String });

const vtuberStore = useVtuberStore();
const calendarStore = useCalendarStore();
const configStore = useImageRenderConfig();
const screen = useScreen();
const isDark = useDark();

const loadData = (uid: number) => {
  ics.value = `webcal://${window.location.host}/api/ics/${props.uid}`;
  calendarStore.clearCalendar();
  calendarStore.listCalendar({ uid: [uid.toString()] }).then(() => {
    state.total = calendarStore.calendars.length;
  });
  vtuberStore.getVtuber(uid).then(() => {
    const id = vtuberStore.curentVtuber?.id as number;
    eventModel.cid = id;
    configStore.list(id);
  });
  calendarStore.listTags();
};

onMounted(() => {
  loadData(+(props.uid as string));
});

const zeroPad = (num: number, places: number) =>
  String(num).padStart(places, "0");

watch(
  () => props.uid,
  (newUID, oldUID) => {
    if (newUID !== "" && newUID !== undefined && newUID !== oldUID) {
      const uid = +(newUID as string);
      loadData(uid);
    }
  }
);

const ics = ref("");
const dialogVisible = ref(false);

const defaultEvent: ICalendar = {
  title: "",
  start_time: moment().toDate(),
  end_time: moment().add(2, "h").toDate(),
  cid: vtuberStore.curentVtuber?.id as number,
  tag_id: 0,
  is_active: true,
};

let eventModel = reactive<ICalendar>(defaultEvent);

const onEdit = (event: ICalendar) => {
  eventModel = {
    ...event,
  };
  dialogVisible.value = true;
};

const onNew = () => {
  eventModel = {
    ...defaultEvent,
  };
  const d = moment();
  eventModel.start_time = d.toDate();
  eventModel.end_time = d.add(2, "h").toDate();
  dialogVisible.value = true;
};

const onDelete = (id: number) => {
  calendarStore.deleteCalendar(id).then(() => {
    ElMessage({
      message: "????????????",
      type: "success",
    });
    loadData(+(props.uid as string));
  });
};

const state = reactive({
  page: 1,
  limit: 10,
  total: 0,
});

const isReverse = ref(true);

const tableData = computed(() => {
  const data = calendarStore.calendars
    .sort((a, b) => (a.start_time > b.start_time ? 1 : -1))
    .filter((event, index) => {
      const date = new Date(event.start_time);
      const [from, to] = pickDateRange.value;
      const res = date.getTime() > from.getTime();
      return res;
    });
  return isReverse.value ? data.reverse() : data;
});

const tableDataReverse = computed(() => tableData.value.reverse());

const size = ref<"default" | "large" | "small">("default");

const pickDateRange = ref<[Date, Date]>([
  moment().startOf("week").toDate(),
  moment().endOf("week").toDate(),
]);

const shortcuts = [
  {
    text: "??????",
    value: () => {
      return [
        moment().startOf("week").toDate(),
        moment().endOf("week").toDate(),
      ];
    },
  },
  {
    text: "??????",
    value: () => {
      return [
        moment().startOf("month").toDate(),
        moment().endOf("month").toDate(),
      ];
    },
  },
  {
    text: "?????????",
    value: () => {
      return [
        moment().startOf("quarter").toDate(),
        moment().endOf("quarter").toDate(),
      ];
    },
  },
  {
    text: "??????",
    value: () => {
      return [
        moment().startOf("year").toDate(),
        moment().endOf("year").toDate(),
      ];
    },
  },
];

const tagName = (tag_id: number) => {
  const tag = calendarStore.tags.find((t) => t.id === tag_id);
  return tag ? tag.name : "-";
};

const renderPanel = ref(false);
const onRender = () => {
  renderPanel.value = true;
};
</script>

<template>
  <main>
    <div style="padding: 40px">
      <el-row>
        <el-col style="flex: 0 1 400px">
          <Avatar :uid="+(uid  as  string)" />
        </el-col>
        <el-col :span="2"> </el-col>
        <el-col style="flex: auto">
          <el-skeleton style="width: 100%" :row="10"> </el-skeleton
        ></el-col>
      </el-row>
      <el-row style="margin-top: 20px">
        <el-col :span="2">
          <el-form-item label="??????">
            <el-switch v-model="isReverse"></el-switch>
          </el-form-item>
        </el-col>
        <el-col :span="2"></el-col>
        <el-col :span="10">
          <el-date-picker
            v-model="pickDateRange"
            type="daterange"
            unlink-panels
            range-separator="To"
            start-placeholder="Start date"
            end-placeholder="End date"
            :shortcuts="shortcuts"
            :size="size"
          />
        </el-col>
        <el-col :span="4"></el-col>
        <el-col :span="4">
          <el-button type="primary" @click="onNew">??????</el-button>
          <el-button @click="onRender">??????</el-button>
        </el-col>
      </el-row>

      <el-table :data="tableData" style="width: 100%">
        <el-table-column label="Active" width="70">
          <template #default="{ row }">
            <el-icon v-if="row.is_active"><View /></el-icon>
            <el-icon v-else><Hide /></el-icon>
          </template>
        </el-table-column>
        <el-table-column label="Datetime" width="180">
          <el-table-column label="Date" width="120">
            <template #default="{ row }">
              {{ moment(row.start_time).format("YYYY-MM-DD") }}
            </template>
          </el-table-column>
          <el-table-column label="Time" width="80">
            <template #default="{ row }">
              {{ moment(row.start_time).format("HH:mm") }}
            </template>
          </el-table-column>

          <el-table-column label="Duration" width="100">
            <template #default="{ row }">
              {{
                moment(row.end_time).diff(moment(row.start_time), "hour")
              }}??????
            </template>
          </el-table-column>
        </el-table-column>

        <el-table-column prop="title" label="Title" />
        <el-table-column prop="tag_id" label="Tag" width="80">
          <template #default="{ row }">
            <el-tag> {{ tagName(row.tag_id) }} </el-tag>
          </template>
        </el-table-column>
        <el-table-column fixed="right" label="Action" width="180">
          <template #default="{ row }">
            <el-button
              type="primary"
              @click="
                () => {
                  onEdit(row);
                }
              "
            >
              ??????</el-button
            >
            <el-button
              type="danger"
              @click="
                () => {
                  onDelete(row.id);
                }
              "
              >??????</el-button
            >
          </template>
        </el-table-column>
      </el-table>
    </div>
    <EditDialog
      v-model="dialogVisible"
      :event="eventModel"
      @update="() => {
        loadData(+(props.uid as string))
      }"
    >
    </EditDialog>
    <el-dialog
      v-if="renderPanel"
      title="??????"
      width="90%"
      v-model="renderPanel"
    >
      <ImageRender :data="tableDataReverse" />
    </el-dialog>
  </main>
</template>

<style scoped>
.button {
  padding: 10px 20px;
  border: 1px solid #ddd;
  color: #333;
  background-color: #fff;
  border-radius: 4px;
  font-size: 14px;
  font-family: "????????????", arail;
  cursor: pointer;
}
</style>
