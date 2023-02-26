<script setup>
import {computed, ref} from "vue";
import axios from "axios";
import FormInputItem from "./FormInputItem.vue";
import {User, Key} from '@element-plus/icons-vue'

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

  <div>
  <FormInputItem label="User Name">
    <template v-slot:icon>
      <User style="width: 1em; height: 1em; margin-right: 8px"/>
    </template>
  </FormInputItem>

  <FormInputItem label="Password">
    <template v-slot:icon>
      <Key style="width: 1em; height: 1em; margin-right: 8px"/>
    </template>
  </FormInputItem>
  </div>

</template>


<style scoped>

</style>