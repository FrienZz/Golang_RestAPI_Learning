<script setup>
import { ref, watch } from "vue";
import axios from "axios";
import Swal from "sweetalert2";

const emit = defineEmits(["submit"]);
const title = ref("");
const description = ref("");
const date = ref("");
const img_url = ref("");
const location = ref("");

defineProps({
  isOpen: Boolean,
});

const submitForm = async () => {
  try {
    const token = localStorage.getItem("token");
    const response = await axios.post(
      "http://localhost:8080/events",
      {
        title: title.value,
        description: description.value,
        image_url: img_url.value,
        event_date: getFormattedDate(),
        location: location.value,
      },
      {
        headers: {
          Authorization: token,
        },
      }
    );
    emit("submit");
    Swal.fire({
      title: "Event created successfully!",
      icon: "success",
    });
  } catch (err) {
    console.log(err);
  }
};

const getFormattedDate = () => {
  return new Date(date.value).toISOString();
};
</script>

<template>
  <form class="form-group" @submit.prevent="submitForm">
    <label for="title">Title</label><br />
    <input
      type="text"
      id="title"
      name="title"
      v-model="title"
      placeholder="Title"
      required
    /><br />
    <label for="description">Description</label><br />
    <textarea
      id="description"
      name="description"
      rows="5"
      cols="50"
      v-model="description"
      placeholder="Description"
    ></textarea
    ><br />
    <label for="date">Date</label><br />
    <input type="date" id="date" name="date" v-model="date" required /><br />
    <label for="img_url">Image Url</label><br />
    <input
      type="text"
      id="img_url"
      name="img_url"
      v-model="img_url"
      placeholder="Image Url"
      required
    /><br />
    <label for="location">Location</label><br />
    <input
      type="text"
      id="location"
      name="location"
      v-model="location"
      placeholder="Location"
      required
    /><br />
    <button type="submit">Submit</button>
  </form>
</template>

<style scoped>
.form-group {
  position: absolute;
  top: 50%;
  left: 50%;
  transform: translate(-50%, -50%);
  background-color: white;
  border-radius: 10px;
  padding: 40px;
  box-shadow: 15px 15px 30px #bebebe;
}

label {
  margin-left: 10px;
  font-weight: bold;
}

input {
  width: 385px;
  height: 40px;
  border-radius: 5px;
  margin: 10px 0 15px;
  font-size: 15px;
  padding-left: 10px;
  box-shadow: 2px 2px 2px 1px rgb(0 0 0 / 15%);
  border: 1px solid #ccc;
  transition: border-color 0.5s ease-out;
}

input:focus {
  border: 2px solid #3d74b6;
  outline: none;
  background-color: #f2f9ff;
}

textarea {
  margin: 10px 0 15px;
  padding: 10px 0 0 10px;
  border-radius: 5px;
  font-size: 15px;
  box-shadow: 2px 2px 2px 1px rgb(0 0 0 / 15%);
  border: 1px solid #ccc;
  transition: border-color 0.5s ease-out;
}

textarea:focus {
  border: 2px solid #3d74b6;
  outline: none;
  background-color: #f2f9ff;
}

button {
  width: 400px;
  height: 45px;
  border-radius: 5px;
  margin: 20px 0 0;
  font-size: 15px;
  background-color: #4a9782;
  color: white;
  font-weight: bold;
  box-shadow: 2px 2px 2px 1px rgb(0 0 0 / 15%);
  border: 1px solid #ccc;
  cursor: pointer;
}

button:hover {
  background-color: #001831;
  border: 2.5px solid white;
}
</style>
