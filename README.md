[![Go Report Card](https://goreportcard.com/badge/clem109/go-wechat-work)](https://goreportcard.com/report/clem109/go-wechat-work)
[![License](https://img.shields.io/badge/License-MIT-blue.svg)](https://github.com/clem109/go-wechat-work/blob/master/LICENSE)

# WeChat Work

Use the WeChat work API in your Golang apps to send a Text Card (文本卡片消息)

API reference for [WeChat Work](https://work.weixin.qq.com).

```bash
go get github.com/clem109/go-wechat-work
```

## Usage

```go
package main

import ("github.com/clem109/go-wechat-work")

func main() {
	wechatMessage := wechat.Plugin{
    // Fill this in with your own details, whether that is via
    // env var or a YAML file.
    // Method, Safe, ContentType, Debug and SkipVerify best to leave
    // how it is.
		Config: wechat.Config{
			Method:      "POST",
			CorpID:      "corpid",
			CorpSecret:  "corp-secret",
			Agentid:     123242,
			MsgType:     "msgtype",
			MsgURL:      "msgurl",
			BtnTxt:      "btntxt",
			ToUser:      "touser",
			ToParty:     "toparty",
			ToTag:       "tostring",
			Title:       "Title",
			Description: "Description",
			Safe:        0,
			ContentType: "application/json",
			Debug:       false,
			SkipVerify:  true,
		},
	}
  err = wechatMessage.Exec()
  if err != nil {
    // deal with the error
  }
}
```
