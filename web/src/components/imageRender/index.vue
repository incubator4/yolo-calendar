<script setup lang="ts">
import { groupBy } from "lodash";
import moment from "moment";
import { useScreen } from "vue-screen";
import type konva from "konva";
const props = defineProps<{ data: ICalendar[] }>();
const image2 = ref<HTMLImageElement>();
const screen = useScreen();

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

const load = () => {
  const image = new window.Image();

  image.setAttribute("crossOrigin", "anonymous");
  image.src =
    "https://yolo-1256553639.cos.ap-shanghai.myqcloud.com/calendars/d7d83d42a87ed0ce78306cdc764748fe.jpg";
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

onMounted(() => {
  load();
  fontSize.value -= 1;
});

const fontSize = ref(56);
const fontFamily = ref("灵动指书");
const fontColor = ref("#E77279");
const fontStyle = ref<Array<String>>([]);

const curFontStyle = computed(() => {
  const style = fontStyle.value;
  return style && style.length > 0 ? style.join("   ") : "normal";
});

const offset = reactive({
  x: 850,
  y: 50,
});

const space = reactive({
  x: 300,
  y: 138,
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
      <el-slider v-model="percent" @change="resize" show-input />
    </el-header>
    <el-container>
      <el-aside width="200px">
        <el-collapse v-model="activeNames" @change="handleChange">
          <el-collapse-item title="字体设置" name="1">
            <el-form-item label="字体">
              <el-select v-model="fontFamily" placeholder="">
                <el-option label="灵动指书" value="灵动指书" />
              </el-select>
            </el-form-item>
            <el-form-item label="字号">
              <el-input-number v-model="fontSize" placeholder="" />
            </el-form-item>
            <el-form-item label="颜色">
              <el-color-picker v-model="fontColor" show-alpha />
            </el-form-item>
            <el-form-item label="样式">
              <el-checkbox-group v-model="fontStyle">
                <el-checkbox label="italic" />
                <el-checkbox label="bold" />
              </el-checkbox-group>
            </el-form-item>
          </el-collapse-item>
          <el-collapse-item title="间距设置" name="2">
            <el-form-item label="横向偏移">
              <el-input-number size="small" v-model="offset.x" placeholder="" />
            </el-form-item>
            <el-form-item label="纵向偏移">
              <el-input-number size="small" v-model="offset.y" placeholder="" />
            </el-form-item>
            <el-form-item label="横向间距">
              <el-input-number size="small" v-model="space.x" placeholder="" />
            </el-form-item>
            <el-form-item label="纵向间距">
              <el-input-number size="small" v-model="space.y" placeholder="" />
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
            <v-layer :config="{ ...offset }">
              <v-group
                v-for="(dayCal, index) in computedData"
                :config="{
                  x: 0,
                  y: space.y * +index,
                }"
              >
                <v-text
                  v-for="(cal, i) in computedData[index]"
                  :config="{
                    text: timeFormat(cal.start_time) + cal.title,
                    fontSize,
                    fontFamily,
                    fontStyle: curFontStyle,
                    fill: fontColor,
                    x: space.x * i,
                    y: 0,
                  }"
                ></v-text>
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
            <v-layer :config="{ ...offset }">
              <v-group
                v-for="(dayCal, index) in computedData"
                :config="{
                  x: 0,
                  y: space.y * +index,
                }"
              >
                <v-text
                  v-for="(cal, i) in computedData[index]"
                  :config="{
                    text: timeFormat(cal.start_time) + cal.title,
                    fontSize,
                    fontFamily,
                    fontStyle: curFontStyle,
                    fill: fontColor,
                    x: space.x * i,
                    y: 0,
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
