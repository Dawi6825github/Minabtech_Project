<template>
    <form @submit.prevent="submitRecipe">
      <input v-model="recipe.title" placeholder="Recipe Title" required />
      <textarea v-model="recipe.description" placeholder="Recipe Description" required></textarea>
      <input type="file" @change="onImageUpload" />
      <button type="submit">Submit</button>
    </form>
  </template>
  
  <script>
  export default {
    data() {
      return {
        recipe: {
          title: '',
          description: '',
          image: null
        }
      };
    },
    methods: {
      submitRecipe() {
        this.$store.dispatch('recipe/createRecipe', this.recipe);
      },
      onImageUpload(event) {
        this.recipe.image = event.target.files[0];
      }
    }
  };
  </script>
  