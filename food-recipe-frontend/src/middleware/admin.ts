// middleware/admin.js

export default function ({ store, redirect }: { store: any; redirect: any }) {
  // Check if the user is authenticated and has an admin role
  if (!store.state.auth.isAuthenticated || store.state.auth.user.role !== 'admin') {
    // If not an admin, redirect to the homepage or login page
    return redirect('/');
  }
}