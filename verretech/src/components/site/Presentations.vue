<template>
  <div class="container">
    <div class="vx-row">
      <div class="w-1/2">
        <h1>Notre expertise</h1>
        <br>
        <p class="pr-5">
          Spécialisée et leader depuis 1955 dans la fabrication et la
          distribution en magasins de produits en verre techniques et
          décoratifs.
        </p>
      </div>
      <div class="w-1/2">
        <div class="vx-row">
          <div
            class="w-1/2 presImg"
            :style="{
              backgroundImage:
                'url(https://www.activitrerie.com/wp-content/uploads/Vitrerie.jpg)',
            }"
          ></div>

          <div
            class="w-1/2 presImg"
            :style="{
              backgroundImage:
                'url(https://upload.wikimedia.org/wikipedia/commons/thumb/c/cc/Vitrail_roccella_x.JPG/1200px-Vitrail_roccella_x.JPG)',
            }"
          ></div>
          <div
            class="w-full presImg mt-5"
            :style="{
              backgroundImage:
                'url(https://www.lavitrerie.fr/wp-content/uploads/2020/11/verre-640x430-2.jpg)',
            }"
          ></div>
        </div>
      </div>
    </div>
  </div>
</template>
<script>
export default {};
</script>
<style>
.presImg {
  width: 100%;
  height: 200px;
  background-size: cover;
  background-position: center;
}
</style>
