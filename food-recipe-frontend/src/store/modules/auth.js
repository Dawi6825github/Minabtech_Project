// store/modules/auth.js
export const state = () => ({
  user: null,  // Current authenticated user
  token: null, // Authentication token (e.g., JWT)
});

export const mutations = {
  setUser(state, user) {
    state.user = user;
  },
  setToken(state, token) {
    state.token = token;
  },
  logout(state) {
    state.user = null;
    state.token = null;
  },
};

export const actions = {
  async login({ commit }, credentials) {
    // Call API to login and get token
    const response = await this.$axios.post('/api/auth/login', credentials);
    commit('setUser', response.data.user);
    commit('setToken', response.data.token);
    // Save token in localStorage
    localStorage.setItem('auth_token', response.data.token);
  },

  logout({ commit }) {
    commit('logout');
    localStorage.removeItem('auth_token');
  },

  // Check if the user is already authenticated on app load
  initializeAuth({ commit }) {
    const token = localStorage.getItem('auth_token');
    if (token) {
      commit('setToken', token);
      // Optionally, fetch user info from API using the token
    }
  },
};

export const getters = {
  isAuthenticated(state) {
    return !!state.token; // If token exists, user is authenticated
  },
  getUser(state) {
    return state.user;
  },
};
