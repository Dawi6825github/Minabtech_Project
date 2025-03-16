import { defineNuxtPlugin } from "../../node_modules/nuxt/dist/app/nuxt.mjs";
import { ApolloClient } from "../../node_modules/_apollo/client/core/ApolloClient.mjs";
import { InMemoryCache } from "../../node_modules/_apollo/client/cache/inmemory/inMemoryCache.mjs";
import { HttpLink } from "../../node_modules/_apollo/client/link/http/HttpLink.mjs";
const apollo_client_iAM6nuJIHOGYamF3EUYO_DcTYtKIC1LHfzGAj8Byz1k = defineNuxtPlugin(() => {
  const apolloClient = new ApolloClient({
    link: new HttpLink({ uri: "https://your-graphql-endpoint.com/graphql" }),
    cache: new InMemoryCache()
  });
  return {
    provide: {
      apollo: apolloClient
    }
  };
});
export {
  apollo_client_iAM6nuJIHOGYamF3EUYO_DcTYtKIC1LHfzGAj8Byz1k as default
};
//# sourceMappingURL=apollo-client.mjs.map
