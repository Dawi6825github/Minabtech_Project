<template>
    <div>
      <h2 class="text-2xl font-bold">Recipes</h2>
      <ul>
        <li v-for="recipe in recipes" :key="recipe.id" @click="goToRecipe(recipe.id)">
          {{ recipe.title }}
        </li>
      </ul>
    </div>
  </template>
  
  <script>
  import { useQuery } from '@vue/apollo-composable';
  import { GET_RECIPES } from '../graphql/queries';
  import { useRouter } from 'vue-router';
  
  export default {
    setup() {
      const { result } = useQuery(GET_RECIPES);
      const router = useRouter();
  
      function goToRecipe(id) {
        router.push(`/recipe/${id}`);
      }
  
      return { recipes: result };
    },
  };
  </script>
  