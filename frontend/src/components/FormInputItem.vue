<script setup xmlns="">
import '@material/web/textfield/outlined-text-field'

defineProps(['label', 'type', 'modelValue'])
defineEmits(['update:modelValue'])

</script>

<template>
  <div class="w-1/2 flex justify-center">

  <md-outlined-text-field :label="label" :type="type" :value="modelValue" @input="$emit('update:modelValue', $event.target.value)" class="w-full">
    <div style="font-size: 20px" slot="leadingicon">
      <slot name="icon"/>
    </div>
  </md-outlined-text-field>

  </div>
</template>

<style scoped>

</style>