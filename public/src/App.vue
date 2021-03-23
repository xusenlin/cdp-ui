<template>
  <AppBar />
  <Container />
</template>

<script>
import AppBar from "./components/AppBar.vue";
import Container from "./components/Container.vue";

export default {
  components: { AppBar, Container }
};
</script>

<style>
#app {
  font-family: Avenir, Helvetica, Arial, sans-serif;
  -webkit-font-smoothing: antialiased;
  -moz-osx-font-smoothing: grayscale;
  width: 100vw;
  height: 100vh;
  overflow: auto;
  background-color: #f7f7f7;
}
</style>
