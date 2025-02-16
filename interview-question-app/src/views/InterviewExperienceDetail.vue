<template>
    <div>
      <h1>{{ experience.title }}</h1>
      <p>{{ experience.content }}</p>
    </div>
  </template>
  
  <script setup>
  import { ref, onMounted } from 'vue';
  import { useRoute } from 'vue-router';
  
  const route = useRoute();
  const experience = ref({});
  
  onMounted(() => {
    // 模拟根据 id 获取面经详情
    const id = Number(route.params.id);
    // 这里可以替换为实际的 API 请求
    experience.value = {
      id,
      title: `面经 ${id} 详情`,
      content: `这是面经 ${id} 的详细内容`
    };
  });
  </script>
  
  <style scoped>
  /* 面经详情样式 */
  </style>