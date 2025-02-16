<template>
    <div class="question-list">
      <h2>热门面试题</h2>
      <ul>
        <li v-for="question in questions" :key="question.id">
          <h3>{{ question.title }}</h3>
          <p>{{ question.description }}</p>
        </li>
      </ul>
    </div>
  </template>
  
  <script setup>
  import { ref } from 'vue';
  
  const questions = ref([
    { id: 1, title: 'Vue 3 组合式 API 有哪些优势？', description: 'Vue 3 的组合式 API 提供了更灵活的逻辑复用和代码组织方式...' },
    { id: 2, title: '如何优化 React 应用的性能？', description: '可以通过使用 React.memo、shouldComponentUpdate 等方法来优化性能...' },
    { id: 3, title: '简述 SQL 中的 JOIN 操作', description: 'JOIN 操作用于从多个表中获取相关的数据...' }
  ]);
  </script>
  
  <style scoped>
  .question-list {
    padding: 20px;
  }
  
  .question-list h2 {
    font-size: 24px;
    margin-bottom: 10px;
  }
  
  .question-list ul {
    list-style-type: none;
    padding: 0;
  }
  
  .question-list li {
    border: 1px solid #ccc;
    border-radius: 4px;
    padding: 10px;
    margin-bottom: 10px;
  }
  
  .question-list h3 {
    font-size: 18px;
    margin-bottom: 5px;
  }
  </style>