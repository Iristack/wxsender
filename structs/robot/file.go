package robot

import (
    "wxsender/structs"
)

// File 群聊天机器人消息结构 - 文件类型
type File struct {
    structs.Message // 消息基类
    File            struct {
        MediaId string `json:"media_id"` // 文件id，通过下文的文件上传接口获取
    } `json:"file"` // 文件消息
}

func (m *File) New() {
    m.Msgtype = "file"
    m.File = struct {
        MediaId string `json:"media_id"`
    }{
        MediaId: "",
    }
}

func (m *File) SetMediaId(mediaId string) {
    m.File.MediaId = mediaId
}

func (m *File) Complete() string {
    return structs.ToJson(m)
}
