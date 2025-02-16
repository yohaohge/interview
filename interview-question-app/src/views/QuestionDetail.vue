<template>
    <div>
      <h1>{{ question.title }}</h1>
      <p>{{ question.content }}</p>
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
  /* 问题详情样式 */
  </style>