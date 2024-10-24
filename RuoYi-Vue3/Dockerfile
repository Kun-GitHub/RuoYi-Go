# Copyright (c) [2024] K. All rights reserved.
# Use of this source code is governed by a MIT license that can be found in the LICENSE file.
# Author: K. See：https://github.com/Kun-GitHub/RuoYi-Go or https://gitee.com/gitee_kun/RuoYi-Go
# Email: hot_kun@hotmail.com or 867917691@qq.com
FROM nginx:latest

MAINTAINER K.

# 将项目根目录下dist文件夹下的所有文件复制到镜像中 /usr/share/nginx/html/ 目录下
COPY dist/ /usr/share/nginx/html/

COPY default.conf /etc/nginx/conf.d/default.conf

# 暴露端口
EXPOSE 80