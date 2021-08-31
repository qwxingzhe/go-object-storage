package drives

type ObjectStorageDrive interface {
	Upload()
	PutFile(localFile string, key string) error
}
