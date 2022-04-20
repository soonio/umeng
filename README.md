# Umeng推送
> 简单整合一下umeng的推送

## 用法示例

```go
package main

import "fmt"
import "github.com/soonio/umeng"

func main() {

	client := &umeng.Client{
		Config: map[umeng.Platform]*umeng.Passport{
			umeng.Android: &umeng.Passport{
				Key:    "111",
				Secret: "222",
			},
			umeng.IOS: &umeng.Passport{
				Key:    "111",
				Secret: "222",
			},
		},
		Product: true,
	}
	res := client.Request(umeng.IOS, &umeng.PushData{
		Type:      "customizedcast",
		AliasType: "phone",
		Alias:     "17600001111",
		Payload: umeng.IOSPayload(&umeng.IOSAps{
			//Alert: "您有一条新的专家问答反馈记录，查看详情",
			Alert: "您预约的课程将在三天后（周五）上午10:00如期举行，请提前做好准备，按时参加，查看详情",
		}),
		ProductionMode: true,
		Description:    "message-push",
	}).Send()

	fmt.Printf("%+v\n", res)

	res = client.IOSSimpleMessageByPhone("17600001111", "带结构体的消息", map[string]interface{}{
		"type": 10,
		"data": map[string]int{
			"id": 20,
		},
	})
	fmt.Printf("%+v\n", res)

	res = client.AndroidSimpleMessageByPhone("17600001111", "这是一个标题", "带结构体的消息", map[string]interface{}{
		"type": 10,
		"data": map[string]int{
			"id": 20,
		},
	})

	fmt.Printf("%+v\n", res)
}
```