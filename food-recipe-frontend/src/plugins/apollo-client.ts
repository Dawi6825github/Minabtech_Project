import { ApolloClient, InMemoryCache, HttpLink } from '@apollo/client/core'
export default defineNuxtPlugin(() => {
  const apolloClient = new ApolloClient({
    link: new HttpLink({ uri: 'https://your-graphql-endpoint.com/graphql' }),
    cache: new InMemoryCache(),
  })

  return {
    provide: {
      apollo: apolloClient
    }
  }

})