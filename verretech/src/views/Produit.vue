<template>
  <div>
    <vx-card>
      <div class="vx-row">
        <div style="background-repeat: no-repeat; background-position:center" :style="{
        backgroundImage:
          'url('+produit.Photos[0].Url+')',
      }" class="vx-col w-2/5 flex ">
          <!-- <img
            :src="produit.Photos[0].Url"
            class="responsive"
            alt=""
          /> -->
        </div>
        <div class="vs-col w-3/5">
          <h3>            {{produit.Nom}}
</h3>
          <p class="my-2"><span class="mr-2">by</span><span>VerreTech</span></p>
          <p class="flex items-center flex-wrap">
            <span
              class="text-2xl leading-none font-medium text-primary mr-4 mt-2"
              >{{produit.Prix}}€</span
            >
     
               <star-rating  
              class="pl-4 mr-2 mt-2 border border-solid d-theme-border-grey-light border-t-0 border-b-0 border-r-0"
                :read-only="true" :increment="0.01"  :star-size="20" :rating="4.4" :fixed-points="2"  :show-rating="false"></star-rating>  <span class="cursor-pointer leading-none mt-2 ml-2"> 424 ratings</span>

          </p>
          <vs-divider></vs-divider>
          <p>
            {{produit.Description}}


          </p>
          <vs-list>
            <vs-list-item
              icon-pack="feather"
              icon="icon-truck"
              title="Livraison offerte"
            ></vs-list-item>
            <vs-list-item
              icon-pack="feather"
              icon="icon-sliders"
              title="Personnalisation disponible"
            ></vs-list-item>
          </vs-list>
          <vs-divider></vs-divider>
          <div class="vx-row">
            <div class="vx-col w-full">
              <p class="my-2">
                <span>Disponible</span><span class="mx-2">-</span
                ><span class="text-success">En stock</span>
              </p>
            </div>
            <div class="vx-col w-full">
              <div class="flex flex-wrap items-start mb-4">
                <vs-button
                  color="primary"
                  type="filled"
                  icon-pack="feather"
                  icon="icon-cart"
                  >Ajouter au panier</vs-button
                >
                <vs-button
                  color="danger"
                  type="border"
                  icon-pack="feather"
                  icon="icon-heart"
                  class="ml-4"
                  >Ajouter à la liste de souhait</vs-button
                >
              </div>
            </div>
          </div>
        </div>
      </div>
      <div class="mt-5">
        <vs-divider>Avis des clients</vs-divider>
       <commentsProduct ></commentsProduct>
      </div>
    </vx-card>
  </div>
</template>
<script>
import StarRating from 'vue-star-rating'
import axios from "axios"
import CommentsProduct from '../components/site/commentaires/CommentsProduct.vue';
export default {
  data() {
    return {
      produit : {}
    }
  },
  components: {

    CommentsProduct,
    StarRating
  },
  created() {
     axios.get('http://localhost:10000/produit/'+this.$route.params.ref).then((res)=>{
        this.produit = res.data
      })
  },
};
</script>
