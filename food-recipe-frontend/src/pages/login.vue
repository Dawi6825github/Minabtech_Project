<template>
  <div class="min-h-screen bg-gradient-to-br from-orange-100 via-orange-200 to-orange-300 flex items-center justify-center p-4">
    
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

    
    <div class="bg-white/90 backdrop-blur-sm p-8 rounded-2xl shadow-[0_8px_30px_rgb(0,0,0,0.12)] w-96 transform hover:scale-105 transition-all duration-300 border border-orange-100">
      <div class="text-center mb-8">
        <div class="animate-bounce mb-4">
          <i class="fas fa-utensils text-4xl text-orange-600"></i>
        </div>
        <h1 class="text-3xl font-bold bg-gradient-to-r from-orange-600 to-red-600 bg-clip-text text-transparent">Welcome Back!</h1>
        <p class="text-gray-600 mt-2">Sign in to your recipe paradise</p>
      </div>

      <form @submit.prevent="handleLogin" class="space-y-6">
        <div class="space-y-2 group">
          <label class="block text-sm font-medium text-gray-700">Email</label>
          <div class="relative">
            <i class="fas fa-envelope absolute left-3 top-3 text-gray-400"></i>
            <input 
              v-model="email"
              type="email"
              class="w-full pl-10 pr-4 py-2 border border-gray-300 rounded-lg focus:ring-2 focus:ring-orange-500 focus:border-transparent transition-all duration-300 group-hover:border-orange-400"
              placeholder="your@email.com"
              required
            />
          </div>
        </div>

        <div class="space-y-2 group">
          <label class="block text-sm font-medium text-gray-700">Password</label>
          <div class="relative">
            <i class="fas fa-lock absolute left-3 top-3 text-gray-400"></i>
            <input 
              v-model="password"
              :type="showPassword ? 'text' : 'password'"
              class="w-full pl-10 pr-12 py-2 border border-gray-300 rounded-lg focus:ring-2 focus:ring-orange-500 focus:border-transparent transition-all duration-300 group-hover:border-orange-400"
              placeholder="••••••••"
              required
            />
            <button 
              type="button" 
              @click="showPassword = !showPassword"
              class="absolute right-3 top-2.5 text-gray-500 hover:text-orange-600 transition-colors"
            >
              <i :class="showPassword ? 'fas fa-eye-slash' : 'fas fa-eye'"></i>
            </button>
          </div>
          <div class="text-xs text-gray-500">
            <i class="fas fa-info-circle mr-1"></i>
            Password must be at least 8 characters
          </div>
        </div>

        <div class="flex items-center justify-between">
          <label class="flex items-center group cursor-pointer">
            <input type="checkbox" v-model="rememberMe" class="rounded text-orange-600 group-hover:ring-2 group-hover:ring-orange-300"/>
            <span class="ml-2 text-sm text-gray-600 group-hover:text-orange-600 transition-colors">Remember me</span>
          </label>
          <a href="#" class="text-sm text-orange-600 hover:text-orange-800 hover:underline transition-all">Forgot password?</a>
        </div>

        <button 
          type="submit" 
          class="w-full bg-gradient-to-r from-orange-600 to-red-600 text-white py-3 rounded-lg hover:from-orange-700 hover:to-red-700 transform hover:scale-105 transition-all duration-200 shadow-lg hover:shadow-orange-300/50"
          :class="{ 'opacity-75 cursor-not-allowed': loading }"
          :disabled="loading"
        >
          <span v-if="!loading">Sign In</span>
          <span v-else class="flex items-center justify-center">
            <svg class="animate-spin h-5 w-5 mr-2" viewBox="0 0 24 24">
              <circle class="opacity-25" cx="12" cy="12" r="10" stroke="currentColor" stroke-width="4" fill="none"/>
              <path class="opacity-75" fill="currentColor" d="M4 12a8 8 0 018-8V0C5.373 0 0 5.373 0 12h4zm2 5.291A7.962 7.962 0 014 12H0c0 3.042 1.135 5.824 3 7.938l3-2.647z"/>
            </svg>
            Signing in...
          </span>
        </button>

        <div class="relative my-6">
          <div class="absolute inset-0 flex items-center">
            <div class="w-full border-t border-gray-300"></div>
          </div>
          <div class="relative flex justify-center text-sm">
            <span class="px-2 bg-white text-gray-500">Or continue with</span>
          </div>
        </div>

        <div class="grid grid-cols-3 gap-4">
          <button class="flex items-center justify-center p-2 border border-gray-300 rounded-lg hover:bg-blue-50 hover:border-blue-400 transition-all group">
            <i class="fab fa-facebook-f text-blue-600 group-hover:scale-110 transition-transform"></i>
          </button>
          <button class="flex items-center justify-center p-2 border border-gray-300 rounded-lg hover:bg-red-50 hover:border-red-400 transition-all group">
            <i class="fab fa-google text-red-600 group-hover:scale-110 transition-transform"></i>
          </button>
          <button class="flex items-center justify-center p-2 border border-gray-300 rounded-lg hover:bg-gray-50 hover:border-gray-600 transition-all group">
            <i class="fab fa-apple text-gray-800 group-hover:scale-110 transition-transform"></i>
          </button>
        </div>
      </form>

      <p class="text-center mt-8 text-gray-600">
        Don't have an account? 
        <router-link to="/register" class="text-orange-600 hover:text-orange-800 font-medium hover:underline transition-all">
          Sign up now
        </router-link>
      </p>

      <div class="text-center mt-4 text-xs text-gray-500">
        <p>By signing in, you agree to our</p>
        <div class="space-x-2">
          <a href="#" class="hover:text-orange-600 hover:underline transition-colors">Terms of Service</a>
          <span>•</span>
          <a href="#" class="hover:text-orange-600 hover:underline transition-colors">Privacy Policy</a>
        </div>
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
const email = ref('')
const password = ref('')
const showPassword = ref(false)
const rememberMe = ref(false)
const loading = ref(false)

const handleLogin = async () => {
  if (password.value.length < 8) {
    toast.error('Password must be at least 8 characters long')
    return
  }

  try {
    loading.value = true
    // Simulate API call
    await new Promise(resolve => setTimeout(resolve, 1500))
    
    // Add your authentication logic here
    console.log('Login attempt:', {
      email: email.value,
      password: password.value,
      rememberMe: rememberMe.value
    })
    
    toast.success('Successfully logged in!')
    router.push('/dashboard')
  } catch (error) {
    console.error('Login failed:', error)
    toast.error('Login failed. Please try again.')
  } finally {
    loading.value = false
  }
}
</script>

<style scoped>
@import url('https://cdnjs.cloudflare.com/ajax/libs/font-awesome/6.0.0/css/all.min.css');

.fade-enter-active,
.fade-leave-active {
  transition: opacity 0.3s ease;
}

.fade-enter-from,
.fade-leave-to {
  opacity: 0;
}

@keyframes gradient {
  0% { background-position: 0% 50%; }
  50% { background-position: 100% 50%; }
  100% { background-position: 0% 50%; }
}

.bg-gradient-animate {
  background-size: 200% 200%;
  animation: gradient 15s ease infinite;
}
</style>
