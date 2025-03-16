<!-- src/pages/admin/users/index.vue -->
<script setup>
import { ref, onMounted } from 'vue';

const users = ref([]);

// Fetch users from API or mock data
const fetchUsers = async () => {
  try {
    // Simulate an API call or use an actual API to fetch users
    users.value = [
      { id: 1, name: 'John Doe', email: 'john.doe@example.com', role: 'Admin', status: 'Active' },
      { id: 2, name: 'Jane Smith', email: 'jane.smith@example.com', role: 'User', status: 'Inactive' },
    ];
  } catch (error) {
    console.error('Failed to fetch users:', error);
  }
};

onMounted(() => {
  fetchUsers();
});

const editUser = (userId) => {
  // Navigate to the edit user page
  navigateTo(`/admin/users/${userId}`);
};

const resetPassword = (userId) => {
  // Reset the password for the user
  alert(`Password reset for user ID: ${userId}`);
};

const toggleUserStatus = (userId) => {
  // Toggle the user account status (active/inactive)
  alert(`Toggling status for user ID: ${userId}`);
};
</script>

<template>
  <div>
    <h2>User Management</h2>
    <table>
      <thead>
        <tr>
          <th>Name</th>
          <th>Email</th>
          <th>Role</th>
          <th>Status</th>
          <th>Actions</th>
        </tr>
      </thead>
      <tbody>
        <tr v-for="user in users" :key="user.id">
          <td>{{ user.name }}</td>
          <td>{{ user.email }}</td>
          <td>{{ user.role }}</td>
          <td>{{ user.status }}</td>
          <td>
            <button @click="editUser(user.id)">Edit</button>
            <button @click="resetPassword(user.id)">Reset Password</button>
            <button @click="toggleUserStatus(user.id)">Toggle Status</button>
          </td>
        </tr>
      </tbody>
    </table>
  </div>
</template>

<style scoped>
table {
  width: 100%;
  border-collapse: collapse;
}

table th, table td {
  padding: 10px;
  border: 1px solid #ddd;
}

button {
  margin-right: 5px;
  padding: 5px 10px;
  background-color: #007bff;
  color: white;
  border: none;
  cursor: pointer;
}

button:hover {
  background-color: #0056b3;
}
</style>
