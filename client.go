package main

import (
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"strings"
	"wxsender/structs"
	"wxsender/types"
)

const (
	DefaultClientAppEndpoint = "https://qyapi.weixin.qq.com/cgi-bin/message/send?access_token="
	DefaultClientBotEndpoint = "https://qyapi.weixin.qq.com/cgi-bin/webhook/send?key="
)

type Client struct {
	Type        types.MessageType // 发送的消息类型
	AccessToken string            // 若选择APP发送, 请填写APP的access_token
	Key         string            // 若选择Bot发送, 请填写Bot的key
	UrlPrefix   string            // 请填写企业微信的url前缀, Bot默认为https://qyapi.weixin.qq.com/cgi-bin/webhook/send?key=, App则默认为https://qyapi.weixin.qq.com/cgi-bin/message/send?access_token=

}

func (c *Client) Create() {
	if c.Type == types.ROBOT {
		c.UrlPrefix = DefaultClientBotEndpoint + c.Key
	} else {
		c.UrlPrefix = DefaultClientAppEndpoint + c.AccessToken
	}

}

func (c *Client) Submit(m structs.Completable) (structs.WxResponse, error) {
	content := m.Complete()

	resp, err := http.Post(c.UrlPrefix, "application/json", strings.NewReader(content))
	if err != nil {
		return structs.WxResponse{}, err
	}
	defer resp.Body.Close()

	body, err := io.ReadAll(resp.Body)
	if err != nil {
		return structs.WxResponse{}, err
	}

	var wxResp structs.WxResponse
	err = json.Unmarshal(body, &wxResp)
	if err != nil {
		return structs.WxResponse{}, fmt.Errorf("解析响应失败: %v, 原始返回: %s", err, body)
	}

	return wxResp, nil
}
