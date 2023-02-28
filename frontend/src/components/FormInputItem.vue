<script setup xmlns="">
import '@material/web/textfield/outlined-text-field'

const props = defineProps(['label'])


</script>

<template>
  <div class="my-3">

  <md-outlined-text-field :label="props.label">
    <div style="font-size: 20px" slot="leadingicon">
      <slot name="icon"/>
    </div>
  </md-outlined-text-field>

  </div>
</template>

<style scoped>

</style>