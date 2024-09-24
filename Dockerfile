FROM golang:1.22 AS builder
ADD ./ /project
WORKDIR /project
ARG APP=proxy_server
ARG VERSION=1.0.0
LABEL version=$VERSION
ENV GO111MODULE=on \
    CGO_ENABLED=0 \
    GOOS=linux \
    GOARCH=arm64 \
    GOPROXY=https://goproxy.io,direct

RUN go build -mod=vendor -ldflags "-X 'main.Version=$VERSION'" -a -v -o "server" "./cmd/$APP"
EXPOSE 8080

FROM alpine:3.12
WORKDIR /app
# 环境
ARG Prod=prod
ENV service_manager=$Prod
# 时区
ENV TZ=Asia/Shanghai
COPY --from=builder /usr/share/zoneinfo/$TZ /usr/share/zoneinfo/$TZ
RUN ln -sf /usr/share/zoneinfo/Asia/Shanghai /etc/localtime && echo $TZ > /etc/timezone
# 程序
COPY --from=builder /project/server /app/server
COPY --from=builder /project/resource /app/resource

ENTRYPOINT ["/app/server"]