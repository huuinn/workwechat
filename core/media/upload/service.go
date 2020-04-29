package upload

import (
	"bytes"
	"encoding/json"
	"errors"
	"fmt"
	"io"
	"io/ioutil"
	"mime/multipart"
	"net/http"
	"os"
	"path"
)

type Service interface {
	UploadImage(token, filePath string) (*UploadResp, error)
	UploadVoice(token, filePath string) (*UploadResp, error)
	UploadVideo(token, filePath string) (*UploadResp, error)
	UploadFile(token, filePath string) (*UploadResp, error)
}

type service struct {
}

func NewService() Service {
	return &service{}
}

func (s *service) UploadImage(token, filePath string) (*UploadResp, error) {
	url := fmt.Sprintf("https://qyapi.weixin.qq.com/cgi-bin/media/upload?access_token=%s&type=%s", token, "image")
	return _PostRequest(url, filePath)
}

func (s *service) UploadVoice(token, filePath string) (*UploadResp, error) {
	url := fmt.Sprintf("https://qyapi.weixin.qq.com/cgi-bin/media/upload?access_token=%s&type=%s", token, "voice")
	return _PostRequest(url, filePath)
}

func (s *service) UploadVideo(token, filePath string) (*UploadResp, error) {
	url := fmt.Sprintf("https://qyapi.weixin.qq.com/cgi-bin/media/upload?access_token=%s&type=%s", token, "video")
	return _PostRequest(url, filePath)
}

func (s *service) UploadFile(token, filePath string) (*UploadResp, error) {
	url := fmt.Sprintf("https://qyapi.weixin.qq.com/cgi-bin/media/upload?access_token=%s&type=%s", token, "file")
	return _PostRequest(url, filePath)
}

func _PostRequest(url, filePath string) (*UploadResp, error) {
	_, fileName := path.Split(filePath)

	file, _ := os.Open(filePath)
	defer file.Close()

	body := &bytes.Buffer{}
	writer := multipart.NewWriter(body)
	part, _ := writer.CreateFormFile("file", fileName)
	len, err := io.Copy(part, file)

	if err != nil {
		return nil, err
	}

	if len <= 0 {
		return nil, errors.New("文件不能为空")
	}

	writer.Close()

	req, _ := http.NewRequest("POST", url, body)
	req.Header.Add("Content-Type", writer.FormDataContentType())
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

	var ret *UploadResp

	err = json.Unmarshal(bodyBytes, &ret)
	if err != nil {
		return nil, err
	}

	return ret, nil
}
