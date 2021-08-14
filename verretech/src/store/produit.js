export const state = () => ({
    produtList: [],
   
})

export const getters = {
    produtList: (state) => state.produtList,
}

export const actions = {
    getProductList: ({ commit }, {list}) => {
            commit('getProductList')
    },

}

export const mutations = {
    getProductList: (state, {list}) => {
        state.produtList = true
    },
}