package robot

import "wxsender/structs"

// Markdown 群聊天机器人消息结构 - Markdown类型
type Markdown struct {
	structs.Message
	Markdown struct {
		Content string `json:"content"` // markdown内容，最长不超过4096个字节，必须是utf8编码
	} `json:"markdown"`
}

func (m *Markdown) New() {
	m.Msgtype = "markdown"
}

func (m *Markdown) SetContent(content string) {
	m.Markdown.Content = content
}

func (m *Markdown) Complete() string {
	return structs.ToJson(m)
}
