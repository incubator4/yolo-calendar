<script setup lang="ts">
import { groupBy } from "lodash";
import moment from "moment";
import { useScreen } from "vue-screen";
import { useImageRenderConfig } from "@/stores";
import type konva from "konva";
const props = defineProps<{ data: ICalendar[] }>();
const image2 = ref<HTMLImageElement>();
const screen = useScreen();

const configStore = useImageRenderConfig();

const stage = ref();

const stageSize = reactive({
  width: 1920,
  height: 1180,
});

const percent = ref(20);

const computeSize = computed(() => {
  const ratio = (stageSize.width * percent.value) / 50 / stageSize.width;
  return {
    width: stageSize.width * ratio,
    height: stageSize.height * ratio,
    scale: {
      x: ratio,
      y: ratio,
    },
  };
});

const resize = () => {
  const width = (stageSize.width * percent.value) / 50;
  const ratio = width / stageSize.width;
};

const render = ref(false);

const computedData = computed(() =>
  groupBy(
    props.data
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

const timeFormat = (d: Date) => {
  const m = moment(d);
  return m.format("HH:mm");
};

const load = () => {
  const image = new window.Image();

  image.setAttribute("crossOrigin", "anonymous");
  image.src = currentConfig.value.image.url;
  stageSize.height = currentConfig.value.image.height;
  stageSize.width = currentConfig.value.image.width;
  image.onload = () => {
    // set image only when it is loaded
    const width = image.width;
    const ratio = width / image.width;
    image.width = width;
    image.height = image.height * ratio;
    image2.value = image;
    render.value = true;
  };
};

onMounted(() => {
  // load();
});

let currentConfig = ref<ImageRenderConfig>({
  name: "",
  image: {
    url: "",
    width: 0,
    height: 0,
  },
  text_offset: {
    x: 0,
    y: 0,
  },
  text_group: [[]],
  text_group_offset: {
    x: 0,
    y: 0,
  },
  row: {
    x: 0,
    y: 0,
    prefix: "",
    suffix: "",
  },
  col: {
    x: 0,
    y: 0,
    prefix: "",
    suffix: "",
  },
  font: {
    family: "",
    size: 0,
    color: "",
    style: "",
    layout: "",
  },
});

const currentConfigId = ref();

const configChange = (id: number) => {
  const config = configStore.configs.find((c) => c.id === id);
  if (config) {
    currentConfig.value = { ...config };
  }
  load();
};

const fontStyle = ref<Array<String>>([]);

const curFontStyle = computed(() => {
  const style = fontStyle.value;
  return style && style.length > 0 ? style.join("   ") : "normal";
});

const activeNames = ref(["1"]);
const handleChange = (val: string[]) => {
  console.log(val);
};

const onSave = () => {
  const s = stage.value.getNode() as unknown as konva.Stage;
  const dataURL = s.toDataURL();
  const link = document.createElement("a");
  link.download = "image.png";
  link.href = dataURL;
  document.body.appendChild(link);
  link.click();

  //   stage.value?.toImage();
};
</script>
<template>
  <el-container>
    <el-header>
      <el-row>
        <el-col :span="6">
          <el-select
            v-model="currentConfigId"
            @change="configChange"
            placeholder=""
          >
            <el-option
              v-for="config in configStore.configs"
              :label="config.name"
              :value="config.id"
            >
            </el-option>
          </el-select>
        </el-col>
        <el-col :span="0"></el-col>
        <el-col :span="16">
          <el-slider v-model="percent" @change="resize" show-input />
        </el-col>
      </el-row>
    </el-header>
    <el-container>
      <el-aside width="220px">
        <el-collapse v-model="activeNames" @change="handleChange">
          <el-collapse-item title="字体设置" name="1">
            <el-form-item label="字体">
              <el-select v-model="currentConfig.font.family" placeholder="">
                <el-option label="默认" value="sans-serif"></el-option>
                <el-option label="灵动指书" value="灵动指书" />
              </el-select>
            </el-form-item>
            <el-form-item label="字号">
              <el-input-number
                style="width: 100%"
                v-model="currentConfig.font.size"
                placeholder=""
              />
            </el-form-item>
            <el-row>
              <el-col :span="12">
                <el-form-item label="颜色">
                  <el-color-picker
                    v-model="currentConfig.font.color"
                    show-alpha
                  />
                </el-form-item>
              </el-col>
              <el-col :span="12">
                <el-form-item v-show="false" label="方向">
                  <el-select v-model="currentConfig.font.layout" placeholder="">
                    <el-option label="横向" value="vertical"></el-option>
                    <el-option label="纵向" value="horizontal"></el-option>
                  </el-select>
                </el-form-item>
              </el-col>
            </el-row>
            <el-form-item label="样式">
              <el-checkbox-group v-model="fontStyle">
                <el-checkbox label="italic" />
                <el-checkbox label="bold" />
              </el-checkbox-group>
            </el-form-item>
          </el-collapse-item>
          <el-collapse-item title="文本设置" name="2">
            <template #title> 文本设置 </template>
            <el-form-item label="横向偏移">
              <el-input-number
                size="small"
                v-model="currentConfig.text_offset.x"
                placeholder=""
              />
            </el-form-item>
            <el-form-item label="纵向偏移">
              <el-input-number
                size="small"
                v-model="currentConfig.text_offset.y"
                placeholder=""
              />
            </el-form-item>
            <el-form-item label="横向间距">
              <el-input-number
                size="small"
                v-model="currentConfig.text_group_offset.x"
                placeholder=""
              />
            </el-form-item>
            <el-form-item label="纵向间距">
              <el-input-number
                size="small"
                v-model="currentConfig.text_group_offset.y"
                placeholder=""
              />
            </el-form-item>
          </el-collapse-item>
          <el-collapse-item title="每日间距设置" name="3">
            <el-form-item label="横向间距">
              <el-input-number
                size="small"
                v-model="currentConfig.row.x"
                placeholder=""
              />
            </el-form-item>
            <el-form-item label="纵向间距">
              <el-input-number
                size="small"
                v-model="currentConfig.row.y"
                placeholder=""
              />
            </el-form-item>
          </el-collapse-item>
          <el-collapse-item title="事件间距设置" name="4">
            <el-form-item label="横向间距">
              <el-input-number
                size="small"
                v-model="currentConfig.col.x"
                placeholder=""
              />
            </el-form-item>
            <el-form-item label="纵向间距">
              <el-input-number
                size="small"
                v-model="currentConfig.col.y"
                placeholder=""
              />
            </el-form-item>
          </el-collapse-item>
        </el-collapse>
        <el-button type="primary" @click="onSave">保存</el-button>
      </el-aside>
      <el-main>
        <div style="text-align: center">
          <v-stage v-if="render" ref="previewStage" :config="computeSize">
            <v-layer ref="layer">
              <v-image
                ref="image"
                :config="{
                  image: image2,
                }"
              />
            </v-layer>
            <v-layer :config="{ ...currentConfig.text_offset }">
              <v-group
                v-for="(subGroup, gIndex) in currentConfig.text_group"
                :config="{
                  x: currentConfig.text_group_offset.x * gIndex,
                  y: currentConfig.text_group_offset.y * gIndex,
                }"
              >
                <v-group
                  v-for="(group, index) in subGroup"
                  :config="{
                    x: currentConfig.row.x * +index,
                    y: currentConfig.row.y * +index,
                  }"
                >
                  <v-text
                    v-for="(cal, i) in computedData[group]"
                    :config="{
                      text: timeFormat(cal.start_time) + cal.title,
                      fontSize: currentConfig.font.size,
                      fontFamily: currentConfig.font.family,
                      fontStyle: curFontStyle,
                      fill: currentConfig.font.color,
                      x: currentConfig.col.x * i,
                      y: currentConfig.col.y * i,
                    }"
                  ></v-text>
                </v-group>
              </v-group>
            </v-layer>
          </v-stage>
          <v-stage v-show="false" v-if="render" ref="stage" :config="stageSize">
            <v-layer ref="layer">
              <v-image
                ref="image"
                :config="{
                  image: image2,
                }"
              />
            </v-layer>
            <v-layer :config="{ ...currentConfig.text_offset }">
              <v-group
                v-for="(dayCal, index) in computedData"
                :config="{
                  x: currentConfig.row.x * +index,
                  y: currentConfig.row.y * +index,
                }"
              >
                <v-text
                  v-for="(cal, i) in computedData[index]"
                  :config="{
                    text: timeFormat(cal.start_time) + cal.title,
                    fontSize: currentConfig.font.size,
                    fontFamily: currentConfig.font.family,
                    fontStyle: curFontStyle,
                    fill: currentConfig.font.color,
                    x: currentConfig.col.x * i,
                    y: currentConfig.col.y * i,
                  }"
                ></v-text>
              </v-group>
            </v-layer>
          </v-stage>
        </div>
      </el-main>
    </el-container>
  </el-container>
</template>

<style scoped>
.title-tooltip {
  width: 110px;
}
</style>
