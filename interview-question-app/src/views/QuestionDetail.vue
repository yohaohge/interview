<template>
    <div class="question-detail-container">
      <h1>{{ question.title }}</h1>
      <p class="question-content">{{ question.content }}</p>
    </div>
  </template>
  
  <script setup>
  import { ref, onMounted } from 'vue';
  import { useRoute } from 'vue-router';
  
  const route = useRoute();
  const question = ref({});
  
  onMounted(() => {
    // 模拟根据 id 获取问题详情
    const id = Number(route.params.id);
    // 这里可以替换为实际的 API 请求
    question.value = {
      id,
      title: `问题 ${id} 详情`,
      content: `这是问题 ${id} 的详细内容`
    };
  });
  </script>
  
  <style scoped>
  /* 问题详情容器样式 */
  .question-detail-container {
    max-width: 800px;
    margin: 0 auto;
    padding: 20px;
  }

  /* 问题内容样式 */
  .question-content {
    font-size: 1.1em;
    line-height: 1.6;
  }
  </style>