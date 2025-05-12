import axios, { type AxiosInstance } from 'axios';
import { useUserStore } from '../stores/user';

// 创建 axios 实例
const instance: AxiosInstance = axios.create({
  baseURL: 'https://api.github.com',
  timeout: 10000,
  headers: {
    'Accept': 'application/vnd.github.v3+json'
  }
});

// 请求拦截器
instance.interceptors.request.use(
  (config) => {
    const userStore = useUserStore();
    const token = userStore.accessToken;
    if (token) {
      config.headers.Authorization = `Bearer ${token}`;
    }else{
      window.location.href = '/login';
    }
    return config;
  },
  (error) => {
    return Promise.reject(error);
  }
);

// 响应拦截器
instance.interceptors.response.use(
  (response) => {
    return response;
  },
  (error) => {
    if (error.response) {
      switch (error.response.status) {
        case 401:
          // 未授权，可以在这里处理登出逻辑
          const userStore = useUserStore();
          userStore.clearUser();
          window.location.href = '/login';
          break;
        case 403:
          console.error('没有权限访问该资源');
          break;
        case 404:
          console.error('请求的资源不存在');
          break;
        default:
          console.error('服务器错误');
      }
    } else if (error.request) {
      console.error('网络请求失败');
    } else {
      console.error('请求配置错误');
    }
    return Promise.reject(error);
  }
);

// 获取用户信息
export const getUserInfo = () => {
  return instance.get('/user');
};
// 创建 axios 实例
export const GetInstance: AxiosInstance = axios.create({
  baseURL: 'https://github.com',
  timeout: 10000,
  headers: {
    'Accept': 'application/vnd.github.v3+json'
  }
});

// 请求拦截器
GetInstance.interceptors.request.use(
  (config) => {
    const userStore = useUserStore();
    const token = userStore.accessToken;
    if (token) {
      config.headers.Authorization = `Bearer ${token}`;
    }else{
      window.location.href = '/login';
    }
    return config;
  },
  (error) => {
    return Promise.reject(error);
  }
);

// 响应拦截器
GetInstance.interceptors.response.use(
  (response) => {
    return response;
  },
  (error) => {
    if (error.response) {
      switch (error.response.status) {
        case 401:
          // 未授权，可以在这里处理登出逻辑
          const userStore = useUserStore();
          userStore.clearUser();
          window.location.href = '/login';
          break;
        case 403:
          console.error('没有权限访问该资源');
          break;
        case 404:
          console.error('请求的资源不存在');
          break;
        default:
          console.error('服务器错误');
      }
    } else if (error.request) {
      console.error('网络请求失败');
    } else {
      console.error('请求配置错误');
    }
    return Promise.reject(error);
  }
);
export default instance;