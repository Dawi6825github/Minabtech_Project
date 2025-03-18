<template>
  <div>
    <h1 class="text-3xl font-bold">Sign Up</h1>
    <form @submit.prevent="signup">
      <div>
        <label for="email">Email:</label>
        <input type="email" v-model="email" required />
      </div>
      <div>
        <label for="password">Password:</label>
        <input type="password" v-model="password" required />
      </div>
      <button type="submit">Sign Up</button>
    </form>
  </div>
</template>

<script>
export default {
  data() {
    return {
      email: '',
      password: '',
    };
  },
  methods: {
    async signup() {
      try {
        const response = await this.$axios.$post('/api/auth/signup', {
          email: this.email,
          password: this.password,
        });
        this.$auth.setToken('local', response.token); // Set the token for authentication
        this.$router.push('/'); // Redirect to homepage after signup
      } catch (error) {
        console.error('Signup failed:', error);
      }
    },
  },
};
</script>

<style scoped>
/* Add any necessary styles here */
</style>
