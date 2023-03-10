import { defineStore } from "pinia";
import request from "@/tools/axios";
import qs from "qs";

export const useImageRenderConfig = defineStore("image-redner-config", () => {
  const configs = ref<Array<ImageRenderConfig>>([]);

  const list = (owner: number) =>
    request
      .get<Array<ImageRenderConfig>>(`/image_render`, {
        params: {
          owner,
        },
      })
      .then((res) => {
        configs.value = res.data;
      });

  return {
    configs,
    list,
  };
});
