<script setup>
defineProps({
  title: String,
  description: String,
  date: String,
  imgUrl: String,
  location: String,
});
</script>

<template>
  <div class="card">
    <div class="img" :style="{ backgroundImage: `url(${imgUrl})` }"></div>
    <div class="title">{{ title }}</div>
    <div class="description">{{ description }}</div>
    <div class="date">
      <i class="fa-regular fa-calendar-days" style="color: #a5aeb6"></i>
      {{ date }}
    </div>
    <div class="location">
      <i class="fa-solid fa-location-dot" style="color: #a5aeb6"></i>
      {{ location }}
    </div>
  </div>
</template>

<style scoped>
.card {
  width: 250px;
  height: 350px;
  background: white;
  border-radius: 30px;
  box-shadow: 15px 15px 30px #bebebe;
  transition: 0.2s ease-in-out;
  margin: 25px 25px;
}

.card:hover {
  cursor: pointer;
  box-shadow: 0px 10px 20px rgba(0, 0, 0, 0.1);
}

.img {
  width: 100%;
  height: 40%;
  background-size: cover;
  background-position: center;
  border-radius: 30px 30px 0px 0px;
}

.title {
  margin: 20px 20px 0px;
  font-size: 16px;
  font-weight: bold;
  overflow: hidden;
  text-overflow: ellipsis;
  white-space: nowrap;
}

.description {
  display: -webkit-box;
  -webkit-line-clamp: 3;
  -webkit-box-orient: vertical;
  margin: 10px 20px 0px;
  text-overflow: ellipsis;
  overflow: hidden;
  color: #787a91;
  font-size: 14px;
}

.date {
  margin: 20px 20px 0px;
  font-size: 14px;
}

.date i {
  margin-right: 6px;
}

.location {
  margin: 20px 20px 0px;
  font-size: 14px;
}

.location i {
  margin-right: 6px;
}
</style>
