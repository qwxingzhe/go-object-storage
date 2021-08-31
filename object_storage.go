package go_object_storage

import "github.com/qwxingzhe/go-object-storage/drives"

type ObjectStorage struct {
	drive drives.ObjectStorageDrive
}

// PutFile 上传本地文件
func (receiver ObjectStorage) PutFile(localFile string, key string) error {
	return receiver.drive.PutFile(localFile, key)
}

// PutStr 上传文本内容
func (receiver ObjectStorage) PutStr(content string, key string) error {
	return receiver.drive.PutStr(content, key)
}

// PutNetFile 上传网络文件
func (receiver ObjectStorage) PutNetFile(fileUrl string, key string) error {
	return receiver.drive.PutNetFile(fileUrl, key)
}
