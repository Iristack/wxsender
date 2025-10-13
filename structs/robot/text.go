package robot

import "wxsender/structs"

// Text 群聊天机器人消息结构 - 文本类型
type Text struct {
	structs.Message
	Text struct {
		Content             string   `json:"content"`                         // 文本内容，最长不超过2048个字节，必须是utf8编码
		MentionedList       []string `json:"mentioned_list,omitempty"`        // userid的列表，提醒群中的指定成员(@某个成员)，@all表示提醒所有人，如果开发者获取不到userid，可以使用mentioned_mobile_list
		MentionedMobileList []string `json:"mentioned_mobile_list,omitempty"` // 手机号列表，提醒手机号对应的群成员(@某个成员)，@all表示提醒所有人
	} `json:"text"`
}

func (m *Text) New() {
	m.Msgtype = "text"
	m.Text = struct {
		Content             string   `json:"content"`
		MentionedList       []string `json:"mentioned_list,omitempty"`
		MentionedMobileList []string `json:"mentioned_mobile_list,omitempty"`
	}{
		Content:             "",
		MentionedList:       []string{},
		MentionedMobileList: []string{},
	}
}

func (m *Text) SetContent(content string) {
	m.Text.Content = content
}

func (m *Text) SetAtAll() {
	m.Text.MentionedList = append(m.Text.MentionedList, "@all")
}

func (m *Text) SetAts(ats []string) {
	m.Text.MentionedList = ats
}

func (m *Text) AddAt(at string) {
	m.Text.MentionedList = append(m.Text.MentionedList, at)
}

func (m *Text) SetAtMobiles(mobileList []string) {
	m.Text.MentionedMobileList = mobileList
}

func (m *Text) AddAtMobile(mobile string) {
	m.Text.MentionedMobileList = append(m.Text.MentionedMobileList, mobile)
}

func (m *Text) Complete() string {
	return structs.ToJson(m)
}
