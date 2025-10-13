package robot

import "wxsender/structs"

// Image 群聊天机器人消息结构 - 图片类型
type Image struct {
    structs.Message
    Image struct {
        Base64 string `json:"base64"`
        Md5    string `json:"md5"`
    } `json:"image"`
}

func (m *Image) New() {
    m.Msgtype = "image"
    m.Image = struct {
        Base64 string `json:"base64"`
        Md5    string `json:"md5"`
    }{
        Base64: "",
        Md5:    "",
    }
}

func (m *Image) SetBase64(base64 string) {
    m.Image.Base64 = base64
}
func (m *Image) SetMd5(md5 string) {
    m.Image.Md5 = md5
}

func (m *Image) SetPicture(picurl string) error {
    base64Str, err := structs.ToBase64(picurl)
    if err != nil {
        return err
    }
    md5Str, err := structs.CalculateMD5(picurl)
    if err != nil {
        return err
    }
    
    m.SetBase64(base64Str)
    m.SetMd5(md5Str)
    return nil
}

func (m *Image) Complete() string {
    return structs.ToJson(m)
}
