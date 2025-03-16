<template>
  <header class="bg-white shadow-sm">
    <nav class="container mx-auto px-4 py-3">
      <div class="flex items-center justify-between">
        <!-- Logo and Brand -->
        <div class="flex items-center">
          <router-link to="/" class="flex items-center">
            <img src="../assets/logo.png" alt="Food Recipe Logo" class="h-10 w-10 mr-2">
            <span class="text-2xl font-bold text-emerald-600">FoodieHub</span>
          </router-link>
        </div>

        <!-- Navigation Links -->
        <div class="hidden md:flex space-x-6">
          <router-link to="/" class="nav-link">Home</router-link>
          <router-link to="/recipes" class="nav-link">Recipes</router-link>
          <router-link to="/categories" class="nav-link">Categories</router-link>
          <router-link to="/meal-planner" class="nav-link">Meal Planner</router-link>
          <router-link to="/favorites" class="nav-link">Favorites</router-link>
        </div>
      </div>
    </nav>
  </header>
</template>
