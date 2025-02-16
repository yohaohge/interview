<template>
  <div id="app">
    <!-- 引入导航栏组件 -->
    <Navbar />
    <!-- 渲染当前路由匹配的组件 -->
    <router-view></router-view>
  </div>
</template>

<script setup>
import Navbar from './components/Navbar.vue';
</script>

<style scoped>
/* App 组件样式 */
body {
  margin: 0;
  font-family: Arial, sans-serif;
}
</style>