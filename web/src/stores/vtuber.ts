import { defineStore } from "pinia";
import request from "@/tools/axios";
import qs from "qs";

export const useVtuberStore = defineStore("vtuber", () => {
  const vtubers = ref<Array<IVtuber>>([]);

  const curentVtuber = ref<IVtuber>();

  const fetchAll = () => {
    return request.get<IVtuber[]>("/vtubers").then((res) => {
      vtubers.value = res.data;
    });
  };

  const getVtuber = (uid: number) => {
    return request.get<IVtuber>(`/vtubers/${uid}`).then((res) => {
      curentVtuber.value = res.data;
    });
  };

  return {
    vtubers,
    curentVtuber,
    fetchAll,
    getVtuber,
  };
});
