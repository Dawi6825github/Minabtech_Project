import { resolveComponent, mergeProps, useSSRContext } from "vue";
import { ssrRenderAttrs, ssrRenderComponent } from "vue/server-renderer";
/* empty css                   */
import _export_sfc from "../../_virtual/_plugin-vue_export-helper.mjs";
const _sfc_main = {
  data() {
    return {
      user: null,
      // User details (fetched from Vuex or API)
      userRecipes: [],
      // List of recipes created by the user
      userStats: {},
      // User-specific stats (e.g., number of recipes, followers)
      isEditing: false
      // Flag to control whether the edit profile form is visible
    };
  },
  methods: {
    fetchUserData() {
    },
    updateProfile(updatedUser) {
    },
    deleteAccount() {
    }
  },
  created() {
    this.fetchUserData();
  }
};
function _sfc_ssrRender(_ctx, _push, _parent, _attrs, $props, $setup, $data, $options) {
  const _component_UserInfo = resolveComponent("UserInfo");
  const _component_UserRecipes = resolveComponent("UserRecipes");
  const _component_UserStats = resolveComponent("UserStats");
  const _component_EditUserProfile = resolveComponent("EditUserProfile");
  const _component_DeleteAccount = resolveComponent("DeleteAccount");
  _push(`<div${ssrRenderAttrs(mergeProps({ class: "user-profile" }, _attrs))} data-v-f48e221e>`);
  _push(ssrRenderComponent(_component_UserInfo, { user: $data.user }, null, _parent));
  _push(ssrRenderComponent(_component_UserRecipes, { recipes: $data.userRecipes }, null, _parent));
  _push(ssrRenderComponent(_component_UserStats, { stats: $data.userStats }, null, _parent));
  if ($data.isEditing) {
    _push(ssrRenderComponent(_component_EditUserProfile, {
      user: $data.user,
      onUpdate: $options.updateProfile
    }, null, _parent));
  } else {
    _push(`<!---->`);
  }
  _push(ssrRenderComponent(_component_DeleteAccount, { onConfirmDelete: $options.deleteAccount }, null, _parent));
  _push(`</div>`);
}
const _sfc_setup = _sfc_main.setup;
_sfc_main.setup = (props, ctx) => {
  const ssrContext = useSSRContext();
  (ssrContext.modules || (ssrContext.modules = /* @__PURE__ */ new Set())).add("pages/user-profile.vue");
  return _sfc_setup ? _sfc_setup(props, ctx) : void 0;
};
const userProfile = /* @__PURE__ */ _export_sfc(_sfc_main, [["ssrRender", _sfc_ssrRender], ["__scopeId", "data-v-f48e221e"]]);
export {
  userProfile as default
};
//# sourceMappingURL=user-profile.vue.mjs.map
