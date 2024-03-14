#!/bin/bash


# 1. 启动前端服务
cd ../web
npm run serve & node rtsp-streaming.js

# 脚本执行完毕
echo "脚本执行完毕！"
