# 简介

企业微信 - 企业内部开发 API（go）
主要实现发送各种类型消息到人或群组。

# 功能

- 自动更新 access_token
- 单推
  - 文件
  - 图片
  - 语音
  - 视频
  - 文件
  - 文本卡片
  - 图文
  - markdown
- 群聊
  - 创建群组
  - 获取群组
  - 文件
  - 图片
  - 语音
  - 视频
  - 文件
  - 文本卡片
  - 图文
  - markdown
- 上传临时素材
- 待补充

# 依赖组件

- redis（ps. 多服务器共享 access_token）

# 安装

```
go get github.com/huuinn/workwechat
```
