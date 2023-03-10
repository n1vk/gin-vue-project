<script setup>
import {HomeFilled} from '@element-plus/icons-vue'
</script>

<template>

  <div class="flex object-top p-5 bg-blue-400 shadow-md -mb-10">

    <router-link to="/" v-slot="{navigate}">
      <HomeFilled style="width: 1.875rem; height: 2.25rem; margin-right: 30px" color="white" @click="navigate"/>
    </router-link>

    <h1 class="text-3xl text-white"> 欢迎使用 </h1>
  </div>

  <router-view />

</template>

<style>
  ::-webkit-scrollbar {
    display: none;
  }
</style>
