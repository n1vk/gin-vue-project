<script setup>
import {computed, ref} from "vue";
import axios from "axios";
import FormInputItem from "./FormInputItem.vue";
import {User, Key} from '@element-plus/icons-vue'
import '@material/web/button/filled-button'
import '@material/web/button/outlined-button'

const account = ref('')
const password = ref('')

async function register() {
  try {
    const response = await axios.get('/reg', {
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

</script>

<template>

  <div class="flex flex-col justify-center min-h-screen items-center gap-y-5 ">

    <FormInputItem label="User Name" type="text" v-model="account">
      <template v-slot:icon>
        <User style="width: 1em; height: 1em; margin-right: 8px"/>
      </template>
    </FormInputItem>

    <FormInputItem label="Password" type="password" v-model="password">
      <template v-slot:icon>
        <Key style="width: 1em; height: 1em; margin-right: 8px"/>
      </template>
    </FormInputItem>

    <div class="flex flex-row w-1/2 justify-center gap-x-10">
      <md-filled-button label="注册" @click="register" class="flex-1"></md-filled-button>
    </div>

  </div>

</template>

<style scoped>

</style>