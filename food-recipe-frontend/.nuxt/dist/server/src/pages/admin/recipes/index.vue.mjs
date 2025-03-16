import { resolveComponent, withCtx, createTextVNode, useSSRContext } from "vue";
import { ssrRenderAttrs, ssrRenderList, ssrRenderAttr, ssrInterpolate, ssrRenderComponent } from "vue/server-renderer";
/* empty css            */
import _export_sfc from "../../../../_virtual/_plugin-vue_export-helper.mjs";
const _sfc_main = {
  data() {
    return {
      recipes: []
    };
  },
  async mounted() {
    this.recipes = await this.$store.dispatch("recipe/fetchRecipes");
  },
  methods: {
    async deleteRecipe(id) {
      const confirmed = confirm("Are you sure you want to delete this recipe?");
      if (confirmed) {
        await this.$store.dispatch("recipe/deleteRecipe", id);
        this.recipes = this.recipes.filter((recipe) => recipe.id !== id);
      }
    }
  }
};
function _sfc_ssrRender(_ctx, _push, _parent, _attrs, $props, $setup, $data, $options) {
  const _component_router_link = resolveComponent("router-link");
  _push(`<div${ssrRenderAttrs(_attrs)} data-v-3591c5c2><h1 data-v-3591c5c2>All Recipes</h1>`);
  if ($data.recipes.length === 0) {
    _push(`<div data-v-3591c5c2>No recipes available.</div>`);
  } else {
    _push(`<ul data-v-3591c5c2><!--[-->`);
    ssrRenderList($data.recipes, (recipe) => {
      _push(`<li data-v-3591c5c2><div class="recipe-item" data-v-3591c5c2><img${ssrRenderAttr("src", recipe.image)} alt="Recipe Image" class="recipe-image" data-v-3591c5c2><div data-v-3591c5c2><h3 data-v-3591c5c2>${ssrInterpolate(recipe.title)}</h3><p data-v-3591c5c2>${ssrInterpolate(recipe.description)}</p>`);
      _push(ssrRenderComponent(_component_router_link, {
        to: `/admin/recipes/edit/${recipe.id}`
      }, {
        default: withCtx((_, _push2, _parent2, _scopeId) => {
          if (_push2) {
            _push2(`Edit`);
          } else {
            return [
              createTextVNode("Edit")
            ];
          }
        }),
        _: 2
      }, _parent));
      _push(`<button data-v-3591c5c2>Delete</button></div></div></li>`);
    });
    _push(`<!--]--></ul>`);
  }
  _push(`</div>`);
}
const _sfc_setup = _sfc_main.setup;
_sfc_main.setup = (props, ctx) => {
  const ssrContext = useSSRContext();
  (ssrContext.modules || (ssrContext.modules = /* @__PURE__ */ new Set())).add("pages/admin/recipes/index.vue");
  return _sfc_setup ? _sfc_setup(props, ctx) : void 0;
};
const index = /* @__PURE__ */ _export_sfc(_sfc_main, [["ssrRender", _sfc_ssrRender], ["__scopeId", "data-v-3591c5c2"]]);
export {
  index as default
};
//# sourceMappingURL=index.vue.mjs.map
