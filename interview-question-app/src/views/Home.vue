<template>
    <div>
        <h1>主页</h1>
        <!-- 轮播图部分 -->
        <Swiper :slides-per-view="3" spaceBetween="30" :pagination="{ clickable: true }" class="swiper-container">
            <SwiperSlide v-for="(slide, index) in slides" :key="index">
                <img :src="slide.image" alt="Slide Image">
            </SwiperSlide>
        </Swiper>
        <!-- 热点面试题部分 -->
        <QuestionList />
        <router-link to="/questions">查看问题列表</router-link>
        <router-link to="/interview-experiences">查看面经列表</router-link>
    </div>
</template>

<script setup>
import { ref } from 'vue';
import { Swiper, SwiperSlide } from 'swiper/vue';
import 'swiper/css';
import QuestionList from '../components/QuestionList.vue';

// 模拟轮播图数据
const slides = ref([
    { image: 'https://example.com/image1.jpg' },
    { image: 'https://example.com/image2.jpg' },
    { image: 'https://example.com/image3.jpg' }
]);
</script>

<style scoped>
/* 主页样式 */
.swiper-container {
    margin: 20px 0;
}
.swiper-container img {
    width: 100%;
    height: auto;
}
</style>