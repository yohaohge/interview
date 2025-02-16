<template>
  <nav class="navbar">
    <div class="navbar-container">
      <!-- 模拟 logo -->
      <div class="navbar-logo">
        <router-link to="/">面试题库</router-link>
      </div>
      <!-- 导航链接 -->
      <div class="navbar-links">
        <router-link to="/questions">问题列表</router-link>
        <router-link to="/interview-experiences">面经列表</router-link>
      </div>
      <!-- 模拟搜索框 -->
      <div class="navbar-search">
        <input type="text" placeholder="搜索问题或面经">
        <button>搜索</button>
      </div>
    </div>
  </nav>
</template>

<script setup>
// 导航栏组件逻辑，目前可留空
</script>

<style scoped>
/* 导航栏整体样式 */
.navbar {
  position: fixed;
  top: 0;
  left: 0;
  width: 100%;
  background-color: #fff;
  box-shadow: 0 2px 4px rgba(0, 0, 0, 0.1);
  z-index: 1000;
  padding: 10px 0;
}

/* 导航栏容器样式 */
.navbar-container {
  display: flex;
  justify-content: space-between;
  align-items: center;
  max-width: 1200px;
  margin: 0 auto;
  padding: 0 20px;
}

/* logo 样式 */
.navbar-logo {
  font-size: 1.8em;
  font-weight: bold;
  color: #333;
}

.navbar-logo router-link {
  color: inherit;
  text-decoration: none;
}

/* 导航链接样式 */
.navbar-links {
  display: flex;
  gap: 30px;
}

.navbar-links router-link {
  color: #333;
  text-decoration: none;
  font-size: 1.1em;
  transition: color 0.3s ease;
}

.navbar-links router-link:hover {
  color: #007bff;
}

/* 搜索框样式 */
.navbar-search {
  display: flex;
  align-items: center;
  gap: 10px;
}

.navbar-search input {
  padding: 10px 15px;
  border: 1px solid #ccc;
  border-radius: 4px;
  outline: none;
  font-size: 1em;
}

.navbar-search button {
  padding: 10px 20px;
  background-color: #007bff;
  color: #fff;
  border: none;
  border-radius: 4px;
  cursor: pointer;
  font-size: 1em;
  transition: background-color 0.3s ease;
}

.navbar-search button:hover {
  background-color: #0056b3;
}
</style>