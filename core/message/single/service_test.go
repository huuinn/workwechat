package single

import (
	"os"
	"testing"
)

var token string

func TestMain(m *testing.M) {
	token = "1kgise5YIImoXSdL8Ozw0Sc4-9vOgXx5w5Tw_-Isr37nTRy7EpCxxSxLA4Pj0igqxtE8hY3Lze9QLfXlvfm8yGJ-L_ZVtXIsZl7FkhPUWQ_taLoEImEQ9x_HyU5RaHtbfwoykRKSM5w7Zm0FRVbdRvUsJWsYNsVSx6dZ2VuJcfQ_PHJIprZ58eK-lTu0A1ETAbj5DqAiD5WTLV8KuNAypg"
	runTests := m.Run()
	os.Exit(runTests)
}

func Test_service_SendText(t *testing.T) {
	service := NewService(1000002, 0, 0, 0, 1800)

	t.Log(service.SendText(token,
		"wangmingfeng", "", "", "测试一下"))
}

func Test_service_SendImage(t *testing.T) {
	service := NewService(1000002, 0, 0, 0, 1800)

	t.Log(service.SendImage(token,
		"wangmingfeng", "", "", "3vHXD0a1oJCR9vlL5Ah07h6PMw-70GIapdEKrQrtrI-rpB929jbp7UekQAN1heDV-"))
}

func Test_service_SendVideo(t *testing.T) {
	service := NewService(1000002, 0, 0, 0, 1800)

	t.Log(service.SendVideo(token,
		"wangmingfeng", "", "", "3brePNyB5mHgrlVdVOs_s6oMIPSxYG18moBPxAdCK_K7yyWxCF3-wJi5LSrGseTGH", "3d", "测试视频"))
}

func Test_service_SendVoice(t *testing.T) {
	service := NewService(1000002, 0, 0, 0, 1800)

	t.Log(service.SendVoice(token,
		"wangmingfeng", "", "", "3fP2dp0GlgIpx1Zeo7Q-JbqGU8m3j3qqho9x3pVLSAqAXC_0Y5FROeGYgSOjSQ8C3"))
}

func Test_service_SendFile(t *testing.T) {
	service := NewService(1000002, 0, 0, 0, 1800)

	t.Log(service.SendFile(token,
		"wangmingfeng", "", "", "3NlMFz8iqrqgp1b8m7b-EV5UYyxjYTY12IZis0qcmV8Y"))
}
