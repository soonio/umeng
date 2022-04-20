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

type data interface {
	Path() string
	Set(key string, timestamp int64)
}

const (
	Host       = "https://msgapi.umeng.com"
	PathUpload = "/upload"
	PathStatus = "/api/status"
	PathStat   = "/api/task/stat"
	PathCancel = "/api/cancel"
	PathSend   = "/api/send"
)

// 结构参考 https://developer.umeng.com/docs/67966/detail/68343#h1-u6D88u606Fu53D1u90014

type Base struct {
	AppKey    string `json:"appkey"`    // 必填，应用唯一标识
	TimeStamp int64  `json:"timestamp"` // 必填，时间戳，10位或者13位均可，时间戳有效期为10分钟
}

func (b *Base) Set(key string, timestamp int64) {
	b.AppKey = key
	b.TimeStamp = timestamp
}

type Policy struct {
	StartTime  string `json:"start_time,omitempty"`
	ExpireTime string `json:"expire_time,omitempty"`
	MaxSendNum int    `json:"max_send_num,omitempty"`
	OutBizNo   string `json:"out_biz_no,omitempty"`
}

// PushData 推送接口数据
type PushData struct {
	Base
	Type           string      `json:"type,omitempty"`            // 必填，详情参考文档
	DeviceTokens   string      `json:"device_tokens,omitempty"`   // 选填，详情参考文档
	AliasType      string      `json:"alias_type,omitempty"`      // 选填，详情参考文档
	Alias          string      `json:"alias,omitempty"`           // 选填，详情参考文档
	FileId         string      `json:"file_id,omitempty"`         // 选填，详情参考文档
	Filter         string      `json:"filter,omitempty"`          // 选填，详情参考文档
	Payload        interface{} `json:"payload,omitempty"`         // 必填，消息内容
	Policy         *Policy     `json:"policy,omitempty"`          // 可选，发送策略
	ProductionMode bool        `json:"production_mode,omitempty"` // 可选，true正式模式，false测试模式。默认为true
	Description    string      `json:"description,omitempty"`     // 可选，发送消息描述，建议填写
}

func (p *PushData) Path() string {
	return PathSend
}

// BaseWithTask 消息状态查询、任务送达数据查询、消息撤回查询
type BaseWithTask struct {
	Base
	TaskId string `json:"task_id"`
}

// StateData 消息状态查询接口
type StateData BaseWithTask

func (p *StateData) Path() string {
	return PathStatus
}

// StatData 任务送达数据查询接口
type StatData BaseWithTask

func (p *StatData) Path() string {
	return PathStat
}

// CancelData 取消接口数据
type CancelData BaseWithTask

func (p *CancelData) Path() string {
	return PathCancel
}

// UploadData 文件上传
type UploadData struct {
	Base
	FileContent string `json:"content"`
}

func (p *UploadData) Path() string {
	return PathUpload
}

// Result 结果对象
type Result struct {
	Code string `json:"ret,omitempty"`
	Data map[string]string
}

// Send 对数据进行编码后发送
func Send(path, key string, data interface{}) *Result {
	bs, _ := json.Marshal(data)
	return POST(Url(path, key, bs), bs)
}

// Url 构建请求URL
func Url(path, secret string, bodyBytes []byte) string {
	h := md5.New()
	h.Write([]byte(fmt.Sprintf("POST%s%s%s%s", Host, path, string(bodyBytes), secret)))
	sign := hex.EncodeToString(h.Sum(nil))
	return fmt.Sprintf("%s%s?sign=%s", Host, path, sign)
}

// POST 发送http请求
func POST(url string, bodyBytes []byte) (result *Result) {
	req, err := http.NewRequest("POST", url, bytes.NewBuffer(bodyBytes))
	req.Header.Set("Content-Type", "application/json")
	client := &http.Client{}
	resp, err := client.Do(req)
	if err != nil {
		panic(err)
	}
	defer func() {
		_ = resp.Body.Close()
	}()

	bs, _ := ioutil.ReadAll(resp.Body)
	_ = json.Unmarshal(bs, &result)
	return
}
