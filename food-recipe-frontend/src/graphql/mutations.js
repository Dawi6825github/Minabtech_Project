import { gql } from '@apollo/client/core';

export const CREATE_RECIPE = gql`
  mutation CreateRecipe($title: String!, $ingredients: [String!]!, $instructions: String!, $category: String!) {
    createRecipe(title: $title, ingredients: $ingredients, instructions: $instructions, category: $category) {
      id
      title
    }
  }
`;
