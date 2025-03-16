<template>
    <nav class="navbar">
      <router-link to="/">Home</router-link>
      <router-link to="/recipes">Recipes</router-link>
      <router-link to="/login">Login</router-link>
      <router-link to="/signup">Sign Up</router-link>
    </nav>
  </template>
  
  <style scoped>
  .navbar {
    display: flex;
    justify-content: space-between;
    padding: 10px 20px;
    background-color: #333;
    color: white;
  }
  .navbar a {
    color: white;
    text-decoration: none;
    margin-right: 20px;
  }
  </style>
  