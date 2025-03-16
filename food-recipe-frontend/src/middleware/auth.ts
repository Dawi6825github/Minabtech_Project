// middleware/auth.js

export default function ({ store, redirect }: { store: { state: { auth: { isAuthenticated: boolean } } }, redirect: (path: string) => void }) {
  // Check if the user is logged in by checking the authentication state
  if (!store.state.auth.isAuthenticated) {
    // If not authenticated, redirect to login page
    return redirect('/login');
  }
}