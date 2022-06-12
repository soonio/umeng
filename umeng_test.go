package umeng

import (
	"fmt"
	"testing"
	"time"
)

const (
	IosKey        = "IosKey"
	IosSecret     = "IosSecret"
	AndroidKey    = "AndroidKey"
	AndroidSecret = "AndroidSecret"
	Phone         = "17611110000"
	Test          = false
)

var exts = map[string]interface{}{
	"type":  "news",
	"event": "update",
	"id":    10000021,
}

func TestIOSSend(t *testing.T) {

	var payload = make(Body)
	payload.
		Set("appkey", IosKey).
		Set("timestamp", time.Now().Unix()).
		Set("type", "customizedcast").
		Set("alias_type", "phone").
		Set("alias", Phone).
		Child("payload", func(payload Body) {
			payload.Child("aps", func(payload Body) {
				payload.Set("alert", "This is tips message.").Set("sound", "chime")
			})
			for k, v := range exts {
				payload.Set(k, v)
			}
		}).
		Set("production_mode", Test).
		Set("description", "message-push")

	result, err := Send(IosSecret, payload)

	fmt.Println(result)

	if err != nil {
		t.Fail()
		return
	}

	if result.Ok() {
		return
	}
	t.Fail()
}

func TestAndroidSend(t *testing.T) {
	var payload = make(Body)
	payload.
		Set("appkey", AndroidKey).
		Set("timestamp", time.Now().Unix()).
		Set("type", "customizedcast").
		Set("alias_type", "phone").
		Set("alias", Phone).
		Child("payload", func(payload Body) {
			payload.
				Set("display_type", "notification").
				Child("body", func(payload Body) {
					payload.
						Set("title", "title").
						Set("text", "This is tips message.").
						Set("play_vibrate", true).
						Set("play_lights", true).
						Set("play_sound", true).
						Set("after_open", "go_app").
						Set("custom", exts)
				}).
				Set("extra", exts)
		}).
		Set("production_mode", Test).
		Set("description", "message-push")

	result, err := Send(AndroidSecret, payload)
	if err != nil {
		t.Fail()
		return
	}

	if result.Ok() {
		return
	}
	t.Fail()
}
