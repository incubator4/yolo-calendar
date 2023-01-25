import { defineStore } from "pinia";
import request from "@/tools/axios";

interface ICharacter {
  id: number;
  name: string;
  uid: number;
  live_id: number;
}

interface ICalendar {
  title: string;
  dateTime: Date;
  cid: number;
}

export const useCharacterStore = defineStore("character", () => {
  const characters = ref<Array<ICharacter>>([]);

  const calendars = ref<Array<ICalendar>>([]);

  const fetchAll = () => {
    request.get<ICharacter[]>("/characters").then((res) => {
      characters.value = res.data;
    });
  };

  const getCalendar = (uid: number) => {
    return request.get<ICalendar[]>(`/cal/${uid}`).then((res) => {
      calendars.value = res.data;
    });
  };

  const listCalendar = () => {
    return request.get<ICalendar[]>(`/cal`).then((res) => {
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
