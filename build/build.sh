#!/bin/bash

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