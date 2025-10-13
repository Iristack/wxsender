package robot

import "github.com/iristack/wxsender/structs"

// MarkdownV2 群聊天机器人消息结构 - MarkdownV2类型
type MarkdownV2 struct {
	structs.Message
	MarkdownV2 struct {
		// markdown_v2内容，最长不超过4096个字节，必须是utf8编码。
		// 特殊的要求:
		// 1. markdown_v2不支持字体颜色、@群成员的语法， 具体支持的语法可参考下面说明
		// 2. 消息内容在客户端 4.1.36 版本以下(安卓端为4.1.38以下) 消息表现为纯文本，建议使用最新客户端版本体验
		// link: https://developer.work.weixin.qq.com/document/path/99110#markdown-v2%E7%B1%BB%E5%9E%8B
		Content string `json:"content"`
	} `json:"markdown_v2"`
}

func (m *MarkdownV2) New() {
	m.Msgtype = "markdown_v2"
}

func (m *MarkdownV2) SetContent(content string) {
	m.MarkdownV2.Content = content
}

func (m *MarkdownV2) Complete() string {
	return structs.ToJson(m)
}
