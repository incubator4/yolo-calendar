<script setup lang="ts">
import { useCharacterStore } from "@/stores";
import { Calendar } from "v-calendar";
import { useDark } from "@vueuse/core";
import Dot from "@/components/Dot.vue";
import { useScreen } from "vue-screen";
import useClipboard from "vue-clipboard3";
import Avatar from "@/components/icons/Avatar.vue";
const props = defineProps({ uid: String });

const store = useCharacterStore();
const screen = useScreen();
const isDark = useDark();

const loadData = (uid: number) => {
  ics.value = "webcal://yolo.incubator4.com/api/ics/" + props.uid;
  store.clearCalendar();
  store.listCalendar({ uid: [uid.toString()] });
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

const today = computed(() => {
  if (store.calendars && store.calendars.length > 0) {
    return [
      {
        key: "today",
        highlight: {
          fillMode: "light",
          style: {
            backgroundColor: "black",
          },
        },
        dates: new Date(),
      },
    ];
  } else {
    return [];
  }
});

const attrs = computed(() => [
  ...store.calendars.map((c) => {
    const dates = new Date(c.start_time);
    return {
      key: `${c.start_time} - ${c.title}`,
      dates,
      popover: {
        label: c.title,
        visibility: "focus",
      },
      customData: {
        title: c.title,
        time: dates.getHours(),
        color: c.main_color,
      },
      dot: {
        style: {
          marginTop: "10px",
          backgroundColor: c.main_color,
        },
      },
    };
  }),
  ...today.value,
]);

const ics = ref("");
const { toClipboard } = useClipboard();
const onClipboard = () => {
  toClipboard(ics.value);
  alert("复制成功");
};
</script>

<template>
  <main>
    <div style="padding: 40px">
      <el-row>
        <el-col style="flex: 0 1 400px">
          <Avatar :uid="+(uid  as  string)" />
          <el-skeleton :row="10"> </el-skeleton>
        </el-col>
        <el-col :span="2"> </el-col>
        <el-col style="flex: auto">
          <Calendar
            :columns="1"
            :rows="1"
            :is-dark="isDark"
            :attributes="attrs"
          >
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
                  {{ zeroPad(attr.customData.time, 2) }} -
                  {{ attr.customData.title }}
                </div>
              </div>
            </template>
          </Calendar>
        </el-col>
      </el-row>
    </div>
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
  font-family: "微软雅黑", arail;
  cursor: pointer;
}
</style>
