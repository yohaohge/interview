// router/index.js
import { createRouter, createWebHistory } from 'vue-router';

// 引入各个页面组件，这里假设组件存放在 views 文件夹中
import Home from '../views/Home.vue';
import QuestionList from '../views/QuestionList.vue';
import QuestionDetail from '../views/QuestionDetail.vue';
import InterviewExperienceList from '../views/InterviewExperienceList.vue';
import InterviewExperienceDetail from '../views/InterviewExperienceDetail.vue';

const routes = [
  {
    path: '/',
    name: 'Home',
    component: Home
  },
  {
    path: '/questions',
    name: 'QuestionList',
    component: QuestionList
  },
  {
    path: '/questions/:id',
    name: 'QuestionDetail',
    component: QuestionDetail,
    props: true // 允许将路由参数作为组件的 props 传入
  },
  {
    path: '/interview-experiences',
    name: 'InterviewExperienceList',
    component: InterviewExperienceList
  },
  {
    path: '/interview-experiences/:id',
    name: 'InterviewExperienceDetail',
    component: InterviewExperienceDetail,
    props: true
  }
];

const router = createRouter({
  history: createWebHistory(),
  routes
});

export default router;