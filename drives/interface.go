package drives

type ObjectStorageDrive interface {
	PutFile(localFile string, key string) error
	PutStr(content string, key string) error
	PutNetFile(fileUrl string, key string) error
}
