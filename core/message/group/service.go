package group

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
)

type Service interface {
	CreateChat(token string, userList []string, name, owner, chatId string) (*CreateChatResp, error)
	GetChat(token, chatId string) (*GetChatResp, error)
	SendText(token, chatId, context string, opts ...OptionFunc) (*MessageResp, error)
	SendImage(token, chatId, mediaId string, opts ...OptionFunc) (*MessageResp, error)
	SendVoice(token, chatId, mediaId string, opts ...OptionFunc) (*MessageResp, error)
	SendVideo(token, chatId, mediaId, title, description string, opts ...OptionFunc) (*MessageResp, error)
	SendFile(token, chatId, mediaId string, opts ...OptionFunc) (*MessageResp, error)
	SendTextCard(token, chatId, title, description, url, btntxt string, opts ...OptionFunc) (*MessageResp, error)
	SendNews(token, chatId, title, description, url, picurl string, opts ...OptionFunc) (*MessageResp, error)
	SendMpNews(token, chatId, title, thumbMediaId, author, contentSourceUrl, context, digest string, opts ...OptionFunc) (*MessageResp, error)
	SendMarkDown(token, chatId, context string, opts ...OptionFunc) (*MessageResp, error)
}

type service struct {
	Safe int
}

func NewService(safe int) Service {
	return &service{
		Safe: safe,
	}
}

func _CreateChatUrl(token string) string {
	return fmt.Sprintf("https://qyapi.weixin.qq.com/cgi-bin/appchat/create?access_token=%s", token)
}

func _GetChatUrl(token, chatId string) string {
	return fmt.Sprintf("https://qyapi.weixin.qq.com/cgi-bin/appchat/get?access_token=%s&chatid=%s", token, chatId)
}

func _MessageUrl(token string) string {
	return fmt.Sprintf("https://qyapi.weixin.qq.com/cgi-bin/appchat/send?access_token=%s", token)
}

func (s *service) GetDefaultOption() *Options {
	return &Options{
		Safe: s.Safe,
	}
}

func (s *service) CreateChat(token string, userList []string, name, owner, chatId string) (*CreateChatResp, error) {
	url := _CreateChatUrl(token)
	params := make(map[string]interface{})
	params["userlist"] = userList
	if name != "" {
		params["name"] = name
	}

	if owner != "" {
		params["owner"] = owner
	}

	if chatId != "" {
		params["chatid"] = chatId
	}

	resp, err := _CreateChatPostRequest(url, params)

	if err != nil {
		return nil, err
	}

	return resp, nil
}

func (s *service) GetChat(token, chatId string) (*GetChatResp, error) {
	url := _GetChatUrl(token, chatId)

	resp, err := _GetChatRequest(url)

	if err != nil {
		return nil, err
	}

	return resp, nil
}

