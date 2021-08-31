package go_object_storage

import (
	"fmt"
	"github.com/qwxingzhe/go-object-storage/drives"
	"testing"
)

func QiNiuObjectStorage() ObjectStorage {
	return ObjectStorage{drive: drives.Kodo{
		AccessKey: kodoConfig.AccessKey,
		SecretKey: kodoConfig.SecretKey,
		Bucket:    kodoConfig.Bucket,
	}}
}

func TestUnit_QiNiu_PutNetFile(t *testing.T) {
	//url1 := "https://imgconvert.csdnimg.cn/aHR0cDovL3AxLnBzdGF0cC5jb20vbGFyZ2UvcGdjLWltYWdlLzA0Zjk5ZjQ1ZTA2NTQyZDA5NjI4Y2E3MjFjNGZiYmVl?x-oss-process=image/format,png"
	url2 := "https://img-home.csdnimg.cn/images/20210831092422.png"
	key := "temp/win11.png"
	err := QiNiuObjectStorage().PutNetFile(url2, key)
	if err != nil {
		fmt.Println(err.Error())
	}
}

//func TestUnit_QiNiu_PutStr(t *testing.T) {
//	localFile := ".111111111111111"
//	key := "temp/11111111111.txt"
//	err := QiNiuObjectStorage().PutStr(localFile, key)
//	if err!=nil{
//		fmt.Println(err.Error())
//	}
//}

//func TestUnit_QiNiu_PutFile(t *testing.T) {
//	localFile := "./golang.jpg"
//	key := "temp/golang.jpg"
//	err := QiNiuObjectStorage().PutFile(localFile, key)
//	if err!=nil{
//		fmt.Println(err.Error())
//	}
//}
