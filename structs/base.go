package structs

type Completable interface {
	Complete() string
}

type Message struct {
	Msgtype string `json:"msgtype"`
}

func (m *Message) New() {
	m.Msgtype = "text"
}

func (m *Message) Complete() string {
	return ""
}

// -- 通用返回 --
type WxResponse struct {
	Errcode int    `json:"errcode"`
	Errmsg  string `json:"errmsg"`
}
