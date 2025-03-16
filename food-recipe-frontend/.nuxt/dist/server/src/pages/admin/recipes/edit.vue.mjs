import { ssrRenderAttrs, ssrRenderAttr, ssrInterpolate } from "vue/server-renderer";
import { useSSRContext } from "vue";
import _export_sfc from "../../../../_virtual/_plugin-vue_export-helper.mjs";
const _sfc_main = {
  data() {
    return {
      recipe: {
        id: null,
        title: "",
        description: "",
        image: null,
        ingredients: "",
        instructions: ""
      }
    };
  },
  async mounted() {
    const recipeId = this.$route.params.id;
    this.recipe = await this.$store.dispatch("recipe/fetchRecipe", recipeId);
  },
  methods: {
    onImageUpload(event) {
      this.recipe.image = event.target.files[0];
    },
    async updateRecipe() {
      await this.$store.dispatch("recipe/updateRecipe", this.recipe);
      this.$router.push("/admin/recipes");
    }
  }
};
function _sfc_ssrRender(_ctx, _push, _parent, _attrs, $props, $setup, $data, $options) {
  _push(`<div${ssrRenderAttrs(_attrs)}><h1>Edit Recipe</h1><form><div><label for="title">Title</label><input${ssrRenderAttr("value", $data.recipe.title)} id="title" required></div><div><label for="description">Description</label><textarea id="description" required>${ssrInterpolate($data.recipe.description)}</textarea></div><div><label for="image">Recipe Image</label><input type="file"></div><div><label for="ingredients">Ingredients</label><textarea id="ingredients" required>${ssrInterpolate($data.recipe.ingredients)}</textarea></div><div><label for="instructions">Instructions</label><textarea id="instructions" required>${ssrInterpolate($data.recipe.instructions)}</textarea></div><button type="submit">Update Recipe</button></form></div>`);
}
const _sfc_setup = _sfc_main.setup;
_sfc_main.setup = (props, ctx) => {
  const ssrContext = useSSRContext();
  (ssrContext.modules || (ssrContext.modules = /* @__PURE__ */ new Set())).add("pages/admin/recipes/edit.vue");
  return _sfc_setup ? _sfc_setup(props, ctx) : void 0;
};
const edit = /* @__PURE__ */ _export_sfc(_sfc_main, [["ssrRender", _sfc_ssrRender]]);
export {
  edit as default
};
//# sourceMappingURL=edit.vue.mjs.map
