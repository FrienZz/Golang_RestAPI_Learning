<script setup>
import EventCard from "./EventCard.vue";
import { onMounted, ref } from "vue";
import axios from "axios";

const events = ref([]);

const getEvents = async () => {
  try {
    const token = localStorage.getItem("token");
    const response = await axios.get("http://localhost:8080/events", {
      headers: {
        Authorization: token,
      },
    });
    events.value = response.data;
  } catch (err) {
    console.log(err);
  }
};
onMounted(() => {
  getEvents();
});
</script>

<template>
  <div class="container">
    <EventCard
      v-for="event in events"
      :key="event.id"
      imgUrl="https://cdn-s.cdprojektred.com/concert/gallery/9a9a114aa2a4cf8681ac5ec602c9d2e3_q90_1920x1280.jpeg"
      :title="event.name"
      :description="event.description"
    />
  </div>
</template>

<style scoped>
.container {
  display: flex;
  justify-content: flex-start;
  margin: 80px 100px;
  flex-wrap: wrap;
}
</style>
