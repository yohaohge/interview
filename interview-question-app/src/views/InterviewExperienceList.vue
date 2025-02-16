<template>
    <div class="question-list-container">
      <h1 class="list-title">问题列表</h1>
      <ul class="question-list">
        <!-- 这里使用 v-for 循环渲染问题列表 -->
        <li v-for="question in questions" :key="question.id" class="question-item">
          <router-link :to="`/questions/${question.id}`" class="question-link">
            <div class="question-title">{{ question.title }}</div>
            <div class="question-tags">
              <span v-for="tag in question.tags" :key="tag" class="tag">{{ tag }}</span>
            </div>
          </router-link>
        </li>
      </ul>
    </div>
  </template>
  
  <script setup>
  import { ref } from 'vue';
  
  // 模拟数据
  const questions = ref([
    { id: 1, title: 'Vue 3 组合式 API 有哪些优势？', tags: ['Vue 3', '组合式 API'] },
    { id: 2, title: '如何优化 React 应用的性能？', tags: ['React', '性能优化'] },
    { id: 3, title: '简述 SQL 中的 JOIN 操作', tags: ['SQL', 'JOIN'] }
  ]);
  </script>
  
  <style scoped>
  /* 整体容器样式 */
  .question-list-container {
    padding: 20px;
    font-family: Arial, sans-serif;
    display: flex;
    flex-direction: column;
    align-items: center; /* 居中 */
  }
  
  /* 列表标题样式 */
  .list-title {
    font-size: 24px;
    font-weight: bold;
    margin-bottom: 20px;
    color: #333;
  }
  
  /* 列表样式 */
  .question-list {
    list-style-type: none;
    padding: 0;
    margin: 0;
    width: 80%; /* 适当调整宽度 */
  }
  
  /* 列表项样式 */
  .question-item {
    border: 1px solid #ddd;
    border-radius: 5px;
    padding: 15px;
    margin-bottom: 15px;
    background-color: #fff;
    box-shadow: 0 2px 4px rgba(0, 0, 0, 0.1);
    transition: box-shadow 0.3s ease;
  }
  
  .question-item:hover {
    box-shadow: 0 4px 8px rgba(0, 0, 0, 0.2);
  }
  
  /* 链接样式 */
  .question-link {
    text-decoration: none;
    color: #333;
    display: block;
  }
  
  /* 问题标题样式 */
  .question-title {
    font-size: 18px;
    font-weight: bold;
    margin-bottom: 10px;
  }
  
  /* 标签列样式 */
  .question-tags {
    display: flex;
    flex-wrap: wrap;
    gap: 5px;
  }
  
  .tag {
    background-color: #f0f0f0;
    padding: 3px 8px;
    border-radius: 3px;
    font-size: 12px;
    color: #666;
  }
  </style>