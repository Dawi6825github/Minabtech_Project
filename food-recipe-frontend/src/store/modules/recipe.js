// store/modules/recipe.js
export const state = () => ({
  recipes: [],  // List of recipes
  currentRecipe: null, // Current recipe for editing or viewing details
});

export const mutations = {
  setRecipes(state, recipes) {
    state.recipes = recipes;
  },
  setCurrentRecipe(state, recipe) {
    state.currentRecipe = recipe;
  },
  clearCurrentRecipe(state) {
    state.currentRecipe = null;
  },
};

export const actions = {
  async fetchRecipes({ commit }) {
    const response = await this.$axios.get('/api/recipes');
    commit('setRecipes', response.data);
  },

  async fetchRecipe({ commit }, id) {
    const response = await this.$axios.get(`/api/recipes/${id}`);
    commit('setCurrentRecipe', response.data);
  },

  async createRecipe({ commit }, recipeData) {
    const response = await this.$axios.post('/api/recipes', recipeData);
    commit('setRecipes', [...state.recipes, response.data]);
  },

  async updateRecipe({ commit }, { id, updatedData }) {
    const response = await this.$axios.put(`/api/recipes/${id}`, updatedData);
    commit('setCurrentRecipe', response.data);
  },

  async deleteRecipe({ commit }, id) {
    await this.$axios.delete(`/api/recipes/${id}`);
    commit('setRecipes', state.recipes.filter(recipe => recipe.id !== id));
  },
};

export const getters = {
  getRecipes(state) {
    return state.recipes;
  },
  getCurrentRecipe(state) {
    return state.currentRecipe;
  },
};
