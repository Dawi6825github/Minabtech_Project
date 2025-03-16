<template>
    <div class="admin-dashboard">
      <h1>Admin Dashboard</h1>
      <h2>Users</h2>
  
      <!-- User Table -->
      <table v-if="users.length">
        <thead>
          <tr>
            <th>ID</th>
            <th>Name</th>
            <th>Email</th>
            <th>Actions</th>
          </tr>
        </thead>
        <tbody>
          <tr v-for="user in users" :key="user.id">
            <td>{{ user.id }}</td>
            <td>{{ user.name }}</td>
            <td>{{ user.email }}</td>
            <td>
              <router-link :to="`/admin/users/${user.id}`">View</router-link>
              <button @click="deleteUser(user.id)">Delete</button>
            </td>
          </tr>
        </tbody>
      </table>
  
      <p v-else>No users found.</p>
    </div>
  </template>
  
  <script>
  export default {
    data() {
      return {
        users: []
      };
    },
    async mounted() {
      // Fetch users from the store or API
      this.users = await this.$store.dispatch('admin/fetchUsers');
    },
    methods: {
      async deleteUser(userId) {
        const confirmDelete = confirm('Are you sure you want to delete this user?');
        if (confirmDelete) {
          await this.$store.dispatch('admin/deleteUser', userId);
          this.users = this.users.filter(user => user.id !== userId); // Remove the user from the list after deletion
        }
      }
    }
  };
  </script>
  
  <style scoped>
  .admin-dashboard {
    padding: 20px;
  }
  
  table {
    width: 100%;
    border-collapse: collapse;
  }
  
  table th, table td {
    border: 1px solid #ddd;
    padding: 8px;
    text-align: left;
  }
  
  button {
    background-color: red;
    color: white;
    border: none;
    padding: 5px 10px;
    cursor: pointer;
  }
  
  button:hover {
    background-color: darkred;
  }
  </style>
  