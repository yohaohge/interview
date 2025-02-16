<template>
    <div class="question-list-container">
      <div class="question-list-header">
        <h1>问题列表</h1>
      </div>
      <div class="question-list-body">
        <ul>
          <li v-for="question in questions" :key="question.id" class="question-item">
            <div class="question-title">
              <router-link :to="`/questions/${question.id}`">{{ question.title }}</router-link>
            </div>
            <div class="question-description">{{ question.description }}</div>
            <div class="question-tags">
              <span v-for="tag in question.tags" :key="tag" class="tag">{{ tag }}</span>
            </div>
          </li>
        </ul>
      </div>
      <div class="pagination">
        <button @click="prevPage" :disabled="currentPage === 1">上一页</button>
        <span>{{ currentPage }}</span>
        <button @click="nextPage" :disabled="currentPage === totalPages">下一页</button>
      </div>
    </div>
  </template>
  
  <script setup>
  import { ref } from 'vue';
  
  // Mock数据
  const mockQuestions = [
    { id: 1, title: 'Vue 3 组合式 API 有哪些优势？', description: 'Vue 3 的组合式 API 提供了更灵活的逻辑复用和代码组织方式...', tags: ['Vue 3', '组合式 API'] },
    { id: 2, title: '如何优化 React 应用的性能？', description: '可以通过使用 React.memo、shouldComponentUpdate 等方法来优化性能...', tags: ['React', '性能优化'] },
    { id: 3, title: '简述 SQL 中的 JOIN 操作', description: 'JOIN 操作用于从多个表中获取相关的数据...', tags: ['SQL', 'JOIN'] },
    { id: 4, title: 'JavaScript 中的闭包是什么？', description: '闭包是指有权访问另一个函数作用域中的变量的函数...', tags: ['JavaScript', '闭包'] },
    { id: 5, title: 'Python 中的装饰器有什么作用？', description: '装饰器可以在不修改原函数代码的情况下，为函数添加额外的功能...', tags: ['Python', '装饰器'] },
    { id: 6, title: 'Java 中的垃圾回收机制是如何工作的？', description: 'Java 的垃圾回收机制会自动回收不再使用的对象所占用的内存...', tags: ['Java', '垃圾回收'] },
    { id: 7, title: 'CSS 中的 Flexbox 布局有哪些优点？', description: 'Flexbox 布局可以更方便地实现弹性布局，适应不同的屏幕尺寸...', tags: ['CSS', 'Flexbox'] },
    { id: 8, title: 'Go 语言的并发模型是怎样的？', description: 'Go 语言通过 goroutine 和 channel 实现了高效的并发编程...', tags: ['Go', '并发模型'] },
    { id: 9, title: '什么是 RESTful API？', description: 'RESTful API 是一种基于 HTTP 协议的 API 设计风格...', tags: ['RESTful API'] },
    { id: 10, title: '如何设计一个高可用的数据库架构？', description: '可以通过主从复制、读写分离、集群等方式来设计高可用的数据库架构...', tags: ['数据库', '高可用架构'] },
    { id: 11, title: 'Docker 的优势有哪些？', description: 'Docker 可以实现容器化部署，提高应用的可移植性和部署效率...', tags: ['Docker', '容器化'] },
    { id: 12, title: 'Kubernetes 是如何实现容器编排的？', description: 'Kubernetes 通过 Pod、Service、Deployment 等资源对象来实现容器编排...', tags: ['Kubernetes', '容器编排'] },
    { id: 13, title: '机器学习中的监督学习和无监督学习有什么区别？', description: '监督学习需要有标注数据，而无监督学习不需要标注数据...', tags: ['机器学习', '监督学习', '无监督学习'] },
    { id: 14, title: '深度学习中的神经网络有哪些类型？', description: '深度学习中的神经网络有前馈神经网络、卷积神经网络、循环神经网络等...', tags: ['深度学习', '神经网络'] },
    { id: 15, title: '区块链的核心技术有哪些？', description: '区块链的核心技术包括分布式账本、加密算法、共识机制等...', tags: ['区块链', '核心技术'] },
    { id: 16, title: '大数据处理中的 Hadoop 生态系统包含哪些组件？', description: 'Hadoop 生态系统包含 HDFS、MapReduce、YARN 等组件...', tags: ['大数据', 'Hadoop'] },
    { id: 17, title: '云计算的三种服务模式是什么？', description: '云计算的三种服务模式是 IaaS、PaaS、SaaS...', tags: ['云计算', '服务模式'] },
    { id: 18, title: '什么是微服务架构？', description: '微服务架构是一种将应用拆分成多个小型服务的架构风格...', tags: ['微服务架构'] },
    { id: 19, title: 'DevOps 的核心思想是什么？', description: 'DevOps 的核心思想是打破开发、测试、运维之间的壁垒，实现快速迭代和持续交付...', tags: ['DevOps', '核心思想'] },
    { id: 20, title: '敏捷开发的原则有哪些？', description: '敏捷开发的原则包括客户满意、团队合作、快速响应变化等...', tags: ['敏捷开发', '原则'] }
  ];
  
  const pageSize = 20;
  const questions = ref([]);
  const currentPage = ref(1);
  const totalPages = ref(Math.ceil(mockQuestions.length / pageSize));
  
  const updateQuestions = () => {
    const startIndex = (currentPage.value - 1) * pageSize;
    const endIndex = startIndex + pageSize;
    questions.value = mockQuestions.slice(startIndex, endIndex);
  };
  
  const prevPage = () => {
    if (currentPage.value > 1) {
      currentPage.value--;
      updateQuestions();
    }
  };
  
  const nextPage = () => {
    if (currentPage.value < totalPages.value) {
      currentPage.value++;
      updateQuestions();
    }
  };
  
  // 初始化问题列表
  updateQuestions();
  </script>
  
  <style scoped>
  .question-list-container {
    font-family: Arial, sans-serif;
    max-width: 800px;
    margin: 0 auto;
    padding: 20px;
  }
  
  .question-list-header {
    margin-bottom: 20px;
  }
  
  .question-list-header h1 {
    font-size: 24px;
    font-weight: bold;
  }
  
  .question-list-body {
    border: 1px solid #ddd;
    border-radius: 4px;
    padding: 10px;
  }
  
  .question-item {
    border-bottom: 1px solid #eee;
    padding: 10px 0;
  }
  
  .question-title {
    font-size: 18px;
    font-weight: bold;
    margin-bottom: 5px;
  }
  
  .question-title a {
    color: #333;
    text-decoration: none;
  }
  
  .question-title a:hover {
    color: #007bff;
  }
  
  .question-description {
    color: #666;
    margin-bottom: 10px;
  }
  
  .question-tags {
    display: flex;
    gap: 5px;
  }
  
  .tag {
    background-color: #f0f0f0;
    color: #666;
    padding: 2px 5px;
    border-radius: 3px;
    font-size: 12px;
  }
  
  .pagination {
    display: flex;
    justify-content: center;
    align-items: center;
    margin-top: 20px;
  }
  
  .pagination button {
    background-color: #007bff;
    color: #fff;
    border: none;
    padding: 5px 10px;
    border-radius: 3px;
    cursor: pointer;
    margin: 0 10px;
  }
  
  .pagination button:disabled {
    background-color: #ccc;
    cursor: not-allowed;
  }
  </style>