package robot

import "github.com/iristack/wxsender/structs"

type TemplateCard struct {
	structs.Message
	TemplateCard _TemplateCard `json:"template_card"`
}

type _TemplateCard struct {
	// 模版卡片的模版类型，文本通知模版卡片的类型为text_notice, 图文展示模版卡片的类型为news_notice
	CardType string `json:"card_type" validate:"required"`

	// 卡片来源样式信息，不需要来源样式可不填写
	Source Source `json:"source,omitempty"`

	// 模版卡片的主要内容，包括一级标题和标题辅助信息
	MainTitle MainTitle `json:"main_title,omitempty" validate:"required"`

	// 关键数据样式
	EmphasisContent EmphasisContent `json:"emphasis_content,omitempty"`

	// 引用文献样式，建议不与关键数据共用
	QuoteArea QuoteArea `json:"quote_area,omitempty"`

	// 二级普通文本，建议不超过112个字。模版卡片主要内容的一级标题main_title.title和二级普通文本sub_title_text必须有一项填写
	SubTitleText string `json:"sub_title_text"`

	// 二级标题+文本列表，该字段可为空数组，但有数据的话需确认对应字段是否必填，列表长度不超过6
	HorizontalContentList []HorizontalContent `json:"horizontal_content_list,omitempty"`

	// 跳转指引样式的列表，该字段可为空数组，但有数据的话需确认对应字段是否必填，列表长度不超过3
	JumpList []Jump `json:"jump_list,omitempty"`

	// 整体卡片的点击跳转事件，text_notice模版卡片中该字段为必填项
	CardAction CardAction `json:"card_action,omitempty"`

	// ---以下是文图展示的特有字段--

	// 图片样式
	CardImage CardImage `json:"card_image,omitempty"`

	// 左图右文样式
	ImageTextArea ImageTextArea `json:"image_text_area,omitempty"`

	// 卡片二级垂直内容，该字段可为空数组，但有数据的话需确认对应字段是否必填，列表长度不超过4
	VerticalContentList []VerticalContent `json:"vertical_content_list,omitempty"`
}

type Source struct {
	IconUrl   string `json:"icon_url,omitempty"`   // 来源图片的url
	Desc      string `json:"desc,omitempty"`       // 来源图片的描述，建议不超过13个字
	DescColor int    `json:"desc_color,omitempty"` // 来源文字的颜色，目前支持：0(默认) 灰色，1 黑色，2 红色，3 绿色
}

type MainTitle struct {
	// 一级标题，建议不超过26个字。模版卡片主要内容的一级标题main_title.title和二级普通文本sub_title_text必须有一项填写
	Title string `json:"title,omitempty"`
	// 标题辅助信息，建议不超过30个字
	Desc string `json:"desc,omitempty"`
}

type EmphasisContent struct {
	Title string `json:"title,omitempty"` // 关键数据样式的数据内容，建议不超过10个字
	Desc  string `json:"desc,omitempty"`  // 关键数据样式的数据描述内容，建议不超过15个字
}

type QuoteArea struct {
	Type      int    `json:"type,omitempty"`       // 引用文献样式区域点击事件，0或不填代表没有点击事件，1 代表跳转url，2 代表跳转小程序
	QuoteText string `json:"quote_text,omitempty"` // 引用文献样式的引用文案
	Url       string `json:"url,omitempty"`        // 点击跳转的url，quote_area.type是1时必填
	AppId     string `json:"appid,omitempty"`      // 点击跳转的小程序的appid，quote_area.type是2时必填
	PagePath  string `json:"pagepath,omitempty"`   // 点击跳转的小程序的pagepath，quote_area.type是2时选填
	Title     string `json:"title,omitempty"`      // 引用文献样式的标题
}

type HorizontalContent struct {
	KeyName string `json:"keyname" validate:"required"` // 二级标题，建议不超过5个字
	Value   string `json:"value,omitempty"`             // 二级文本，如果horizontal_content_list.type是2，该字段代表文件名称（要包含文件类型），建议不超过26个字
	Type    int    `json:"type,omitempty"`              // 模版卡片的二级标题信息内容支持的类型，1是url，2是文件附件，3 代表点击跳转成员详情
	Url     string `json:"url,omitempty"`               // 链接跳转的url，horizontal_content_list.type是1时必填
	MediaId string `json:"media_id,omitempty"`          // 附件的media_id，horizontal_content_list.type是2时必填
}

type VerticalContent struct {
	Title string `json:"title" validate:"required"` // 卡片二级标题，建议不超过26个字
	Desc  string `json:"desc,omitempty"`            // 二级普通文本，建议不超过112个字
}

type Jump struct {
	Title    string `json:"title" validate:"required"` // 跳转链接样式的文案内容，建议不超过13个字
	Url      string `json:"url,omitempty"`             // 跳转链接的url，jump_list.type是1时必填
	AppId    string `json:"appid,omitempty"`           // 跳转链接的小程序的appid，jump_list.type是2时必填
	PagePath string `json:"pagepath,omitempty"`        // 跳转链接的小程序的pagepath，jump_list.type是2时选填
	Type     int    `json:"type,omitempty"`            // 跳转链接类型，0或不填代表不是链接，1 代表跳转url，2 代表跳转小程序
}

