package drives

import (
	go_file_type "github.com/qwxingzhe/go-file-type"
	"io/ioutil"
	"net/http"
	"path"
	"strings"
)

type ObjectStorageDrive interface {
	// PutFile 上传本地文件
	PutFile(localFile string, key string) error
	// PutContent 上传字符串到对象存储
	PutContent(fileInfo FileInfo, key string) error
}

type FileInfo struct {
	Content []byte
	DataLen int64
	Ext     string
}

// GetNetFileInfo 读取网路文件基础信息
func GetNetFileInfo(fileUrl string) FileInfo {
	res, err := http.Get(fileUrl)

	if err != nil {
		panic(err)
	}

	defer func() {
		if ferr := res.Body.Close(); ferr != nil {
			err = ferr
		}
	}()

	if err != nil {
		panic(err)
	}

	dataLen := res.ContentLength
	bytes, _ := ioutil.ReadAll(res.Body)

	// 获取文件后缀
	Ext := go_file_type.GetFileType(bytes[:10])

	return FileInfo{
		Content: bytes,
		DataLen: dataLen,
		Ext:     Ext,
	}
}

// GetContentInfo 读取字符串基础信息
func GetContentInfo(content string) FileInfo {
	return FileInfo{
		Content: []byte(content),
		DataLen: int64(len(content)),
	}
}

// GetLocalFileInfo 读取本地文件基础信息
func GetLocalFileInfo(localFile string) FileInfo {
	Ext := path.Ext(localFile)
	Ext = strings.Replace(Ext, ".", "", 1)
	return FileInfo{
		Ext: Ext,
	}
}
