<script setup lang="ts">
import WelcomeItem from "./WelcomeItem.vue";
import Avatar from "./icons/Avatar.vue";
import Calendar from "./Calendar.vue";

import { useCharacterStore } from "@/stores";

const store = useCharacterStore();

store.fetchAll();
</script>

<template>
  <WelcomeItem v-for="character in store.characters">
    <template #icon>
      <Avatar :uid="character.uid" />
    </template>
    <template #heading>
      <div>
        {{ character.name }}
        <Calendar />
      </div>
    </template>
  </WelcomeItem>
</template>
