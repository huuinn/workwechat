package single

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
)

type Service interface {
	SendText(token, toUser, toParty, toTag, context string) (*MessageResp, error)
	SendImage(token, toUser, toParty, toTag, mediaId string) (*MessageResp, error)
	SendVoice(token, toUser, toParty, toTag, mediaId string) (*MessageResp, error)
	SendVideo(token, toUser, toParty, toTag, mediaId, title, description string) (*MessageResp, error)
	SendFile(token, toUser, toParty, toTag, mediaId string) (*MessageResp, error)
}

type service struct {
	AgentId                int
	Safe                   int
	EnableIdTrans          int
	EnableDuplicateCheck   int
	DuplicateCheckInterval int
}

func NewService(agentId, safe, enableIdTrans, enableDuplicateCheck, duplicateCheckInterval int) Service {
	return &service{
		AgentId:                agentId,
		Safe:                   safe,
		EnableIdTrans:          enableIdTrans,
		EnableDuplicateCheck:   enableDuplicateCheck,
		DuplicateCheckInterval: duplicateCheckInterval,
	}
}

func _GetUrl(token string) string {
	return fmt.Sprintf("https://qyapi.weixin.qq.com/cgi-bin/message/send?access_token=%s", token)
}

func (s *service) SendText(token, toUser, toParty, toTag, context string) (*MessageResp, error) {
	url := _GetUrl(token)

	body := TextMessage{}
	body.ToUser = toUser
	body.ToParty = toParty
	body.ToTag = toTag
	body.MsgType = "text"
	body.AgentId = s.AgentId
	body.Safe = s.Safe
	body.EnableIdTrans = s.EnableIdTrans
	body.EnableDuplicateCheck = s.EnableDuplicateCheck
	body.DuplicateCheckInterval = s.DuplicateCheckInterval
	body.Text.Content = context

	resp, err := _PostRequest(url, body)

	if err != nil {
		return nil, err
	}

	return resp, nil
}

func (s *service) SendImage(token, toUser, toParty, toTag, mediaId string) (*MessageResp, error) {
	url := _GetUrl(token)

	body := ImageMessage{}
	body.ToUser = toUser
	body.ToParty = toParty
	body.ToTag = toTag
	body.MsgType = "image"
	body.AgentId = s.AgentId
	body.Safe = s.Safe
	body.EnableIdTrans = s.EnableIdTrans
	body.EnableDuplicateCheck = s.EnableDuplicateCheck
	body.DuplicateCheckInterval = s.DuplicateCheckInterval
	body.Image.MediaId = mediaId

	resp, err := _PostRequest(url, body)

	if err != nil {
		return nil, err
	}

	return resp, nil
}

func (s *service) SendVoice(token, toUser, toParty, toTag, mediaId string) (*MessageResp, error) {
	url := _GetUrl(token)

	body := VoiceMessage{}
	body.ToUser = toUser
	body.ToParty = toParty
	body.ToTag = toTag
	body.MsgType = "voice"
	body.AgentId = s.AgentId
	body.Safe = s.Safe
	body.EnableIdTrans = s.EnableIdTrans
	body.EnableDuplicateCheck = s.EnableDuplicateCheck
	body.DuplicateCheckInterval = s.DuplicateCheckInterval
	body.Voice.MediaId = mediaId

	resp, err := _PostRequest(url, body)

	if err != nil {
		return nil, err
	}

	return resp, nil
}

func (s *service) SendVideo(token, toUser, toParty, toTag, mediaId, title, description string) (*MessageResp, error) {
	url := _GetUrl(token)

	body := VideoMessage{}
	body.ToUser = toUser
	body.ToParty = toParty
	body.ToTag = toTag
	body.MsgType = "video"
	body.AgentId = s.AgentId
	body.Safe = s.Safe
	body.EnableIdTrans = s.EnableIdTrans
	body.EnableDuplicateCheck = s.EnableDuplicateCheck
	body.DuplicateCheckInterval = s.DuplicateCheckInterval
	body.Video.MediaId = mediaId
	body.Video.Title = title
	body.Video.Description = description

	resp, err := _PostRequest(url, body)

	if err != nil {
		return nil, err
	}

	return resp, nil
}

func (s *service) SendFile(token, toUser, toParty, toTag, mediaId string) (*MessageResp, error) {
	url := _GetUrl(token)

	body := FileMessage{}
	body.ToUser = toUser
	body.ToParty = toParty
	body.ToTag = toTag
	body.MsgType = "file"
	body.AgentId = s.AgentId
	body.Safe = s.Safe
	body.EnableIdTrans = s.EnableIdTrans
	body.EnableDuplicateCheck = s.EnableDuplicateCheck
	body.DuplicateCheckInterval = s.DuplicateCheckInterval
	body.File.MediaId = mediaId

	resp, err := _PostRequest(url, body)

	if err != nil {
		return nil, err
	}

	return resp, nil
}

func _PostRequest(url string, data interface{}) (*MessageResp, error) {
	inJson, _ := json.Marshal(data)
	body := bytes.NewBuffer(inJson)
	req, err := http.NewRequest("POST", url, body)
	if err != nil {
		return nil, err
	}
	client := &http.Client{}
	resp, err := client.Do(req)

	if err != nil {
		return nil, err
	}

	defer resp.Body.Close()

	bodyBytes, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return nil, err
	}

	var ret *MessageResp

	err = json.Unmarshal(bodyBytes, &ret)
	if err != nil {
		return nil, err
	}

	return ret, nil
}