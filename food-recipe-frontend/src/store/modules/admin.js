export const state = () => ({
    users: []
  });
  
  export const mutations = {
    SET_USERS(state, users) {
      state.users = users;
    },
    SET_USER(state, user) {
      const index = state.users.findIndex(u => u.id === user.id);
      if (index !== -1) {
        state.users.splice(index, 1, user);
      }
    }
  };
  
  export const actions = {
    async fetchUsers({ commit }) {
      // Simulating an API call to fetch users
      const users = await this.$axios.get('/api/users'); // Adjust this based on your API
      commit('SET_USERS', users.data);
    },
    async fetchUserById({ commit }, userId) {
      const user = await this.$axios.get(`/api/users/${userId}`);
      return user.data;
    },
    async deleteUser({ commit }, userId) {
      // Simulate deleting a user via API
      await this.$axios.delete(`/api/users/${userId}`);
      commit('SET_USERS', state.users.filter(user => user.id !== userId));
    },
    async updateUserStatus({ commit }, { id, active }) {
      // Simulate updating user status via API
      const updatedUser = await this.$axios.put(`/api/users/${id}`, { active });
      commit('SET_USER', updatedUser.data);
    }
  };
  