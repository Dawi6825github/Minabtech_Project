import validate from "../node_modules/nuxt/dist/pages/runtime/validate.mjs";
import manifest_45route_45rule from "../node_modules/nuxt/dist/app/middleware/manifest-route-rule.mjs";
const globalMiddleware = [
  validate,
  manifest_45route_45rule
];
const namedMiddleware = {
  admin: () => import("../src/middleware/admin.mjs"),
  auth: () => import("../src/middleware/auth.mjs")
};
export {
  globalMiddleware,
  namedMiddleware
};
//# sourceMappingURL=virtual_nuxt__workspace_Minabtech_Project_food-recipe-frontend_.nuxt_middleware.mjs.map
