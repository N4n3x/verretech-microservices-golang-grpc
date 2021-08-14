<template>
  <div>
    <vs-input
      v-validate="'required|email|min:3'"
      data-vv-validate-on="blur"
      name="email"
      icon-no-border
      icon="icon icon-user"
      icon-pack="feather"
      label-placeholder="Email"
      v-model="email"
      class="w-full"
    />
    <span class="text-danger text-sm">{{ errors.first('email') }}</span>

    <vs-input
      data-vv-validate-on="blur"
      v-validate="'required|min:6|max:10'"
      type="password"
      name="password"
      icon-no-border
      icon="icon icon-lock"
      icon-pack="feather"
      label-placeholder="Password"
      v-model="password"
      class="w-full mt-6"
    />
    <span class="text-danger text-sm">{{ errors.first('password') }}</span>

    <div class="flex flex-wrap justify-between my-5">
      <vs-checkbox v-model="checkbox_remember_me" class="mb-3">Se souvenir de moi</vs-checkbox>
      <router-link to="/pages/forgot-password">Mot de passe oublié ?</router-link>
    </div>
    <vs-button
      class="float-right"
      :disabled="!validateForm"
      @click="login()"
    >Login</vs-button>
  </div>
</template>

<script>
// FIREBASE
import "firebase/auth";
const firebase = require("firebase");
require("firebase/firestore");
require("firebase/auth");
var db = firebase.firestore();
import { firebaseConfig } from "@/firebase/firebaseConfig";

export default {
  data() {
    return {
      email: "",
      password: "",
      checkbox_remember_me: false,
    };
  },
  computed: {
    validateForm() {
      return !this.errors.any() && this.email !== "" && this.password !== "";
    }
  },
  methods: {
    checkLogin() {
      // If user is already logged in notify
      if (this.$store.state.auth.isUserLoggedIn()) {
        // Close animation if passed as payload
        // this.$vs.loading.close()

        this.$vs.notify({
          title: "Tentative de connexion",
          text: "Vous êtes déjà connecté",
          iconPack: "feather",
          icon: "icon-alert-circle",
          color: "warning"
        });

        return false;
      }
      return true;
    },
    login() {
      // Loading
      this.$vs.loading();
      const payload = {
        checkbox_remember_me: this.checkbox_remember_me,
        userDetails: {
          email: this.email,
          password: this.password
        },
        notify: this.$vs.notify,
        closeAnimation: this.$vs.loading.close
      };
      this.$store.dispatch("auth/loginAttempt", payload);
    },

    // Google login
    loginWithGoogle() {
      this.$store.dispatch("auth/loginWithGoogle", { notify: this.$vs.notify });
    },

    // Facebook login
    loginWithFacebook() {
      this.$store.dispatch("auth/loginWithFacebook", {
        notify: this.$vs.notify
      });
    },

    // Twitter login
    loginWithTwitter() {
      this.$store.dispatch("auth/loginWithTwitter", {
        notify: this.$vs.notify
      });
    },

    // Github login
    loginWithGithub() {
      this.$store.dispatch("auth/loginWithGithub", { notify: this.$vs.notify });
    },
    registerUser() {
      if (!this.checkLogin()) return;
      this.$router.push("/pages/register").catch(() => {});
    }
  }
};
</script>


<style lang="scss">
#page-login {
  .social-login-buttons {
    .bg-facebook {
      background-color: #1551b1;
    }
    .bg-twitter {
      background-color: #00aaff;
    }
    .bg-google {
      background-color: #4285f4;
    }
    .bg-github {
      background-color: #333;
    }
  }
}
</style>
