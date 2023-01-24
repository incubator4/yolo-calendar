import axios from "axios";

const request = axios.create({
  baseURL: "/api", // 所有请求的公共地址部分
  timeout: 3000, // 请求超时时间 这里的意思是当请求时间超过5秒还未取得结果时 提示用户请求超时
});

request.interceptors.request.use(
  (config) => {
    // config 请求的所有信息
    // console.log(config);
    return config; // 将配置完成的config对象返回出去 如果不返回 请求讲不会进行
  },
  (err) => {
    // 请求发生错误时的相关处理 抛出错误
    Promise.reject(err);
  }
);

request.interceptors.response.use(
  (res) => {
    // 我们一般在这里处理，请求成功后的错误状态码 例如状态码是500，404，403
    // res 是所有相应的信息
    // console.log(res);
    return Promise.resolve(res.data);
  },
  (err) => {
    // 服务器响应发生错误时的处理
    Promise.reject(err);
  }
);

export default request;
