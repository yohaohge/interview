<template>
    <div>
      <h1>面经列表</h1>
      <ul>
        <!-- 这里可以使用 v-for 循环渲染面经列表 -->
        <li v-for="experience in experiences" :key="experience.id">
          <router-link :to="`/interview-experiences/${experience.id}`">{{ experience.title }}</router-link>
        </li>
      </ul>
    </div>
  </template>
  
  <script setup>
  import { ref } from 'vue';
  
  const experiences = ref([
    { id: 1, title: '面经 1' },
    { id: 2, title: '面经 2' },
    { id: 3, title: '面经 3' }
  ]);
  </script>
  
  <style scoped>
  /* 面经列表样式 */
  </style>