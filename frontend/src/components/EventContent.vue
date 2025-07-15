<script setup>
import EventCard from "./EventCard.vue";
import EventForm from "./EventForm.vue";
import { onMounted, ref } from "vue";
import axios from "axios";

const events = ref([]);
const openForm = ref(false);

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

const toggleForm = () => {
  openForm.value = !openForm.value;
};

const handleFormSubmit = () => {
  toggleForm();
  getEvents();
};

onMounted(() => {
  getEvents();
});
</script>

<template>
  <div>
    <div class="button-container">
      <button class="add-event" @click="toggleForm">+ Add event</button>
    </div>
    <transition name="form-popup">
      <EventForm
        :isOpen="openForm"
        v-if="openForm"
        @submit="handleFormSubmit"
      />
    </transition>
    <div class="event-container">
      <EventCard
        v-for="event in events"
        :key="event.id"
        :title="event.title"
        :description="event.description"
        :date="event.event_date.split('T')[0]"
        :imgUrl="event.image_url"
        :location="event.location"
      />
    </div>
  </div>
</template>

<style scoped>
.event-container {
  display: flex;
  justify-content: flex-start;
  margin: 10px 10px 80px 300px;
  flex-wrap: wrap;
}

.button-container {
  display: flex;
  justify-content: flex-end;
  margin: 30px 60px 0px 300px;
}

.add-event {
  font-weight: bold;
  font-size: 16px;
  border-radius: 5px;
  border: 1px solid;
  cursor: pointer;
  background-color: #3d74b6;
  color: white;
  padding: 10px;
}

.add-event:hover {
  background-color: #5a96d6;
}

.form-popup-enter-active,
.form-popup-leave-active {
  transition: opacity 0.3s ease, transform 0.3s ease;
}
.form-popup-enter-from,
.form-popup-leave-to {
  opacity: 0;
  transform: translate(-50%, -60%) scale(0.95);
}
.form-popup-enter-to,
.form-popup-leave-from {
  opacity: 1;
  transform: translate(-50%, -50%) scale(1);
}
</style>
