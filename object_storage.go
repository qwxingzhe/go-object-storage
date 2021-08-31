package go_object_storage

import "github.com/qwxingzhe/go-object-storage/drives"

type ObjectStorage struct {
	drive drives.ObjectStorageDrive
}

func (receiver ObjectStorage) Upload() {
	receiver.drive.Upload()
}

// PutFile 上传本地文件
func (receiver ObjectStorage) PutFile(localFile string, key string) error {
	return receiver.drive.PutFile(localFile, key)
}
