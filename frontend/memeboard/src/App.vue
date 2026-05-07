<script setup>
import MemeHeader from './components/MemeHeader.vue'
import MemeBackground from './components/MemeBackground.vue'
import { useMemeStore } from './store/store';
const store = useMemeStore()
store.getMemes()
</script>

<template>
  <div class="app-main">
    <MemeHeader/>
    <MemeBackground/>
  </div>
</template>


<style scoped>
.app-main {
  display: flex;
  flex-direction: column;
  align-items: center;
  justify-content: center;
  height: 100vh;
  width: 100%;
}
</style>