func (s *service) SendText(token, chatId, context string, opts ...OptionFunc) (*MessageResp, error) {
	defaultOptions := s.GetDefaultOption()
	for _, o := range opts {
		o(defaultOptions)
	}

	url := _MessageUrl(token)

	body := TextMessage{}
	body.ChatId = chatId
	body.MsgType = "text"
	body.Safe = defaultOptions.Safe
	body.Text.Content = context

	resp, err := _PostRequest(url, body)

	if err != nil {
		return nil, err
	}

	return resp, nil
}
func (s *service) SendImage(token, chatId, mediaId string, opts ...OptionFunc) (*MessageResp, error) {
	defaultOptions := s.GetDefaultOption()
	for _, o := range opts {
		o(defaultOptions)
	}

	url := _MessageUrl(token)

	body := ImageMessage{}
	body.ChatId = chatId
	body.MsgType = "image"
	body.Safe = defaultOptions.Safe
	body.Image.MediaId = mediaId

	resp, err := _PostRequest(url, body)

	if err != nil {
		return nil, err
	}

	return resp, nil
}
func (s *service) SendVoice(token, chatId, mediaId string, opts ...OptionFunc) (*MessageResp, error) {
	defaultOptions := s.GetDefaultOption()
	for _, o := range opts {
		o(defaultOptions)
	}

	url := _MessageUrl(token)

	body := VoiceMessage{}
	body.ChatId = chatId
	body.MsgType = "voice"
	body.Safe = defaultOptions.Safe
	body.Voice.MediaId = mediaId

	resp, err := _PostRequest(url, body)

	if err != nil {
		return nil, err
	}

	return resp, nil
}
func (s *service) SendVideo(token, chatId, mediaId, title, description string, opts ...OptionFunc) (*MessageResp, error) {
	defaultOptions := s.GetDefaultOption()
	for _, o := range opts {
		o(defaultOptions)
	}

	url := _MessageUrl(token)

	body := VideoMessage{}
	body.ChatId = chatId
	body.MsgType = "video"
	body.Safe = defaultOptions.Safe
	body.Video.MediaId = mediaId
	body.Video.Title = title
	body.Video.Description = description

	resp, err := _PostRequest(url, body)

	if err != nil {
		return nil, err
	}

	return resp, nil
}
func (s *service) SendFile(token, chatId, mediaId string, opts ...OptionFunc) (*MessageResp, error) {
	defaultOptions := s.GetDefaultOption()
	for _, o := range opts {
		o(defaultOptions)
	}

	url := _MessageUrl(token)

	body := FileMessage{}
	body.ChatId = chatId
	body.MsgType = "file"
	body.Safe = defaultOptions.Safe
	body.File.MediaId = mediaId

	resp, err := _PostRequest(url, body)

	if err != nil {
		return nil, err
	}

	return resp, nil
}
func (s *service) SendTextCard(token, chatId, title, description, url, btntxt string, opts ...OptionFunc) (*MessageResp, error) {
	defaultOptions := s.GetDefaultOption()
	for _, o := range opts {
		o(defaultOptions)
	}

	reqUrl := _MessageUrl(token)

	body := TextCardMessage{}
	body.ChatId = chatId
	body.MsgType = "textcard"
	body.Safe = defaultOptions.Safe
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
func (s *service) SendNews(token, chatId, title, description, url, picurl string, opts ...OptionFunc) (*MessageResp, error) {
	defaultOptions := s.GetDefaultOption()
	for _, o := range opts {
		o(defaultOptions)
	}

	reqUrl := _MessageUrl(token)

	body := NewsMessage{}
	body.ChatId = chatId
	body.MsgType = "news"
	body.Safe = defaultOptions.Safe
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
func (s *service) SendMpNews(token, chatId, title, thumbMediaId, author, contentSourceUrl, context, digest string, opts ...OptionFunc) (*MessageResp, error) {
	defaultOptions := s.GetDefaultOption()
	for _, o := range opts {
		o(defaultOptions)
	}

	url := _MessageUrl(token)

	body := MpNewsMessage{}
	body.ChatId = chatId
	body.MsgType = "mpnews"
	body.Safe = defaultOptions.Safe
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
func (s *service) SendMarkDown(token, chatId, context string, opts ...OptionFunc) (*MessageResp, error) {
	defaultOptions := s.GetDefaultOption()
	for _, o := range opts {
		o(defaultOptions)
	}

	url := _MessageUrl(token)

	body := MarkdownMessage{}
	body.ChatId = chatId
	body.MsgType = "markdown"
	body.Safe = defaultOptions.Safe
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

func _CreateChatPostRequest(url string, data interface{}) (*CreateChatResp, error) {
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

	var ret *CreateChatResp

	err = json.Unmarshal(bodyBytes, &ret)
	if err != nil {
		return nil, err
	}

	return ret, nil
}

func _GetChatRequest(url string) (*GetChatResp, error) {
	resp, err := http.Get(url)
	if err != nil {
		return nil, err
	}

	if err != nil {
		return nil, err
	}

	defer resp.Body.Close()

	bodyBytes, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return nil, err
	}

	var ret *GetChatResp

	err = json.Unmarshal(bodyBytes, &ret)
	if err != nil {
		return nil, err
	}

	return ret, nil
}
