<template>
    <div>
      <h1>问题列表</h1>
      <ul>
        <!-- 这里可以使用 v-for 循环渲染问题列表 -->
        <li v-for="question in questions" :key="question.id">
          <router-link :to="`/questions/${question.id}`">{{ question.title }}</router-link>
        </li>
      </ul>
    </div>
  </template>
  
  <script setup>
  import { ref } from 'vue';
  
  const questions = ref([
    { id: 1, title: '问题 1' },
    { id: 2, title: '问题 2' },
    { id: 3, title: '问题 3' }
  ]);
  </script>
  
  <style scoped>
  /* 问题列表样式 */
  </style>