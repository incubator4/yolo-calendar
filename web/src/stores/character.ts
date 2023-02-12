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
  start_time: Date;
  end_time: Date;
  cid: number;
}

type CharacterCalendar = Omit<ICalendar & ICharacter, "id">;

export const useCharacterStore = defineStore("character", () => {
  const characters = ref<Array<ICharacter>>([]);

  const calendars = ref<Array<CharacterCalendar>>([]);

  const fetchAll = () => {
    request.get<ICharacter[]>("/characters").then((res) => {
      characters.value = res.data;
    });
  };

  const getCalendar = (uid: number) => {
    return request.get<CharacterCalendar[]>(`/cal/${uid}`).then((res) => {
      calendars.value = res.data;
    });
  };

  const listCalendar = () => {
    return request.get<CharacterCalendar[]>(`/cal`).then((res) => {
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
