#!/bin/bash

# 1. 编译 Go 程序
cd server && go build main.go

# 2. 启动前端服务
cd ../web && npm run serve

# 3. 启动 RTSP 流媒体服务器
node rtsp-streaming.js

# 脚本执行完毕
echo "脚本执行完毕！"
