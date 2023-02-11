<script setup lang="ts">
import { useCharacterStore } from "@/stores";
import Avatar from "@/components/icons/Avatar.vue";
import { useScreen } from "vue-screen";
import colorMatrix from "@/tools/color";

const qqGroup = [781289057, 659081677, 374946805, 298018245, 736532313];

const size = ref("");
const screen = useScreen();
const store = useCharacterStore();
store.fetchAll();
</script>
<template>
  <div :style="{ height: `${screen.height}px`, width: 'auto' }">
    <el-scrollbar>
      <el-row>
        <el-col
          :span="24"
          v-for="{ uid, name, id, live_id } in store.characters"
          :key="id"
          class="scrollbar-demo-item"
          :style="{
            color: colorMatrix(id),
            border: `1px solid ${colorMatrix(id)}`,
          }"
        >
          <el-container>
            <el-header height="60px" style="margin-bottom: 10px; padding: 0">
              <el-container style="padding: 20px">
                <el-aside
                  style="margin-right: 20px"
                  v-if="screen.width > 800"
                  width="60px"
                >
                  <Avatar class="avatar" :uid="uid"></Avatar>
                </el-aside>
                <el-main style="padding: 0">
                  <h1
                    :style="{
                      background: colorMatrix(id),
                      textAlign: 'center',
                      fontSize: '36px',
                      borderRadius: '8px',
                    }"
                    class="title"
                  >
                    {{ name }}
                  </h1>
                </el-main>
              </el-container>
            </el-header>
            <el-main>
              <el-descriptions
                class="margin-top"
                :column="1"
                :size="size"
                border
              >
                <el-descriptions-item>
                  <template #label>
                    <div class="cell-item">个人主页</div>
                  </template>
                  <div class="cell-context">
                    <el-link
                      :href="`https://space.bilibili.com/${uid}`"
                      style="font-size: 18px"
                      target="_blank"
                      >{{ uid }}</el-link
                    >
                  </div>
                </el-descriptions-item>
                <el-descriptions-item>
                  <template #label>
                    <div class="cell-item">直播间</div>
                  </template>
                  <div class="cell-context">
                    <el-link
                      :href="`https://live.bilibili.com/${live_id}`"
                      style="font-size: 18px"
                      target="_blank"
                      >{{ live_id }}</el-link
                    >
                  </div>
                </el-descriptions-item>
                <el-descriptions-item>
                  <template #label>
                    <div class="cell-item">生日</div>
                  </template>
                  <div class="cell-context">-</div>
                </el-descriptions-item>
                <el-descriptions-item>
                  <template #label>
                    <div class="cell-item">Remarks</div>
                  </template>
                  <div class="cell-context">
                    <el-tag size="small">Tag</el-tag>
                  </div>
                </el-descriptions-item>
                <el-descriptions-item>
                  <template #label>
                    <div class="cell-item">粉丝群</div>
                  </template>
                  <div class="cell-context">
                    {{ qqGroup[id - 1] }}
                  </div>
                </el-descriptions-item>
              </el-descriptions>
            </el-main>
          </el-container>
        </el-col>
      </el-row>
    </el-scrollbar>
  </div>
</template>
<style scoped>
.avatar {
  width: 60px;
  height: 60px;
}

.title {
  height: 60px;
  padding: 0;
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
  width: 100%;
  max-width: 400px;
  margin: 10px;
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
</style>
