<template>
    <div>
      <h1>Create New Recipe</h1>
      <form @submit.prevent="createRecipe">
        <div>
          <label for="title">Title</label>
          <input v-model="recipe.title" id="title" placeholder="Recipe Title" required />
        </div>
        <div>
          <label for="description">Description</label>
          <textarea v-model="recipe.description" id="description" placeholder="Recipe Description" required></textarea>
        </div>
        <div>
          <label for="image">Recipe Image</label>
          <input type="file" @change="onImageUpload" />
        </div>
        <div>
          <label for="ingredients">Ingredients</label>
          <textarea v-model="recipe.ingredients" id="ingredients" placeholder="Ingredients" required></textarea>
        </div>
        <div>
          <label for="instructions">Instructions</label>
          <textarea v-model="recipe.instructions" id="instructions" placeholder="Instructions" required></textarea>
        </div>
        <button type="submit">Create Recipe</button>
      </form>
    </div>
  </template>
  
  <script>
  export default {
    data() {
      return {
        recipe: {
          title: '',
          description: '',
          image: null,
          ingredients: '',
          instructions: ''
        }
      };
    },
    methods: {
      onImageUpload(event) {
        this.recipe.image = event.target.files[0];
      },
      async createRecipe() {
        await this.$store.dispatch('recipe/createRecipe', this.recipe);
        this.$router.push('/admin/recipes');
      }
    }
  };
  </script>
  