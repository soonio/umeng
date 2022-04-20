# Umeng推送
> 简单整合一下umeng的推送

## 使用

```bash
go get -u github.com/soonio/umeng
```

## 实现参考

https://developer.umeng.com/docs/67966/detail/68343#h1-u6D88u606Fu53D1u90014

## 用法示例

```go
package main

import "fmt"
import "github.com/soonio/umeng"

func main() {

	client := &umeng.Client{
		Config: map[umeng.Platform]*umeng.Passport{
			umeng.Android: {
				Key:    "111",
				Secret: "222",
			},
			umeng.IOS: {
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
			Alert: "你收到了这条消息，快看看吧",
		}),
		ProductionMode: true,
		Description:    "message-push",
	}).Send()

	fmt.Printf("%+v\n", res)

	res = client.IOSSimpleMessageByPhone("17600001111", "你收到了这条消息，快看看吧", map[string]interface{}{
		"type": 10,
		"data": map[string]int{
			"id": 20,
		},
	})
	fmt.Printf("%+v\n", res)

	res = client.AndroidSimpleMessageByPhone("17600001111", "你收到了这条消息，快看看吧", "带结构体的消息", map[string]interface{}{
		"type": 10,
		"data": map[string]int{
			"id": 20,
		},
	})

	fmt.Printf("%+v\n", res)
}
```