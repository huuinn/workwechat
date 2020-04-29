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

# 方法介绍

- TokenService() //token 管理服务
- GetAccessToken() //获取 access_token
- UploadService() //素材管理服务
- SingleService() //单发消息服务
- GroupService() //群发消息服务

# 使用范例

```
package main

import (
	"github.com/go-redis/redis/v7"
	"github.com/huuinn/workwechat"
)

func main() {
	cache := redis.NewClient(&redis.Options{
		Addr:     "localhost:6379",
		Password: "",
		DB:       0,
	})

	cnf := &workwechat.Config{
		CorpId:                 "", //企业ID
		CorpSecrt:              "", //Secret
		AgentId:                1000002, //对应agentId
		Safe:                   0, //表示是否是保密消息，0表示否，1表示是，默认0
		EnableIdTrans:          0,//表示是否开启id转译，0表示否，1表示是，默认0
		EnableDuplicateCheck:   0,//表示是否开启重复消息检查，0表示否，1表示是，默认0
		DuplicateCheckInterval: 1800,//表示是否重复消息检查的时间间隔，默认1800s，最大不超过4小时
	}

	client := workwechat.NewClient(cnf, cache)
	token := client.GetAccessToken()

	client.SingleService().SendText(token, "wangmingfeng", "", "", "test")
}

```
