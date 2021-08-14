<template>
  <router-link :to="{ name: 'Produit', params: {ref: produit.Ref } }">
    <div
      class="imgPreview"
      :style="{
        backgroundImage:
          'url('+produit.Photos[0].Url+')',
      }"
    >
    <div class="hoverProduct">
        <p>{{produit.Description}}</p>
    </div>
    </div>
    <div class="vx-row m-1">
      <div class="w-full titleProduct">
        {{produit.Nom}}
      </div>
      <div class="w-full">
        {{produit.Prix}} €
      </div>
    </div>

  </router-link> 
</template>
<script>
export default {
  data() {
    return {
      id: 1
    }
  },
  props:{
    produit:{
      type: Object
    }
  },
  methods: {
    product(){
      this.$router.push("Produit")
    }
  },
};
</script>
<style>
.imgPreview {
    position: relative;
  width: 250px;
  height: 200px;
  background-size: cover;
  background-position: center;
  border: solid 3px #2e578c;
}
.hoverProduct{
    height: 0;
}
.imgPreview:hover > .hoverProduct {
    position: absolute;
    height: 70%;
    width: 101%;
    bottom:0;
    backdrop-filter: blur(2px);
    background-color: #00000082;
     transition: all 0.2s;
}
.hoverProduct p{
    display: none;
    color:white;
    padding:5px;
}
.imgPreview:hover > .hoverProduct p {
    display:contents
}
.titleProduct {
  font-weight: bold;
}
</style>
