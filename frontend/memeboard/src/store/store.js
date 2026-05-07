import {defineStore} from 'pinia'
import {ref} from 'vue'

export const useMemeStore = defineStore('meme',()=>{
    let memeIdCount = 0
    const API_URL = "http://127.0.0.1:8080"
    const showEditWindow = ref(false)
    const editingMeme = ref(null)
    const Memes = ref([])
    const getMemes = async ()=>{
        try{
            const res = await fetch(`${API_URL}/api/memeinfo`,{
                method : "GET"
            })
            const data = await res.json()
            Memes.value = data
        }catch(e){
           console.error('抓取 memes 失敗',e)
        }
    }
    const showUpdateMemeCard = (M)=>{
        showEditWindow.value = true
        editingMeme.value = M 
    }
    const deleteMemeCard = async (M)=>{
        try{
            const res = await fetch(`${API_URL}/api/memeinfo/${M.id}`,{
                method : "DELETE"
            })
            if(res.status===200){
                Memes.value = Memes.value.filter(m => m!==M)
            }else{
                window.alert(`刪除失敗 http status code = ${res.status}`)
            }
        }catch(e){
            console.log(e)
        }
    }
    const escEditWindow = ()=>{
        showEditWindow.value = false
        editingMeme.value = null
    }
    const saveAndEscEditWindow = async (M) =>{
        if(Memes.value.find(item => item.id === M.id) == undefined){
            try{
                const res = await fetch(`${API_URL}/api/memeinfo`,{
                    method : 'POST',
                    headers : {
                        'Content-Type' : 'application/json'
                    },
                    body : JSON.stringify(M)
                })
                if(res.status===201){
                    const data = await res.json()
                    Memes.value.push(data)
                }else{
                    window.alert(`新增貼文失敗 http status code = ${res.status}`)
                }
            }catch(e){
                console.log(e)
            }
        }else{
            try{
                const res = await fetch(`${API_URL}/api/memeinfo`,{
                    method : "PUT",
                    headers : {"Content-Type" : "application/json"},
                    body : JSON.stringify(M)
                })
                if(res.status === 200){
                    const data = await res.json()
                    Object.assign(editingMeme.value, data)
                }else{
                    const data = await res.json()
                    window.alert(`修改失敗 http status code = ${res.status} ${data.message}`)
                }
            }catch(e){
                console.log(e)
            }
        }
        showEditWindow.value = false
        editingMeme.value = null
    }
    return {
        memeIdCount,
        showEditWindow,
        editingMeme,
        Memes,
        getMemes,
        showUpdateMemeCard,
        deleteMemeCard,
        escEditWindow,
        saveAndEscEditWindow,
    }
})