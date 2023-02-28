<script setup>
import {computed, ref} from "vue";
import axios from "axios";
import FormInputItem from "./FormInputItem.vue";
import {User, Key} from '@element-plus/icons-vue'
import '@material/web/button/filled-button'
import '@material/web/button/outlined-button'


const userName = ref('')
const password = ref('')

const loginDisabled = computed(() => userName.value === "" || password.value === "")


async function login() {
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

const prompt = "请输入用户名"
</script>

<template>
  <div class="flex flex-col justify-center min-h-screen items-center">

    <FormInputItem label="User Name" :disabled="loginDisabled">
      <template v-slot:icon>
        <User style="width: 1em; height: 1em; margin-right: 8px"/>
      </template>
    </FormInputItem>

    <FormInputItem label="Password">
      <template v-slot:icon>
        <Key style="width: 1em; height: 1em; margin-right: 8px"/>
      </template>
    </FormInputItem>

    <div class="flex flex-row">
      <md-filled-button label="登录" @click="login" class="m-5"></md-filled-button>

      <router-link to="/register" v-slot="{navigate}">
        <md-outlined-button label="注册" @click="navigate" class="m-5"></md-outlined-button>
      </router-link>

    </div>



  </div>


</template>

<style scoped>

</style>