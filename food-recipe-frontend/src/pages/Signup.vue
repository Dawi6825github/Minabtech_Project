<template>
  <div class="h-screen overflow-hidden bg-gradient-to-br from-orange-100 via-orange-200 to-orange-300 flex items-center justify-center p-2 bg-pattern-food relative">
    <!-- Animated Background Elements -->
    <div class="absolute inset-0 overflow-hidden pointer-events-none">
      <div class="animate-pulse-slow absolute top-0 left-0 w-full h-full bg-gradient-to-r from-orange-200/30 via-transparent to-orange-200/30"></div>
      <div class="grid grid-cols-4 gap-8 absolute inset-0">
        <i v-for="(icon, index) in ['pizza-slice', 'hamburger', 'cookie', 'ice-cream', 'carrot', 'apple-alt', 'cheese', 'drumstick-bite']" 
           :key="index"
           :class="`fas fa-${icon} text-orange-300/20 text-5xl absolute animate-float-${index % 3 === 0 ? 'slow' : index % 2 === 0 ? 'medium' : 'fast'}`"
           :style="`top: ${Math.random() * 100}%; left: ${Math.random() * 100}%`">
        </i>
      </div>
    </div>

    <div class="relative bg-white/95 backdrop-blur-lg p-8 rounded-2xl shadow-2xl w-96 transform hover:scale-105 transition-all duration-300 border-2 border-orange-200">
      <!-- Decorative Header -->
      <div class="absolute -top-6 left-1/2 transform -translate-x-1/2 bg-orange-500 text-white px-8 py-2 rounded-full shadow-lg">
        <h2 class="text-2xl font-bold">Join Us</h2>
      </div>

      <form @submit.prevent="handleSignup" class="space-y-6 mt-8">
        <!-- Social Sign Up Options -->
        <div class="flex gap-4 mb-6">
          <button type="button" class="flex-1 py-2 px-4 bg-blue-600 text-white rounded-lg hover:bg-blue-700 transition-colors">
            <i class="fab fa-facebook-f mr-2"></i>Facebook
          </button>
          <button type="button" class="flex-1 py-2 px-4 bg-red-600 text-white rounded-lg hover:bg-red-700 transition-colors">
            <i class="fab fa-google mr-2"></i>Google
          </button>
        </div>

        <div class="relative">
          <span class="absolute inset-y-0 left-0 pl-3 flex items-center">
            <i class="fas fa-user text-orange-500"></i>
          </span>
          <input 
            v-model="username"
            type="text"
            required
            placeholder="Username"
            class="pl-10 w-full px-4 py-3 bg-orange-50 border-2 border-orange-200 rounded-lg focus:ring-2 focus:ring-orange-500 focus:border-transparent transition-all duration-200"
          >
        </div>

        <div class="relative">
          <span class="absolute inset-y-0 left-0 pl-3 flex items-center">
            <i class="fas fa-envelope text-orange-500"></i>
          </span>
          <input 
            v-model="email"
            type="email"
            required
            placeholder="Email"
            class="pl-10 w-full px-4 py-3 bg-orange-50 border-2 border-orange-200 rounded-lg focus:ring-2 focus:ring-orange-500 focus:border-transparent transition-all duration-200"
          >
        </div>

        <div class="relative">
          <span class="absolute inset-y-0 left-0 pl-3 flex items-center">
            <i class="fas fa-lock text-orange-500"></i>
          </span>
          <input 
            v-model="password"
            :type="showPassword ? 'text' : 'password'"
            required
            placeholder="Password"
            class="pl-10 w-full px-4 py-3 bg-orange-50 border-2 border-orange-200 rounded-lg focus:ring-2 focus:ring-orange-500 focus:border-transparent transition-all duration-200"
          >
          <button 
            type="button"
            @click="showPassword = !showPassword"
            class="absolute right-3 top-1/2 transform -translate-y-1/2 text-orange-500 hover:text-orange-600"
          >
            <i :class="showPassword ? 'fas fa-eye-slash' : 'fas fa-eye'"></i>
          </button>
        </div>

        <div class="relative">
          <span class="absolute inset-y-0 left-0 pl-3 flex items-center">
            <i class="fas fa-utensils text-orange-500"></i>
          </span>
          <select 
            v-model="experience"
            class="pl-10 w-full px-4 py-3 bg-orange-50 border-2 border-orange-200 rounded-lg focus:ring-2 focus:ring-orange-500 focus:border-transparent transition-all duration-200"
          >
            <option value="">Select your cooking level</option>
            <option value="beginner">🥄 Beginner Chef</option>
            <option value="intermediate">🍳 Home Cook</option>
            <option value="advanced">👨‍🍳 Advanced Cook</option>
            <option value="professional">🎖️ Professional Chef</option>
          </select>
        </div>

        <div class="flex items-center gap-2 bg-orange-50 p-3 rounded-lg">
          <input 
            type="checkbox"
            v-model="newsletter"
            class="h-5 w-5 text-orange-600 focus:ring-orange-500 border-orange-300 rounded"
          >
          <label class="text-sm text-gray-700">
            Send me delicious recipes and cooking tips! 🍽️
          </label>
        </div>

        <button
          type="submit"
          :disabled="loading"
          class="w-full py-3 px-4 text-white bg-gradient-to-r from-orange-500 to-orange-600 rounded-lg font-medium shadow-lg hover:from-orange-600 hover:to-orange-700 focus:outline-none focus:ring-2 focus:ring-offset-2 focus:ring-orange-500 disabled:opacity-50 disabled:cursor-not-allowed transform hover:scale-102 transition-all duration-200"
        >
          <div class="flex items-center justify-center">
            <i v-if="loading" class="fas fa-spinner fa-spin mr-2"></i>
            <span>{{ loading ? 'Creating your account...' : 'Start Cooking! 🚀' }}</span>
          </div>
        </button>

        <!-- Progress Indicator -->
        <div v-if="loading" class="w-full bg-orange-100 rounded-full h-1.5 mt-4">
          <div class="bg-orange-500 h-1.5 rounded-full animate-progress"></div>
        </div>
      </form>

      <!-- Decorative Elements -->
      <div class="absolute -left-16 top-1/2 transform -translate-y-1/2 space-y-4">
        <i class="fas fa-star text-4xl text-yellow-400 animate-pulse"></i>
        <i class="fas fa-award text-5xl text-orange-400 animate-bounce"></i>
        <i class="fas fa-medal text-4xl text-yellow-500 animate-pulse"></i>
      </div>

      <div class="absolute -right-16 top-1/2 transform -translate-y-1/2 space-y-4">
        <i class="fas fa-crown text-4xl text-yellow-400 animate-pulse"></i>
        <i class="fas fa-certificate text-5xl text-orange-400 animate-bounce"></i>
        <i class="fas fa-trophy text-4xl text-yellow-500 animate-pulse"></i>
      </div>
    </div>
  </div>
