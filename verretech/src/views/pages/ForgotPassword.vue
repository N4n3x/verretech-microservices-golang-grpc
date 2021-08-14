<!-- =========================================================================================
    File Name: ForgotPassword.vue
    Description: FOrgot Password Page
    ----------------------------------------------------------------------------------------
    Item Name: Vuexy - Vuejs, HTML & Laravel Admin Dashboard Template
      Author: Pixinvent
    Author URL: http://www.themeforest.net/user/pixinvent
========================================================================================== -->


<template>
  <div class="h-screen flex w-full bg-img" id="page-forgot-password">
    <div class="vx-col w-4/5 sm:w-4/5 md:w-3/5 lg:w-3/4 xl:w-3/5 mx-auto self-center">
      <vx-card>
        <div slot="no-body" class="full-page-bg-color">
          <div class="vx-row">
            <div class="vx-col hidden sm:hidden md:hidden lg:block lg:w-1/2 mx-auto self-center">
              <img
                id="logoForgotPassword"
                src="@/assets/images/pages/forgot-password.png"
                alt="login"
                class="mx-auto"
              />
            </div>
            <div class="vx-col sm:w-full md:w-full lg:w-1/2 mx-auto self-center d-theme-dark-bg">
              <div class="p-8">
                <div class="vx-card__title mb-8">
                  <h4 class="mb-4">Mot de passe oublié</h4>
                  <p>Entrez votre email, vous recevrez les instructions pour réinitialiser votre mot de passe.</p>
                </div>

                <vs-input
                  type="email"
                  label-placeholder="Email"
                  v-model="email"
                  class="w-full mb-8"
                />
                <vs-button type="border" to="/pages/login" class="px-4 w-full md:w-auto">Retour</vs-button>
                <vs-button
                  class="float-right px-4 w-full md:w-auto mt-3 mb-8 md:mt-0 md:mb-0"
                  @click="sendEmailToResetPassword(email)"
                >Envoyer</vs-button>
              </div>
            </div>
          </div>
        </div>
      </vx-card>
    </div>
  </div>
</template>

<script>
// FIREBASE
const firebase = require("firebase");
require("firebase/firestore");
require("firebase/auth");
var db = firebase.firestore();

export default {
  data() {
    return {
      email: "",
      actionCodeSettings: ""
    };
  },
  methods: {
    sendEmailToResetPassword(email, actionCodeSettings) {
      var status = false;
      firebase
        .auth()
        .sendPasswordResetEmail(email, actionCodeSettings)
        .then(response => this.displayMessage("OK"))
        .catch(error => this.displayMessage("KO", error));
    },
    displayMessage(state, error = null) {
      if (state == "OK") {
        this.$vs.notify({
          title: "Email envoyé",
          text: "Un email vous a été envoyé par email",
          iconPack: "feather",
          icon: "icon-check-circle",
          color: "success"
        });
        setTimeout(() => {
          this.$router.push("/pages/login");
        }, 2000);
      } else if (state == "KO") {
        console.error(error);
        this.$vs.notify({
          title: "Erreur",
          text: "Une erreur est survenu, veuillez contacter le support",
          iconPack: "feather",
          icon: "icon-warning-circle",
          color: "danger"
        });
        setTimeout(() => {
          this.$router.push("/pages/login");
        }, 2000);
      }
    }
  }
};
</script>
