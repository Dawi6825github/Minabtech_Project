<template>
  <div class="user-profile">
    <!-- User Information Section -->
    <UserInfo :user="user" />
    
    <!-- User Recipe Management Section -->
    <UserRecipes :recipes="userRecipes" />

    <!-- User Stats Section -->
    <UserStats :stats="userStats" />

    <!-- Edit Profile Section (Optional, maybe conditionally shown) -->
    <EditUserProfile v-if="isEditing" :user="user" @update="updateProfile" />

    <!-- Delete Account Option (Optional, confirmation needed) -->
    <DeleteAccount @confirmDelete="deleteAccount" />
  </div>
</template>

<script>
import UserInfo from '~/components/user_profile/UserInfo.vue';
import UserRecipes from '~/components/user_profile/UserRecipes.vue';
import EditUserProfile from '~/components/user_profile/EditUserProfile.vue';
import UserStats from '~/components/user_profile/UserStats.vue';
import DeleteAccount from '~/components/user_profile/DeleteAccount.vue';

export default {
  data() {
    return {
      user: null,            // User details (fetched from Vuex or API)
      userRecipes: [],       // List of recipes created by the user
      userStats: {},         // User-specific stats (e.g., number of recipes, followers)
      isEditing: false,      // Flag to control whether the edit profile form is visible
    };
  },
  methods: {
    fetchUserData() {
      // Fetch user data (profile info, recipes, stats) from API or Vuex store
    },
    updateProfile(updatedUser) {
      // Update user profile logic
    },
    deleteAccount() {
      // Delete user account logic (show confirmation first)
    },
  },
  created() {
    this.fetchUserData(); // Fetch user data when the profile page is created
  },
};
</script>

<style scoped>
.user-profile {
  display: flex;
  flex-direction: column;
  align-items: center;
  padding: 20px;
}
</style>
