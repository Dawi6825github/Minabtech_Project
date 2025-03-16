import { resolveComponent, mergeProps, withCtx, createTextVNode, useSSRContext } from "vue";
import { ssrRenderAttrs, ssrRenderAttr, ssrRenderComponent } from "vue/server-renderer";
/* empty css             */
import _export_sfc from "../../_virtual/_plugin-vue_export-helper.mjs";
const _sfc_main = {
  name: "Signup",
  data() {
    return {
      username: "",
      email: "",
      password: ""
    };
  },
  methods: {
    handleSignup() {
      console.log("Signup attempted", {
        username: this.username,
        email: this.email,
        password: this.password
      });
    }
  }
};
function _sfc_ssrRender(_ctx, _push, _parent, _attrs, $props, $setup, $data, $options) {
  const _component_router_link = resolveComponent("router-link");
  _push(`<div${ssrRenderAttrs(mergeProps({ class: "signup-container" }, _attrs))} data-v-5fbde5b7><div class="signup-box" data-v-5fbde5b7><h2 data-v-5fbde5b7>Sign Up</h2><form data-v-5fbde5b7><div class="form-group" data-v-5fbde5b7><label data-v-5fbde5b7>Username</label><input type="text"${ssrRenderAttr("value", $data.username)} placeholder="Enter username" required data-v-5fbde5b7></div><div class="form-group" data-v-5fbde5b7><label data-v-5fbde5b7>Email</label><input type="email"${ssrRenderAttr("value", $data.email)} placeholder="Enter email" required data-v-5fbde5b7></div><div class="form-group" data-v-5fbde5b7><label data-v-5fbde5b7>Password</label><input type="password"${ssrRenderAttr("value", $data.password)} placeholder="Enter password" required data-v-5fbde5b7></div><button type="submit" class="signup-btn" data-v-5fbde5b7>Sign Up</button></form><p class="login-link" data-v-5fbde5b7> Already have an account? `);
  _push(ssrRenderComponent(_component_router_link, { to: "/login" }, {
    default: withCtx((_, _push2, _parent2, _scopeId) => {
      if (_push2) {
        _push2(`Login here`);
      } else {
        return [
          createTextVNode("Login here")
        ];
      }
    }),
    _: 1
  }, _parent));
  _push(`</p></div></div>`);
}
const _sfc_setup = _sfc_main.setup;
_sfc_main.setup = (props, ctx) => {
  const ssrContext = useSSRContext();
  (ssrContext.modules || (ssrContext.modules = /* @__PURE__ */ new Set())).add("pages/Signup.vue");
  return _sfc_setup ? _sfc_setup(props, ctx) : void 0;
};
const Signup = /* @__PURE__ */ _export_sfc(_sfc_main, [["ssrRender", _sfc_ssrRender], ["__scopeId", "data-v-5fbde5b7"]]);
export {
  Signup as default
};
//# sourceMappingURL=Signup.vue.mjs.map
