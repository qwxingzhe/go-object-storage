package drives

import (
	"bufio"
	go_file_type "github.com/qwxingzhe/go-file-type"
	"net/http"
	"path"
)

type ObjectStorageDrive interface {
	// PutFile 上传本地文件
	PutFile(localFile string, key string) error
	// PutStr 上传字符串到对象存储
	PutStr(content string, key string) error
	// PutNetFile 将网络文件转存到对象存储
	PutNetFile(fileUrl FileInfo, key string) error
}

type FileInfo struct {
	Response *http.Response
	Reader   *bufio.Reader
	DataLen  int64
	Ext      string
}

// GetNetFileInfo 读取网路文件基础信息
func GetNetFileInfo(fileUrl string) FileInfo {
	Res, errGet := http.Get(fileUrl)
	if errGet != nil {
		panic(errGet)
	}
	//defer res.Body.Close()
	// 获得get请求响应的reader对象
	Reader := bufio.NewReaderSize(Res.Body, 32*1024)
	DataLen := Res.ContentLength

	// 获取文件后缀
	bytes := make([]byte, 10)
	Reader.Read(bytes)
	Ext := go_file_type.GetFileType(bytes)

	return FileInfo{
		Response: Res,
		Reader:   Reader,
		DataLen:  DataLen,
		Ext:      Ext,
	}
}

// GetLocalFileInfo 读取本地文件基础信息
func GetLocalFileInfo(localFile string) FileInfo {
	Ext := path.Ext(localFile)
	return FileInfo{
		Ext: Ext,
	}
}
