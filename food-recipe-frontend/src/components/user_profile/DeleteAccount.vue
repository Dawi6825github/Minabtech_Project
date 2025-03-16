<template>
    <div class="delete-account">
      <button @click="confirmDelete">Delete Account</button>
    </div>
  </template>
  
  <script>
  export default {
    methods: {
      confirmDelete() {
        if (confirm('Are you sure you want to delete your account? This action cannot be undone.')) {
          this.$emit('confirmDelete'); // Emit event to delete account
        }
      },
    },
  };
  </script>
  
  <style scoped>
  .delete-account button {
    color: red;
    border: 1px solid red;
    background-color: white;
    padding: 10px;
    cursor: pointer;
  }
  </style>
  