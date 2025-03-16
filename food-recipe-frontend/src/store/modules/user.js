// store/modules/user.js
export const state = () => ({
    profile: null, // User profile data
  });
  
  export const mutations = {
    setProfile(state, profile) {
      state.profile = profile;
    },
    updateProfile(state, updatedData) {
      state.profile = { ...state.profile, ...updatedData };
    },
  };
  
  export const actions = {
    async fetchUserProfile({ commit }) {
      const response = await this.$axios.get('/api/user/profile');
      commit('setProfile', response.data);
    },
  
    async updateUserProfile({ commit }, updatedData) {
      const response = await this.$axios.put('/api/user/profile', updatedData);
      commit('updateProfile', response.data);
    },
  };
  
  export const getters = {
    getUserProfile(state) {
      return state.profile;
    },
  };
  