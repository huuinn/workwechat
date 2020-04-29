package single

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
)

type Service interface {
	SendText(token, toUser, toParty, toTag, context string, opts ...OptionFunc) (*MessageResp, error)
	SendImage(token, toUser, toParty, toTag, mediaId string, opts ...OptionFunc) (*MessageResp, error)
	SendVoice(token, toUser, toParty, toTag, mediaId string, opts ...OptionFunc) (*MessageResp, error)
	SendVideo(token, toUser, toParty, toTag, mediaId, title, description string, opts ...OptionFunc) (*MessageResp, error)
	SendFile(token, toUser, toParty, toTag, mediaId string, opts ...OptionFunc) (*MessageResp, error)
	SendTextCard(token, toUser, toParty, toTag, title, description, url, btntxt string, opts ...OptionFunc) (*MessageResp, error)
	SendNews(token, toUser, toParty, toTag, title, description, url, picurl string, opts ...OptionFunc) (*MessageResp, error)
	SendMpNews(token, toUser, toParty, toTag, title, thumbMediaId, author, contentSourceUrl, context, digest string, opts ...OptionFunc) (*MessageResp, error)
	SendMarkDown(token, toUser, toParty, toTag, context string, opts ...OptionFunc) (*MessageResp, error)
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

func (s *service) GetDefaultOption() *Options {
	return &Options{
		Safe:                   s.Safe,
		EnableIdTrans:          s.EnableIdTrans,
		EnableDuplicateCheck:   s.EnableDuplicateCheck,
		DuplicateCheckInterval: s.DuplicateCheckInterval,
	}
}

func (s *service) SendText(token, toUser, toParty, toTag, context string, opts ...OptionFunc) (*MessageResp, error) {
	defaultOptions := s.GetDefaultOption()
	for _, o := range opts {
		o(defaultOptions)
	}

	url := _GetUrl(token)

	body := TextMessage{}
	body.ToUser = toUser
	body.ToParty = toParty
	body.ToTag = toTag
	body.MsgType = "text"
	body.AgentId = s.AgentId
	body.Safe = defaultOptions.Safe
	body.EnableIdTrans = defaultOptions.EnableIdTrans
	body.EnableDuplicateCheck = defaultOptions.EnableDuplicateCheck
	body.DuplicateCheckInterval = defaultOptions.DuplicateCheckInterval
	body.Text.Content = context

	resp, err := _PostRequest(url, body)

	if err != nil {
		return nil, err
	}

	return resp, nil
}

func (s *service) SendImage(token, toUser, toParty, toTag, mediaId string, opts ...OptionFunc) (*MessageResp, error) {
	defaultOptions := s.GetDefaultOption()
	for _, o := range opts {
		o(defaultOptions)
	}

	url := _GetUrl(token)

	body := ImageMessage{}
	body.ToUser = toUser
	body.ToParty = toParty
	body.ToTag = toTag
	body.MsgType = "image"
	body.AgentId = s.AgentId
	body.Safe = defaultOptions.Safe
	body.EnableIdTrans = defaultOptions.EnableIdTrans
	body.EnableDuplicateCheck = defaultOptions.EnableDuplicateCheck
	body.DuplicateCheckInterval = defaultOptions.DuplicateCheckInterval
	body.Image.MediaId = mediaId

	resp, err := _PostRequest(url, body)

	if err != nil {
		return nil, err
	}

	return resp, nil
}

func (s *service) SendVoice(token, toUser, toParty, toTag, mediaId string, opts ...OptionFunc) (*MessageResp, error) {
	defaultOptions := s.GetDefaultOption()
	for _, o := range opts {
		o(defaultOptions)
	}

	url := _GetUrl(token)

	body := VoiceMessage{}
	body.ToUser = toUser
	body.ToParty = toParty
	body.ToTag = toTag
	body.MsgType = "voice"
	body.AgentId = s.AgentId
	body.Safe = defaultOptions.Safe
	body.EnableIdTrans = defaultOptions.EnableIdTrans
	body.EnableDuplicateCheck = defaultOptions.EnableDuplicateCheck
	body.DuplicateCheckInterval = defaultOptions.DuplicateCheckInterval
	body.Voice.MediaId = mediaId

	resp, err := _PostRequest(url, body)

	if err != nil {
		return nil, err
	}

	return resp, nil
}

func (s *service) SendVideo(token, toUser, toParty, toTag, mediaId, title, description string, opts ...OptionFunc) (*MessageResp, error) {
	defaultOptions := s.GetDefaultOption()
	for _, o := range opts {
		o(defaultOptions)
	}

	url := _GetUrl(token)

	body := VideoMessage{}
	body.ToUser = toUser
	body.ToParty = toParty
	body.ToTag = toTag
	body.MsgType = "video"
	body.AgentId = s.AgentId
	body.Safe = defaultOptions.Safe
	body.EnableIdTrans = defaultOptions.EnableIdTrans
	body.EnableDuplicateCheck = defaultOptions.EnableDuplicateCheck
	body.DuplicateCheckInterval = defaultOptions.DuplicateCheckInterval
	body.Video.MediaId = mediaId
	body.Video.Title = title
	body.Video.Description = description

	resp, err := _PostRequest(url, body)

	if err != nil {
		return nil, err
	}

	return resp, nil
}

func (s *service) SendFile(token, toUser, toParty, toTag, mediaId string, opts ...OptionFunc) (*MessageResp, error) {
	defaultOptions := s.GetDefaultOption()
	for _, o := range opts {
		o(defaultOptions)
	}

	url := _GetUrl(token)

	body := FileMessage{}
	body.ToUser = toUser
	body.ToParty = toParty
	body.ToTag = toTag
	body.MsgType = "file"
	body.AgentId = s.AgentId
	body.Safe = defaultOptions.Safe
	body.EnableIdTrans = defaultOptions.EnableIdTrans
	body.EnableDuplicateCheck = defaultOptions.EnableDuplicateCheck
	body.DuplicateCheckInterval = defaultOptions.DuplicateCheckInterval
	body.File.MediaId = mediaId

	resp, err := _PostRequest(url, body)

	if err != nil {
		return nil, err
	}

	return resp, nil
}

func (s *service) SendTextCard(token, toUser, toParty, toTag, title, description, url, btntxt string, opts ...OptionFunc) (*MessageResp, error) {
	defaultOptions := s.GetDefaultOption()
	for _, o := range opts {
		o(defaultOptions)
	}

	reqUrl := _GetUrl(token)

	body := TextCardMessage{}
	body.ToUser = toUser
	body.ToParty = toParty
	body.ToTag = toTag
	body.MsgType = "textcard"
	body.AgentId = s.AgentId
	body.Safe = defaultOptions.Safe
	body.EnableIdTrans = defaultOptions.EnableIdTrans
	body.EnableDuplicateCheck = defaultOptions.EnableDuplicateCheck
	body.DuplicateCheckInterval = defaultOptions.DuplicateCheckInterval
	body.TextCard.Title = title
	body.TextCard.Description = description
	body.TextCard.Url = url
	body.TextCard.BtnTxt = btntxt

	resp, err := _PostRequest(reqUrl, body)

	if err != nil {
		return nil, err
	}

	return resp, nil
}

func (s *service) SendNews(token, toUser, toParty, toTag, title, description, url, picurl string, opts ...OptionFunc) (*MessageResp, error) {
	defaultOptions := s.GetDefaultOption()
	for _, o := range opts {
		o(defaultOptions)
	}

	reqUrl := _GetUrl(token)

	body := NewsMessage{}
	body.ToUser = toUser
	body.ToParty = toParty
	body.ToTag = toTag
	body.MsgType = "news"
	body.AgentId = s.AgentId
	body.Safe = defaultOptions.Safe
	body.EnableIdTrans = defaultOptions.EnableIdTrans
	body.EnableDuplicateCheck = defaultOptions.EnableDuplicateCheck
	body.DuplicateCheckInterval = defaultOptions.DuplicateCheckInterval
	article := Article{
		Title:       title,
		Description: description,
		Url:         url,
		PicUrl:      picurl,
	}
	body.News.Articles = append(body.News.Articles, article)

	resp, err := _PostRequest(reqUrl, body)

	if err != nil {
		return nil, err
	}

	return resp, nil
}
func (s *service) SendMpNews(token, toUser, toParty, toTag, title, thumbMediaId, author, contentSourceUrl, context, digest string, opts ...OptionFunc) (*MessageResp, error) {
	defaultOptions := s.GetDefaultOption()
	for _, o := range opts {
		o(defaultOptions)
	}

	url := _GetUrl(token)

	body := MpNewsMessage{}
	body.ToUser = toUser
	body.ToParty = toParty
	body.ToTag = toTag
	body.MsgType = "mpnews"
	body.AgentId = s.AgentId
	body.Safe = defaultOptions.Safe
	body.EnableIdTrans = defaultOptions.EnableIdTrans
	body.EnableDuplicateCheck = defaultOptions.EnableDuplicateCheck
	body.DuplicateCheckInterval = defaultOptions.DuplicateCheckInterval
	mpArticle := MpArticle{
		Title:            title,
		ThumbMediaId:     thumbMediaId,
		Author:           author,
		ContentSourceUrl: contentSourceUrl,
		Content:          context,
		Digest:           digest,
	}
	body.MpNews.Articles = append(body.MpNews.Articles, mpArticle)

	resp, err := _PostRequest(url, body)

	if err != nil {
		return nil, err
	}

	return resp, nil
}
func (s *service) SendMarkDown(token, toUser, toParty, toTag, context string, opts ...OptionFunc) (*MessageResp, error) {
	defaultOptions := s.GetDefaultOption()
	for _, o := range opts {
		o(defaultOptions)
	}

	url := _GetUrl(token)

	body := MarkdownMessage{}
	body.ToUser = toUser
	body.ToParty = toParty
	body.ToTag = toTag
	body.MsgType = "markdown"
	body.AgentId = s.AgentId
	body.Safe = defaultOptions.Safe
	body.EnableIdTrans = defaultOptions.EnableIdTrans
	body.EnableDuplicateCheck = defaultOptions.EnableDuplicateCheck
	body.DuplicateCheckInterval = defaultOptions.DuplicateCheckInterval
	body.Markdown.Content = context

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
