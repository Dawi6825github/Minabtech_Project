// /src/graphql/queries.js
import { gql } from '@apollo/client/core';

export const GET_RECIPES = gql`
  query GetRecipes {
    recipes {
      id
      title
      ingredients
      instructions
      recipeCategory
    }
  }
`;

export const GET_RECIPE_BY_ID = gql`
  query GetRecipeById($id: ID!) {
    recipe(id: $id) {
      id
      title
      ingredients
      instructions
      recipeCategory
    }
  }
`;
