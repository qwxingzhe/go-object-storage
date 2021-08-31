package drives

import (
	"bufio"
	"bytes"
	"context"
	"fmt"
	"github.com/qiniu/go-sdk/v7/auth/qbox"
	"github.com/qiniu/go-sdk/v7/storage"
	"net/http"
)

type Kodo struct {
	AccessKey    string
	SecretKey    string
	Bucket       string
	upToken      string
	formUploader *storage.FormUploader
	putRet       storage.PutRet
	putExtra     storage.PutExtra
}

func (receiver *Kodo) init() {
	putPolicy := storage.PutPolicy{
		Scope: receiver.Bucket,
	}
	mac := qbox.NewMac(receiver.AccessKey, receiver.SecretKey)
	receiver.upToken = putPolicy.UploadToken(mac)

	cfg := storage.Config{}
	// 空间对应的机房
	cfg.Zone = &storage.ZoneHuadong
	// 是否使用https域名
	cfg.UseHTTPS = true
	// 上传是否使用CDN上传加速
	cfg.UseCdnDomains = false
	// 构建表单上传的对象
	receiver.formUploader = storage.NewFormUploader(&cfg)
	receiver.putRet = storage.PutRet{}
	// 可选配置
	receiver.putExtra = storage.PutExtra{
		Params: map[string]string{
			"x:name": "github logo",
		},
	}
}

func (receiver Kodo) PutNetFile(fileUrl string, key string) error {
	receiver.init()
	//data := []byte(fileUrl)
	//dataLen := int64(len(data))

	//+-------------------------------------------------------
	res, errGet := http.Get(fileUrl)
	if errGet != nil {
		return errGet
	}
	defer res.Body.Close()
	// 获得get请求响应的reader对象
	reader := bufio.NewReaderSize(res.Body, 32*1024)
	dataLen := res.ContentLength
	//+-------------------------------------------------------

	err := receiver.formUploader.Put(context.Background(), &receiver.putRet, receiver.upToken, key, reader, dataLen, &receiver.putExtra)
	if err != nil {
		fmt.Println(err)
		return err
	}
	fmt.Println(receiver.putRet.Key, receiver.putRet.Hash)

	return nil
}

func (receiver Kodo) PutStr(content string, key string) error {
	receiver.init()
	data := []byte(content)
	dataLen := int64(len(data))
	err := receiver.formUploader.Put(context.Background(), &receiver.putRet, receiver.upToken, key, bytes.NewReader(data), dataLen, &receiver.putExtra)
	if err != nil {
		fmt.Println(err)
		return err
	}
	fmt.Println(receiver.putRet.Key, receiver.putRet.Hash)

	return nil
}

func (receiver Kodo) PutFile(localFile string, key string) error {
	receiver.init()
	err := receiver.formUploader.PutFile(context.Background(), &receiver.putRet, receiver.upToken, key, localFile, &receiver.putExtra)
	if err != nil {
		return err
	}
	fmt.Println(receiver.putRet.Key, receiver.putRet.Hash)

	return nil
}
