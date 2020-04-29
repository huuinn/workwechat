package group

import (
	"os"
	"testing"
)

var token string
var chatId string

func TestMain(m *testing.M) {
	token = "1kgise5YIImoXSdL8Ozw0Sc4-9vOgXx5w5Tw_-Isr37nTRy7EpCxxSxLA4Pj0igqxtE8hY3Lze9QLfXlvfm8yGJ-L_ZVtXIsZl7FkhPUWQ_taLoEImEQ9x_HyU5RaHtbfwoykRKSM5w7Zm0FRVbdRvUsJWsYNsVSx6dZ2VuJcfQ_PHJIprZ58eK-lTu0A1ETAbj5DqAiD5WTLV8KuNAypg"
	chatId = "wrp7CLCQAALDZ4gfjm54di5gtmvt81LA"
	runTests := m.Run()
	os.Exit(runTests)
}

func Test_service_CreateChat(t *testing.T) {
	service := NewService(0)

	t.Log(service.CreateChat(token, []string{"wangmingfeng", "duoduomashineisheji"}, "", "wangmingfeng", ""))
}

func Test_service_GetChat(t *testing.T) {
	service := NewService(0)

	t.Log(service.GetChat(token, "wrp7CLCQAASEylJD032d3f-U0E5DAERA"))
}

func Test_service_SendText(t *testing.T) {
	service := NewService(0)

	t.Log(service.SendText(token,
		chatId, "测试一下1"))
}

func Test_service_SendImage(t *testing.T) {
	service := NewService(0)

	t.Log(service.SendImage(token,
		chatId, "3vHXD0a1oJCR9vlL5Ah07h6PMw-70GIapdEKrQrtrI-rpB929jbp7UekQAN1heDV-", WithOptionSafe(1)))
}

func Test_service_SendVideo(t *testing.T) {
	service := NewService(0)

	t.Log(service.SendVideo(token,
		chatId, "3brePNyB5mHgrlVdVOs_s6oMIPSxYG18moBPxAdCK_K7yyWxCF3-wJi5LSrGseTGH", "3d2", "测试视频", WithOptionSafe(1)))
}

func Test_service_SendVoice(t *testing.T) {
	service := NewService(0)

	t.Log(service.SendVoice(token,
		chatId, "3fP2dp0GlgIpx1Zeo7Q-JbqGU8m3j3qqho9x3pVLSAqAXC_0Y5FROeGYgSOjSQ8C3"))
}

func Test_service_SendFile(t *testing.T) {
	service := NewService(0)

	t.Log(service.SendFile(token,
		chatId, "3NlMFz8iqrqgp1b8m7b-EV5UYyxjYTY12IZis0qcmV8Y", WithOptionSafe(1)))
}

func Test_service_SendTextCard(t *testing.T) {
	service := NewService(0)

	t.Log(service.SendTextCard(token,
		chatId, "标题", "这是描述", "https://www.huuinn.com", "点击"))
}

func Test_service_SendNews(t *testing.T) {
	service := NewService(0)

	t.Log(service.SendNews(token,
		chatId, "标题", "这是描述", "https://www.huuinn.com", "https://www.huuinn.com/wp-content/uploads/2018/04/ubuntu_login_bg.png"))
}

func Test_service_SendMpNews(t *testing.T) {
	service := NewService(0)

	t.Log(service.SendMpNews(token,
		chatId, "标题", "3vHXD0a1oJCR9vlL5Ah07h6PMw-70GIapdEKrQrtrI-rpB929jbp7UekQAN1heDV-", "作者", "https://www.huuinn.com", "测试内容", "测试描述"))
}

func Test_service_SendMarkdown(t *testing.T) {
	service := NewService(0)

	context := `您的会议室已经预定，稍后会同步到邮箱
                >**事项详情** 
                >事　项：<font color="info">开会</font> 
                >组织者：@miglioguan 
                >参与者：@miglioguan、@kunliu、@jamdeezhou、@kanexiong、@kisonwang 
                > 
                >会议室：<font color="info">广州TIT 1楼 301</font> 
                >日　期：<font color="warning">2018年5月18日</font> 
                >时　间：<font color="comment">上午9:00-11:00</font> 
                > 
                >请准时参加会议。 
                > 
				>如需修改会议信息，请点击：[修改会议信息](https://work.weixin.qq.com)`

	t.Log(service.SendMarkDown(token,
		chatId, context))
}
