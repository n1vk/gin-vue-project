<script setup>
import '@material/web/button/filled-button'
import axios from "axios";

async function send() {
  try {
    const response = await axios.get('/home/hello', {
      params : {
        token: localStorage.token
      }
    })
    console.log(response.data)

  } catch (e) {
    console.log(e)
  }
}

</script>

<template>

  <div class="flex flex-col justify-center min-h-screen items-center">
    <div class="flex flex-row w-1/2 justify-center">
      <md-filled-button label="发消息" @click="send" class="flex-1"></md-filled-button>
    </div>
  </div>
</template>



<style scoped>

</style>