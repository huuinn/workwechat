# 简介

企业微信 - 企业内部开发 API（go）
主要实现发送各种类型消息到人或群组。

# 功能

- 自动更新 access_token - done
- 单推
  - 文件 - done
  - 图片 - done
  - 语音 - done
  - 视频 - done
  - 文件 - done
  - 文本卡片 - done
  - 图文 - done
  - markdown -done
- 群聊
  - 创建群组 - done
  - 获取群组 - done
  - 文件 - done
  - 图片 - done
  - 语音 - done
  - 视频 - done
  - 文件 - done
  - 文本卡片 - done
  - 图文 - done
  - markdown -done
- 上传临时素材 - done
- 待补充

# 依赖组件

- redis（ps. 多服务器共享 access_token）

# 安装

```
go get github.com/huuinn/workwechat
```
