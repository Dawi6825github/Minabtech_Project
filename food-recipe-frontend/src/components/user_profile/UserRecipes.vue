<template>
    <div class="user-recipes">
      <h2>My Recipes</h2>
      <ul>
        <li v-for="recipe in recipes" :key="recipe.id">
          <router-link :to="`/recipes/${recipe.id}`">{{ recipe.title }}</router-link>
          <button @click="editRecipe(recipe.id)">Edit</button>
          <button @click="deleteRecipe(recipe.id)">Delete</button>
        </li>
      </ul>
    </div>
  </template>
  
  <script>
  export default {
    props: {
      recipes: Array, // List of recipes passed from parent component
    },
    methods: {
      editRecipe(id) {
        this.$router.push(`/recipes/edit/${id}`);
      },
      deleteRecipe(id) {
        // Handle recipe deletion
      },
    },
  };
  </script>
  
  <style scoped>
  .user-recipes {
    width: 100%;
    max-width: 800px;
    margin: 20px 0;
  }
  
  button {
    margin-left: 10px;
  }
  </style>
  