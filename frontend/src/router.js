import { createWebHistory, createRouter } from "vue-router";
import LoginPage from "./components/LoginPage.vue";
import EventPage from "./components/EventPage.vue";

const routes = [
  { path: "/", redirect: "/login" },
  { path: "/logout", redirect: "/login" },
  { path: "/login", name: "Login", component: LoginPage },
  { path: "/event", name: "Event", component: EventPage },
  { path: "/:notFound(.*)", redirect: "/login" },
];

const router = createRouter({ history: createWebHistory(), routes });

const isAuthenticated = () => {
  return !!localStorage.getItem("token");
};

router.beforeEach((to, from, next) => {
  if (!isAuthenticated() && to.name !== "Login") {
    next({ name: "Login" });
  } else {
    next();
  }
});

export default router;
