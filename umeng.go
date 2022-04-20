package umeng

import "time"

type Passport struct {
	Key    string `mapstructure:"key" json:"key" yaml:"key"`
	Secret string `mapstructure:"secret" json:"secret" yaml:"secret"`
}

type Platform int32

const (
	Android Platform = 1
	IOS     Platform = 2
)

// 文档地址 https://developer.umeng.com/docs/67966/detail/68343#h1-u6D88u606Fu53D1u90014

type Client struct {
	Config  map[Platform]*Passport // 配置
	Product bool                   // 是否真实发送，仅支持对简单推送方法有效
}

type Request struct {
	p        Platform
	Passport *Passport
	data     interface{}
	path     string
}

func (c *Client) Request(p Platform, payload interface{}) *Request {
	var r = &Request{p: p, Passport: c.Config[p], data: payload}

	if v, ok := payload.(data); ok {
		r.path = v.Path()
		v.Set(c.Config[p].Key, time.Now().Unix())
	} else {
		panic("Payload format error.")
	}
	return r
}

func (r *Request) Send() *Result {
	return Send(r.path, r.Passport.Secret, r.data)
}

// IOSSimpleMessageByPhone 发送简单的消息给IOS客户端
func (c *Client) IOSSimpleMessageByPhone(phone, message string, extras ...map[string]interface{}) *Result {
	return c.Request(IOS, &PushData{
		Type:      "customizedcast",
		AliasType: "phone",
		Alias:     phone,
		Payload: IOSPayload(&IOSAps{
			Alert: message,
			Sound: "chime",
		}, extras...),
		ProductionMode: c.Product,
		Description:    "message-push",
	}).Send()
}

func (c *Client) AndroidSimpleMessageByPhone(phone, title, message string, extras ...map[string]interface{}) *Result {
	return c.Request(Android, &PushData{
		Type:      "customizedcast",
		AliasType: "phone",
		Alias:     phone,
		Payload: &AndroidPayload{
			DisplayType: "notification",
			Body: &AndroidBody{
				Title:       title,
				Text:        message,
				PlayVibrate: true,
				PlayLights:  true,
				PlaySound:   true,
				AfterOpen:   "go_app",
				Custom:      extras[0],
			},
			Extra: extras[0],
		},
		ProductionMode: c.Product,
		Description:    "message-push",
	}).Send()
}
