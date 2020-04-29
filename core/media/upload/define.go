package upload

type UploadResp struct {
	ErrCode   int    `json:"errcode"`
	ErrMsg    string `json:"errmsg"`
	MediaId   string `json:"media_id"`
	CreatedAt string `json:"created_at"`
}
