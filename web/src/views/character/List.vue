<script setup lang="ts">
import { useCharacterStore } from "@/stores";
import Avatar from "@/components/icons/Avatar.vue";
import { useScreen } from "vue-screen";
import useClipboard from "vue-clipboard3";
import colorMatrix from "@/tools/color";
import {
  Calendar,
  ChatDotRound,
  CopyDocument,
  More,
} from "@element-plus/icons-vue";

const qqGroup = [781289057, 659081677, 374946805, 298018245, 736532313];

const { toClipboard } = useClipboard();

const onClipboard = (uid: number) => {
  toClipboard("webcal://yolo.incubator4.com/api/ics/" + uid);
  alert("日历链接复制成功");
};

const onQQGroup = (group: number) => {
  toClipboard("" + group);
  alert(`粉丝群 ${group} 复制成功`);
};

const size = ref("");
const screen = useScreen();
const store = useCharacterStore();
store.fetchAll();
</script>
<template>
  <div :style="{ height: `${screen.height}px`, width: 'auto' }">
    <el-scrollbar>
      <el-row>
        <template
          v-for="{ uid, name, id, live_id } in store.characters"
          :key="id"
          class="scrollbar-demo-item"
          :style="{
            color: colorMatrix(id),
            border: `1px solid ${colorMatrix(id)}`,
          }"
        >
          <el-col :lg="12" :span="24">
            <el-card style="margin: 10px">
              <el-row>
                <el-col :span="8">
                  <a
                    :href="`https://space.bilibili.com/${uid}`"
                    style="font-size: 16px"
                    target="_blank"
                    class="title"
                    ><Avatar class="avatar" :uid="uid"></Avatar
                  ></a>

                  <el-button style="width: 120px" :color="colorMatrix(id)">
                    <a
                      :href="`https://live.bilibili.com/${live_id}`"
                      style="font-size: 16px"
                      target="_blank"
                      class="title"
                      >{{ name }}</a
                    ></el-button
                  >
                </el-col>
                <el-col :span="16">
                  <ul class="infinite-list" style="overflow: auto">
                    <li class="scrollbar-demo-item">
                      <el-button
                        class="button"
                        :icon="ChatDotRound"
                        @click="() => onQQGroup(qqGroup[id - 1])"
                        >粉丝群</el-button
                      >
                    </li>

                    <li class="scrollbar-demo-item">
                      <el-button class="button">
                        <a
                          :href="`webcal://yolo.incubator4.com/api/ics/${uid}`"
                          style="margin: 10px; color: black"
                        >
                          <el-icon><Calendar /></el-icon>
                          订阅
                        </a>
                      </el-button>
                    </li>
                    <li class="scrollbar-demo-item">
                      <el-button
                        class="button"
                        :icon="CopyDocument"
                        @click="() => onClipboard(uid)"
                      >
                        复制
                      </el-button>
                    </li>
                    <li class="scrollbar-demo-item">
                      <el-button class="button" :icon="More"
                        >更多信息</el-button
                      >
                    </li>
                  </ul>
                </el-col>
              </el-row>
            </el-card>
          </el-col>
        </template>
      </el-row>
    </el-scrollbar>
  </div>
</template>
<style scoped>
.avatar {
  width: 120px;
  height: 120px;
}

.title {
  /* height: 60px; */
  /* padding: 0; */
  color: white;
}
.scrollbar-demo-item {
  border-radius: 4px;
  display: flex;
  align-items: center;
  justify-content: center;
  text-align: center;
  /* height: 180px; */
  /* min-width: 350px; */
  max-width: 350px;
  margin: 5px;
  text-align: center;
  border-radius: 4px;
  /* background: var(--el-color-primary-light-9); */
  color: var(--el-color-primary);
}

.cell-item {
  display: flex;
  min-width: 60px;
  text-align: center;
  align-items: center;
}
.cell-context {
  text-align: center;
}
.button {
  width: 180px;
}
</style>
