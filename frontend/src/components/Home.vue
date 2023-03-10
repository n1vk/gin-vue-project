<script setup>
import '@material/web/button/filled-button'
import '@material/web/dialog/dialog'
import axios from "axios";
import {ref} from "vue";

const qualified = ref(false)

async function send() {
  try {
    const response = await axios.get('/home/hello', {
      params: {
        // 如果这里 token 是 undefined 就会返回403
        token: localStorage.token
      }
    })
    console.log(response.data)

  } catch (e) {
    qualified.value = true
    console.log(e)
  }
}

</script>

<template>

  <div class="flex flex-col justify-center min-h-screen items-center gap-y-5">
    <md-dialog :open="qualified" class="" v-if="qualified">
      <div class="mb-14 flex justify-center">
        <h2 class="text-lg font-bold">您没有权限发消息，请登录</h2>
      </div>
      <router-link to="/" v-slot="{navigate}" class="flex-1 flex justify-center">
        <div class="flex justify-center">
          <button @click="qualified = false" class="bg-blue-500 hover:bg-blue-700 text-white font-bold py-2 px-4 rounded">
            好的
          </button>
        </div>
      </router-link>
    </md-dialog>

    <div class="flex flex-row w-1/2 justify-center">
      <md-filled-button label="发消息" @click="send" class="flex-1"></md-filled-button>
    </div>

  </div>
</template>


<style scoped>

</style>