</template>

<script setup>
import { ref } from 'vue'
import { useRouter } from 'vue-router'
import { useToast } from 'vue-toastification'

const router = useRouter()
const toast = useToast()
const username = ref('')
const email = ref('')
const password = ref('')
const showPassword = ref(false)
const loading = ref(false)
const experience = ref('')
const newsletter = ref(false)

const handleSignup = async () => {
  if (password.value.length < 8) {
    toast.error('Password must be at least 8 characters long')
    return
  }

  try {
    loading.value = true
    await new Promise(resolve => setTimeout(resolve, 1500))
    console.log('Signup:', {
      username: username.value,
      email: email.value,
      password: password.value,
      experience: experience.value,
      newsletter: newsletter.value
    })
    toast.success('Welcome to our cooking community! 🎉')
    router.push('/dashboard')
  } catch (error) {
    console.error('Signup failed:', error)
    toast.error('Oops! Something went wrong')
  } finally {
    loading.value = false
  }
}
</script>

<style scoped>
@import url('https://cdnjs.cloudflare.com/ajax/libs/font-awesome/6.0.0/css/all.min.css');

.fade-enter-active, .fade-leave-active {
  transition: opacity 0.2s ease;
}

.fade-enter-from, .fade-leave-to {
  opacity: 0;
}

@keyframes gradient {
  0% { background-position: 0% 50%; }
  50% { background-position: 100% 50%; }
  100% { background-position: 0% 50%; }
}

.bg-gradient-animate {
  background-size: 150% 150%;
  animation: gradient 10s ease infinite;
}

.bg-pattern-food {
  background-image: url("data:image/svg+xml,%3Csvg width='40' height='40' viewBox='0 0 60 60' xmlns='http://www.w3.org/2000/svg'%3E%3Cg fill='none' fill-rule='evenodd'%3E%3Cg fill='%23FFA500' fill-opacity='0.1'%3E%3Cpath d='M36 34v-4h-2v4h-4v2h4v4h2v-4h4v-2h-4zm0-30V0h-2v4h-4v2h4v4h2V6h4V4h-4zM6 34v-4H4v4H0v2h4v4h2v-4h4v-2H6zM6 4V0H4v4H0v2h4v4h2V6h4V4H6z'/%3E%3C/g%3E%3C/g%3E%3C/svg%3E");
}

@keyframes float {
  0% { transform: translateY(0px); }
  50% { transform: translateY(-5px); }
  100% { transform: translateY(0px); }
}

.animate-float {
  animation: float 2s ease-in-out infinite;
}

@keyframes float-slow {
  0% { transform: translateY(0px) rotate(0deg); }
  50% { transform: translateY(-15px) rotate(5deg); }
  100% { transform: translateY(0px) rotate(0deg); }
}

@keyframes float-medium {
  0% { transform: translateY(0px) rotate(0deg); }
  50% { transform: translateY(-10px) rotate(-5deg); }
  100% { transform: translateY(0px) rotate(0deg); }
}

@keyframes float-fast {
  0% { transform: translateY(0px) rotate(0deg); }
  50% { transform: translateY(-8px) rotate(5deg); }
  100% { transform: translateY(0px) rotate(0deg); }
}

.animate-float-slow {
  animation: float-slow 6s ease-in-out infinite;
}

.animate-float-medium {
  animation: float-medium 4s ease-in-out infinite;
}

.animate-float-fast {
  animation: float-fast 3s ease-in-out infinite;
}

@keyframes progress {
  0% { width: 0% }
  100% { width: 100% }
}

.animate-progress {
  animation: progress 1.5s linear;
}

.animate-pulse-slow {
  animation: pulse 4s cubic-bezier(0.4, 0, 0.6, 1) infinite;
}

@keyframes pulse {
  0%, 100% { opacity: 1; }
  50% { opacity: .5; }
}
</style>
