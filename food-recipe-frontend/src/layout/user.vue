layouts/user.vue
<template>
    <div>
      <header class="bg-green-600 text-white p-4">User Dashboard</header>
      <main class="p-4">
        <slot /> <!-- Inject user-specific content here -->
      </main>
    </div>
  </template>
  