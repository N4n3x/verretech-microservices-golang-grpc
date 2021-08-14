<template>
  <div>
    <div class="vx-row">
      <div class="col-md-3">
        <h4>{{user}}</h4>
        <span>Il y a 5 heures</span>
      </div>
      <div class="col-md-3">
        <star-rating
          class="pl-4 mr-2 mt-2 mb-5 border border-solid d-theme-border-grey-light border-t-0 border-b-0 border-r-0 float-right ml-5"
          :read-only="true"
          :increment="0.01"
          :star-size="20"
          :rating="rate"
          :fixed-points="2"
          :show-rating="false"
        ></star-rating>
      </div>
    </div>
    <div class="vx-row">
      <div class="w-full">
          {{comment}}
      </div>
    </div>
    <vs-divider></vs-divider>
  </div>
</template>
<script>
import StarRating from "vue-star-rating";

export default {
    props:{
        comment: {
            type: String
        },
          user: {
            type: String
        },
        rate:{
            type: Number
        }
    },
  data() {
    return {};
  },
  components: {
    StarRating,
  },
};
</script>
