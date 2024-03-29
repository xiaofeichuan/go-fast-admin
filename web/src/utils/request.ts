import axios, { AxiosInstance, AxiosRequestConfig } from 'axios';
import { ElMessage, ElMessageBox } from 'element-plus';
import qs from 'qs';
import cache from '/@/utils/cache';

// 配置新建一个 axios 实例
const service: AxiosInstance = axios.create({
	baseURL: import.meta.env.VITE_API_URL,
	timeout: 50000,
	headers: { 'Content-Type': 'application/json' },
	paramsSerializer: {
		serialize(params) {
			return qs.stringify(params, { allowDots: true });
		},
	},
});

// 添加请求拦截器
service.interceptors.request.use(
	(config: AxiosRequestConfig) => {
		// 在发送请求之前做些什么 token
		var token = cache.getToken();
		if (token) {
			config.headers!['Authorization'] = `Bearer ${token}`;
		}
		return config;
	},
	(error) => {
		// 对请求错误做些什么
		return Promise.reject(error);
	}
);

// 添加响应拦截器
service.interceptors.response.use(
	(response) => {
		// 对响应数据做点什么
		const res = response.data;
		if (res.code === 200) {
			//请求成功
			if (!res.success) {
				//业务处理错误
				ElMessage.error(res.message);
			}
			return Promise.resolve(res);
		}
		else if (res.code === 401) {
			// token错误
			cache.clearAll();
			ElMessageBox.alert('你已被登出，请重新登录', '提示', {})
				.then(() => {
					// 去登录页
					window.location.href = '/';
				})
				.catch(() => { });
			return Promise.reject(res);
			
		}
		else {
			//其他异常
			ElMessage.error(res.message);
			return Promise.reject(res);
		}
	},
	(error) => {
		// 对响应错误做点什么
		if (error.message.indexOf('timeout') != -1) {
			ElMessage.error('网络超时');
		} else if (error.message == 'Network Error') {
			ElMessage.error('网络连接错误');
		} else {
			if (error.response.data) ElMessage.error(error.response.statusText);
			else ElMessage.error('接口路径找不到');
		}
		return Promise.reject(error);
	}
);

// 导出 axios 实例
export default service;
