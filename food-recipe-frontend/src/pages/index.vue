// /src/pages/index.vue
<template>
  <div class="min-h-screen bg-gray-50">
    <!-- Hero Section with Search -->
    <div class="relative bg-gradient-to-r from-orange-500 to-red-600 py-16">
      <div class="container mx-auto px-4 text-center">
        <h1 class="text-4xl md:text-6xl font-bold text-white mb-6">Discover Amazing Food Recipes</h1>
        <p class="text-xl text-white mb-8">Share your cooking journey and explore authentic recipes from food lovers worldwide</p>
        <div class="max-w-2xl mx-auto mb-8">
          <input v-model="searchQuery" type="text" placeholder="Search for recipes..." 
                 class="w-full px-6 py-3 rounded-full text-gray-700 focus:outline-none focus:ring-2 focus:ring-red-500">
        </div>
        <div class="flex justify-center gap-4">
          <NuxtLink to="/recipes" class="bg-white text-red-600 px-8 py-3 rounded-full font-semibold hover:bg-gray-100 transition">
            Explore Recipes
          </NuxtLink>
          <button @click="showTrendingRecipes" class="bg-red-700 text-white px-8 py-3 rounded-full font-semibold hover:bg-red-800 transition">
            Trending Now
          </button>
        </div>
      </div>
    </div>

    <!-- Popular Categories -->
    <div class="container mx-auto px-4 py-12">
      <h2 class="text-3xl font-bold text-center mb-8">Popular Categories</h2>
      <div class="flex overflow-x-auto gap-4 pb-4">
        <div v-for="category in categories" :key="category.id" 
             class="flex-shrink-0 cursor-pointer hover:transform hover:scale-105 transition"
             @click="filterByCategory(category.id)">
          <div class="w-40 h-40 rounded-lg overflow-hidden relative">
            <img :src="category.image" :alt="category.name" class="w-full h-full object-cover">
            <div class="absolute inset-0 bg-black bg-opacity-40 flex items-center justify-center">
              <span class="text-white font-semibold">{{ category.name }}</span>
            </div>
          </div>
        </div>
      </div>
    </div>

    <!-- Featured Testimonials with Interaction -->
    <div class="container mx-auto px-4 py-16">
      <h2 class="text-3xl font-bold text-center mb-12">What Our Food Lovers Say</h2>
      <div class="grid grid-cols-1 md:grid-cols-3 gap-8">
        <div v-for="(testimony, index) in testimonials" :key="index" 
             class="bg-white p-6 rounded-lg shadow-lg hover:shadow-xl transition">
          <div class="flex items-center mb-4">
            <img :src="testimony.avatar" :alt="testimony.name" class="w-12 h-12 rounded-full">
            <div class="ml-4">
              <h3 class="font-semibold">{{ testimony.name }}</h3>
              <p class="text-gray-600">{{ testimony.location }}</p>
            </div>
          </div>
          <p class="text-gray-700">{{ testimony.content }}</p>
          <div class="mt-4 flex items-center justify-between">
            <div class="flex items-center">
              <span class="text-yellow-400">★★★★★</span>
              <span class="ml-2 text-gray-600">{{ testimony.date }}</span>
            </div>
            <button @click="likeTestimonial(index)" 
                    class="text-red-500 hover:text-red-600 transition">
              <i class="fas fa-heart"></i> {{ testimony.likes }}
            </button>
          </div>
        </div>
      </div>
    </div>

    <!-- Interactive Features Section -->
    <div class="bg-gray-100 py-16">
      <div class="container mx-auto px-4">
        <h2 class="text-3xl font-bold text-center mb-12">Why Choose Our Platform</h2>
        <div class="grid grid-cols-1 md:grid-cols-4 gap-8">
          <div v-for="feature in features" :key="feature.id" 
               class="text-center transform hover:scale-105 transition cursor-pointer"
               @click="showFeatureDetails(feature)">
            <div class="bg-white p-6 rounded-full w-20 h-20 mx-auto mb-4 flex items-center justify-center">
              <i :class="feature.icon + ' text-3xl text-red-600'"></i>
            </div>
            <h3 class="font-semibold mb-2">{{ feature.title }}</h3>
            <p class="text-gray-600">{{ feature.description }}</p>
          </div>
        </div>
      </div>
    </div>

    <!-- Enhanced Call to Action -->
    <div class="bg-red-600 py-16">
      <div class="container mx-auto px-4 text-center">
        <h2 class="text-3xl font-bold text-white mb-6">Ready to Share Your Recipe?</h2>
        <p class="text-white mb-8">Join {{ userCount.toLocaleString() }} food lovers and share your culinary masterpieces</p>
        <div class="flex justify-center gap-4">
          <NuxtLink to="/submit-recipe" class="bg-white text-red-600 px-8 py-3 rounded-full font-semibold hover:bg-gray-100 transition">
            Share Now
          </NuxtLink>
          <button @click="showDemo" class="bg-transparent border-2 border-white text-white px-8 py-3 rounded-full font-semibold hover:bg-white hover:text-red-600 transition">
            Watch Demo
          </button>
        </div>
      </div>
    </div>
  </div>
