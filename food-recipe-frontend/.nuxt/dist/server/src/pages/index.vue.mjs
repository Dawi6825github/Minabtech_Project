import __nuxt_component_0 from "../../node_modules/nuxt/dist/app/components/nuxt-link.mjs";
import { ref, mergeProps, withCtx, createTextVNode, useSSRContext } from "vue";
import { ssrRenderAttrs, ssrRenderAttr, ssrRenderComponent, ssrRenderList, ssrInterpolate, ssrRenderClass } from "vue/server-renderer";
/* empty css            */
import _export_sfc from "../../_virtual/_plugin-vue_export-helper.mjs";
const _sfc_main = {
  __name: "index",
  __ssrInlineRender: true,
  setup(__props) {
    const searchQuery = ref("");
    const userCount = ref(15420);
    const categories = [
      { id: 1, name: "Italian", image: "/categories/italian.jpg" },
      { id: 2, name: "Asian", image: "/categories/asian.jpg" },
      { id: 3, name: "Mexican", image: "/categories/mexican.jpg" },
      { id: 4, name: "Desserts", image: "/categories/desserts.jpg" },
      { id: 5, name: "Vegetarian", image: "/categories/vegetarian.jpg" },
      { id: 6, name: "Seafood", image: "/categories/seafood.jpg" }
    ];
    const testimonials = [
      {
        name: "Sarah Johnson",
        location: "New York, USA",
        avatar: "/avatars/sarah.jpg",
        content: "I found the most amazing pasta recipe here! The community is so supportive and the tips are invaluable.",
        date: "March 15, 2024",
        likes: 124
      },
      {
        name: "Marco Rivera",
        location: "Barcelona, Spain",
        avatar: "/avatars/marco.jpg",
        content: "As a professional chef, I love sharing my recipes here. The feedback from the community helps me improve.",
        date: "March 12, 2024",
        likes: 98
      },
      {
        name: "Emily Chen",
        location: "Singapore",
        avatar: "/avatars/emily.jpg",
        content: "This platform has transformed my cooking journey. The recipes are authentic and easy to follow.",
        date: "March 10, 2024",
        likes: 156
      }
    ];
    const features = [
      {
        id: 1,
        icon: "fas fa-utensils",
        title: "Authentic Recipes",
        description: "Discover recipes from real home cooks"
      },
      {
        id: 2,
        icon: "fas fa-users",
        title: "Community Driven",
        description: "Join a community of food enthusiasts"
      },
      {
        id: 3,
        icon: "fas fa-star",
        title: "Verified Reviews",
        description: "Honest feedback from real users"
      },
      {
        id: 4,
        icon: "fas fa-mobile-alt",
        title: "Mobile Friendly",
        description: "Cook from anywhere, anytime"
      }
    ];
    return (_ctx, _push, _parent, _attrs) => {
      const _component_NuxtLink = __nuxt_component_0;
      _push(`<div${ssrRenderAttrs(mergeProps({ class: "min-h-screen bg-gray-50" }, _attrs))} data-v-73b71f6a><div class="relative bg-gradient-to-r from-orange-500 to-red-600 py-16" data-v-73b71f6a><div class="container mx-auto px-4 text-center" data-v-73b71f6a><h1 class="text-4xl md:text-6xl font-bold text-white mb-6" data-v-73b71f6a>Discover Amazing Food Recipes</h1><p class="text-xl text-white mb-8" data-v-73b71f6a>Share your cooking journey and explore authentic recipes from food lovers worldwide</p><div class="max-w-2xl mx-auto mb-8" data-v-73b71f6a><input${ssrRenderAttr("value", searchQuery.value)} type="text" placeholder="Search for recipes..." class="w-full px-6 py-3 rounded-full text-gray-700 focus:outline-none focus:ring-2 focus:ring-red-500" data-v-73b71f6a></div><div class="flex justify-center gap-4" data-v-73b71f6a>`);
      _push(ssrRenderComponent(_component_NuxtLink, {
        to: "/recipes",
        class: "bg-white text-red-600 px-8 py-3 rounded-full font-semibold hover:bg-gray-100 transition"
      }, {
        default: withCtx((_, _push2, _parent2, _scopeId) => {
          if (_push2) {
            _push2(` Explore Recipes `);
          } else {
            return [
              createTextVNode(" Explore Recipes ")
            ];
          }
        }),
        _: 1
      }, _parent));
      _push(`<button class="bg-red-700 text-white px-8 py-3 rounded-full font-semibold hover:bg-red-800 transition" data-v-73b71f6a> Trending Now </button></div></div></div><div class="container mx-auto px-4 py-12" data-v-73b71f6a><h2 class="text-3xl font-bold text-center mb-8" data-v-73b71f6a>Popular Categories</h2><div class="flex overflow-x-auto gap-4 pb-4" data-v-73b71f6a><!--[-->`);
      ssrRenderList(categories, (category) => {
        _push(`<div class="flex-shrink-0 cursor-pointer hover:transform hover:scale-105 transition" data-v-73b71f6a><div class="w-40 h-40 rounded-lg overflow-hidden relative" data-v-73b71f6a><img${ssrRenderAttr("src", category.image)}${ssrRenderAttr("alt", category.name)} class="w-full h-full object-cover" data-v-73b71f6a><div class="absolute inset-0 bg-black bg-opacity-40 flex items-center justify-center" data-v-73b71f6a><span class="text-white font-semibold" data-v-73b71f6a>${ssrInterpolate(category.name)}</span></div></div></div>`);
      });
      _push(`<!--]--></div></div><div class="container mx-auto px-4 py-16" data-v-73b71f6a><h2 class="text-3xl font-bold text-center mb-12" data-v-73b71f6a>What Our Food Lovers Say</h2><div class="grid grid-cols-1 md:grid-cols-3 gap-8" data-v-73b71f6a><!--[-->`);
      ssrRenderList(testimonials, (testimony, index2) => {
        _push(`<div class="bg-white p-6 rounded-lg shadow-lg hover:shadow-xl transition" data-v-73b71f6a><div class="flex items-center mb-4" data-v-73b71f6a><img${ssrRenderAttr("src", testimony.avatar)}${ssrRenderAttr("alt", testimony.name)} class="w-12 h-12 rounded-full" data-v-73b71f6a><div class="ml-4" data-v-73b71f6a><h3 class="font-semibold" data-v-73b71f6a>${ssrInterpolate(testimony.name)}</h3><p class="text-gray-600" data-v-73b71f6a>${ssrInterpolate(testimony.location)}</p></div></div><p class="text-gray-700" data-v-73b71f6a>${ssrInterpolate(testimony.content)}</p><div class="mt-4 flex items-center justify-between" data-v-73b71f6a><div class="flex items-center" data-v-73b71f6a><span class="text-yellow-400" data-v-73b71f6a>★★★★★</span><span class="ml-2 text-gray-600" data-v-73b71f6a>${ssrInterpolate(testimony.date)}</span></div><button class="text-red-500 hover:text-red-600 transition" data-v-73b71f6a><i class="fas fa-heart" data-v-73b71f6a></i> ${ssrInterpolate(testimony.likes)}</button></div></div>`);
      });
      _push(`<!--]--></div></div><div class="bg-gray-100 py-16" data-v-73b71f6a><div class="container mx-auto px-4" data-v-73b71f6a><h2 class="text-3xl font-bold text-center mb-12" data-v-73b71f6a>Why Choose Our Platform</h2><div class="grid grid-cols-1 md:grid-cols-4 gap-8" data-v-73b71f6a><!--[-->`);
      ssrRenderList(features, (feature) => {
        _push(`<div class="text-center transform hover:scale-105 transition cursor-pointer" data-v-73b71f6a><div class="bg-white p-6 rounded-full w-20 h-20 mx-auto mb-4 flex items-center justify-center" data-v-73b71f6a><i class="${ssrRenderClass(feature.icon + " text-3xl text-red-600")}" data-v-73b71f6a></i></div><h3 class="font-semibold mb-2" data-v-73b71f6a>${ssrInterpolate(feature.title)}</h3><p class="text-gray-600" data-v-73b71f6a>${ssrInterpolate(feature.description)}</p></div>`);
      });
      _push(`<!--]--></div></div></div><div class="bg-red-600 py-16" data-v-73b71f6a><div class="container mx-auto px-4 text-center" data-v-73b71f6a><h2 class="text-3xl font-bold text-white mb-6" data-v-73b71f6a>Ready to Share Your Recipe?</h2><p class="text-white mb-8" data-v-73b71f6a>Join ${ssrInterpolate(userCount.value.toLocaleString())} food lovers and share your culinary masterpieces</p><div class="flex justify-center gap-4" data-v-73b71f6a>`);
      _push(ssrRenderComponent(_component_NuxtLink, {
        to: "/submit-recipe",
        class: "bg-white text-red-600 px-8 py-3 rounded-full font-semibold hover:bg-gray-100 transition"
      }, {
        default: withCtx((_, _push2, _parent2, _scopeId) => {
          if (_push2) {
            _push2(` Share Now `);
          } else {
            return [
              createTextVNode(" Share Now ")
            ];
          }
        }),
        _: 1
      }, _parent));
      _push(`<button class="bg-transparent border-2 border-white text-white px-8 py-3 rounded-full font-semibold hover:bg-white hover:text-red-600 transition" data-v-73b71f6a> Watch Demo </button></div></div></div></div>`);
    };
  }
};
const _sfc_setup = _sfc_main.setup;
_sfc_main.setup = (props, ctx) => {
  const ssrContext = useSSRContext();
  (ssrContext.modules || (ssrContext.modules = /* @__PURE__ */ new Set())).add("pages/index.vue");
  return _sfc_setup ? _sfc_setup(props, ctx) : void 0;
};
const index = /* @__PURE__ */ _export_sfc(_sfc_main, [["__scopeId", "data-v-73b71f6a"]]);
export {
  index as default
};
//# sourceMappingURL=index.vue.mjs.map
