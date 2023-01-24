import { defineStore } from "pinia";
import request from "@/tools/axios";

interface ICharacter {
  id: number;
  name: string;
  uid: number;
  live_id: number;
}

export const useCharacterStore = defineStore("character", () => {
  const characters = ref<Array<ICharacter>>([]);

  const fetchAll = () => {
    request.get<ICharacter[]>("/characters").then((res) => {
      characters.value = res.data;
    });
  };

  return { characters, fetchAll };
});