</template>

<script setup>
import { ref } from 'vue';

const searchQuery = ref('');
const userCount = ref(15420);

const categories = [
  { id: 1, name: 'Italian', image: '/categories/italian.jpg' },
  { id: 2, name: 'Asian', image: '/categories/asian.jpg' },
  { id: 3, name: 'Mexican', image: '/categories/mexican.jpg' },
  { id: 4, name: 'Desserts', image: '/categories/desserts.jpg' },
  { id: 5, name: 'Vegetarian', image: '/categories/vegetarian.jpg' },
  { id: 6, name: 'Seafood', image: '/categories/seafood.jpg' }
];

const testimonials = [
  {
    name: 'Sarah Johnson',
    location: 'New York, USA',
    avatar: '/avatars/sarah.jpg',
    content: 'I found the most amazing pasta recipe here! The community is so supportive and the tips are invaluable.',
    date: 'March 15, 2024',
    likes: 124
  },
  {
    name: 'Marco Rivera',
    location: 'Barcelona, Spain',
    avatar: '/avatars/marco.jpg',
    content: 'As a professional chef, I love sharing my recipes here. The feedback from the community helps me improve.',
    date: 'March 12, 2024',
    likes: 98
  },
  {
    name: 'Emily Chen',
    location: 'Singapore',
    avatar: '/avatars/emily.jpg',
    content: 'This platform has transformed my cooking journey. The recipes are authentic and easy to follow.',
    date: 'March 10, 2024',
    likes: 156
  }
];

const features = [
  {
    id: 1,
    icon: 'fas fa-utensils',
    title: 'Authentic Recipes',
    description: 'Discover recipes from real home cooks'
  },
  {
    id: 2,
    icon: 'fas fa-users',
    title: 'Community Driven',
    description: 'Join a community of food enthusiasts'
  },
  {
    id: 3,
    icon: 'fas fa-star',
    title: 'Verified Reviews',
    description: 'Honest feedback from real users'
  },
  {
    id: 4,
    icon: 'fas fa-mobile-alt',
    title: 'Mobile Friendly',
    description: 'Cook from anywhere, anytime'
  }
];

const showTrendingRecipes = () => {
  // Implementation for showing trending recipes
};

const filterByCategory = (categoryId) => {
  // Implementation for filtering by category
};

const likeTestimonial = (index) => {
  testimonials[index].likes++;
};

const showFeatureDetails = (feature) => {
  // Implementation for showing feature details
};

const showDemo = () => {
  // Implementation for showing demo
};
</script>

<style scoped>
.container {
  max-width: 1200px;
}

/* Add smooth scrolling */
html {
  scroll-behavior: smooth;
}

/* Add hover animations */
.transform {
  transition: transform 0.3s ease;
}

/* Custom scrollbar for categories */
.overflow-x-auto {
  scrollbar-width: thin;
  scrollbar-color: #ef4444 #f3f4f6;
}

.overflow-x-auto::-webkit-scrollbar {
  height: 6px;
}

.overflow-x-auto::-webkit-scrollbar-track {
  background: #f3f4f6;
}

.overflow-x-auto::-webkit-scrollbar-thumb {
  background-color: #ef4444;
  border-radius: 6px;
}
</style>