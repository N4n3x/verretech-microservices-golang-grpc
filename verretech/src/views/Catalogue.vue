<template>
  <div class="vx-row container">
  
    <div class="w-1/3">
        <panel-filter class="panelFilter"></panel-filter>
    </div>
    <div class="w-2/3">
      <div class="vx-row">
        <div class="w-1/3" v-for="(produit, index) in productList" :key="index">
          <produit-preview :produit="produit"></produit-preview>
        </div>
      </div>
    </div>
  </div>
</template>
<script>
import panelFilter from "@/components/site/PanelFilter";
import produitPreview from "@/components/site/ProduitPreview";

export default {
  data() {
    return {
      length: 12,
    };
  },
  created() {
    this.$store.dispatch('updateProductList')
    
  },
   computed: {
    productList() {
      return this.$store.state.productList;
    }
  },
  components: {
    produitPreview,
    panelFilter
  },
};
</script>
<style >
.vx-row .container{
  margin-top: 200px;
}
    .panelFilter{
        width: 80%;
    }
</style>