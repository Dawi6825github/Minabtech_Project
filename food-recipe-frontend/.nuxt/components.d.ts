
import type { DefineComponent, SlotsType } from 'vue'
type IslandComponent<T extends DefineComponent> = T & DefineComponent<{}, {refresh: () => Promise<void>}, {}, {}, {}, {}, {}, {}, {}, {}, {}, {}, SlotsType<{ fallback: { error: unknown } }>>
type HydrationStrategies = {
  hydrateOnVisible?: IntersectionObserverInit | true
  hydrateOnIdle?: number | true
  hydrateOnInteraction?: keyof HTMLElementEventMap | Array<keyof HTMLElementEventMap> | true
  hydrateOnMediaQuery?: string
  hydrateAfter?: number
  hydrateWhen?: boolean
  hydrateNever?: true
}
type LazyComponent<T> = (T & DefineComponent<HydrationStrategies, {}, {}, {}, {}, {}, {}, { hydrated: () => void }>)
interface _GlobalComponents {
      'Footer': typeof import("../src/components/Footer.vue")['default']
    'Header': typeof import("../src/components/Header.vue")['default']
    'Navbar': typeof import("../src/components/Navbar.vue")['default']
    'RecipeCard': typeof import("../src/components/RecipeCard.vue")['default']
    'RecipeDetail': typeof import("../src/components/RecipeDetail.vue")['default']
    'RecipeForm': typeof import("../src/components/RecipeForm.vue")['default']
    'RecipeList': typeof import("../src/components/RecipeList.vue")['default']
    'UserProfileDeleteAccount': typeof import("../src/components/user_profile/DeleteAccount.vue")['default']
    'UserProfileEditUserProfile': typeof import("../src/components/user_profile/EditUserProfile.vue")['default']
    'UserProfileUserInfo': typeof import("../src/components/user_profile/UserInfo.vue")['default']
    'UserProfileUserRecipes': typeof import("../src/components/user_profile/UserRecipes.vue")['default']
    'UserProfileUserStats': typeof import("../src/components/user_profile/UserStats.vue")['default']
    'NuxtWelcome': typeof import("../node_modules/nuxt/dist/app/components/welcome.vue")['default']
    'NuxtLayout': typeof import("../node_modules/nuxt/dist/app/components/nuxt-layout")['default']
    'NuxtErrorBoundary': typeof import("../node_modules/nuxt/dist/app/components/nuxt-error-boundary")['default']
    'ClientOnly': typeof import("../node_modules/nuxt/dist/app/components/client-only")['default']
    'DevOnly': typeof import("../node_modules/nuxt/dist/app/components/dev-only")['default']
    'ServerPlaceholder': typeof import("../node_modules/nuxt/dist/app/components/server-placeholder")['default']
    'NuxtLink': typeof import("../node_modules/nuxt/dist/app/components/nuxt-link")['default']
    'NuxtLoadingIndicator': typeof import("../node_modules/nuxt/dist/app/components/nuxt-loading-indicator")['default']
    'NuxtRouteAnnouncer': typeof import("../node_modules/nuxt/dist/app/components/nuxt-route-announcer")['default']
    'NuxtImg': typeof import("../node_modules/nuxt/dist/app/components/nuxt-stubs")['NuxtImg']
    'NuxtPicture': typeof import("../node_modules/nuxt/dist/app/components/nuxt-stubs")['NuxtPicture']
    'ErrorMessage': typeof import("vee-validate")['ErrorMessage']
    'Field': typeof import("vee-validate")['Field']
    'FieldArray': typeof import("vee-validate")['FieldArray']
    'Form': typeof import("vee-validate")['Form']
    'NuxtPage': typeof import("../node_modules/nuxt/dist/pages/runtime/page")['default']
    'NoScript': typeof import("../node_modules/nuxt/dist/head/runtime/components")['NoScript']
    'Link': typeof import("../node_modules/nuxt/dist/head/runtime/components")['Link']
    'Base': typeof import("../node_modules/nuxt/dist/head/runtime/components")['Base']
    'Title': typeof import("../node_modules/nuxt/dist/head/runtime/components")['Title']
    'Meta': typeof import("../node_modules/nuxt/dist/head/runtime/components")['Meta']
    'Style': typeof import("../node_modules/nuxt/dist/head/runtime/components")['Style']
    'Head': typeof import("../node_modules/nuxt/dist/head/runtime/components")['Head']
    'Html': typeof import("../node_modules/nuxt/dist/head/runtime/components")['Html']
    'Body': typeof import("../node_modules/nuxt/dist/head/runtime/components")['Body']
    'NuxtIsland': typeof import("../node_modules/nuxt/dist/app/components/nuxt-island")['default']
    'NuxtRouteAnnouncer': IslandComponent<typeof import("../node_modules/nuxt/dist/app/components/server-placeholder")['default']>
      'LazyFooter': LazyComponent<typeof import("../src/components/Footer.vue")['default']>
    'LazyHeader': LazyComponent<typeof import("../src/components/Header.vue")['default']>
    'LazyNavbar': LazyComponent<typeof import("../src/components/Navbar.vue")['default']>
    'LazyRecipeCard': LazyComponent<typeof import("../src/components/RecipeCard.vue")['default']>
    'LazyRecipeDetail': LazyComponent<typeof import("../src/components/RecipeDetail.vue")['default']>
    'LazyRecipeForm': LazyComponent<typeof import("../src/components/RecipeForm.vue")['default']>
    'LazyRecipeList': LazyComponent<typeof import("../src/components/RecipeList.vue")['default']>
    'LazyUserProfileDeleteAccount': LazyComponent<typeof import("../src/components/user_profile/DeleteAccount.vue")['default']>
    'LazyUserProfileEditUserProfile': LazyComponent<typeof import("../src/components/user_profile/EditUserProfile.vue")['default']>
    'LazyUserProfileUserInfo': LazyComponent<typeof import("../src/components/user_profile/UserInfo.vue")['default']>
    'LazyUserProfileUserRecipes': LazyComponent<typeof import("../src/components/user_profile/UserRecipes.vue")['default']>
    'LazyUserProfileUserStats': LazyComponent<typeof import("../src/components/user_profile/UserStats.vue")['default']>
    'LazyNuxtWelcome': LazyComponent<typeof import("../node_modules/nuxt/dist/app/components/welcome.vue")['default']>
    'LazyNuxtLayout': LazyComponent<typeof import("../node_modules/nuxt/dist/app/components/nuxt-layout")['default']>
    'LazyNuxtErrorBoundary': LazyComponent<typeof import("../node_modules/nuxt/dist/app/components/nuxt-error-boundary")['default']>
    'LazyClientOnly': LazyComponent<typeof import("../node_modules/nuxt/dist/app/components/client-only")['default']>
    'LazyDevOnly': LazyComponent<typeof import("../node_modules/nuxt/dist/app/components/dev-only")['default']>
    'LazyServerPlaceholder': LazyComponent<typeof import("../node_modules/nuxt/dist/app/components/server-placeholder")['default']>
    'LazyNuxtLink': LazyComponent<typeof import("../node_modules/nuxt/dist/app/components/nuxt-link")['default']>
    'LazyNuxtLoadingIndicator': LazyComponent<typeof import("../node_modules/nuxt/dist/app/components/nuxt-loading-indicator")['default']>
    'LazyNuxtRouteAnnouncer': LazyComponent<typeof import("../node_modules/nuxt/dist/app/components/nuxt-route-announcer")['default']>
    'LazyNuxtImg': LazyComponent<typeof import("../node_modules/nuxt/dist/app/components/nuxt-stubs")['NuxtImg']>
    'LazyNuxtPicture': LazyComponent<typeof import("../node_modules/nuxt/dist/app/components/nuxt-stubs")['NuxtPicture']>
    'LazyErrorMessage': LazyComponent<typeof import("vee-validate")['ErrorMessage']>
    'LazyField': LazyComponent<typeof import("vee-validate")['Field']>
    'LazyFieldArray': LazyComponent<typeof import("vee-validate")['FieldArray']>
    'LazyForm': LazyComponent<typeof import("vee-validate")['Form']>
    'LazyNuxtPage': LazyComponent<typeof import("../node_modules/nuxt/dist/pages/runtime/page")['default']>
    'LazyNoScript': LazyComponent<typeof import("../node_modules/nuxt/dist/head/runtime/components")['NoScript']>
    'LazyLink': LazyComponent<typeof import("../node_modules/nuxt/dist/head/runtime/components")['Link']>
    'LazyBase': LazyComponent<typeof import("../node_modules/nuxt/dist/head/runtime/components")['Base']>
    'LazyTitle': LazyComponent<typeof import("../node_modules/nuxt/dist/head/runtime/components")['Title']>
    'LazyMeta': LazyComponent<typeof import("../node_modules/nuxt/dist/head/runtime/components")['Meta']>
    'LazyStyle': LazyComponent<typeof import("../node_modules/nuxt/dist/head/runtime/components")['Style']>
    'LazyHead': LazyComponent<typeof import("../node_modules/nuxt/dist/head/runtime/components")['Head']>
    'LazyHtml': LazyComponent<typeof import("../node_modules/nuxt/dist/head/runtime/components")['Html']>
    'LazyBody': LazyComponent<typeof import("../node_modules/nuxt/dist/head/runtime/components")['Body']>
    'LazyNuxtIsland': LazyComponent<typeof import("../node_modules/nuxt/dist/app/components/nuxt-island")['default']>
    'LazyNuxtRouteAnnouncer': LazyComponent<IslandComponent<typeof import("../node_modules/nuxt/dist/app/components/server-placeholder")['default']>>
}

