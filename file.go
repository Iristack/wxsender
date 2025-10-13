package main

import (
	"bytes"
	"encoding/json"
	"io"
	"mime/multipart"
	"net/http"
	"os"
	"path/filepath"
)

type Type string

const (
	File  Type = "file"
	Voice Type = "voice"
)

const baseUrl = "https://qyapi.weixin.qq.com/cgi-bin/webhook/upload_media?"

type FileHelper struct {
	Type      Type   // 文件类型
	FilePath  string // 文件路径
	Key       string // 文件名
	UrlPrefix string // 文件上传地址前缀, 默认为: https://qyapi.weixin.qq.com/cgi-bin/webhook/upload_media?
	Response  FileResponse
	Err       error
}

type FileResponse struct {
	Errcode  int    `json:"errcode"`
	Errmsg   string `json:"errmsg"`
	Type     string `json:"type"`
	MediaId  string `json:"media_id"`
	CreateAt int64  `json:"create_at"`
}

func (f *FileHelper) Upload() *FileHelper {
	if f.UrlPrefix == "" {
		f.UrlPrefix = baseUrl
	}

	url := f.UrlPrefix + "key=" + f.Key + "&type=" + string(f.Type)

	file, err := os.Open(f.FilePath)
	if err != nil {
		f.Err = err
		return f
	}
	defer file.Close()

	body := &bytes.Buffer{}
	writer := multipart.NewWriter(body)

	part, err := writer.CreateFormFile("media", filepath.Base(f.FilePath))
	if err != nil {
		f.Err = err
		return f
	}

	_, err = io.Copy(part, file)
	if err != nil {
		f.Err = err
		return f
	}

	err = writer.Close()
	if err != nil {
		f.Err = err
		return f
	}

	req, err := http.NewRequest("POST", url, body)
	if err != nil {
		f.Err = err
		return f
	}

	req.Header.Set("Content-Type", writer.FormDataContentType())

	client := &http.Client{}
	resp, err := client.Do(req)
	if err != nil {
		f.Err = err
		return f
	}
	defer resp.Body.Close()

	respBody, err := io.ReadAll(resp.Body)
	if err != nil {
		f.Err = err
		return f
	}

	var fileResp FileResponse
	err = json.Unmarshal(respBody, &fileResp)
	if err != nil {
		f.Err = err
		return f
	}

	f.Response = fileResp
	return f
}

func (f *FileHelper) IsOk() bool {
	return f.Response.Errcode == 0
}
