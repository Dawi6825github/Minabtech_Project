import { ref, resolveComponent, mergeProps, withCtx, createTextVNode, useSSRContext } from "vue";
import { ssrRenderAttrs, ssrRenderAttr, ssrRenderDynamicModel, ssrRenderClass, ssrIncludeBooleanAttr, ssrLooseContain, ssrRenderComponent } from "vue/server-renderer";
import { useRouter } from "vue-router";
import { useToast } from "vue-toastification";
/* empty css            */
import _export_sfc from "../../_virtual/_plugin-vue_export-helper.mjs";
const _sfc_main = {
  __name: "login",
  __ssrInlineRender: true,
  setup(__props) {
    useRouter();
    useToast();
    const email = ref("");
    const password = ref("");
    const showPassword = ref(false);
    const rememberMe = ref(false);
    const loading = ref(false);
    return (_ctx, _push, _parent, _attrs) => {
      const _component_router_link = resolveComponent("router-link");
      _push(`<div${ssrRenderAttrs(mergeProps({ class: "min-h-screen bg-gradient-to-br from-orange-100 via-orange-200 to-orange-300 flex items-center justify-center p-4" }, _attrs))} data-v-5dd820e1><div class="bg-white/90 backdrop-blur-sm p-8 rounded-2xl shadow-[0_8px_30px_rgb(0,0,0,0.12)] w-96 transform hover:scale-105 transition-all duration-300 border border-orange-100" data-v-5dd820e1><div class="text-center mb-8" data-v-5dd820e1><div class="animate-bounce mb-4" data-v-5dd820e1><i class="fas fa-utensils text-4xl text-orange-600" data-v-5dd820e1></i></div><h1 class="text-3xl font-bold bg-gradient-to-r from-orange-600 to-red-600 bg-clip-text text-transparent" data-v-5dd820e1>Welcome Back!</h1><p class="text-gray-600 mt-2" data-v-5dd820e1>Sign in to your recipe paradise</p></div><form class="space-y-6" data-v-5dd820e1><div class="space-y-2 group" data-v-5dd820e1><label class="block text-sm font-medium text-gray-700" data-v-5dd820e1>Email</label><div class="relative" data-v-5dd820e1><i class="fas fa-envelope absolute left-3 top-3 text-gray-400" data-v-5dd820e1></i><input${ssrRenderAttr("value", email.value)} type="email" class="w-full pl-10 pr-4 py-2 border border-gray-300 rounded-lg focus:ring-2 focus:ring-orange-500 focus:border-transparent transition-all duration-300 group-hover:border-orange-400" placeholder="your@email.com" required data-v-5dd820e1></div></div><div class="space-y-2 group" data-v-5dd820e1><label class="block text-sm font-medium text-gray-700" data-v-5dd820e1>Password</label><div class="relative" data-v-5dd820e1><i class="fas fa-lock absolute left-3 top-3 text-gray-400" data-v-5dd820e1></i><input${ssrRenderDynamicModel(showPassword.value ? "text" : "password", password.value, null)}${ssrRenderAttr("type", showPassword.value ? "text" : "password")} class="w-full pl-10 pr-12 py-2 border border-gray-300 rounded-lg focus:ring-2 focus:ring-orange-500 focus:border-transparent transition-all duration-300 group-hover:border-orange-400" placeholder="••••••••" required data-v-5dd820e1><button type="button" class="absolute right-3 top-2.5 text-gray-500 hover:text-orange-600 transition-colors" data-v-5dd820e1><i class="${ssrRenderClass(showPassword.value ? "fas fa-eye-slash" : "fas fa-eye")}" data-v-5dd820e1></i></button></div><div class="text-xs text-gray-500" data-v-5dd820e1><i class="fas fa-info-circle mr-1" data-v-5dd820e1></i> Password must be at least 8 characters </div></div><div class="flex items-center justify-between" data-v-5dd820e1><label class="flex items-center group cursor-pointer" data-v-5dd820e1><input type="checkbox"${ssrIncludeBooleanAttr(Array.isArray(rememberMe.value) ? ssrLooseContain(rememberMe.value, null) : rememberMe.value) ? " checked" : ""} class="rounded text-orange-600 group-hover:ring-2 group-hover:ring-orange-300" data-v-5dd820e1><span class="ml-2 text-sm text-gray-600 group-hover:text-orange-600 transition-colors" data-v-5dd820e1>Remember me</span></label><a href="#" class="text-sm text-orange-600 hover:text-orange-800 hover:underline transition-all" data-v-5dd820e1>Forgot password?</a></div><button type="submit" class="${ssrRenderClass([{ "opacity-75 cursor-not-allowed": loading.value }, "w-full bg-gradient-to-r from-orange-600 to-red-600 text-white py-3 rounded-lg hover:from-orange-700 hover:to-red-700 transform hover:scale-105 transition-all duration-200 shadow-lg hover:shadow-orange-300/50"])}"${ssrIncludeBooleanAttr(loading.value) ? " disabled" : ""} data-v-5dd820e1>`);
      if (!loading.value) {
        _push(`<span data-v-5dd820e1>Sign In</span>`);
      } else {
        _push(`<span class="flex items-center justify-center" data-v-5dd820e1><svg class="animate-spin h-5 w-5 mr-2" viewBox="0 0 24 24" data-v-5dd820e1><circle class="opacity-25" cx="12" cy="12" r="10" stroke="currentColor" stroke-width="4" fill="none" data-v-5dd820e1></circle><path class="opacity-75" fill="currentColor" d="M4 12a8 8 0 018-8V0C5.373 0 0 5.373 0 12h4zm2 5.291A7.962 7.962 0 014 12H0c0 3.042 1.135 5.824 3 7.938l3-2.647z" data-v-5dd820e1></path></svg> Signing in... </span>`);
      }
      _push(`</button><div class="relative my-6" data-v-5dd820e1><div class="absolute inset-0 flex items-center" data-v-5dd820e1><div class="w-full border-t border-gray-300" data-v-5dd820e1></div></div><div class="relative flex justify-center text-sm" data-v-5dd820e1><span class="px-2 bg-white text-gray-500" data-v-5dd820e1>Or continue with</span></div></div><div class="grid grid-cols-3 gap-4" data-v-5dd820e1><button class="flex items-center justify-center p-2 border border-gray-300 rounded-lg hover:bg-blue-50 hover:border-blue-400 transition-all group" data-v-5dd820e1><i class="fab fa-facebook-f text-blue-600 group-hover:scale-110 transition-transform" data-v-5dd820e1></i></button><button class="flex items-center justify-center p-2 border border-gray-300 rounded-lg hover:bg-red-50 hover:border-red-400 transition-all group" data-v-5dd820e1><i class="fab fa-google text-red-600 group-hover:scale-110 transition-transform" data-v-5dd820e1></i></button><button class="flex items-center justify-center p-2 border border-gray-300 rounded-lg hover:bg-gray-50 hover:border-gray-600 transition-all group" data-v-5dd820e1><i class="fab fa-apple text-gray-800 group-hover:scale-110 transition-transform" data-v-5dd820e1></i></button></div></form><p class="text-center mt-8 text-gray-600" data-v-5dd820e1> Don&#39;t have an account? `);
      _push(ssrRenderComponent(_component_router_link, {
        to: "/register",
        class: "text-orange-600 hover:text-orange-800 font-medium hover:underline transition-all"
      }, {
        default: withCtx((_, _push2, _parent2, _scopeId) => {
          if (_push2) {
            _push2(` Sign up now `);
          } else {
            return [
              createTextVNode(" Sign up now ")
            ];
          }
        }),
        _: 1
      }, _parent));
      _push(`</p><div class="text-center mt-4 text-xs text-gray-500" data-v-5dd820e1><p data-v-5dd820e1>By signing in, you agree to our</p><div class="space-x-2" data-v-5dd820e1><a href="#" class="hover:text-orange-600 hover:underline transition-colors" data-v-5dd820e1>Terms of Service</a><span data-v-5dd820e1>•</span><a href="#" class="hover:text-orange-600 hover:underline transition-colors" data-v-5dd820e1>Privacy Policy</a></div></div></div></div>`);
    };
  }
};
const _sfc_setup = _sfc_main.setup;
_sfc_main.setup = (props, ctx) => {
  const ssrContext = useSSRContext();
  (ssrContext.modules || (ssrContext.modules = /* @__PURE__ */ new Set())).add("pages/login.vue");
  return _sfc_setup ? _sfc_setup(props, ctx) : void 0;
};
const login = /* @__PURE__ */ _export_sfc(_sfc_main, [["__scopeId", "data-v-5dd820e1"]]);
export {
  login as default
};
//# sourceMappingURL=login.vue.mjs.map
