import axios from 'axios';
import { getCookie, deleteCookie, setCookie } from 'cookies-next';
import { _ROUTES } from './constants';

const unProtectedRoutes = [
  _ROUTES.LADING_PAGE,
  _ROUTES.LOGIN_PAGE,
  _ROUTES.REGISTER_PAGE,
];

const axiosProtectedAPI = axios.create({
  baseURL: process.env.NEXT_PUBLIC_BACKEND_DOMAIN,
  withCredentials: true,
});

axiosProtectedAPI.interceptors.request.use(
  async (config: any) => {
    const token = getCookie('access_token');
    if (token) {
      config.headers['Authorization'] = `Bearer ${token}`;
    }
    return config;
  },
  (error: any) => {
    return error;
  }
);

axiosProtectedAPI.interceptors.response.use(
  (response) => response,
  async (error: any) => {
    const config = error?.config;
    try {
      if (
        error.response.status === 401 &&
        !config?.sent &&
        error.response.data.message === 'Unauthorized access'
      ) {
        config.sent = true;
        const res = await newToken();
        if (res?.data) {
          setCookie('access_token', res.data.access_token);
          config.headers['Authorization'] = `Bearer ${res.data.access_token}`;
          return axiosProtectedAPI(config)
        }
        return config;
      }
    } catch (err) {
      deleteCookie('refresh_token');
      deleteCookie('access_token');
    }
    // deleteCookie('refresh_token');
    // deleteCookie('access_token');
    // if (!unProtectedRoutes.includes(window.location.pathname)) {
    //   window.location.href = '/auth/login';
    // }
    return error.response;
  }
);

const newToken = async () => {
  const token = getCookie('refresh_token');
  if (token) {
    try {
      const res = await axios.get(
        `${process.env.NEXT_PUBLIC_BACKEND_DOMAIN}auth/token`,
        {
          headers: {
            'X-Refresh-Token': token,
          },
        }
      );
      return res;
    } catch (error: any) {
      return error;
    }
  }
};

export default axiosProtectedAPI;
