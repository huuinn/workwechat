package upload

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

func Test_service_UploadImage(t *testing.T) {
	service := NewService()
	t.Log(service.UploadImage(token,
		"../../../test/ubuntu_login_bg.png"))
}

func Test_service_UploadVideo(t *testing.T) {
	service := NewService()
	t.Log(service.UploadVideo(token,
		"../../../test/family_00.mp4"))
}

func Test_service_UploadVoice(t *testing.T) {
	service := NewService()
	t.Log(service.UploadVoice(token,
		"../../../test/gs-16b-1c-8000hz.amr"))
}

func Test_service_UploadFile(t *testing.T) {
	service := NewService()
	t.Log(service.UploadFile(token,
		"../../../test/test.md"))
}
