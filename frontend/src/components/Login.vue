<template>

  <div>
      <span> 用户名　</span>
      <input type="text" id="acc_input" class="input_box" :value="accContent" />
      <br><br>
      <span> 密　码　</span>
      <input type="password" id="pwd_input" class="input_box"/>
  </div>

  <button class="login_button" >登录</button>
</template>

<script>
export default {
  name: "Login.vue",
}
</script>

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