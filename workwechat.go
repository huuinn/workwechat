package workwechat

import (
	"github.com/go-redis/redis/v7"
	"github.com/huuinn/workwechat/core/common/token"
	"github.com/huuinn/workwechat/core/media/upload"
	"github.com/huuinn/workwechat/core/message/group"
	"github.com/huuinn/workwechat/core/message/single"
)

type Config struct {
	CorpId                 string
	CorpSecrt              string
	AgentId                int
	Safe                   int
	EnableIdTrans          int
	EnableDuplicateCheck   int
	DuplicateCheckInterval int
}

type Client struct {
	config        *Config
	Cache         *redis.Client
	tokenService  token.Service
	uploadService upload.Service
	singleService single.Service
	groupService  group.Service
}

func NewClient(cnf *Config, cache *redis.Client) *Client {
	return &Client{
		config: cnf,
		Cache:  cache,
	}
}

func (c *Client) TokenService() token.Service {
	if c.tokenService == nil {
		c.tokenService = token.NewService(c.config.CorpId, c.config.CorpSecrt, c.Cache)
	}

	return c.tokenService
}

func (c *Client) GetAccessToken() string {
	accessToken, err := c.TokenService().Get()
	if err != nil {
		panic("get access token error!!!")
	}

	return accessToken
}

func (c *Client) DeleteAccessToken(oldToken string) error {
	return c.TokenService().Delete(oldToken)
}

func (c *Client) UploadService() upload.Service {
	if c.uploadService == nil {
		c.uploadService = upload.NewService()
	}

	return c.uploadService
}

func (c *Client) SingleService() single.Service {
	if c.singleService == nil {
		c.singleService = single.NewService(c.config.AgentId, c.config.Safe, c.config.EnableIdTrans, c.config.EnableDuplicateCheck, c.config.DuplicateCheckInterval)
	}

	return c.singleService
}

func (c *Client) GroupService() group.Service {
	if c.groupService == nil {
		c.groupService = group.NewService(c.config.Safe)
	}

	return c.groupService
}
