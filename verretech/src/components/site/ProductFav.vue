<template>
  <div class="bandeau ">
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
    <div class="justify-center">
      <div class="box  space-around">
        <div class="divBg">
          <div class="bg"></div>
          <div class="p-5 ml-35">
            <span class="bold">Cardosa</span><br />
            <span>449.00€</span>
          </div>
        </div>
        <div class="divBg">
          <div class="bg"></div>
          <div class="p-5 ml-35">
            <span class="bold">Cardosa</span><br />
            <span>11 999.99€</span>
          </div>
        </div>
        <div class="divBg">
          <div class="bg"></div>
          <div class="p-5 ml-35">
            <span class="bold">Cardosa</span> <br />
            <span>3 999.99€</span>
          </div>
        </div>
      </div>
    </div>
    <div class="justify-center mt-5">
      <vs-button>Voir les produits</vs-button>
    </div>
  </div>
</template>
<script>
export default {
  data() {
    return {};
  },
};
</script>
<style scoped>
.bandeau {
  border-top: solid 8px #c2c2c2;
  background: rgba(238, 118, 95, 0.75);
  padding-bottom: 50px;
  /* min-height: 700px; */
}
.justify-center {
  display: flex;
  justify-content: center;
}
h1 {
  font-size: 3em;
  font-weight: bold;
}
.fs-1-5 {
  font-size: 1.5em;
}
.box {
  background: rgba(255, 255, 255, 0.3);
  /* height: 250px; */
  width: 80%;
}
.space-around {
  display: flex;
  justify-content: space-around;
}
.bg {
  background-image: url("https://www.anywaydoors.be/sites/default/files/styles/galery_big/public/fotogalerij/cloison-en-verre-style-atelier.jpg?itok=YABCHcIm");
  height: 300px;
  margin: 50px;
  margin-bottom: 0px;
  background-size: cover;
  background-position: center;
  border-radius: 4px;
}
.divBg {
  width: 50%;
  display: flex;
  flex-direction: column;
}
.ml-35 {
  margin-left: 35px;
}
span {
  font-size: 1.2em;
}
.bold {
  font-weight: bold;
}
</style>
