# 服务管理中心（scp-service-manager）
## 项目介绍
利用一个 worker 复用 goroutine，减轻 runtime 调度 goroutine 的压力

## 系统设计


## 快速上手
### 开发环境搭建
~~~ shell
# 1、检查go,版本>=1.19
git clone http://git.gtmap.cn/pdc/cab/scp-service-manager.git
cd scp-service-manager && go mod tidy

# 2、安装make

# 3、安装 wire
go install github.com/google/wire/cmd/wire@latest
# 把GO_PATH/bin添加到环境变量中

# 4、安装protoc
# 下载 https://github.com/protocolbuffers/protobuf/releases 解压后把路径添加到环境变量中
protoc --version

# 5、安装protoc-gen-go
go  install github.com/golang/protobuf/protoc-gen-go@latest
# 把GO_PATH/bin添加到环境变量中

# 6、生成IOC代码
make wire

# 7、生成PB协议
make pb
~~~

### window打包部署
~~~shell
# 检查go、zip
go version

# 安装zip
# 直接将本仓库 docs/zip.exe 下载到本地，然后将路径添加到环境变量中

# 打包
make window

# 部署
# 解压zip双击启动.exe
~~~

### linux打包部署
~~~shell
# 检查go
go version

# 打包
make linux version=1.x.x

# 部署linux包
tar -zxvf ./service-manager-v1.x.x.tar.gz
cd service-manager
nohup ./server > output.log 2>1&
~~~

### docker
```shell
make docker name=scp-service-manager version=1.x.x
make guangdong name=scp-service-manager version=1.x.x
```

### k8s打包部署
~~~shell
# 检查是否安装机器是否安装docker
docker -v

# 打包
make docker version=1.x.x
#如果要构建arm64的镜像
make docker-arm64 version=1.x.x

# 将镜像推送到harbor
docker login --username=your_loginname_xxxx 192.168.2.171
docker tag service-manager:1.x.x 192.168.2.171/gtmap_gl/service-manager:1.x.x
docker push 192.168.2.171/gtmap_gl/service-manager:1.x.x
# 注意如果登录出现：failed to verify certificate: x509: cannot validate certificate for 192.168.2.171 because it doesn't contain any IP SA:
vi  /usr/lib/systemd/system/docker.service
# 修改
# ExecStart=/usr/bin/dockerd --insecure-registry https://192.168.2.171
# 重启
systemctl   daemon-reload &&  systemctl   restart docker.service

# 部署
# 1、修改scp-service-manager/k8s中yaml版本号和端口
# 2、将目录 scp-service-manager/k8s上传到k8s主节点
# 3、启动
kubectl create -f ./k8s/deployment.yaml
kubectl create -f ./k8s/service.yaml
# 4、查看
kubectl get pod,svc -n default
# 或者
kubectl describe pod service-manager
# 5、查看日志
kubectl logs -f service-manager-xxx(pod-name)
# 6、访问
# http://master-node:node-port/home
# 7、删除
kubectl delete deployment service-manager
~~~

## Feature