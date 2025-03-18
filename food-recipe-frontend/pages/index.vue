<template>
  <div>
    <h1 class="text-3xl font-bold">Welcome to the Food Recipe Website</h1>
    <div>
      <h2 class="text-2xl">Browse Recipes</h2>
      <div v-for="recipe in recipes" :key="recipe.id" class="recipe-card">
        <h3>{{ recipe.title }}</h3>
        <p>{{ recipe.description }}</p>
        <nuxt-link :to="`/recipe/${recipe.id}`">View Recipe</nuxt-link>
      </div>
    </div>
  </div>
</template>

<script>
export default {
  data() {
    return {
      recipes: [], // This will hold the list of recipes
    };
  },
    async mounted() {
      try {
        const response = await this.$axios.$get('/api/recipes'); // Adjust the endpoint as necessary
        this.recipes = response.data;
      } catch (error) {
        console.error('Error fetching recipes:', error);
      }

  },
};
</script>

<style scoped>
.recipe-card {
  border: 1px solid #ccc;
  padding: 16px;
  margin: 8px 0;
}
</style>
