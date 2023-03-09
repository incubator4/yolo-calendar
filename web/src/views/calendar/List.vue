<script setup lang="ts">
import { useCalendarStore } from "@/stores";
import { groupBy } from "lodash";
import moment from "moment";
const store = useCalendarStore();
store.listCalendar({
  end: moment().format("yyyy-MM-DD"),
});
</script>
<template></template>

<style></style>
