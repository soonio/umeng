package umeng

import (
	"bytes"
	"crypto/md5"
	"encoding/hex"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
)

const (
	PathUpload = "/upload"
	PathStatus = "/api/status"
	PathStat   = "/api/task/stat"
	PathCancel = "/api/cancel"
	PathSend   = "/api/send"
)

// Send 消息发送
func Send(key string, data Body) (result *Result, err error) {
	return Post(PathSend, key, data)
}

// Upload 文件上传
func Upload(key string, data Body) (result *Result, err error) {
	return Post(PathUpload, key, data)
}

// Status 消息状态查询
func Status(key string, data Body) (result *Result, err error) {
	return Post(PathStatus, key, data)
}

// Stat 任务送达数据查询
func Stat(key string, data Body) (result *Result, err error) {
	return Post(PathStat, key, data)
}

// Cancel 消息撤销
func Cancel(key string, data Body) (result *Result, err error) {
	return Post(PathCancel, key, data)
}

func Post(path, key string, data Body) (result *Result, err error) {
	bs, err := json.Marshal(data)
	if err != nil {
		return nil, err
	}

	resp, err := http.Post(Url(path, key, bs), "application/json", bytes.NewBuffer(bs))
	if err != nil {
		return
	}
	defer resp.Body.Close()

	bs, err = ioutil.ReadAll(resp.Body)
	if err == nil {
		err = json.Unmarshal(bs, &result)
	}
	return
}

// Url 构建请求URL
func Url(path, secret string, bodyBytes []byte) string {
	h := md5.New()
	h.Write([]byte(fmt.Sprintf("POSThttps://msgapi.umeng.com%s%s%s", path, string(bodyBytes), secret)))
	sign := hex.EncodeToString(h.Sum(nil))
	return fmt.Sprintf("https://msgapi.umeng.com%s?sign=%s", path, sign)
}
