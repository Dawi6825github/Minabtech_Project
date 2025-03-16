<template>
    <div class="edit-profile">
      <h2>Edit Profile</h2>
      <form @submit.prevent="submitForm">
        <label for="name">Name</label>
        <input v-model="user.name" type="text" id="name" required />
  
        <label for="email">Email</label>
        <input v-model="user.email" type="email" id="email" required />
  
        <label for="bio">Bio</label>
        <textarea v-model="user.bio" id="bio"></textarea>
  
        <button type="submit">Update Profile</button>
      </form>
    </div>
  </template>
  
  <script>
  export default {
    props: {
      user: Object, // User data passed from parent
    },
    methods: {
      submitForm() {
        this.$emit('update', this.user); // Emit event to parent to update profile
      },
    },
  };
  </script>
  
  <style scoped>
  .edit-profile form {
    display: flex;
    flex-direction: column;
    gap: 10px;
  }
  
  button {
    margin-top: 20px;
  }
  </style>
  