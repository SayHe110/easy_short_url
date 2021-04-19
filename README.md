
### EasyShortUrl

---
### 项目简介
**EasyShortUrl** 为短链接服务，在我们将网址分享给其他人时，有时分享的 URL 会过长，该项目将长链接转换为短链接，直接分享短链接即可，在用户请求短链接时会自动重定向到之前的长链接。

#### 生成长链接
```shell
curl --location --request POST 'localhost:8080/api/gen_short_url' \
--form 'full_url="https://baidu.com"'
```
响应数据：
```json
{
    "Code": 200,
    "Message": "Success",
    "Data": {
        "full_url": "https://baidu.com",
        "short_url": "ABbIbm"
    }
}
```
#### 短链接转发/重定向
```shell
curl --location --request GET 'localhost:8080/ABbIbm'
```

### 项目启动
1. 项目使用 `module` 对包的管理，所以需要使用 `go mod tidy` 安装所依赖的包
2. 修改 `pkg/client/redis_client.go` 文件中连接 `redis` 配置

### 项目结构
```shell
├── http                                  // 路由及接口
├── pkg                                   // 自定义的包
│   ├── client                            // 相关客户端（Redis）
│   ├── response                          // 统一响应格式
│   └── short_url                         // 生成短链接 
└── runtime                               // 运行时产生的临时文件
    └── log                                     // 日志
```
