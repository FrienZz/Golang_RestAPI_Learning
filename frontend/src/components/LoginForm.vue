<script setup>
import { ref } from "vue";
import axios from "axios";
import { useRouter } from "vue-router";
import Swal from "sweetalert2";

const router = useRouter();
const emit = defineEmits(["getToken"]);
const email = ref("");
const password = ref("");
const confirmPassword = ref("");
const errMessage = ref("");
const isSignUp = ref(false);

const submitForm = async () => {
  try {
    const path = isSignUp.value ? "register" : "login";
    if (isSignUp.value && confirmPassword.value != password.value) {
      throw new Error("Passwords do not match.");
    }
    const response = await axios.post(`http://localhost:8080/${path}`, {
      email: email.value,
      password: password.value,
    });
    resetForm();
    if (isSignUp.value) {
      Swal.fire({
        title: "Registration successful!",
        text: "Your account has been successfully created.",
        icon: "success",
      });
      isSignUp.value = false;
    } else {
      Swal.fire({
        title: "Login successful!",
        text: "You have successfully logged in.",
        icon: "success",
      });
      //emit("getToken", response.data.token);
      localStorage.setItem("token", response.data.token);
      goEventPage();
    }
  } catch (err) {
    errMessage.value = isSignUp.value ? err.message : err.response.data.message;
    Swal.fire({
      title: "Error!",
      text: errMessage.value,
      icon: "error",
    });
  }
};

const toggleSignUp = () => {
  isSignUp.value = !isSignUp.value;
  resetForm();
};

const resetForm = () => {
  email.value = "";
  password.value = "";
  confirmPassword.value = "";
  errMessage.value = "";
};

const goEventPage = () => {
  router.push({ name: "Event" });
};
</script>

<template>
  <form class="form-content" @submit.prevent="submitForm" method="post">
    <h1>{{ isSignUp ? "Sign Up" : "Login" }}</h1>
    <input
      type="email"
      name="email"
      placeholder="✉️ Email address"
      v-model="email"
      required
    />
    <input
      type="password"
      name="password"
      placeholder="🔒 Password"
      v-model="password"
      required
    />
    <input
      :type="isSignUp ? 'password' : 'hidden'"
      name="confirmPassword"
      placeholder="🔒 Confirm Password"
      v-model="confirmPassword"
      required
    />
    <p class="invalid">
      <i
        v-if="errMessage"
        class="fa fa-exclamation-circle"
        aria-hidden="true"
      ></i>
      {{ errMessage }}
    </p>
    <button type="submit">{{ isSignUp ? "Signup" : "Login" }}</button>
    <p>
      {{ isSignUp ? "Already have an account?" : "Don't have an account?" }}
      <a href="" @click.prevent="toggleSignUp">{{
        isSignUp ? "Log In" : "Sign Up"
      }}</a>
    </p>
  </form>
</template>

<style scoped>
.form-content {
  position: absolute;
  top: 40%;
  left: 80%;
  transform: translate(-50%, -50%);
  text-align: center;
}

input {
  width: 300px;
  height: 40px;
  border-radius: 5px;
  margin: 15px 0;
  font-size: 15px;
  padding-left: 10px;
  box-shadow: 2px 2px 2px 1px rgb(0 0 0 / 15%);
  border: 1px solid #ccc;
  transition: border-color 0.5s;
}

input:focus {
  border: 2px solid #00b894;
  outline: none;
}

button {
  width: 315px;
  height: 47px;
  border-radius: 5px;
  margin: 20px 0 0;
  font-size: 15px;
  background-color: #3674b5;
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

h1 {
  color: #333446;
}

p {
  color: gray;
  font-size: 15px;
}

.invalid {
  color: #ed4337;
  text-align: left;
  font-size: 13px;
  margin: 0;
}

@media screen and (min-width: 580px) and (max-width: 992px) {
  .form-content {
    left: 50%;
    background-color: white;
    border-radius: 8px;
    height: 450px;
    width: 500px;
  }

  input {
    width: 400px;
  }

  button {
    width: 415px;
  }

  .invalid {
    margin: 0 45px 0;
  }
}

@media screen and (max-width: 579px) {
  .form-content {
    left: 50%;
    background-color: white;
    border-radius: 8px;
    height: 550px;
    width: 400px;
  }

  .invalid {
    margin: 0 45px 0;
  }
}
</style>
