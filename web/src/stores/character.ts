import { defineStore } from "pinia";
import request from "@/tools/axios";
import qs from "qs";

export const useCharacterStore = defineStore("character", () => {
  const characters = ref<Array<ICharacter>>([]);

  const calendars = ref<Array<CharacterCalendar>>([]);

  const fetchAll = () => {
    return request.get<ICharacter[]>("/characters").then((res) => {
      characters.value = res.data;
    });
  };

  const getCalendar = (uid: number) => {
    return request.get<CharacterCalendar[]>(`/cal/${uid}`).then((res) => {
      calendars.value = res.data;
    });
  };

  const listCalendar = (params?: {
    start?: string;
    end?: string;
    uid?: Array<string>;
  }) => {
    return request
      .get<CharacterCalendar[]>(`/cal`, {
        params,
        paramsSerializer: {
          serialize: (params) => {
            return qs.stringify(params, { arrayFormat: "repeat" });
          },
        },
      })
      .then((res) => {
        calendars.value = res.data;
      });
  };

  const clearCalendar = () => {
    calendars.value = [];
  };

  return {
    characters,
    calendars,
    fetchAll,
    getCalendar,
    clearCalendar,
    listCalendar,
  };
});
