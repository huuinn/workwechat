package token

import (
	"bytes"
	"encoding/json"
	"errors"
	"fmt"
	"io/ioutil"
	"net/http"
	"sync"
	"time"

	"github.com/go-redis/redis/v7"
)

const RedisAccessTokenKey = "work:wechat:access:token"

type Service interface {
	Get() (string, error)
	Delete(oldToken string) error
}

type service struct {
	sync.Mutex
	CorpId     string
	CorpSecret string
	Cache      *redis.Client
}

func NewService(corpid, corpsecret string, cache *redis.Client) Service {
	return &service{
		CorpId:     corpid,
		CorpSecret: corpsecret,
		Cache:      cache,
	}
}

func (s *service) Get() (string, error) {
	s.Lock()
	defer s.Unlock()

	token := s.Cache.Get(RedisAccessTokenKey).Val()

	if token != "" {
		return token, nil
	}

	url := fmt.Sprintf("https://qyapi.weixin.qq.com/cgi-bin/gettoken?corpid=%s&corpsecret=%s", s.CorpId, s.CorpSecret)

	resp, err := http.Get(url)
	if err != nil {
		return "", err
	}

	defer resp.Body.Close()
	body, err := ioutil.ReadAll(resp.Body)

	if err != nil {
		return "", err
	}

	if !bytes.Contains(body, []byte("access_token")) {
		return "", errors.New("response error")
	}

	atr := AccessTokenResp{}
	err = json.Unmarshal(body, &atr)
	if err != nil {
		return "", err
	}
	token = atr.AccessToken

	s.Cache.Set(RedisAccessTokenKey, token, time.Duration(atr.ExpiresIn)*time.Second)
	return token, nil
}

func (s *service) Delete(oldToken string) error {
	s.Lock()
	defer s.Unlock()

	token := s.Cache.Get(RedisAccessTokenKey).Val()
	if token == oldToken {
		return s.Cache.Del(RedisAccessTokenKey).Err()
	}

	return nil
}
