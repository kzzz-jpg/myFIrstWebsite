<script setup>
import { useMemeStore } from '../store/store';
const props = defineProps({
    M : Object
})
const emits = defineEmits(['upd','del'])
const store = useMemeStore()
</script>

<template>
    <div class="meme-card">
        <div class="meme-content">
            {{ M.content }}
        </div>
        <div class="meme-info">
            <div class="meme-author">
                {{M.author}}
            </div>
            <div class="meme-modify">
                <button class="meme-update-btn" @click="store.showUpdateMemeCard(M)">修改</button>
                <button class="meme-delete-btn" @click="store.deleteMemeCard(M)">刪除</button>
            </div>
        </div>
    </div>
</template>

<style scoped>
.meme-card{
    border-radius: 16px;
    border: 1px solid var(--meme-card-border-color,#ADFEDC);
    padding : 1rem;
    background-color: var(--meme-card-color,#7AFEC6);
    display : flex;
    flex-direction: column;
}
.meme-content{
    flex:1;
    padding: 1rem 2rem;
    text-align: left;
    display: block;
    word-wrap: break-word;
    word-break: break-word;
    white-space: normal;
    overflow-y: auto;
    max-height: 80%;
}
.meme-info{
    height:20%;
    display: flex;
    align-items: right;
    text-align: right;
    justify-content: space-between;
    padding :0.5rem 2rem;
}
.meme-author{
    text-align: left;
}

button{
    border : 1px solid var(--btn-border-color,#28FF28);
    background-color : var(--meme-card-btn-color,#FFFF6F);
    color : var(--meme-card-btn-text-color,#A6A6D2);
    font-size: 1rem;
    padding:0.4rem 0.8rem;
    border-radius: 8px;
    cursor:pointer;     
}
</style>