import { resolveComponent, mergeProps, withCtx, createTextVNode, useSSRContext } from "vue";
import { ssrRenderAttrs, ssrRenderList, ssrInterpolate, ssrRenderComponent } from "vue/server-renderer";
/* empty css            */
import _export_sfc from "../../../_virtual/_plugin-vue_export-helper.mjs";
const _sfc_main = {
  data() {
    return {
      users: []
    };
  },
  async mounted() {
    this.users = await this.$store.dispatch("admin/fetchUsers");
  },
  methods: {
    async deleteUser(userId) {
      const confirmDelete = confirm("Are you sure you want to delete this user?");
      if (confirmDelete) {
        await this.$store.dispatch("admin/deleteUser", userId);
        this.users = this.users.filter((user) => user.id !== userId);
      }
    }
  }
};
function _sfc_ssrRender(_ctx, _push, _parent, _attrs, $props, $setup, $data, $options) {
  const _component_router_link = resolveComponent("router-link");
  _push(`<div${ssrRenderAttrs(mergeProps({ class: "admin-dashboard" }, _attrs))} data-v-d4e5877d><h1 data-v-d4e5877d>Admin Dashboard</h1><h2 data-v-d4e5877d>Users</h2>`);
  if ($data.users.length) {
    _push(`<table data-v-d4e5877d><thead data-v-d4e5877d><tr data-v-d4e5877d><th data-v-d4e5877d>ID</th><th data-v-d4e5877d>Name</th><th data-v-d4e5877d>Email</th><th data-v-d4e5877d>Actions</th></tr></thead><tbody data-v-d4e5877d><!--[-->`);
    ssrRenderList($data.users, (user) => {
      _push(`<tr data-v-d4e5877d><td data-v-d4e5877d>${ssrInterpolate(user.id)}</td><td data-v-d4e5877d>${ssrInterpolate(user.name)}</td><td data-v-d4e5877d>${ssrInterpolate(user.email)}</td><td data-v-d4e5877d>`);
      _push(ssrRenderComponent(_component_router_link, {
        to: `/admin/users/${user.id}`
      }, {
        default: withCtx((_, _push2, _parent2, _scopeId) => {
          if (_push2) {
            _push2(`View`);
          } else {
            return [
              createTextVNode("View")
            ];
          }
        }),
        _: 2
      }, _parent));
      _push(`<button data-v-d4e5877d>Delete</button></td></tr>`);
    });
    _push(`<!--]--></tbody></table>`);
  } else {
    _push(`<p data-v-d4e5877d>No users found.</p>`);
  }
  _push(`</div>`);
}
const _sfc_setup = _sfc_main.setup;
_sfc_main.setup = (props, ctx) => {
  const ssrContext = useSSRContext();
  (ssrContext.modules || (ssrContext.modules = /* @__PURE__ */ new Set())).add("pages/admin/index.vue");
  return _sfc_setup ? _sfc_setup(props, ctx) : void 0;
};
const index = /* @__PURE__ */ _export_sfc(_sfc_main, [["ssrRender", _sfc_ssrRender], ["__scopeId", "data-v-d4e5877d"]]);
export {
  index as default
};
//# sourceMappingURL=index.vue.mjs.map
