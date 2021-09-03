package go_object_storage

import (
	"fmt"
	"github.com/qwxingzhe/go-object-storage/drives"
	"testing"
)

func GetObjectStorage(isAutomaticProductionPath bool) *ObjectStorage {
	return &ObjectStorage{
		Drive: drives.Kodo{
			AccessKey: kodoConfig.AccessKey,
			SecretKey: kodoConfig.SecretKey,
			Bucket:    kodoConfig.Bucket,
		},
		IsAutomaticProductionPath: isAutomaticProductionPath,
		FilePathPrefix:            "test/",
		IsAppendExt:               false,
		BaseUrl:                   "http://qynr9haq9.hd-bkt.clouddn.com/",
	}
}

func TestUnit_PutNetFile(t *testing.T) {

	//url := "https://imgconvert.csdnimg.cn/aHR0cDovL3AxLnBzdGF0cC5jb20vbGFyZ2UvcGdjLWltYWdlLzA0Zjk5ZjQ1ZTA2NTQyZDA5NjI4Y2E3MjFjNGZiYmVl?x-oss-process=image/format,png"
	//url := "https://img-home.csdnimg.cn/images/20210831092422.png"
	//url := "https://img-home.csdnimg.cn/images/20210902053813.jpg"
	url := "https://tenfei03.cfp.cn/creative/vcg/veer/800/new/VCG41N813911068.jpg"

	fileInfo, err := GetObjectStorage(true).PutNetFile(url)
	if err != nil {
		t.Errorf(err.Error())
	}
	fmt.Println(fileInfo)
}

func TestUnit_PutStr(t *testing.T) {
	localFile := "111111111111111"
	key := "temp/11111111111.txt"
	fileInfo, err := GetObjectStorage(false).SetFilePath(key).PutStr(localFile)
	if err != nil {
		t.Errorf(err.Error())
	}
	fmt.Println(fileInfo)
}

func TestUnit_PutFile(t *testing.T) {
	localFile := "./golang.jpg"
	fileInfo, err := GetObjectStorage(true).PutFile(localFile)
	if err != nil {
		t.Errorf(err.Error())
	}
	fmt.Println(fileInfo)
}
