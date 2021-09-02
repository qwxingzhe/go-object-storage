package go_object_storage

import (
	"github.com/qwxingzhe/go-object-storage/drives"
	"testing"
)

func QiNiuObjectStorage() *ObjectStorage {
	return &ObjectStorage{
		drive: drives.Kodo{
			AccessKey: kodoConfig.AccessKey,
			SecretKey: kodoConfig.SecretKey,
			Bucket:    kodoConfig.Bucket,
		},
		isAutomaticProductionPath: true,
		filePathPrefix:            "test/",
		//appendExt: true,
	}
}

func TestUnit_QiNiu_PutNetFile(t *testing.T) {

	//url := "https://imgconvert.csdnimg.cn/aHR0cDovL3AxLnBzdGF0cC5jb20vbGFyZ2UvcGdjLWltYWdlLzA0Zjk5ZjQ1ZTA2NTQyZDA5NjI4Y2E3MjFjNGZiYmVl?x-oss-process=image/format,png"
	//url := "https://img-home.csdnimg.cn/images/20210831092422.png"
	//url := "https://img-home.csdnimg.cn/images/20210902053813.jpg"
	url := "https://tenfei03.cfp.cn/creative/vcg/veer/800/new/VCG41N813911068.jpg"

	err := QiNiuObjectStorage().PutNetFile(url)
	if err != nil {
		t.Errorf(err.Error())
	}
}

func TestUnit_QiNiu_PutStr(t *testing.T) {
	localFile := "111111111111111"
	key := "temp/11111111111.txt"
	err := QiNiuObjectStorage().PutStr(localFile, key)
	if err != nil {
		t.Errorf(err.Error())
	}
}

func TestUnit_QiNiu_PutFile(t *testing.T) {
	localFile := "./golang.jpg"
	err := QiNiuObjectStorage().PutFile(localFile)
	if err != nil {
		t.Errorf(err.Error())
	}
}
