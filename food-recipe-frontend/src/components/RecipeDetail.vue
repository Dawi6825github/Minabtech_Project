<template>
    <div v-if="recipe">
      <h1>{{ recipe.title }}</h1>
      <img :src="recipe.image" alt="Recipe Image" />
      <p>{{ recipe.description }}</p>
      <p>Ingredients: {{ recipe.ingredients }}</p>
      <p>Instructions: {{ recipe.instructions }}</p>
    </div>
  </template>
  
  <script>
  export default {
    data() {
      return {
        recipe: null
      };
    },
    async mounted() {
      const recipeId = this.$route.params.id;
      this.recipe = await this.$store.dispatch('recipe/fetchRecipe', recipeId);
    }
  };
  </script>
  