<script setup>
import {computed, ref} from "vue";
import axios from "axios";
import FormInputItem from "./FormInputItem.vue";
import {User, Key} from '@element-plus/icons-vue'
import '@material/web/button/filled-button'
import '@material/web/button/outlined-button'


const account = ref('')
const password = ref('')


const loginDisabled = computed(() => account.value === "" || password.value === "")


async function login() {

  try {
    const response = await axios.get('/log', {
      params: {
        Acc: account.value,
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
  <div class="flex flex-col justify-center min-h-screen items-center gap-y-5 ">

    <FormInputItem label="User Name" type="text" v-model="account" >
      <template v-slot:icon>
        <User style="width: 1em; height: 1em; margin-right: 8px"/>
      </template>
    </FormInputItem>

    <FormInputItem label="Password" type="password" v-model="password">
      <template v-slot:icon>
        <Key style="width: 1em; height: 1em; margin-right: 8px"/>
      </template>
    </FormInputItem>


    <div class="flex flex-col w-1/2 justify-center gap-y-5">
      <md-filled-button label="登录" @click="login" class="flex-1"></md-filled-button>

      <router-link to="/register" v-slot="{navigate}" class="flex-1 flex justify-center">
        <a @click="navigate" class="text-blue-500 hover:text-red-500 underline">
          注册新用户
        </a>
<!--        <md-outlined-button label="注册" @click="navigate" class="w-full"></md-outlined-button>-->
      </router-link>

    </div>

  </div>


</template>

<style scoped>

</style>