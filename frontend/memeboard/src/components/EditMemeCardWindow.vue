<script setup>
import {ref} from 'vue'
import { useMemeStore } from '../store/store';
const store = useMemeStore()
const tempM = ref({...store.editingMeme})
</script>

<template>
    <div class="edit-bg">
        <div class="edit-window-top">
            <button class="esc-edit" @click="store.escEditWindow()">X</button>
        </div>
        <div class="edit-window-mid">
            <textarea class="edit-area" v-model="tempM.content"></textarea>
        </div>
        <div class="edit-window-btm">
            <button class="save-edit" @click="store.saveAndEscEditWindow(tempM)">儲存並退出</button>
        </div>
    </div>
</template>

<style scope>
.edit-bg{
    display: flex;
    flex-direction: column;
    width: 80%;
    height: 90%;
    background-color: var(--edit-window-bg-color,#B9B9FF);
    border-radius: 16px;
}
.edit-window-top{
    display: flex;
    justify-content: end;
    text-align: center;
    padding : 0.8rem 1.5rem;
}
.edit-window-mid{
    flex:1;
    display: flex;
    justify-content: center;
    align-items: center;
    padding:1rem 2rem;
}
.edit-window-btm{
    display: flex;
    justify-content: center;
    align-items: center;
    padding : 1rem 3rem;
}
.esc-edit{
    aspect-ratio: 1.1/1;
    border: 1px solid var(--esc-edit-border-color,#beaeae);
    border-radius: 50%;
    background-color: var(--esc-edit-color,#B87070);
    cursor:pointer;
    overflow: hidden;
}
.save-edit{
    border-radius: 8px;
    border: 1px solid var(--save-edit-btn-border-color,#81C0C0);
    background-color: var(--save-edit-btn-color,#4F9D9D);
    color:white;
    padding : 0.3rem 0.5rem;
    font-size : 0.8rem;
    cursor : pointer;
}
.edit-area{
    width: 100%;
    height:100%;
    resize:none;
    box-sizing : border-box;
    border-radius: 8px;
}
</style>