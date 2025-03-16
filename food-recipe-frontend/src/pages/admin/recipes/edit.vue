<template>
    <div>
      <h1>Edit Recipe</h1>
      <form @submit.prevent="updateRecipe">
        <div>
          <label for="title">Title</label>
          <input v-model="recipe.title" id="title" required />
        </div>
        <div>
          <label for="description">Description</label>
          <textarea v-model="recipe.description" id="description" required></textarea>
        </div>
        <div>
          <label for="image">Recipe Image</label>
          <input type="file" @change="onImageUpload" />
        </div>
        <div>
          <label for="ingredients">Ingredients</label>
          <textarea v-model="recipe.ingredients" id="ingredients" required></textarea>
        </div>
        <div>
          <label for="instructions">Instructions</label>
          <textarea v-model="recipe.instructions" id="instructions" required></textarea>
        </div>
        <button type="submit">Update Recipe</button>
      </form>
    </div>
  </template>
  
  <script>
  export default {
    data() {
      return {
        recipe: {
          id: null,
          title: '',
          description: '',
          image: null,
          ingredients: '',
          instructions: ''
        }
      };
    },
    async mounted() {
      const recipeId = this.$route.params.id;
      this.recipe = await this.$store.dispatch('recipe/fetchRecipe', recipeId);
    },
    methods: {
      onImageUpload(event) {
        this.recipe.image = event.target.files[0];
      },
      async updateRecipe() {
        await this.$store.dispatch('recipe/updateRecipe', this.recipe);
        this.$router.push('/admin/recipes');
      }
    }
  };
  </script>
  