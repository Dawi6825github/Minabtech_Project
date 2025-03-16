function handleHotUpdate(_router, _generateRoutes) {
}
const _routes = [
  {
    name: "admin",
    path: "/admin",
    component: () => import("../src/pages/admin/index.vue.mjs")
  },
  {
    name: "admin-recipes-create",
    path: "/admin/recipes/create",
    component: () => import("../src/pages/admin/recipes/create.vue.mjs")
  },
  {
    name: "admin-recipes-edit",
    path: "/admin/recipes/edit",
    component: () => import("../src/pages/admin/recipes/edit.vue.mjs")
  },
  {
    name: "admin-recipes",
    path: "/admin/recipes",
    component: () => import("../src/pages/admin/recipes/index.vue.mjs")
  },
  {
    name: "admin-users-id",
    path: "/admin/users/:id()",
    component: () => import("../src/pages/admin/users/_id_.vue.mjs")
  },
  {
    name: "admin-users",
    path: "/admin/users",
    component: () => import("../src/pages/admin/users/index.vue.mjs")
  },
  {
    name: "index",
    path: "/",
    component: () => import("../src/pages/index.vue.mjs")
  },
  {
    name: "login",
    path: "/login",
    component: () => import("../src/pages/login.vue.mjs")
  },
  {
    name: "Signup",
    path: "/Signup",
    component: () => import("../src/pages/Signup.vue.mjs")
  },
  {
    name: "user-profile",
    path: "/user-profile",
    component: () => import("../src/pages/user-profile.vue.mjs")
  }
];
export {
  _routes as default,
  handleHotUpdate
};
//# sourceMappingURL=virtual_nuxt__workspace_Minabtech_Project_food-recipe-frontend_.nuxt_routes.mjs.map
