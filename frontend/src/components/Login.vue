<script setup>
import {ref} from "vue";
import axios from "axios";

const userName = ref('')
const password = ref('')

async function login() {
  // console.log(userName.value)
  // console.log(password.value)

    try {
      const response = await axios.get('/log', {
      params: {
        Acc: userName.value,
        Pwd: password.value
      }
    })
      console.log(response.data)
    } catch (e) {
      console.log(e)
    }

}

const prompt = "xxx"
</script>

<template>
  <div>
      <span> 用户名　</span>
      <input class="input_box" v-model="userName" :placeholder="prompt"/>
      <br><br>
      <span> 密　码　</span>
      <input type="password" id="pwd_input" class="input_box" v-model="password"/>
  </div>
  <button class="login_button" @click="login">登录</button>
</template>

<style scoped>
.login_button {
  margin-top: 2em;
  width: 10em;
  padding: 0.5em;
  background-color: #e9e9e9;
  will-change: filter;
  transition: filter 300ms;
}

.input_box {
  border-radius: 0.5em;
}
</style>