// Please check the github Readme for an example of how to use this package. This package will help make
// sending notifications using WeChat Work 企业微信 easier. Make sure to refer to the api docs (chinese)
// to be able to know how to correctly configer each struct before sending the notification.
package wechat

import (
	"bytes"
	"crypto/tls"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
)

const (
	respFormat      = "Webhook %d\n  URL: %s\n  RESPONSE STATUS: %s\n  RESPONSE BODY: %s\n"
	debugRespFormat = "Webhook %d\n  URL: %s\n  METHOD: %s\n  HEADERS: %s\n  REQUEST BODY: %s\n  RESPONSE STATUS: %s\n  RESPONSE BODY: %s\n"
)

// Config is the minimum information you need to send a notification, CorpId and CorpSecret are used to
// get the access token required before sending a notification.
type (
	Config struct {
		CorpID       string
		CorpSecret   string
		SkipVerify   bool
		Debug        bool
		Notification Notification
	}

	// Notification defines information about who to send the notification to, the type of message and the agentid of your
	// wechat work application. The other structs are to define which notification to send. If you define both a TextCard and a
	// Text but set the message type to "text" then only a text will be sent.
	Notification struct {
		ToUser   string   `json:"touser"`
		ToParty  string   `json:"toparty"`
		ToTag    string   `json:"totag"`
		MsgType  string   `json:"msgtype"`
		Agentid  int      `json:"agentid"`
		TextCard TextCard `json:"textcard"`
		Text     Text     `json:"text"`
		Image    Image    `json:"image"`
		Voice    Voice    `json:"voice"`
		Video    Video    `json:"video"`
		File     File     `json:"file"`
		News     News     `json:"news"`
		MpNews   MpNews   `json:"mpnews"`
		Safe     int      `json:"safe"`
	}

	// TextCard is a text card that will be sent as a notification.
	TextCard struct {
		Title       string `json:"title"`
		Description string `json:"description"`
		MsgURL      string `json:"url"`
		BtnTxt      string `json:"btntext"`
	}

	// Text is a text message
	Text struct {
		Content string `json:"content"`
	}

	// Image is  an image to send
	Image struct {
		MediaId string `json:"media_id"`
	}

	// Voice  sends a voice message
	Voice struct {
		MediaId string `json:"media_id"`
	}

	// Video sends a video with a title and description
	Video struct {
		MediaId     string `json:"media_id"`
		Title       string `json:"title"`
		Description string `json:"description"`
	}

	// File sends a file that needs to be uploaded to wechat work
	File struct {
		MediaId string `json:"media_id"`
	}

	// News is an array of articles to be sent
	News struct {
		Articles []Article `json:"articles"`
	}

	// Article defines an article to be sent
	Article struct {
		Title       string `json:"title"`
		Description string `json:"description"`
		URL         string `json:"url"`
		PicURL      string `json:"picurl"`
		BtnTxt      string `json:"btntxt"`
	}

	MpNews struct {
		MpNewsArticles []MpNewsArticle `json:"articles"`
	}

	MpNewsArticle struct {
		Title              string `json:"title"`
		Thumb_Media_ID     string `json:"thumb_media_id"`
		Author             string `json:"author"`
		Content_Source_Url string `json:"content_source_url"`
		Content            string `json:"content"`
		Digest             string `json:"digest"`
	}

	Response struct {
		Errcode     int    `json:"errcode"`
		Errmsg      string `json:"errmsg"`
		AccessToken string `json:"access_token"`
		ExpiresIn   int    `json:"expires_in"`
	}

	WeChatWork struct {
		Config   Config
		Response Response
	}
)

// gets the access token to be used in subsequent requests
func getAccessToken(body []byte) (*Response, error) {
	var s = new(Response)
	err := json.Unmarshal(body, &s)
	if err != nil {
		fmt.Println("Unable to get access token:", err)
	}
	return s, err
}

// Exec on a completed WeChatWork struct sends a message
func (p WeChatWork) Exec() error {

	var b []byte

	// construct URL to get access token
	accessURL := "https://qyapi.weixin.qq.com/cgi-bin/gettoken?corpid=" + p.Config.CorpID + "&corpsecret=" + p.Config.CorpSecret

	req, err := http.NewRequest("GET", accessURL, bytes.NewBuffer(b))
	var client = http.DefaultClient
	if p.Config.SkipVerify {
		client = &http.Client{
			Transport: &http.Transport{
				TLSClientConfig: &tls.Config{InsecureSkipVerify: true},
			},
		}
	}
	resp, err := client.Do(req)

	if err != nil {
		fmt.Printf("Error: Failed to execute the HTTP request. %s\n", err)
		return err
	}

	defer resp.Body.Close()

	fmt.Println("response Status:", resp.Status)
	fmt.Println("response Headers:", resp.Header)
	body, _ := ioutil.ReadAll(resp.Body)
	fmt.Println("response Body:", string(body))

	s, err := getAccessToken(body)

	// POST Request to WeChat work
	// textCard := p.Config.TextCard{p.Config.Title, p.Config.Description, p.Config.MsgURL, p.Config.BtnTxt}

	b, _ = json.Marshal(p.Config.Notification) // []byte(data)

	// POST URL
	url := "https://qyapi.weixin.qq.com/cgi-bin/message/send?access_token=" + s.AccessToken

	request, err := http.NewRequest("POST", url, bytes.NewBuffer(b))
	request.Header.Set("Content-Type", "application/json")

	client = http.DefaultClient
	client = &http.Client{
		Transport: &http.Transport{
			TLSClientConfig: &tls.Config{InsecureSkipVerify: true},
		},
	}
	response, err := client.Do(request)

	if err != nil {
		fmt.Printf("Error: Failed to execute the HTTP request. %s\n", err)
		return err
	}
	defer response.Body.Close()

	fmt.Println("response Status:", response.Status)
	fmt.Println("response Headers:", response.Header)
	responseBody, _ := ioutil.ReadAll(response.Body)
	fmt.Println("response Body:", string(responseBody))

	if p.Config.Debug || resp.StatusCode >= http.StatusBadRequest {
		body, err := ioutil.ReadAll(resp.Body)

		if err != nil {
			fmt.Printf("Error: Failed to read the HTTP response body. %s\n", err)
		}

		if p.Config.Debug {
			fmt.Printf(
				debugRespFormat,
				req.URL,
				req.Method,
				req.Header,
				string(b),
				resp.Status,
				string(body),
			)
		} else {
			fmt.Printf(
				respFormat,
				req.URL,
				resp.Status,
				string(body),
			)
		}
	}
	return nil
}
