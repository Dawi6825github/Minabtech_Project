// plugins/veevalidate.js

import { defineNuxtPlugin } from '#app'
import { configure, defineRule } from 'vee-validate'
import { required, email, min } from '@vee-validate/rules'

export default defineNuxtPlugin(() => {
  defineRule('required', required)
  defineRule('email', email)
  defineRule('min', min)

  configure({
    validateOnBlur: true,
    validateOnChange: true,
    validateOnInput: false,
    validateOnModelUpdate: true,
  })
})
