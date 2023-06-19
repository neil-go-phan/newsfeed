export const _ROUTES = {
  USER_HOME: '/feeds',
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
  ADD_CRAWLER :'/admin/crawler/add',
  ADD_CUSTOM_CRAWLER :'/admin/crawler/add/custom',
  ADMIN_CRONJOB: '/admin/cronjob'
}

export const _REGEX = {
  REGEX_USENAME_PASSWORD: /^[a-z0-9_]*$/,
  REGEX_FULLNAME: /^[a-zA-Z0-9_ ]*$/,
}

export const TOASTIFY_TIME = 1000;