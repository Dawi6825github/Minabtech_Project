<template>
    <div class="user-info">
      <img src="" alt="User Profile Picture" class="profile-img" />
      <h1>{{ user.name }}</h1>
      <p>Email: {{ user.email }}</p>
      <p v-if="user.bio">Bio: {{ user.bio }}</p>
    </div>
  </template>
  
  <script>
  export default {
    props: {
      user: Object, // Passed user data from parent
    },
  };
  </script>
  
  <style scoped>
  .user-info {
    text-align: center;
    margin-bottom: 20px;
  }
  
  .profile-img {
    width: 100px;
    height: 100px;
    border-radius: 50%;
  }
  </style>
  