<script setup>
import {computed, ref} from "vue";
import axios from "axios";
import FormInputItem from "./FormInputItem.vue";
import {User, Key} from '@element-plus/icons-vue'
import '@material/web/button/filled-button'
import '@material/web/button/outlined-button'
import '@material/web/dialog/dialog'

const account = ref('')
const password = ref('')
let showDialog = false

const loginDisabled = computed(() => account.value === "" || password.value === "")


async function login() {

  try {
    const response = await axios.get('/log', {
      params: {
        account: account.value,
        password: password.value
      }
    })
    console.log(response.data)
    showDialog = true

    if(response.data['code'] === 4013 || response.data['code'] === 4012) {
      // TODO: 换成提示框，并要求用户重新登录
      alert("密码错误")
    } else {
      localStorage.token = response.data['token']
      console.log(localStorage.token)
      // TODO: 跳转到/home路径
    }


  } catch (e) {
    console.log(e)
  }

}

</script>

<template>
  <div class="flex flex-col justify-center min-h-screen items-center gap-y-5">

  <!-- TODO: 不能操控显示 -->
    <md-dialog :open="showDialog" class="">

      <div class="mb-14 flex justify-center">
        <h2 class="text-lg font-bold">登录成功</h2>
      </div>
      <router-link to="/home" v-slot="{navigate}" class="flex-1 flex justify-center">
        <div class="flex justify-center">
          <button @click="navigate" class="bg-blue-500 hover:bg-blue-700 text-white font-bold py-2 px-4 rounded">进入系统</button>
        </div>
      </router-link>

    </md-dialog>

    <FormInputItem label="登录用户名" type="text" v-model="account" >
      <template v-slot:icon>
        <User style="width: 1em; height: 1em; margin-right: 8px"/>
      </template>
    </FormInputItem>

    <FormInputItem label="密码" type="password"  v-model="password">
      <template v-slot:icon>
        <Key style="width: 1em; height: 1em; margin-right: 8px"/>
      </template>
    </FormInputItem>


    <div class="flex flex-col w-1/2 justify-center gap-y-5">
      <md-filled-button label="登录" @click="login" :disabled="loginDisabled" class="flex-1"></md-filled-button>

      <router-link to="/register" v-slot="{navigate}" class="flex-1 flex justify-center">
        <a @click="navigate" class="text-blue-500 hover:text-red-500 underline">
          注册新用户
        </a>
      </router-link>

    </div>

  </div>


</template>

<style scoped>

</style>