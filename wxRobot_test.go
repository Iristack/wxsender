package main

import (
	"fmt"
	"testing"
	"wxsender/structs/robot"
	"wxsender/types"
)

var c Client

const key = ""

func TestRobotSuite(t *testing.T) {
	if key == "" {
		t.Skip("请配置 key")
	}

	test := []struct {
		name     string
		function func(*testing.T)
	}{
		{"TestRobotClientCreate", testRobotCreateClient},
		{"TestRobotSendText", testRobotSendText},
		{"TestRobotSendMarkdown", testRobotSendMarkdown},
		{"TestRobotSendMarkdownV2", testRobotSendMarkdownV2},
		{"TestRobotSendNews", testRobotSendNews},
		{"TestRobotSendPic", testRobotSendPic},
		{"TestRobotSendFile", testRobotSendFile},
		{"TestRobotSendVoice", testRobotSendVoice},
		{"TestRobotSendCard", testRobotSendCard},
		{"TestRobotSendNewsCard", testRobotSendNewsCard},
	}

	for _, v := range test {
		t.Run(v.name, v.function)
	}
}

func testRobotCreateClient(t *testing.T) {
	c.Type = types.ROBOT
	c.Key = key
	c.Create()
}

func testRobotSendText(t *testing.T) {
	m := robot.Text{}
	m.New()
	m.Text.Content = "单元测试 Text"
	m.Text.MentionedList = []string{"@all"}
	m.Text.MentionedMobileList = []string{""}
	wxResp, err := c.Submit(&m)
	if err != nil {
		t.Error(err)
	}
	if wxResp.Errcode != 0 {
		t.Error(wxResp.Errmsg)
	}
}

func testRobotSendMarkdown(t *testing.T) {
	m := robot.Markdown{}
	m.New()

	m.SetContent("### 单元测试 Markdown \n #### 单元测试 Markdown")
	wxResp, err := c.Submit(&m)
	if err != nil {
		t.Error(err)
	}
	if wxResp.Errcode != 0 {
		t.Error(wxResp.Errmsg)
	}
}

func testRobotSendMarkdownV2(t *testing.T) {
	m := robot.MarkdownV2{}
	m.New()

	m.SetContent("### 单元测试 MarkdownV2 \n #### 单元测试 MarkdownV2")
	wxResp, err := c.Submit(&m)
	if err != nil {
		t.Error(err)
	}
	if wxResp.Errcode != 0 {
		t.Error(wxResp.Errmsg)
	}
}

func testRobotSendNews(t *testing.T) {
	m := robot.News{}
	m.New()
	m.AddArticle("单元测试 News", "单元测试 News", "https://res.mail.qq.com/node/ww/wwopenmng/images/independent/doc/test_pic_msg1.png", "https://www.baidu.com")

	t.Log(m.Complete())

	wxResp, err := c.Submit(&m)
	if err != nil {
		t.Error(err)
	}
	if wxResp.Errcode != 0 {
		t.Error(wxResp.Errmsg)
	}
}

func testRobotSendPic(t *testing.T) {
	picPath := "./resources/pic.png"

	m := robot.Image{}
	m.New()

	err := m.SetPicture(picPath)
	if err != nil {
		t.Error(err)
	}

	wxResp, err := c.Submit(&m)
	if err != nil {
		t.Error(err)
	}

	if wxResp.Errcode != 0 {
		t.Error(wxResp.Errmsg)
	}
}

func testRobotSendFile(t *testing.T) {
	fh := FileHelper{
		FilePath: "./resources/file",
		Key:      key,
		Type:     File,
	}

	if fh.Upload().Err != nil {
		t.Error(fh.Err)
	}

	if !fh.IsOk() {
		t.Error(fmt.Sprintf("上传失败: %s", fh.Err))
	}

	m := robot.File{}
	m.New()
	m.SetMediaId(fh.Response.MediaId)

	wxResp, err := c.Submit(&m)
	if err != nil {
		t.Error(err)
	}

	if wxResp.Errcode != 0 {
		t.Error(wxResp.Errmsg)
	}
}

func testRobotSendVoice(t *testing.T) {
	fh := FileHelper{
		FilePath: "./resources/tts.amr",
		Key:      key,
		Type:     Voice,
	}

	if fh.Upload().Err != nil {
		t.Error(fh.Err)
	}

	if !fh.IsOk() {
		t.Error(fmt.Sprintf("上传失败: %s", fh.Response.Errmsg))
	}

	m := robot.Voice{}
	m.New()
	m.SetMediaId(fh.Response.MediaId)

	wxResp, err := c.Submit(&m)
	if err != nil {
		t.Error(err)
	}

	if wxResp.Errcode != 0 {
		t.Error(wxResp.Errmsg)
	}
}

