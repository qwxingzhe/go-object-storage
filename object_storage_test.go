package go_object_storage

import (
	"github.com/qwxingzhe/go-object-storage/drives"
	"testing"
)

func TestUnit_Upload(t *testing.T) {
	os := ObjectStorage{drive: drives.Kodo{
		AccessKey: kodoConfig.AccessKey,
		SecretKey: kodoConfig.SecretKey,
		Bucket:    kodoConfig.Bucket,
		//BasePath:  kodoConfig.BasePath,
	}}

	localFile := "./golang.jpg"
	key := "temp/golang.jpg"
	os.PutFile(localFile, key)
}
