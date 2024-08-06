#!/bin/bash
# Copyright (c) [2024] K. All rights reserved.
# Use of this source code is governed by a MIT license that can be found in the LICENSE file.
# Author: K. See：https://github.com/Kun-GitHub/RuoYi-Go or https://gitee.com/gitee_kun/RuoYi-Go
# Email: hot_kun@hotmail.com or 867917691@qq.com

SERVER="RuoYi-Go"

# 获取当前脚本所在的目录
BLDIR="$(dirname "$(readlink -f "$0")")"
if [ -e "${SERVER}" ]; then
    rm "${SERVER}"
fi

# 获取项目根目录
ROOT="$(cd "$(dirname "$0")/.." && pwd)"

echo "Building ${SERVER}"
cd "${ROOT}/cmd/api"
go build -o "${BLDIR}/${SERVER}"

echo "Build done"