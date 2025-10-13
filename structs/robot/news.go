package robot

import "github.com/iristack/wxsender/structs"

// News 群聊天机器人消息结构 - 图文类型
type News struct {
	structs.Message
	News struct {
		Articles []Article `json:"articles"` // 图文消息，一个图文消息支持1到8条图文
	} `json:"news"`
}

type Article struct {
	Title       string `json:"title"`       // 标题，不超过128个字节，超过会自动截断
	Picurl      string `json:"picurl"`      // 图文消息的图片链接，支持JPG、PNG格式，较好的效果为大图 1068*455，小图150*150。
	Description string `json:"description"` // 描述，不超过512个字节，超过会自动截断
	Url         string `json:"url"`         // 点击后跳转的链接。
}

func (m *News) New() {
	m.Msgtype = "news"
	m.News = struct {
		Articles []Article `json:"articles"`
	}{}
}

func (m *News) AddArticle(title, description, picurl, url string) {
	m.News.Articles = append(m.News.Articles, Article{
		Title:       title,
		Description: description,
		Picurl:      picurl,
		Url:         url,
	})
}

func (m *News) SetArticles(articles []Article) {
	m.News.Articles = articles
}

func (m *News) Complete() string {
	return structs.ToJson(m)
}
