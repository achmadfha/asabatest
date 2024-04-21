<template>
  <div class="w-full h-full h-max-full flex justify-center align-center text-center">
    <div class="flex-col m-auto">
      <h1 class="font-bold text-2xl mb-[1rem]">Hello Weaver</h1>
      <h2 class="text-lg text-gray-500">Neque porro quisquam est qui dolorem ipsum quia dolor sit amet, consectetur, adipisci velit...</h2>
      <img src="@/assets/welcome.svg" class="w-1/3 my-[2rem] mx-auto" >
      <el-button :icon="Waves" type="primary" plain><router-link to="/about">About Weaver</router-link></el-button>
      <el-button :icon="Github"><a target="_blank">Explore More</a></el-button>
    </div>
  </div>
</template>

<script setup>
import { Github, Waves } from '@icon-park/vue-next';
</script>
