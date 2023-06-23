export const _ROUTES = {
  LADING_PAGE: '/',
  FEATURE_PAGE: '/feature',
  PRICING_PAGE: '/pricing',
  DISCOVER_PAGE: '/discover',
  LOGIN_PAGE: '/auth/login',
  REGISTER_PAGE: '/auth/register',
  TOKEN_REDIRECT: '/auth/token',
  ADMIN_PAGE: '/admin',
  ADMIN_CRAWLER: '/admin/crawler',
  ADMIN_ARTICLES: '/admin/articles',
  ADMIN_CATEGORIES: '/admin/categories',
  ADMIN_TOPICS: '/admin/topics',
  ADD_CRAWLER :'/admin/crawler/add',
  ADD_CUSTOM_CRAWLER :'/admin/crawler/add/custom',
  ADMIN_CRONJOB: '/admin/cronjob',
  FEEDS: '/feeds',
  FEEDS_SEARCH: '/feeds/search',
  FEEDS_SEARCH_WEBS: '/feeds/search/webs',
  FEEDS_SEARCH_WEBS_RESULT: '/feeds/search/webs/results',
  FEEDS_SEARCH_WEBS_CATEGORY: '/feeds/search/webs/category',
  FEEDS_SEARCH_ARTICLES: '/feeds/search/articles'

}

export const _REGEX = {
  REGEX_USENAME_PASSWORD: /^[a-z0-9_]*$/,
  REGEX_FULLNAME: /^[a-zA-Z0-9_ ]*$/,
}

export const TOASTIFY_TIME = 1000;