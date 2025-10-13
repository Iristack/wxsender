package robot

import "wxsender/structs"

type Voice struct {
	structs.Message
	Voice struct {
		MediaId string `json:"media_id"` // 语音文件id，通过下文的文件上传接口获取
	} `json:"voice"`
}

func (m *Voice) New() {
	m.Msgtype = "voice"
	m.Voice = struct {
		MediaId string `json:"media_id"`
	}{
		MediaId: "",
	}
}

func (m *Voice) SetMediaId(mediaId string) {
	m.Voice.MediaId = mediaId
}

func (m *Voice) Complete() string {
	return structs.ToJson(m)
}
