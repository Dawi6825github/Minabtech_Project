<template>
  <div>
    <h1 class="text-3xl font-bold">{{ recipe.title }}</h1>
    <img :src="recipe.featuredImage" alt="Featured Image" />
    <p>{{ recipe.description }}</p>
    <h2>Ingredients</h2>
    <ul>
      <li v-for="ingredient in recipe.ingredients" :key="ingredient.id">{{ ingredient.name }}</li>
    </ul>
    <h2>Preparation Steps</h2>
    <ol>
      <li v-for="step in recipe.steps" :key="step.id">{{ step.description }}</li>
    </ol>
    <button @click="likeRecipe">Like</button>
    <button @click="bookmarkRecipe">Bookmark</button>
    <div>
      <h3>Comments</h3>
      <div v-for="comment in recipe.comments" :key="comment.id">
        <p><strong>{{ comment.user }}</strong>: {{ comment.text }}</p>
      </div>
      <form @submit.prevent="addComment">
        <input v-model="newComment" placeholder="Add a comment" />
        <button type="submit">Submit</button>
      </form>
    </div>
  </div>
</template>

<script>
export default {
  data() {
    return {
      recipe: {},
      newComment: '',
    };
  },
  async mounted() {
    const recipeId = this.$route.params.id;
    const response = await this.$axios.$get(`/api/recipes/${recipeId}`); // Adjust the endpoint as necessary
    this.recipe = response.data;
  },
  methods: {
    async likeRecipe() {
      // Logic to like the recipe
    },
    async bookmarkRecipe() {
      // Logic to bookmark the recipe
    },
    async addComment() {
      // Logic to add a comment
      const response = await this.$axios.$post(`/api/recipes/${this.recipe.id}/comments`, {
        text: this.newComment,
      });
      this.recipe.comments.push(response.data);
      this.newComment = '';
    },
  },
};
</script>

<style scoped>
/* Add any necessary styles here */
</style>
