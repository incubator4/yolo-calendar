<script setup lang="ts">
import type {
  CalendarOptions,
  CalendarApi,
  EventApi,
  DateSelectArg,
  EventClickArg,
} from "@fullcalendar/core";
import FullCalendar from "@fullcalendar/vue3";
import dayGridPlugin from "@fullcalendar/daygrid";
import timeGridPlugin from "@fullcalendar/timegrid";
import listPlugin from "@fullcalendar/list";
import interactionPlugin from "@fullcalendar/interaction";
import zhCn from "@fullcalendar/core/locales/zh-cn";
import { INITIAL_EVENTS, createEventId } from "./event-util";

import "./style.css";

const props = defineProps<{ calendars: Array<CharacterCalendar> }>();

const fullCalendar = ref();
const currentEvents = ref<EventApi[]>([]);

watch(
  () => props.calendars,
  () => {
    const api: CalendarApi = fullCalendar.value.getApi();
    api.removeAllEvents();
    props.calendars.forEach((event) => {
      api.addEvent({
        id: event.id + "  ",
        title: event.title,
        start: event.start_time,
        end: event.end_time,
        color: event.main_color,
        name: event.name,
      });
    });
  }
);

const handleWeekendsToggle = () => {
  calendarOptions.value.weekends = !calendarOptions.value.weekends; // update a property
};
const handleDateSelect = (selectInfo: DateSelectArg) => {
  let title = prompt("Please enter a new title for your event");
  let calendarApi = selectInfo.view.calendar;

  calendarApi.unselect(); // clear date selection

  if (title) {
    calendarApi.addEvent({
      id: createEventId(),
      title,
      start: selectInfo.startStr,
      end: selectInfo.endStr,
      allDay: selectInfo.allDay,
    });
  }
};
const handleEventClick = (clickInfo: EventClickArg) => {
  if (
    confirm(
      `Are you sure you want to delete the event '${clickInfo.event.title}'`
    )
  ) {
    clickInfo.event.remove();
  }
};
const handleEvents = (events: EventApi[]) => {
  currentEvents.value = events;
};

const calendarOptions = ref<CalendarOptions>({
  locale: zhCn,
  plugins: [
    dayGridPlugin,
    timeGridPlugin,
    listPlugin,
    interactionPlugin, // needed for dateClick
  ],
  headerToolbar: {
    left: "prev,today,next",
    center: "title",
    right: "timeGridWeek,listDay",
    // right: "dayGridMonth,timeGridWeek,timeGridDay",
  },
  height: "auto",
  allDaySlot: false,
  slotDuration: "01:00",
  slotMinTime: "08:00",
  slotMaxTime: "23:59",
  initialView: "listDay",
  // initialEvents: INITIAL_EVENTS, // alternatively, use the `events` setting to fetch from a feed
  editable: true,
  selectable: true,
  selectMirror: true,
  dayMaxEvents: true,
  weekends: true,
  // select: handleDateSelect,
  // eventClick: handleEventClick,
  // eventsSet: handleEvents,
  /* you can update a remote database when these fire:
        eventAdd:
        eventChange:
        eventRemove:
        */
});
</script>

<template>
  <el-card>
    <FullCalendar
      ref="fullCalendar"
      class="demo-app-calendar"
      :options="calendarOptions"
    >
      <template v-slot:eventContent="arg">
        <p>{{ arg.timeText }}</p>

        <p style="font-weight: bold">{{ arg.event.extendedProps.name }}</p>
        <b>{{ arg.event.title }}</b>
      </template>
    </FullCalendar>
  </el-card>
</template>

<style scoped lang="css"></style>
