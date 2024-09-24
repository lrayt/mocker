# Mocker（数据模拟）
## 介绍
通过JSON文件配置一个数据模拟服务，支持文件、API、简单CRUD、http代理

## 快速开始
### 1、基础配置
~~~json
{
  "host": "0.0.0.0",// 主机IP
  "port": 8080,     // 服务端口
  "cors": false     // 是否支持跨域
}
~~~

### 2、文件服务
~~~json
{
  "host": "0.0.0.0",// 主机IP
  "port": 8080,     // 服务端口
  "cors": false,     // 是否支持跨域
  "router": {
    "assets": {
      "type": "fs",              // 类型为文件服务
      "dir": "${WorkDir}/assets" // 文件目录，注：${WorkDir}表示程序同级目录
      // dir: "d://assets"       // 或绝对路径
    }
  }
}
~~~

### 3、代理服务

### 4、API

### 5、CRUD

## Feature