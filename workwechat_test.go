package workwechat

import (
	"os"
	"testing"

	"github.com/go-redis/redis/v7"
)

var cache *redis.Client
var cnf *Config
var client *Client

func TestMain(m *testing.M) {
	cache = redis.NewClient(&redis.Options{
		Addr:     "localhost:6379",
		Password: "", // no password set
		DB:       0,  // use default DB
	})

	cnf = &Config{
		CorpId:                 "",
		CorpSecrt:              "",
		AgentId:                1000002,
		Safe:                   0,
		EnableIdTrans:          0,
		EnableDuplicateCheck:   0,
		DuplicateCheckInterval: 1800,
	}

	client = NewClient(cnf, cache)
	runTests := m.Run()
	os.Exit(runTests)
}

func TestClient_TestSingleSendText(t *testing.T) {
	token := client.GetAccessToken()

	t.Log(client.SingleService().SendText(token, "wangmingfeng", "", "", "test"))
}

func TestClient_TestSingleSendFile(t *testing.T) {
	token := client.GetAccessToken()
	t.Log(token)

	uploadResp, _ := client.UploadService().UploadFile(token, "test/test.md")
	t.Log(uploadResp.MediaId)

	t.Log(client.SingleService().SendFile(token, "wangmingfeng", "", "", uploadResp.MediaId))
}

func TestClient_TestGroupSendVoice(t *testing.T) {
	token := client.GetAccessToken()

	uploadResp, _ := client.UploadService().UploadFile(token, "test/gs-16b-1c-8000hz.amr")
	t.Log(uploadResp.MediaId)

	t.Log(client.SingleService().SendVoice(token, "wangmingfeng", "", "", uploadResp.MediaId))
}
