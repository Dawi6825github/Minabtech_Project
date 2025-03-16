// https://nuxt.com/docs/api/configuration/nuxt-config
export default defineNuxtConfig({
  devtools: { enabled: true },

  modules: [
    '@nuxtjs/tailwindcss',
    '@pinia/nuxt',
    '@vee-validate/nuxt',
  ],

  srcDir: 'src/',

  // Routes should be handled by the pages directory structure in Nuxt 3
  // Remove the manual route configuration
  router: {
    options: {
      linkActiveClass: 'active',
      linkExactActiveClass: 'exact-active'
    }
  },

  dir: {
    public: 'public'  // The public directory to store static files (like images, fonts)
  },

  app: {
    head: {
      title: 'Food Recipe App',
      meta: [
        { charset: 'utf-8' },
        { name: 'viewport', content: 'width=device-width, initial-scale=1' },
        { name: 'description', content: 'Find and share delicious recipes' }
      ],
      link: [
        { rel: 'icon', type: 'image/x-icon', href: '/favicon.ico' }
      ]
    }
  },

  runtimeConfig: {
    public: {
      apiBase: process.env.API_BASE_URL || 'http://localhost:8080',
      graphqlEndpoint: process.env.GRAPHQL_ENDPOINT || 'http://localhost:8080/graphql',
    }
  },

  css: [
    '@/assets/css/tailwind.css'
  ],

  build: {
    transpile: ['@heroicons/vue', '@apollo/client', 'ts-invariant/process']
  },

  postcss: {
    plugins: {
      tailwindcss: {},
      autoprefixer: {},
    }
  },

  vite: {
    server: {
      hmr: {
        host: 'localhost',
        protocol: 'ws',
        port: 0 // This will use a random available port
      },
      allowedHosts: ['3002-dawi6825git-minabtechpr-rta84mbqvq1.ws-eu118.gitpod.io']
    },
    css: {
      preprocessorOptions: {
        scss: {
          additionalData: '@import "@/assets/css/tailwind.css";'
        }
      }
    },
    optimizeDeps: {
      exclude: ['fsevents']
    }
  },

  compatibilityDate: '2025-03-13'
})