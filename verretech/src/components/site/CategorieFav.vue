<template>
  <div class="bandeau">
    <div class="vx-row m-5 p-5">
      <div class="w-full justify-center">
        <h1>Catégories phares</h1>
      </div>
      <div class="w-full justify-center">
        <p class="fs-1-5 w-2/3 text-center mt-5">
          Utilisez nos verres traditionnels pour vos plans de travail ou toute
          autre réalisation en verre pour votre intérieur. Nos verres
          apporteront également du style à votre cloisonnement.
        </p>
      </div>
    </div>
    <div class="mt-5 justify-content-around container">
      <div class=" flex-column align-center">
        <img style="width:70%;height:auto;" src="https://www.anywaydoors.be/sites/default/files/styles/galery_big/public/fotogalerij/cloison-en-verre-style-atelier.jpg?itok=YABCHcIm" alt="">
         <span
          class="fs-1-5 mt-3"
          style="text-wrap:wrap; display:flex;justify-content:center; font-weight:bold"
        >
          Plans de travail
        </span>
      </div>
      <div class="  flex-column align-center">
        <img style="width:70%;height:auto;" src="https://www.anywaydoors.be/sites/default/files/styles/galery_big/public/fotogalerij/cloison-en-verre-style-atelier.jpg?itok=YABCHcIm" alt="">
          <span
          class="fs-1-5 mt-3"
          style="text-wrap:wrap; display:flex;justify-content:center; font-weight:bold"
        >
          Cloisons
        </span>
      </div>
    </div>
  </div>
</template>
<script>
import 'swiper/dist/css/swiper.min.css'
import { swiper, swiperSlide } from 'vue-awesome-swiper'
export default {
  data() {
    return {
      swiperOption: {
        pagination: {
          el: ".swiper-pagination",
        },
      },
    };
  },
  components: {
    swiper,
    swiperSlide,
  },
};
</script>

<style scoped>
.bandeau{
border-top: solid 8px #c2c2c2;

  margin-bottom: 50px;
}
.carousel-example{
  padding:100px
}
.flex-column{
  display: flex;
  flex-direction: column;
}
.swiper-slide {
  width: 50% !important;
}
.justify-center {
  display: flex;
  justify-content: center;
}
.justify-content-around {
  display: flex;
  justify-content: space-around;
}
h1 {
  font-size: 3em;
  font-weight: bold;
}
.fs-1-5 {
  font-size: 1.5em;
}
img{
  border:solid 3px #2c2c2c;
  border-radius: 3px;
}
.align-center{
  align-items: center;
}

</style>
