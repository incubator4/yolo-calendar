<script setup lang="ts">
import { Calendar } from "v-calendar";
const month = new Date().getMonth();
const year = new Date().getFullYear();
const masks = {
  weekdays: "WWW",
};
const attributes = [
  {
    key: 1,
    customData: {
      title: "Lunch with mom.",
      class: "bg-red-600 text-white",
    },
    dates: new Date(year, month, 1),
  },
  {
    key: 2,
    customData: {
      title: "Take Noah to basketball practice",
      class: "bg-blue-500 text-white",
    },
    dates: new Date(year, month, 2),
  },
  {
    key: 3,
    customData: {
      title: "Noah's basketball game.",
      class: "bg-blue-500 text-white",
    },
    dates: new Date(year, month, 5),
  },
  {
    key: 4,
    customData: {
      title: "Take car to the shop",
      class: "bg-indigo-500 text-white",
    },
    dates: new Date(year, month, 5),
  },
  {
    key: 4,
    customData: {
      title: "Meeting with new client.",
      class: "bg-teal-500 text-white",
    },
    dates: new Date(year, month, 7),
  },
  {
    key: 5,
    customData: {
      title: "Mia's gymnastics practice.",
      class: "bg-pink-500 text-white",
    },
    dates: new Date(year, month, 11),
  },
  {
    key: 6,
    customData: {
      title: "Cookout with friends.",
      class: "bg-orange-500 text-white",
    },
    dates: { months: 5, ordinalWeekdays: { 2: 1 } },
  },
  {
    key: 7,
    customData: {
      title: "Mia's gymnastics recital.",
      class: "bg-pink-500 text-white",
    },
    dates: new Date(year, month, 22),
  },
  {
    key: 8,
    customData: {
      title: "Visit great grandma.",
      class: "bg-red-600 text-white",
    },
    dates: new Date(year, month, 25),
  },
];
</script>

<template>
  <div class="text-center section">
    <Calendar :masks="masks" :attributes="attributes"> </Calendar>
  </div>
</template>
