function auth({ store, redirect }) {
  if (!store.state.auth.isAuthenticated) {
    return redirect("/login");
  }
}
export {
  auth as default
};
//# sourceMappingURL=auth.mjs.map
