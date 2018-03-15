[![Go Report Card](https://goreportcard.com/badge/clem109/go-wechat-work)](https://goreportcard.com/report/clem109/go-wechat-work)
[![License](https://img.shields.io/badge/License-MIT-blue.svg)](https://github.com/clem109/go-wechat-work/blob/master/LICENSE)

# WeChat Work

Use the WeChat work API in your Golang apps to send notifications

API reference for [WeChat Work](https://work.weixin.qq.com/api/doc#10167), please check this to know how to configure the settings below.

```bash
go get github.com/clem109/go-wechat-work
```

## Usage

```go
package main

import ("github.com/clem109/go-wechat-work")

func main() {
	wechat := WeChatWork{
		Config: Config{
			CorpID:     "corpid",
			CorpSecret: "somesecret",
			SkipVerify: true,
			Debug:      false,
			Notification: Notification{
				Agentid: 123456,
				MsgType: "news",
				ToUser:  "@all",
				ToParty: "@all",
				ToTag:   "@all",
				Safe:    0,
				TextCard: TextCard{
					Title:       "test",
					Description: "testing",
					MsgURL:      "someurl;",
					BtnTxt:      "p",
				},
				Text: Text{
					Content: "text message test",
				},
				News: News{
					Articles: []Article{
						Article{
							Title:       "Testing News",
							Description: "Star my repo ;)",
							URL:         "https://github.com/clem109/go-wechat-work",
							PicURL:      "https://raw.githubusercontent.com/clem109/glowing-gopher/master/gopher.jpeg",
						},
					},
				},
			},
		},
	}

	wechat.Exec()
}
```
