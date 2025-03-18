<template>
  <div>
    <h1 class="text-3xl font-bold">Login</h1>
    <form @submit.prevent="login">
      <div>
        <label for="email">Email:</label>
        <input type="email" v-model="email" required />
      </div>
      <div>
        <label for="password">Password:</label>
        <input type="password" v-model="password" required />
      </div>
      <button type="submit">Login</button>
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
    async login() {
      try {
        const response = await this.$axios.$post('/api/auth/login', {
          email: this.email,
          password: this.password,
        });
        this.$auth.setToken('local', response.token); // Set the token for authentication
        this.$router.push('/'); // Redirect to homepage after login
      } catch (error) {
        console.error('Login failed:', error);
      }
    },
  },
};
</script>

<style scoped>
/* Add any necessary styles here */
</style>
