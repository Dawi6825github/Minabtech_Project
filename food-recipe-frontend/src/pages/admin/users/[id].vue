<template>
    <div class="user-details">
      <h1>User Profile</h1>
  
      <div v-if="user">
        <h2>{{ user.name }}</h2>
        <p>Email: {{ user.email }}</p>
        <p>Status: {{ user.active ? 'Active' : 'Inactive' }}</p>
  
        <button @click="toggleUserStatus">
          {{ user.active ? 'Deactivate' : 'Activate' }} User
        </button>
        <button @click="deleteUser">Delete User</button>
      </div>
      <p v-else>Loading user data...</p>
    </div>
  </template>
  
  <script>
  export default {
    data() {
      return {
        user: null
      };
    },
    async mounted() {
      const userId = this.$route.params.id;
      this.user = await this.$store.dispatch('admin/fetchUserById', userId);
    },
    methods: {
      async toggleUserStatus() {
        this.user.active = !this.user.active;
        await this.$store.dispatch('admin/updateUserStatus', { id: this.user.id, active: this.user.active });
      },
      async deleteUser() {
        const confirmDelete = confirm('Are you sure you want to delete this user?');
        if (confirmDelete) {
          await this.$store.dispatch('admin/deleteUser', this.user.id);
          this.$router.push('/admin'); // Redirect to the admin dashboard after deletion
        }
      }
    }
  };
  </script>
  
  <style scoped>
  .user-details {
    padding: 20px;
  }
  
  button {
    margin-right: 10px;
    padding: 5px 10px;
    cursor: pointer;
  }
  
  button:hover {
    background-color: darkred;
  }
  </style>
  