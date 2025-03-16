// plugins/apollo.js

import { ApolloClient, InMemoryCache, HttpLink } from '@apollo/client/core';

// Create the Apollo Client instance
const apolloClient = new ApolloClient({
  link: new HttpLink({
    uri: process.env.GRAPHQL_API_URL || 'https://your-graphql-endpoint.com/graphql',  // GraphQL API URL
    headers: {
      Authorization: `Bearer ${process.env.GRAPHQL_API_TOKEN || ''}`, // Optional token for auth
    },
  }),
  cache: new InMemoryCache(),  // Enable caching for improved performance
});

export default ({ app }, inject) => {
  // Inject the Apollo client into the Vue instance so you can access it globally
  inject('apollo', apolloClient);
};
