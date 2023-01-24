<script setup lang="ts">
import WelcomeItem from "./WelcomeItem.vue";
import DocumentationIcon from "./icons/IconDocumentation.vue";

import { useCharacterStore } from "@/stores";

const store = useCharacterStore();

store.fetchAll();
</script>

<template>
  <WelcomeItem style="margin: 100px" v-for="character in store.characters">
    <template #icon>
      <DocumentationIcon />
    </template>
    <template #heading>
      <div>
        {{ character.name }}
      </div>
    </template>
  </WelcomeItem>
</template>
