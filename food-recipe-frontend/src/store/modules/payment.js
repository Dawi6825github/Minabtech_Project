// store/modules/payment.js
export const state = () => ({
    paymentStatus: null, // Status of payment (pending, success, failed)
  });
  
  export const mutations = {
    setPaymentStatus(state, status) {
      state.paymentStatus = status;
    },
  };
  
  export const actions = {
    async makePayment({ commit }, paymentData) {
      const response = await this.$axios.post('/api/payment', paymentData);
      if (response.data.success) {
        commit('setPaymentStatus', 'success');
      } else {
        commit('setPaymentStatus', 'failed');
      }
    },
  
    async checkPaymentStatus({ commit }) {
      const response = await this.$axios.get('/api/payment/status');
      commit('setPaymentStatus', response.data.status);
    },
  };
  
  export const getters = {
    getPaymentStatus(state) {
      return state.paymentStatus;
    },
  };
  