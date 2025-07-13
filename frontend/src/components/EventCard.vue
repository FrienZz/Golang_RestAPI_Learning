<script setup>
defineProps({
  title: String,
  description: String,
  imgUrl: String,
});
</script>

<template>
  <div class="card">
    <div class="img" :style="{ backgroundImage: `url(${imgUrl})` }"></div>
    <div class="title">{{ title }}</div>
    <div class="description">{{ description }}</div>
  </div>
</template>

<style scoped>
.card {
  width: 350px;
  height: 300px;
  background: white;
  border-radius: 30px;
  box-shadow: 15px 15px 30px #bebebe, -15px -15px 30px #ffffff;
  transition: 0.2s ease-in-out;
  margin: 25px 45px;
  transition: 0.2s ease-in-out;
}

.card:hover {
  cursor: pointer;
  box-shadow: 0px 10px 20px rgba(0, 0, 0, 0.1);
}

.img {
  width: 100%;
  height: 50%;
  background-size: cover;
  background-position: center;
  border-radius: 30px 30px 0px 0px;
}

.title {
  margin: 20px 20px 0px;
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
}
</style>
