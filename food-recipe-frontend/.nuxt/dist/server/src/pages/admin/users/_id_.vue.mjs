import { mergeProps, useSSRContext } from "vue";
import { ssrRenderAttrs, ssrInterpolate } from "vue/server-renderer";
/* empty css           */
import _export_sfc from "../../../../_virtual/_plugin-vue_export-helper.mjs";
const _sfc_main = {
  data() {
    return {
      user: null
    };
  },
  async mounted() {
    const userId = this.$route.params.id;
    this.user = await this.$store.dispatch("admin/fetchUserById", userId);
  },
  methods: {
    async toggleUserStatus() {
      this.user.active = !this.user.active;
      await this.$store.dispatch("admin/updateUserStatus", { id: this.user.id, active: this.user.active });
    },
    async deleteUser() {
      const confirmDelete = confirm("Are you sure you want to delete this user?");
      if (confirmDelete) {
        await this.$store.dispatch("admin/deleteUser", this.user.id);
        this.$router.push("/admin");
      }
    }
  }
};
function _sfc_ssrRender(_ctx, _push, _parent, _attrs, $props, $setup, $data, $options) {
  _push(`<div${ssrRenderAttrs(mergeProps({ class: "user-details" }, _attrs))} data-v-6547ec6e><h1 data-v-6547ec6e>User Profile</h1>`);
  if ($data.user) {
    _push(`<div data-v-6547ec6e><h2 data-v-6547ec6e>${ssrInterpolate($data.user.name)}</h2><p data-v-6547ec6e>Email: ${ssrInterpolate($data.user.email)}</p><p data-v-6547ec6e>Status: ${ssrInterpolate($data.user.active ? "Active" : "Inactive")}</p><button data-v-6547ec6e>${ssrInterpolate($data.user.active ? "Deactivate" : "Activate")} User </button><button data-v-6547ec6e>Delete User</button></div>`);
  } else {
    _push(`<p data-v-6547ec6e>Loading user data...</p>`);
  }
  _push(`</div>`);
}
const _sfc_setup = _sfc_main.setup;
_sfc_main.setup = (props, ctx) => {
  const ssrContext = useSSRContext();
  (ssrContext.modules || (ssrContext.modules = /* @__PURE__ */ new Set())).add("pages/admin/users/[id].vue");
  return _sfc_setup ? _sfc_setup(props, ctx) : void 0;
};
const _id_ = /* @__PURE__ */ _export_sfc(_sfc_main, [["ssrRender", _sfc_ssrRender], ["__scopeId", "data-v-6547ec6e"]]);
export {
  _id_ as default
};
//# sourceMappingURL=_id_.vue.mjs.map
