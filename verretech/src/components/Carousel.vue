<template>
  <div class="carousel-example">
    <swiper :dir="$vs.rtl ? 'rtl' : 'ltr'" :key="$vs.rtl">
      <swiper-slide
        class="carousel-item"
        style="background-image: url('https://i.pinimg.com/originals/3a/7f/13/3a7f13c4c4366a60dc414c661f6ddd8c.jpg');
        "
      >
        <!-- <img class="responsive" src="https://picsum.photos/100/50" alt="banner"> -->
        <div
          style="height:100%; display:flex;align-items:center;  
          "
          class="vx-row"
        >
          <div
            class="w-1/2"
            style="display:flex;align-items:center; flex-direction:column
          "
          >
            <h1 style=" font-weight:700; ">
              Apprenez-en plus sur VERRE-TECH
            </h1>
            <p class="w-2/3 mt-3" style="text-align:center; font-weight:600; color:black">
                Lorem ipsum dolor sit, amet consectetur adipisicing elit. Error eum sapiente soluta deleniti aperiam est cum ipsum nemo, et, laborum laudantium voluptatum placeat tenetur! Maxime est assumenda ipsa dolores ducimus?
            </p>
            <vs-button class="mt-3" >En savoir plus</vs-button>
          </div>
          <div class="w-1/2"></div>
        </div>
      </swiper-slide>
    </swiper>
  </div>
</template>

<script>
import "swiper/dist/css/swiper.min.css";
import { swiper, swiperSlide } from "vue-awesome-swiper";

export default {
  components: {
    swiper,
    swiperSlide,
  },
};
</script>

<style>

.carousel-item {
  width: 100%;
  height: 650px;
  background-repeat: no-repeat;
  background-size: cover;
  background-position: center;
  /* filter: blur(5px); */
}
</style>
