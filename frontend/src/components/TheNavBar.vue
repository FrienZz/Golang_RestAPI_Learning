<script setup>
const handleLogout = () => {
  localStorage.removeItem("token");
};
</script>

<template>
  <div class="navbar">
    <router-link to="/logout" @click="handleLogout">Logout</router-link>
  </div>
</template>

<style scoped>
.navbar {
  display: flex;
  justify-content: flex-end;
  background-color: #001831;
  padding: 0px 30px;
}

.navbar a {
  padding: 14px 16px;
  text-decoration: none;
  color: white;
  font-size: 20px;
  font-weight: bold;
}

.navbar a:hover {
  background-color: #3d74b6;
}
</style>
