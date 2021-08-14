<template>
  <div class="container">
    <div class="vx-row ">
      <div class="w-3/4"></div>
      <div class="w-1/4 items-center">
        <span class="cursor-pointer leading-none pt-5 "> Notez ce produit</span>

        <star-rating
          class=" mt-2 mb-5 pl-3 border border-solid d-theme-border-grey-light border-t-0 border-b-0 border-r-0 float-right col-md-3"
          :read-only="false"
          :increment="0.01"
          :star-size="20"
          :rating="commentForm.rate"
          :fixed-points="2"
          :show-rating="false"
        ></star-rating>
      </div>
    </div>
    <div class="vx-row">
      <vs-textarea
        class="mt-5 w-full"
        placeholder="Donnez votre avis"
        v-model="commentForm.comment"
      ></vs-textarea>
    </div>
    <div class="vx-row">
      <div class="w-2/3"></div>
      <div class=" mt-5 mb-5 w-1/3">
        <vs-button @click="addComment" class="success w-full">Publier un commentaire</vs-button>
      </div>
    </div>
    <div class="container">
      <UserComment
        v-for="(comment, i) in comments"
        :key="i"
        :comment="comment.comment"
        :rate="comment.rate"
        :user="comment.user"
        class="mb-5"
      >
      </UserComment>
      
      
    </div>
  </div>
</template>
<script>
import StarRating from "vue-star-rating";
import UserComment from "@/components/site/commentaires/UserComment.vue";

export default {
  data() {
    return {
      commentForm: {
        user: "John Doe",
        comment:
          "",
        rate: 2.5,
      },
      comments: [
        {
          user: "John Doe",
          comment:
            "Lorem ipsum dolor sit, amet consectetur adipisicing elit. Rerum alias inventore dolorem. Non ipsam magnam inventore, nisi eaque nam omnis harum illo vero asperiores! Sequi eaque ducimus est aspernatur excepturi? ",
          rate: 4.3,
        },
      ],
    };
  },
  methods:{
    addComment(){
      this.comments.push(JSON.parse(JSON.stringify(this.commentForm)));
      this.commentForm = {
        user: "John Doe",
        comment:
          "",
        rate: 2.5,
      }
    }
  },
  components: {
    StarRating,
    UserComment,
  },
};
</script>
