

// Couleurs 
let colors = {
  primary : '#2E578C',
  success : '#28C76F',
  danger  : '#EA5455',
  warning : '#FF9F43',
  dark    : '#1E1E1E',
}

// paramêtres d'affichages
const themeConfig = {
  disableCustomizer : false,       // options[Boolean] : true, false(default)
  disableThemeTour  : false,       // options[Boolean] : true, false(default)
  footerType        : "static",    // options[String]  : static(default) / sticky / hidden
  hideScrollToTop   : false,       // options[Boolean] : true, false(default)
  mainLayoutType    : "horizontal",  // options[String]  : vertical(default) / horizontal
  navbarColor       : "#fff",      // options[String]  : HEX color / rgb / rgba 
  navbarType        : "sticky",  // options[String]  : floating(default) / static / sticky / hidden
  routerTransition  : "zoom-fade", // options[String]  : zoom-fade / slide-fade / fade-bottom / fade / zoom-out / none(default)
  rtl               : false,       // options[Boolean] : true, false(default)
  sidebarCollapsed  : false,       // options[Boolean] : true, false(default)
  theme             : "light",     // options[String]  : "light"(default), "dark", "semi-dark"


  userInfoLocalStorageKey: "userInfo",

}

import Vue from 'vue'
import Vuesax from 'vuesax'
Vue.use(Vuesax, { theme:{ colors }, rtl: themeConfig.rtl })

export default themeConfig
