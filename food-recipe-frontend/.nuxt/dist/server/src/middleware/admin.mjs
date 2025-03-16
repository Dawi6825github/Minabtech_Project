function admin({ store, redirect }) {
  if (!store.state.auth.isAuthenticated || store.state.auth.user.role !== "admin") {
    return redirect("/");
  }
}
export {
  admin as default
};
//# sourceMappingURL=admin.mjs.map
