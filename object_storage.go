package go_object_storage

import (
	"fmt"
	"github.com/qwxingzhe/go-object-storage/drives"
)

type ObjectStorage struct {
	// 存储驱动
	drive drives.ObjectStorageDrive
	// 是否依据文件类型自动补充文件后缀
	appendExt bool
}

// PutNetFile 上传网络文件
func (receiver ObjectStorage) PutNetFile(fileUrl string, key string) error {
	fileInfo := drives.GetNetFileInfo(fileUrl)
	if receiver.appendExt {
		key = key + "." + fileInfo.Ext
	}
	fmt.Println("----------------------------->>>>", key)
	return receiver.drive.PutNetFile(fileInfo, key)
}

// PutFile 上传本地文件
func (receiver ObjectStorage) PutFile(localFile string, key string) error {
	return receiver.drive.PutFile(localFile, key)
}

// PutStr 上传文本内容
func (receiver ObjectStorage) PutStr(content string, key string) error {
	return receiver.drive.PutStr(content, key)
}
