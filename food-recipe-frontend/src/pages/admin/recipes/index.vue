<template>
    <div>
      <h1>All Recipes</h1>
      <div v-if="recipes.length === 0">No recipes available.</div>
      <ul v-else>
        <li v-for="recipe in recipes" :key="recipe.id">
          <div class="recipe-item">
            <img :src="recipe.image" alt="Recipe Image" class="recipe-image" />
            <div>
              <h3>{{ recipe.title }}</h3>
              <p>{{ recipe.description }}</p>
              <router-link :to="`/admin/recipes/edit/${recipe.id}`">Edit</router-link>
              <button @click="deleteRecipe(recipe.id)">Delete</button>
            </div>
          </div>
        </li>
      </ul>
    </div>
  </template>
  
  <script>
  export default {
    data() {
      return {
        recipes: []
      };
    },
    async mounted() {
      this.recipes = await this.$store.dispatch('recipe/fetchRecipes');
    },
    methods: {
      async deleteRecipe(id) {
        const confirmed = confirm('Are you sure you want to delete this recipe?');
        if (confirmed) {
          await this.$store.dispatch('recipe/deleteRecipe', id);
          this.recipes = this.recipes.filter(recipe => recipe.id !== id); // Update the list after delete
        }
      }
    }
  };
  </script>
  
  <style scoped>
  .recipe-item {
    display: flex;
    margin-bottom: 20px;
    align-items: center;
  }
  .recipe-image {
    width: 100px;
    height: 100px;
    object-fit: cover;
    margin-right: 20px;
  }
  </style>
  