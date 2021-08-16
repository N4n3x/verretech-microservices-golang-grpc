<!-- =========================================================================================
    File Name: ItemListView.vue
    Description: Item Component - List VIew
    ----------------------------------------------------------------------------------------
    Item Name: Vuexy - Vuejs, HTML & Laravel Admin Dashboard Template
      Author: Pixinvent
    Author URL: http://www.themeforest.net/user/pixinvent
========================================================================================== -->

<template>
  <vx-card class="list-view-item overflow-hidden">
    <template slot="no-body">
      <div class="vx-row item-details no-gutter">
        <!-- IMG COL -->
        <div
          class="vx-col sm:w-1/4 w-full item-img-container bg-white flex items-center justify-center cursor-pointer"
        >
          <img
            :src="item.Photos[0].Url"
            :alt="item.Nom"
            class="grid-view-img p-4"
          />
        </div>

        <!-- ITEM NAME & DESC COL -->
        <div class="vx-col sm:w-1/2">
          <div class="p-4 pt-6">
            <slot name="item-meta">
              <h6
                class="item-name font-semibold mb-1 hover:text-primary cursor-pointer"
              >
                {{ item.Nom }}
              </h6>
              <p class="text-sm mb-4">
                By <span class="font-semibold cursor-pointer"> VerreTech </span>
              </p>
              <p class="item-description text-sm">{{ item.Description }}.</p>
            </slot>
          </div>
        </div>

        <!-- PURCHASE COL -->
        <div
          class="vx-col sm:w-1/4 w-full flex items-center sm:border border-0 border-solid d-theme-border-grey-light border-r-0 border-t-0 border-b-0"
        >
          <div class="p-4 flex flex-col w-full">
            <div class="row">
              <div class="col" v-for="(tag, i) in item.Tags" :key="i">
                <vs-chip color="primary">{{ tag }}</vs-chip>
              </div>
            </div>
            <!-- <div class="text-warning flex self-end border border-solid border-warning py-1 px-2 rounded">
                            <span class="text-sm mr-1">4.2</span>
                            <feather-icon icon="StarIcon" svgClasses="h-4 w-4" />
                        </div> -->

            <div class="my-6">
              <h5 class="font-bold text-center">
                {{ (item.Prix * 1.2 * item.Qte).toFixed(2) }}€
              </h5>
            </div>

            <!-- SLOT: ACTION BUTTONS -->
            <slot name="action-buttons" />
          </div>
        </div>
      </div>
    </template>
  </vx-card>
</template>

<script>
export default {
  props: {
    item: {
      type: Object,
      required: true,
    },
  },
  computed: {},
  methods: {},
};
</script>

<style lang="scss" scoped>
.list-view-item {
  .item-name,
  .item-description {
    display: -webkit-box;
    -webkit-box-orient: vertical;
    overflow: hidden;
  }

  .item-name {
    -webkit-line-clamp: 2;
  }

  .item-description {
    -webkit-line-clamp: 5;
  }

  .grid-view-img {
    max-width: 100%;
    max-height: 100%;
    width: auto;
    transition: 0.35s;
  }

  &:hover {
    transform: translateY(-5px);
    box-shadow: 0px 4px 25px 0px rgba(0, 0, 0, 0.25);

    .grid-view-img {
      opacity: 0.9;
    }
  }
}
</style>
