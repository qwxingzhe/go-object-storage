package go_object_storage

import (
	"github.com/google/uuid"
	"github.com/qwxingzhe/go-object-storage/drives"
	"time"
)

type ObjectStorage struct {
	// 存储驱动
	Drive drives.ObjectStorageDrive
	// 是否依据文件类型自动补充文件后缀
	IsAppendExt bool
	// 路径前缀
	FilePathPrefix string
	// 是否自动生产路径
	IsAutomaticProductionPath bool
	// 文件存储路径
	FilePathKey string
	// 文件存储基础URL地址
	BaseUrl string
}

type UploadFileInfo struct {
	// 存储文件的完整url
	Url string
	// 存储文件的路径
	Key string
	// TODO 文件大小
	// TODO 文件类型
}

// PRIVATE
//+------------------------------------------------------------------------------------------

// 自动生成文件存储路径
func (receiver *ObjectStorage) automaticProductionPath(fileInfo drives.FileInfo) {
	receiver.FilePathKey = receiver.BuildBasePath() + "." + fileInfo.Ext
}

// 获取文件存储路径
func (receiver *ObjectStorage) getFilePath(fileInfo drives.FileInfo) string {
	if receiver.IsAutomaticProductionPath { // 获取动态路径
		receiver.automaticProductionPath(fileInfo)
	} else if receiver.IsAppendExt { // 拼接文件后缀
		receiver.FilePathKey = receiver.FilePathKey + "." + fileInfo.Ext
	}
	if receiver.FilePathPrefix != "" { // 拼接文件前缀
		receiver.FilePathKey = receiver.FilePathPrefix + receiver.FilePathKey
	}
	return receiver.FilePathKey
}

func (receiver *ObjectStorage) getUploadFileInfo() UploadFileInfo {
	return UploadFileInfo{
		Url: receiver.BaseUrl + receiver.FilePathKey,
		Key: receiver.FilePathKey,
	}
}

// PUBLIC
//+------------------------------------------------------------------------------------------

// BuildBasePath 生成基础路径
func (receiver *ObjectStorage) BuildBasePath() string {
	date := time.Unix(time.Now().Unix(), 0).Format("2006/01/02/")
	return date + uuid.New().String()
}

// SetFilePath  设置文件存储路径
func (receiver *ObjectStorage) SetFilePath(filePathKey string) *ObjectStorage {
	receiver.FilePathKey = filePathKey
	return receiver
}

// PutNetFile 上传网络文件
func (receiver *ObjectStorage) PutNetFile(fileUrl string) (UploadFileInfo, error) {
	// 通过文件地址获取基本信息
	fileInfo := drives.GetNetFileInfo(fileUrl)
	// 获取文件存储路径
	key := receiver.getFilePath(fileInfo)
	err := receiver.Drive.PutContent(fileInfo, key)
	return receiver.getUploadFileInfo(), err
}

// PutFile 上传本地文件
func (receiver *ObjectStorage) PutFile(localFile string) (UploadFileInfo, error) {
	// 通过文件地址获取基本信息
	fileInfo := drives.GetLocalFileInfo(localFile)
	key := receiver.getFilePath(fileInfo)
	err := receiver.Drive.PutFile(localFile, key)
	return receiver.getUploadFileInfo(), err
}

// PutStr 上传文本内容
func (receiver *ObjectStorage) PutStr(content string) (UploadFileInfo, error) {
	fileInfo := drives.GetContentInfo(content)
	key := receiver.getFilePath(fileInfo)
	err := receiver.Drive.PutContent(fileInfo, key)
	return receiver.getUploadFileInfo(), err
}