func testRobotSendCard(t *testing.T) {
	m := robot.TemplateCard{}
	m.New()

	m.SetTextNotice()

	m.SetSource(robot.Source{
		IconUrl:   "https://wework.qpic.cn/wwpic/252813_jOfDHtcISzuodLa_1629280209/0",
		Desc:      "测试",
		DescColor: 0,
	})

	m.SetMainTitle(robot.MainTitle{
		Title: "欢迎使用企业微信",
		Desc:  "您的好友正在邀请您加入企业微信",
	})

	m.SetEmphasisContent(robot.EmphasisContent{
		Title: "100",
		Desc:  "数据含义",
	})

	m.SetQuoteArea(robot.QuoteArea{
		Type:      1,
		Title:     "引用文本",
		QuoteText: "Jack：企业微信真的很好用?\\nBalian：超级好的一款软件?",
		AppId:     "APPID",
		PagePath:  "PagePath",
		Url:       "https://work.weixin.qq.com/?from=openApi",
	})

	m.SetSubTitleText("副标题")

	m.AddHorizontalContent(robot.HorizontalContent{
		KeyName: "邀请人",
		Value:   "张三",
	})

	m.AddHorizontalContent(robot.HorizontalContent{
		KeyName: "企微官网",
		Value:   "点击访问",
		Type:    1,
		Url:     "https://work.weixin.qq.com/?from=openApi",
	})

	m.AddJump(robot.Jump{
		Type:  1,
		Title: "跳转链接",
		Url:   "https://work.weixin.qq.com/?from=openApi",
	})

	m.SetCardAction(robot.CardAction{
		Type:     1,
		AppId:    "APPID",
		PagePath: "PagePath",
		Url:      "https://work.weixin.qq.com/?from=openApi",
	})

	wxResp, err := c.Submit(&m)
	if err != nil {
		t.Error(err)
	}

	if wxResp.Errcode != 0 {
		t.Error(wxResp.Errmsg)
	}
}

func testRobotSendNewsCard(t *testing.T) {
	m := robot.TemplateCard{}
	m.New()

	m.SetNewsNotice()

	m.SetSource(robot.Source{
		IconUrl:   "https://wework.qpic.cn/wwpic/252813_jOfDHtcISzuodLa_1629280209/0",
		Desc:      "测试",
		DescColor: 0,
	})

	m.SetMainTitle(robot.MainTitle{
		Title: "欢迎使用企业微信",
		Desc:  "您的好友正在邀请您加入企业微信",
	})

	m.SetCardImage(robot.CardImage{
		PicUrl:      "https://wework.qpic.cn/wwpic/252813_jOfDHtcISzuodLa_1629280209/0",
		AspectRatio: 2.25,
	})

	m.SetQuoteArea(robot.QuoteArea{
		Type:      1,
		Title:     "引用文本",
		QuoteText: "Jack：企业微信真的很好用?\\nBalian：超级好的一款软件?",
		AppId:     "APPID",
		PagePath:  "PagePath",
		Url:       "https://work.weixin.qq.com/?from=openApi",
	})

	m.SetSubTitleText("副标题")

	m.SetImageTextArea(robot.ImageTextArea{
		Type:     1,
		ImageUrl: "https://wework.qpic.cn/wwpic/252813_jOfDHtcISzuodLa_1629280209/0",
		Title:    "图片标题",
		Desc:     "图片描述",
		Url:      "https://work.weixin.qq.com/?from=openApi",
	})

	m.AddHorizontalContent(robot.HorizontalContent{
		KeyName: "邀请人",
		Value:   "张三",
	})

	m.AddHorizontalContent(robot.HorizontalContent{
		KeyName: "企微官网",
		Value:   "点击访问",
		Type:    1,
		Url:     "https://work.weixin.qq.com/?from=openApi",
	})

	m.AddJump(robot.Jump{
		Type:  1,
		Title: "跳转链接",
		Url:   "https://work.weixin.qq.com/?from=openApi",
	})

	m.SetCardAction(robot.CardAction{
		Type:     1,
		AppId:    "APPID",
		PagePath: "PagePath",
		Url:      "https://work.weixin.qq.com/?from=openApi",
	})

	m.AddVerticalContent(robot.VerticalContent{
		Title: "垂直内容标题",
		Desc:  "垂直内容描述",
	})

	wxResp, err := c.Submit(&m)
	if err != nil {
		t.Error(err)
	}

	if wxResp.Errcode != 0 {
		t.Error(wxResp.Errmsg)
	}
}