declare module 'vue' {
  export interface GlobalComponents extends _GlobalComponents { }
}

export const Footer: typeof import("../src/components/Footer.vue")['default']
export const Header: typeof import("../src/components/Header.vue")['default']
export const Navbar: typeof import("../src/components/Navbar.vue")['default']
export const RecipeCard: typeof import("../src/components/RecipeCard.vue")['default']
export const RecipeDetail: typeof import("../src/components/RecipeDetail.vue")['default']
export const RecipeForm: typeof import("../src/components/RecipeForm.vue")['default']
export const RecipeList: typeof import("../src/components/RecipeList.vue")['default']
export const UserProfileDeleteAccount: typeof import("../src/components/user_profile/DeleteAccount.vue")['default']
export const UserProfileEditUserProfile: typeof import("../src/components/user_profile/EditUserProfile.vue")['default']
export const UserProfileUserInfo: typeof import("../src/components/user_profile/UserInfo.vue")['default']
export const UserProfileUserRecipes: typeof import("../src/components/user_profile/UserRecipes.vue")['default']
export const UserProfileUserStats: typeof import("../src/components/user_profile/UserStats.vue")['default']
export const NuxtWelcome: typeof import("../node_modules/nuxt/dist/app/components/welcome.vue")['default']
export const NuxtLayout: typeof import("../node_modules/nuxt/dist/app/components/nuxt-layout")['default']
export const NuxtErrorBoundary: typeof import("../node_modules/nuxt/dist/app/components/nuxt-error-boundary")['default']
export const ClientOnly: typeof import("../node_modules/nuxt/dist/app/components/client-only")['default']
export const DevOnly: typeof import("../node_modules/nuxt/dist/app/components/dev-only")['default']
export const ServerPlaceholder: typeof import("../node_modules/nuxt/dist/app/components/server-placeholder")['default']
export const NuxtLink: typeof import("../node_modules/nuxt/dist/app/components/nuxt-link")['default']
export const NuxtLoadingIndicator: typeof import("../node_modules/nuxt/dist/app/components/nuxt-loading-indicator")['default']
export const NuxtRouteAnnouncer: typeof import("../node_modules/nuxt/dist/app/components/nuxt-route-announcer")['default']
export const NuxtImg: typeof import("../node_modules/nuxt/dist/app/components/nuxt-stubs")['NuxtImg']
export const NuxtPicture: typeof import("../node_modules/nuxt/dist/app/components/nuxt-stubs")['NuxtPicture']
export const ErrorMessage: typeof import("vee-validate")['ErrorMessage']
export const Field: typeof import("vee-validate")['Field']
export const FieldArray: typeof import("vee-validate")['FieldArray']
export const Form: typeof import("vee-validate")['Form']
export const NuxtPage: typeof import("../node_modules/nuxt/dist/pages/runtime/page")['default']
export const NoScript: typeof import("../node_modules/nuxt/dist/head/runtime/components")['NoScript']
export const Link: typeof import("../node_modules/nuxt/dist/head/runtime/components")['Link']
export const Base: typeof import("../node_modules/nuxt/dist/head/runtime/components")['Base']
export const Title: typeof import("../node_modules/nuxt/dist/head/runtime/components")['Title']
export const Meta: typeof import("../node_modules/nuxt/dist/head/runtime/components")['Meta']
export const Style: typeof import("../node_modules/nuxt/dist/head/runtime/components")['Style']
export const Head: typeof import("../node_modules/nuxt/dist/head/runtime/components")['Head']
export const Html: typeof import("../node_modules/nuxt/dist/head/runtime/components")['Html']
export const Body: typeof import("../node_modules/nuxt/dist/head/runtime/components")['Body']
export const NuxtIsland: typeof import("../node_modules/nuxt/dist/app/components/nuxt-island")['default']
export const NuxtRouteAnnouncer: IslandComponent<typeof import("../node_modules/nuxt/dist/app/components/server-placeholder")['default']>
export const LazyFooter: LazyComponent<typeof import("../src/components/Footer.vue")['default']>
export const LazyHeader: LazyComponent<typeof import("../src/components/Header.vue")['default']>
export const LazyNavbar: LazyComponent<typeof import("../src/components/Navbar.vue")['default']>
export const LazyRecipeCard: LazyComponent<typeof import("../src/components/RecipeCard.vue")['default']>
export const LazyRecipeDetail: LazyComponent<typeof import("../src/components/RecipeDetail.vue")['default']>
export const LazyRecipeForm: LazyComponent<typeof import("../src/components/RecipeForm.vue")['default']>
export const LazyRecipeList: LazyComponent<typeof import("../src/components/RecipeList.vue")['default']>
export const LazyUserProfileDeleteAccount: LazyComponent<typeof import("../src/components/user_profile/DeleteAccount.vue")['default']>
export const LazyUserProfileEditUserProfile: LazyComponent<typeof import("../src/components/user_profile/EditUserProfile.vue")['default']>
export const LazyUserProfileUserInfo: LazyComponent<typeof import("../src/components/user_profile/UserInfo.vue")['default']>
export const LazyUserProfileUserRecipes: LazyComponent<typeof import("../src/components/user_profile/UserRecipes.vue")['default']>
export const LazyUserProfileUserStats: LazyComponent<typeof import("../src/components/user_profile/UserStats.vue")['default']>
export const LazyNuxtWelcome: LazyComponent<typeof import("../node_modules/nuxt/dist/app/components/welcome.vue")['default']>
export const LazyNuxtLayout: LazyComponent<typeof import("../node_modules/nuxt/dist/app/components/nuxt-layout")['default']>
export const LazyNuxtErrorBoundary: LazyComponent<typeof import("../node_modules/nuxt/dist/app/components/nuxt-error-boundary")['default']>
export const LazyClientOnly: LazyComponent<typeof import("../node_modules/nuxt/dist/app/components/client-only")['default']>
export const LazyDevOnly: LazyComponent<typeof import("../node_modules/nuxt/dist/app/components/dev-only")['default']>
export const LazyServerPlaceholder: LazyComponent<typeof import("../node_modules/nuxt/dist/app/components/server-placeholder")['default']>
export const LazyNuxtLink: LazyComponent<typeof import("../node_modules/nuxt/dist/app/components/nuxt-link")['default']>
export const LazyNuxtLoadingIndicator: LazyComponent<typeof import("../node_modules/nuxt/dist/app/components/nuxt-loading-indicator")['default']>
export const LazyNuxtRouteAnnouncer: LazyComponent<typeof import("../node_modules/nuxt/dist/app/components/nuxt-route-announcer")['default']>
export const LazyNuxtImg: LazyComponent<typeof import("../node_modules/nuxt/dist/app/components/nuxt-stubs")['NuxtImg']>
export const LazyNuxtPicture: LazyComponent<typeof import("../node_modules/nuxt/dist/app/components/nuxt-stubs")['NuxtPicture']>
export const LazyErrorMessage: LazyComponent<typeof import("vee-validate")['ErrorMessage']>
export const LazyField: LazyComponent<typeof import("vee-validate")['Field']>
export const LazyFieldArray: LazyComponent<typeof import("vee-validate")['FieldArray']>
export const LazyForm: LazyComponent<typeof import("vee-validate")['Form']>
export const LazyNuxtPage: LazyComponent<typeof import("../node_modules/nuxt/dist/pages/runtime/page")['default']>
export const LazyNoScript: LazyComponent<typeof import("../node_modules/nuxt/dist/head/runtime/components")['NoScript']>
export const LazyLink: LazyComponent<typeof import("../node_modules/nuxt/dist/head/runtime/components")['Link']>
export const LazyBase: LazyComponent<typeof import("../node_modules/nuxt/dist/head/runtime/components")['Base']>
export const LazyTitle: LazyComponent<typeof import("../node_modules/nuxt/dist/head/runtime/components")['Title']>
export const LazyMeta: LazyComponent<typeof import("../node_modules/nuxt/dist/head/runtime/components")['Meta']>
export const LazyStyle: LazyComponent<typeof import("../node_modules/nuxt/dist/head/runtime/components")['Style']>
export const LazyHead: LazyComponent<typeof import("../node_modules/nuxt/dist/head/runtime/components")['Head']>
export const LazyHtml: LazyComponent<typeof import("../node_modules/nuxt/dist/head/runtime/components")['Html']>
export const LazyBody: LazyComponent<typeof import("../node_modules/nuxt/dist/head/runtime/components")['Body']>
export const LazyNuxtIsland: LazyComponent<typeof import("../node_modules/nuxt/dist/app/components/nuxt-island")['default']>
export const LazyNuxtRouteAnnouncer: LazyComponent<IslandComponent<typeof import("../node_modules/nuxt/dist/app/components/server-placeholder")['default']>>

export const componentNames: string[]
