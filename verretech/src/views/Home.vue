<template>
	<div >
		<carousel class="pt-5"></carousel>
		<categorie></categorie>
		<CategorieFav></CategorieFav>	
		<ProductFav></ProductFav>	
	</div>
</template>
<script>
import carousel from "@/components/Carousel.vue"
import categorie from "@/components/site/Categorie.vue"
import CategorieFav from "@/components/site/CategorieFav.vue"
import ProductFav from "@/components/site/ProductFav.vue"
export default {
	components:{
		carousel,
		categorie,
		CategorieFav,
		ProductFav
	}
}
</script>