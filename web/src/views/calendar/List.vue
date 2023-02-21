<script setup lang="ts">
import { useCharacterStore } from "@/stores";
import { groupBy } from "lodash";
import moment from "moment";
const store = useCharacterStore();
store.listCalendar({
  end: moment().format("yyyy-MM-DD"),
});
</script>
<template></template>

<style></style>
