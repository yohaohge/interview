<template>
  <footer class="footer">
    <p>&copy; 2024 面试抱佛脚. 保留所有权利</p>
    <ul class="footer-links">
      <li><a href="#" class="footer-link">联系我们</a></li>
      <li><a href="#" class="footer-link">隐私政策</a></li>
      <li><a href="#" class="footer-link">使用条款</a></li>
    </ul>
  </footer>
</template>

<style scoped>
.footer {
  background-color: #333;
  color: white;
  text-align: center;
  padding: 20px;
}

.footer-links {
  list-style-type: none;
  display: flex;
  justify-content: center;
  gap: 20px;
  margin-top: 10px;
}

.footer-link {
  color: white;
  text-decoration: none;
  transition: color 0.3s ease;
}

.footer-link:hover {
  color: #007bff;
}
</style>