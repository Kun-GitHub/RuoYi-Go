# Copyright (c) [2024] K. All rights reserved.
# Use of this source code is governed by a MIT license that can be found in the LICENSE file.
# Author: K. See：https://github.com/Kun-GitHub/RuoYi-Go or https://gitee.com/gitee_kun/RuoYi-Go
# Email: hot_kun@hotmail.com or 867917691@qq.com

# 使用官方的Golang镜像作为构建阶段的基础镜像
FROM golang:1.22-alpine as builder

ENV GOPROXY=https://goproxy.cn,direct
ENV GO111MODULE=on

# 设置工作目录
WORKDIR /app

# 复制go.mod和go.sum文件
COPY go.mod go.sum ./

# 下载依赖
RUN go mod download

# 复制项目源码
COPY . ./

## 设置脚本可执行权限
#RUN chmod +x ./build/build.sh

# 构建Go应用
RUN go build -o RuoYi-Go ./cmd/api

# 使用包含Go运行时的alpine镜像作为最终镜像
FROM alpine:latest

# 设置工作目录
WORKDIR /app

# 将编译好的二进制文件复制到最终镜像中
# 假设build.sh构建出的二进制文件名为RuoYi-Go
COPY --from=builder /app/RuoYi-Go .

# 复制配置文件
COPY --from=builder /app/config/config.yaml ./config/config.yaml

# 暴露端口
EXPOSE 8080

# 启动命令
CMD ["./RuoYi-Go"]