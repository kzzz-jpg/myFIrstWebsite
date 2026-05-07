<script setup>
import {ref} from 'vue'
import { useMemeStore } from '../store/store';
const store = useMemeStore()
const postMeme = ()=>{
    if(store.showEditWindow) return;
    let tempM = ref({
        id : store.memeIdCount++,
        content : "",
        author : "kzzz",
    })
    store.showEditWindow= true
    store.editingMeme= tempM.value
}
</script>

<template>
  <div class="meme-header">
    <div class="header-content">
      MemeBoard
    </div>
    <div class="header-post-btn">
        <button class="post-btn" @click="postMeme">發文</button>
    </div>
  </div>
</template>

<style scoped>
.meme-header {
  background-color: var(--header-bg-color, #FFE6D9);
  padding: 1.5rem 2rem;
  text-align: center;
  align-items: center;
  width: 100vw;
  display:flex;
  height: 16px;
  border-bottom : 1px solid var(--header-bottom-color,#eeb183)
}
.header-content{
    width : 50%;
    text-align: left;
    padding : 2rem;
    font-size: 1.5rem;
    font-weight: 800; 
    letter-spacing: -0.05rem;
    color : var(--header-content-color, #d79d70);
}
.header-post-btn{
    width: 50%;
    padding-right : 5em;
    text-align: right;
}
.post-btn{
    border-radius: 8px;
    border:1px solid var(--post-btn-border-color,#B8B8DC);
    overflow: hidden;
    background-color: var(--post-btn-background-color,#FF8F59);
    color: var(--post-btn-text-color,white);
    cursor : pointer;
    font-size : 1rem;
    padding : 0.5rem 1.5rem;
}
</style>