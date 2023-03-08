import { defineStore } from "pinia";
import request from "@/tools/axios";
import qs from "qs";

export const useCharacterStore = defineStore("character", () => {
  const characters = ref<Array<ICharacter>>([]);

  const calendars = ref<Array<CharacterCalendar>>([]);

  const curentCharacter = ref<ICharacter>();

  const tags = ref<Array<ITag>>([]);

  const fetchAll = () => {
    return request.get<ICharacter[]>("/characters").then((res) => {
      characters.value = res.data;
    });
  };

  const getCharacter = (uid: number) => {
    return request.get<ICharacter>(`/characters/${uid}`).then((res) => {
      curentCharacter.value = res.data;
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

  const listTags = () => {
    return request.get<ITag[]>("/event_tags").then((res) => {
      tags.value = res.data;
    });
  };

  const clearCalendar = () => {
    calendars.value = [];
  };

  return {
    characters,
    calendars,
    tags,
    curentCharacter,
    fetchAll,
    getCalendar,
    getCharacter,
    clearCalendar,
    listCalendar,
    listTags,
  };
});
