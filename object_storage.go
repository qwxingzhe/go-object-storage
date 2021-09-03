package go_object_storage

import (
	"github.com/google/uuid"
	"github.com/qwxingzhe/go-object-storage/drives"
	"time"
)

type ObjectStorage struct {
	// 存储驱动
	drive drives.ObjectStorageDrive
	// 是否依据文件类型自动补充文件后缀
	appendExt bool
	// 路径前缀
	filePathPrefix string
	// 是否自动生产路径
	isAutomaticProductionPath bool
	// 文件存储路径
	filePathKey string
	// 文件存储基础URL地址
	baseUrl string
}

type UploadFileInfo struct {
	// 存储文件的完整url
	url string
	// 存储文件的路径
	key string
	// TODO 文件大小
	// TODO 文件类型
}

// PRIVATE
//+------------------------------------------------------------------------------------------

// 自动生成文件存储路径
func (receiver *ObjectStorage) automaticProductionPath(fileInfo drives.FileInfo) {
	receiver.filePathKey = receiver.BuildBasePath() + "." + fileInfo.Ext
}

// 获取文件存储路径
func (receiver *ObjectStorage) getFilePath(fileInfo drives.FileInfo) string {
	if receiver.isAutomaticProductionPath { // 获取动态路径
		receiver.automaticProductionPath(fileInfo)
	} else if receiver.appendExt { // 拼接文件后缀
		receiver.filePathKey = receiver.filePathKey + "." + fileInfo.Ext
	}
	if receiver.filePathPrefix != "" { // 拼接文件前缀
		receiver.filePathKey = receiver.filePathPrefix + receiver.filePathKey
	}
	return receiver.filePathKey
}

func (receiver *ObjectStorage) getUploadFileInfo() UploadFileInfo {
	return UploadFileInfo{
		url: receiver.baseUrl + receiver.filePathKey,
		key: receiver.filePathKey,
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
	receiver.filePathKey = filePathKey
	return receiver
}

// PutNetFile 上传网络文件
func (receiver *ObjectStorage) PutNetFile(fileUrl string) (UploadFileInfo, error) {
	// 通过文件地址获取基本信息
	fileInfo := drives.GetNetFileInfo(fileUrl)
	// 获取文件存储路径
	key := receiver.getFilePath(fileInfo)
	err := receiver.drive.PutContent(fileInfo, key)
	return receiver.getUploadFileInfo(), err
}

// PutFile 上传本地文件
func (receiver *ObjectStorage) PutFile(localFile string) (UploadFileInfo, error) {
	// 通过文件地址获取基本信息
	fileInfo := drives.GetLocalFileInfo(localFile)
	key := receiver.getFilePath(fileInfo)
	err := receiver.drive.PutFile(localFile, key)
	return receiver.getUploadFileInfo(), err
}

// PutStr 上传文本内容
func (receiver *ObjectStorage) PutStr(content string) (UploadFileInfo, error) {
	fileInfo := drives.GetContentInfo(content)
	key := receiver.getFilePath(fileInfo)
	err := receiver.drive.PutContent(fileInfo, key)
	return receiver.getUploadFileInfo(), err
}
