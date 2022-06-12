# Umeng推送

> 简单整合一下umeng的推送

## 使用

```bash
go get -u github.com/soonio/umeng
```

## 实现参考

https://developer.umeng.com/docs/67966/detail/68343#h1-u6D88u606Fu53D1u90014

## 接口列表

  - Send 消息发送
  - Upload 文件上传
  - Status 消息状态查询
  - Stat 任务送达数据查询
  - Cancel 消息撤销

## 测试

参考[umeng_test.go](./umeng_test.go)