type CardAction struct {
	Type     int    `json:"type" validate:"required"` // 卡片跳转类型，1 代表跳转url，2 代表打开小程序。text_notice模版卡片中该字段取值范围为[1,2]
	Url      string `json:"url,omitempty"`            // 跳转事件的url，card_action.type是1时必填
	AppId    string `json:"appid,omitempty"`          // 跳转事件的小程序的appid，card_action.type是2时必填
	PagePath string `json:"pagepath,omitempty"`       // 跳转事件的小程序的pagepath，card_action.type是2时选填
}

type CardImage struct {
	PicUrl      string  `json:"pic_url" validate:"required"` // 图片的url
	AspectRatio float64 `json:"aspect_ratio,omitempty"`      // 图片的宽高比，宽高比要小于2.25，大于1.3，不填该参数默认1.3
}

type ImageTextArea struct {
	Type     int    `json:"type"`      // 左图右文样式区域点击事件，0或不填代表没有点击事件，1 代表跳转url，2 代表跳转小程序
	Url      string `json:"url"`       // 点击跳转的url，image_text_area.type是1时必填
	AppId    string `json:"appid"`     // 点击跳转的小程序的appid，必须是与当前应用关联的小程序，image_text_area.type是2时必填
	Pagepath string `json:"pagepath"`  // 点击跳转的小程序的pagepath，image_text_area.type是2时选填
	Title    string `json:"title"`     // 左图右文样式的标题
	Desc     string `json:"desc"`      // 左图右文样式的描述
	ImageUrl string `json:"image_url"` // 左图右文样式的图片url
}

func (m *TemplateCard) New() {
	m.Msgtype = "template_card"
	m.TemplateCard = _TemplateCard{}
	m.TemplateCard.CardType = "text_notice"
}

func (m *TemplateCard) SetTextNotice() {
	m.TemplateCard.CardType = "text_notice"
}

func (m *TemplateCard) SetNewsNotice() {
	m.TemplateCard.CardType = "news_notice"
}

func (m *TemplateCard) SetSource(source Source) {
	m.TemplateCard.Source = source
}

func (m *TemplateCard) SetMainTitle(mainTitle MainTitle) {
	m.TemplateCard.MainTitle = mainTitle
}

func (m *TemplateCard) SetEmphasisContent(emphasisContent EmphasisContent) {
	m.TemplateCard.EmphasisContent = emphasisContent
}

func (m *TemplateCard) SetQuoteArea(quoteArea QuoteArea) {
	m.TemplateCard.QuoteArea = quoteArea
}

func (m *TemplateCard) SetHorizontalContentList(horizontalContentList []HorizontalContent) {
	m.TemplateCard.HorizontalContentList = horizontalContentList
}

func (m *TemplateCard) AddHorizontalContent(horizontalContent HorizontalContent) {
	if m.TemplateCard.HorizontalContentList == nil {
		m.TemplateCard.HorizontalContentList = []HorizontalContent{}
	}
	m.TemplateCard.HorizontalContentList = append(m.TemplateCard.HorizontalContentList, horizontalContent)
}

func (m *TemplateCard) SetSubTitleText(subTitleText string) {
	m.TemplateCard.SubTitleText = subTitleText
}

func (m *TemplateCard) SetJumpList(jumpList []Jump) {
	m.TemplateCard.JumpList = jumpList
}

func (m *TemplateCard) AddJump(jump Jump) {
	if m.TemplateCard.JumpList == nil {
		m.TemplateCard.JumpList = []Jump{}
	}
	m.TemplateCard.JumpList = append(m.TemplateCard.JumpList, jump)
}

func (m *TemplateCard) SetCardAction(cardAction CardAction) {
	m.TemplateCard.CardAction = cardAction
}

func (m *TemplateCard) SetCardImage(cardImage CardImage) {
	if m.TemplateCard.CardType != "news_notice" {
		return
	}
	m.TemplateCard.CardImage = cardImage
}

func (m *TemplateCard) SetImageTextArea(imageTextArea ImageTextArea) {
	if m.TemplateCard.CardType != "news_notice" {
		return
	}
	m.TemplateCard.ImageTextArea = imageTextArea
}
func (m *TemplateCard) SetVerticalContentList(verticalContentList []VerticalContent) {
	if m.TemplateCard.CardType != "news_notice" {
		return
	}
	m.TemplateCard.VerticalContentList = verticalContentList
}

func (m *TemplateCard) AddVerticalContent(verticalContent VerticalContent) {
	if m.TemplateCard.CardType != "news_notice" {
		return
	}
	if m.TemplateCard.VerticalContentList == nil {
		m.TemplateCard.VerticalContentList = []VerticalContent{}
	}
	m.TemplateCard.VerticalContentList = append(m.TemplateCard.VerticalContentList, verticalContent)
}

func (m *TemplateCard) Complete() string {
	return structs.ToJson(m)
}
