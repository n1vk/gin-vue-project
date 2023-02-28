<script setup>
import {computed, ref} from "vue";
import axios from "axios";
import FormInputItem from "./FormInputItem.vue";
import {User, Key} from '@element-plus/icons-vue'
import '@material/web/button/filled-button'
import '@material/web/button/outlined-button'

</script>

<template>

  <div class="flex flex-col justify-center min-h-screen items-center">

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

    <div class="flex flex-row">
      <md-filled-button label="注册" @click="register" class="m-5"></md-filled-button>
    </div>

  </div>

</template>

<style scoped>

</style>