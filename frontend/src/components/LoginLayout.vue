<script setup></script>

<template>
  <div class="background">
    <div class="card">
      <img src="../img/loginImg.png" alt="Login" width="400" height="350" />
      <p>
        Create your event and invite your friend easily with
        <strong>Eventra</strong>.
      </p>
    </div>
  </div>
</template>

<style scoped>
.background {
  position: relative;
  background-image: url("../img/eventBg.png");
  height: 100vh;
  width: 60vw;
  font-size: 28px;
  color: #1b1b1b;
  text-align: center;
}

.card {
  position: absolute;
  top: 50%;
  left: 50%;
  transform: translate(-50%, -50%);
}

@media screen and (max-width: 992px) {
  .background {
    display: none;
  }
}
</style>
