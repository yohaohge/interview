<template>
    <div class="swiper">
      <div class="swiper-wrapper">
        <div class="swiper-slide" v-for="(slide, index) in slides" :key="index">
          <img :src="slide.image" alt="Slide">
        </div>
      </div>
      <div class="swiper-pagination"></div>
      <div class="swiper-button-prev"></div>
      <div class="swiper-button-next"></div>
    </div>
  </template>
  
  <script setup>
  import { onMounted } from 'vue';
  import Swiper from 'swiper';
  import { Pagination, Navigation } from 'swiper/modules';
  import 'swiper/css';
  import 'swiper/css/pagination';
  import 'swiper/css/navigation';
  
  const slides = [
    { image: 'https://via.placeholder.com/1920x600?text=Slide+1' },
    { image: 'https://via.placeholder.com/1920x600?text=Slide+2' },
    { image: 'https://via.placeholder.com/1920x600?text=Slide+3' }
  ];
  
  onMounted(() => {
    new Swiper('.swiper', {
      modules: [Pagination, Navigation],
      pagination: {
        el: '.swiper-pagination'
      },
      navigation: {
        nextEl: '.swiper-button-next',
        prevEl: '.swiper-button-prev'
      }
    });
  });
  </script>
  
  <style scoped>
  .swiper {
    width: 100%;
    height: 600px;
  }
  
  .swiper-slide img {
    width: 100%;
    height: 100%;
    object-fit: cover;
  }
  </style>