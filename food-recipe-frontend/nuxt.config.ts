export default {
  // Disable server-side rendering (default in Nuxt 3)
  ssr: true,

  // Target (only for `nuxt generate`)
  target: 'server',

  // Modules
  modules: [
    '@nuxtjs/axios',
    '@nuxtjs/auth-next',
    '@nuxtjs/cloudinary',
    'vee-validate',
    '@nuxtjs/apollo', // Added Apollo module
  ],

  // Apollo module configuration
  apollo: {
    clientConfigs: {
      default: {
        httpEndpoint: 'http://localhost:8080/v1/graphql', // Replace with your Hasura endpoint
      },
    },
  },

  // Axios module configuration
  axios: {
    baseURL: 'http://localhost:3000', // Replace with your API base URL
    timeout: 10000, // Set a timeout for requests
    headers: {
      common: {
        Accept: 'application/json',
        'Content-Type': 'application/json',
      },
    },
  },

  // Auth module configuration
  auth: {
    strategies: {
      local: {
        token: {
          property: 'token',
          global: true,
          required: true,
          type: 'Bearer',
        },
        user: {
          property: 'user',
          autoFetch: true,
        },
        endpoints: {
          login: { url: '/api/auth/login', method: 'post' },
          logout: { url: '/api/auth/logout', method: 'post' },
          user: { url: '/api/auth/user', method: 'get' },
        },
      },
    },
  },
}
