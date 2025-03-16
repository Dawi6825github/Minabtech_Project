import { defineNuxtPlugin } from "../../node_modules/nuxt/dist/app/nuxt.mjs";
import "vue";
import "destr";
import "klona";
import "defu";
import "#internal/nuxt/paths";
import { defineRule, configure } from "../../node_modules/vee-validate/dist/vee-validate.mjs";
import { required, email, min } from "@vee-validate/rules";
const veevalidate_vrEgjDvdfVWQMkSkMc4v54sp8AoC5VgfWBUhez9ksMk = defineNuxtPlugin(() => {
  defineRule("required", required);
  defineRule("email", email);
  defineRule("min", min);
  configure({
    validateOnBlur: true,
    validateOnChange: true,
    validateOnInput: false,
    validateOnModelUpdate: true
  });
});
export {
  veevalidate_vrEgjDvdfVWQMkSkMc4v54sp8AoC5VgfWBUhez9ksMk as default
};
//# sourceMappingURL=veevalidate.mjs.map
