import { ssrRenderAttrs, ssrRenderList, ssrInterpolate } from "vue/server-renderer";
import { ref, useSSRContext } from "vue";
import "hookable";
/* empty css            */
import _export_sfc from "../../../../_virtual/_plugin-vue_export-helper.mjs";
const _sfc_main = {
  __name: "index",
  __ssrInlineRender: true,
  setup(__props) {
    const users = ref([]);
    return (_ctx, _push, _parent, _attrs) => {
      _push(`<div${ssrRenderAttrs(_attrs)} data-v-6400c8dc><h2 data-v-6400c8dc>User Management</h2><table data-v-6400c8dc><thead data-v-6400c8dc><tr data-v-6400c8dc><th data-v-6400c8dc>Name</th><th data-v-6400c8dc>Email</th><th data-v-6400c8dc>Role</th><th data-v-6400c8dc>Status</th><th data-v-6400c8dc>Actions</th></tr></thead><tbody data-v-6400c8dc><!--[-->`);
      ssrRenderList(users.value, (user) => {
        _push(`<tr data-v-6400c8dc><td data-v-6400c8dc>${ssrInterpolate(user.name)}</td><td data-v-6400c8dc>${ssrInterpolate(user.email)}</td><td data-v-6400c8dc>${ssrInterpolate(user.role)}</td><td data-v-6400c8dc>${ssrInterpolate(user.status)}</td><td data-v-6400c8dc><button data-v-6400c8dc>Edit</button><button data-v-6400c8dc>Reset Password</button><button data-v-6400c8dc>Toggle Status</button></td></tr>`);
      });
      _push(`<!--]--></tbody></table></div>`);
    };
  }
};
const _sfc_setup = _sfc_main.setup;
_sfc_main.setup = (props, ctx) => {
  const ssrContext = useSSRContext();
  (ssrContext.modules || (ssrContext.modules = /* @__PURE__ */ new Set())).add("pages/admin/users/index.vue");
  return _sfc_setup ? _sfc_setup(props, ctx) : void 0;
};
const index = /* @__PURE__ */ _export_sfc(_sfc_main, [["__scopeId", "data-v-6400c8dc"]]);
export {
  index as default
};
//# sourceMappingURL=index.vue.mjs.map
