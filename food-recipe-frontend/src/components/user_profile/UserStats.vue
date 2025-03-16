<template>
    <div class="user-stats">
      <h2>My Stats</h2>
      <p>Recipes Created: {{ stats.recipesCreated }}</p>
      <p>Followers: {{ stats.followers }}</p>
      <p>Comments: {{ stats.comments }}</p>
    </div>
  </template>
  
  <script>
  export default {
    props: {
      stats: Object, // User statistics data passed from parent
    },
  };
  </script>
  
  <style scoped>
  .user-stats {
    margin-top: 30px;
    font-size: 1.2rem;
  }
  </style>
  