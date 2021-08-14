<template>
  <div class="panelFilter">
    <!-- RECHERCHE -->
    <div class="vx-row mb-6">
      <div class="vx-col w-full">
        <vs-input
          class="w-full"
          icon-pack="feather"
          icon="icon-search"
          icon-no-border
          label="Rechercher"
          v-model="search"
        />
      </div>
    </div>
    <!-- CATEGORIE -->
    <div class="vx-row mb-6">
      <div class="vx-col w-full">
        <span>Catégories :</span>
        <div class="vx-row">
          <div
            :key="index"
            v-for="(categorie, index) in categories"
            class="vx-col w-full"
          >
            <vs-checkbox class="categorie" v-model="categorie.value">{{
              categorie.label
            }}</vs-checkbox>
          </div>
        </div>
      </div>
    </div>
    <vs-divider>Prix</vs-divider>
    <!-- PRIX -->
    <div class="vx-row mb-6">
      <div class="vx-col w-1/6">{{ minMaxPrice[0] }}€</div>
      <div class="vx-col w-2/3">
        <vue-slider v-model="minMaxPrice" :order="false"></vue-slider>
      </div>
      <div class="vx-col w-1/6">{{ minMaxPrice[1] }}€</div>
    </div>
    <vs-divider></vs-divider>
    <div class="vx-col w-full">
      <span>Tags :</span>
      <div class="vx-row">
        <div class="vx-col w-full">
          <v-select
            multiple
            :closeOnSelect="false"
            v-model="selected"
            :options="options"
            :dir="$vs.rtl ? 'rtl' : 'ltr'"
          /><br />
        </div>
      </div>
    </div>
  </div>
</template>
<script>
import VueSlider from "vue-slider-component";
import "vue-slider-component/theme/default.css";
import vSelect from "vue-select";
export default {
  data() {
    return {
      selected: ["Transparant", "Gravure"],
      options: ["Coloré", "Transparant", "Gravure", "SDB"],
      minMaxPrice: [0, 50],
      search: "",
      categories: [
        {
          value: false,
          label: "Verre décoratif",
        },
        {
          value: false,
          label: "Verre décoratif",
        },
        {
          value: false,
          label: "Verre décoratif",
        },
        {
          value: false,
          label: "Verre décoratif",
        },
        {
          value: false,
          label: "Verre décoratif",
        },
      ],
      prixMax: 0,
    };
  },
  components: {
    VueSlider,
    "v-select": vSelect,
  },
};
</script>
<style>
.panelFilter {
  border: solid 3px #2e578c;
  background: #f1f1f1;
  padding: 20px;
}
.categorie {
  float: right;
  margin: 5px;
}
</style>
