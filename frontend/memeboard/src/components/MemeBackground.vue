<script setup>

import Memepreview from './Memepreview.vue'
import EdidMemeCardWindow from './EditMemeCardWindow.vue'
import { useMemeStore } from '../store/store'
import {ref} from 'vue'
const store = useMemeStore()



//getMemes()


</script>

<template>
<div class="main-bg">
    <div class="meme-cards">
        <Memepreview v-for="m in store.Memes" :key="m.id" :M="m"/>
    </div>
    <div class="edit-window" v-if="store.showEditWindow">
        <EdidMemeCardWindow/>
    </div>
</div>
</template>

<style scoped>
.main-bg{
    position: relative;
    background-color: var(--meme-bg-color,#C4E1E1);
    flex:1;
    width:100%;
    display: flex;
    justify-content: center;
    padding:2rem;
}
.meme-cards{
    width: 90%;
    display: grid;
    grid-template-columns: repeat(auto-fill  ,20rem);
    grid-template-rows: repeat(auto-fill  ,15rem);
    gap :1rem;
}
.edit-window{
    position : absolute;
    top:0;
    left: 0;
    display: flex;
    height : 100%;
    width: 100%;
    background: rgba(0,0,0,0.3  );
    align-items: center;
    justify-content: center;
}
</style>


