import axios, {AxiosError, AxiosRequestConfig, AxiosResponse} from 'axios';


// 创建 axios 实例
const instance = axios.create({
    baseURL: import.meta.env.PUBLIC_URL,
    timeout: 10000,
    validateStatus: function (status) {
        return status >= 200 && status < 600; // 默认的
    }
});


instance.interceptors.response.use(
    (response: AxiosResponse) => response,
    (error: AxiosError) => {
        return Promise.reject(error)
    }
);

export function POST<T = any>(url: string, data?: any, config?: AxiosRequestConfig): Promise<AxiosResponse<T>> {
    return request<AxiosResponse<T>>({
        method: 'post',
        url,
        data,
        ...config,
    });
}

function request<T>(config: AxiosRequestConfig): Promise<T> {
    return instance.request<T, T>(config)

}

export function GET<T = any>(url: string, config?: AxiosRequestConfig): Promise<T> {
    return request<T>({method: 'get', url, ...config});
}


export function PUT<T = any>(url: string, data?: any, config?: AxiosRequestConfig): Promise<T> {
    return request<T>({method: 'put', url, data, ...config});
}

export function DELETE<T = any>(url: string, config?: AxiosRequestConfig): Promise<T> {
    return request<T>({method: 'delete', url, ...config});
}