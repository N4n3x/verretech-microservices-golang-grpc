<template>
  <form-wizard
    color="rgba(var(--vs-primary), 1)"
    :title="null"
    nextButtonText="Suivant"
    :subtitle="null"
    finishButtonText="Confimer ma commande"
     @on-complete="formSubmitted"
  >

    <tab-content title="Panier" class="mb-5" icon="feather icon-home">
      <!-- tab 1 content -->
      <div class="vx-row">
        <div class="vx-col lg:w-2/3 w-full relative">
          <div class="items-list-view">
            <item-list-view
              v-for="(produit, index) in articles"
              :key="index"
              class="mb-base"
              :item="produit"
            >
              <!-- SLOT: ITEM META -->
              <template slot="item-meta">
                <h6
                  class="item-name font-semibold mb-1 cursor-pointer hover:text-primary"
                >
                  {{produit.Nom}}
                </h6>
                <p class="text-sm mb-2">
                  Par
                  <span class="font-semibold cursor-pointer">VerreTech</span>
                </p>
                <p class="text-success text-sm">En stock</p>

                <p class="mt-4 font-bold text-sm">Quantité</p>
                <vs-input-number
                  min="1"
                
                  :value="produit.Qte"
                  class="inline-flex"
                />

                <!-- <p class="font-medium text-grey mt-4">
                  Disponible le 30/04/2021
                </p>
                <p class="text-success font-medium">
                  10% offert
                </p> -->
              </template>

              <!-- SLOT: ACTION BUTTONS -->
              <template slot="action-buttons">
                <!-- PRIMARY BUTTON: REMOVE -->
                <div
                  class="item-view-primary-action-btn p-3 rounded-lg flex flex-grow items-center justify-center cursor-pointer mb-3"
                >
                  <feather-icon icon="XIcon" svgClasses="h-4 w-4" />
                  <span class="text-sm font-semibold ml-2">Supprimer</span>
                </div>

                <!-- SECONDARY BUTTON: MOVE-TO/VIEW-IN WISHLIST -->
                <div
                  class="item-view-secondary-action-btn bg-primary p-3 rounded-lg flex flex-grow items-center justify-center text-white cursor-pointer"
                >
                  <feather-icon icon="HeartIcon" />
                  <span class="text-sm font-semibold ml-2"
                    >Liste de souhait</span
                  >
                </div>
              </template>
            </item-list-view>
          </div>
        </div>

        <!-- RIGHT COL -->
        <div class="vx-col lg:w-1/3 w-full">
          <vx-card>
           

            

            <p class="font-semibold mb-3">Détails du prix</p>
            <div class="flex justify-between mb-2">
              <span class="text-grey">Total HT</span>
              <span>{{htPrice}}€</span>
            </div>
         
            <div class="flex justify-between mb-2">
              <span class="text-grey">TVA</span>
              <span>{{tva}}€</span>
            </div>
            <!-- <div class="flex justify-between mb-2">
              <span class="text-grey">EMI Eligibility</span>
              <a href="#" class="text-primary">Details</a>
            </div> -->
            <!-- <div class="flex justify-between mb-2">
              <span class="text-grey">Frais de livraison</span>
              <span class="text-success">Gratuit</span>
            </div> -->

            <vs-divider />

            <div class="flex justify-between font-semibold mb-3">
              <span>Total</span>
              <span>{{totalPrice}}€</span>
            </div>
          </vx-card>
        </div>
      </div>
    </tab-content>

    <!-- tab 2 content -->
    <!-- <tab-content
      title="Vos informations"
      class="mb-5"
      icon="feather icon-briefcase"
    >
      <div class="vx-row">
        <div class="vx-col md:w-1/2">
          <vs-input label="Prénom" class="w-full" />
          <vs-input label="Adresse de facturation" class="w-full" />
        </div>
        <div class="vx-col md:w-1/2">
          <vs-input label="Nom" class="w-full" v-model="jobTitle" />
          <vs-input
            label="Complément d'adresse"
            class="w-full"
            v-model="jobTitle"
          />
        </div>
      </div>
      <div class="vx-row">
        <div class="vx-col md:w-1/2">
          <label for="">Lieu de retrait</label>
          <v-select :options="[{ label: 'Rouen', value: 'rouen' }]" />
        </div>
      </div>
    </tab-content> -->

    <!-- tab 3 content -->
  
  </form-wizard>
</template>

<script>
import { FormWizard, TabContent } from "vue-form-wizard";
import "vue-form-wizard/dist/vue-form-wizard.min.css";
const ItemListView = () => import("./components/ItemListView.vue");
import vSelect from "vue-select";
import axios from "axios";
export default {
  data() {
    return {
      firstName: "",
      panier: {},
      articles: [],
      commandValid: false
      /// other data....
    };
  },
  methods: {
    formSubmitted() {
      console.log('sbmited')
        axios
        .get("http://localhost:10000/commande/validation", {
          auth: {
            username:"un@mail.com",
            password:"motdepasse"
          }
        })
        .then((res) => {
          console.log(res)
          this.commandValid  = true;
        }).catch(err=>{
          console.error(err)
        });
    },
  },
  created() {
    console.log(JSON.parse(localStorage.panier).Articles);
    this.panier = JSON.parse(localStorage.panier);
    for (let i = 0; i < this.panier.Articles.length; i++) {
      const element = this.panier.Articles[i];
      axios
        .get("http://localhost:10000/produit/" + element.ProduitRef)
        .then((res) => {
          var produit = res.data;
          produit.Qte = element.Qte
          this.articles.push(produit)
        });
    }
  },
  computed:{
    tva(){
      
      return (this.totalPrice-this.htPrice).toFixed(2)
    },
    htPrice(){
       var total = 0;
      for (let i = 0; i < this.articles.length; i++) {
        const element = this.articles[i];
        total+= (element.Prix)*element.Qte
      }
      return total.toFixed(2)
    },
    totalPrice(){
      var total = 0;
      for (let i = 0; i < this.articles.length; i++) {
        const element = this.articles[i];
               total+= (element.Prix*1.2)*element.Qte

      }
      return total.toFixed(2)
    }
  },
  components: {
    FormWizard,
    TabContent,
    ItemListView,

    "v-select": vSelect,
  },
};
</script>
