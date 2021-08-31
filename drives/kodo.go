package drives

import (
	"context"
	"fmt"
	"github.com/qiniu/go-sdk/v7/auth/qbox"
	"github.com/qiniu/go-sdk/v7/storage"
)

type Kodo struct {
	AccessKey string
	SecretKey string
	Bucket    string
}

func (receiver Kodo) PutFile(localFile string, key string) error {
	fmt.Println("Kodo Upload |--------------------------")

	putPolicy := storage.PutPolicy{
		Scope: receiver.Bucket,
	}
	mac := qbox.NewMac(receiver.AccessKey, receiver.SecretKey)
	upToken := putPolicy.UploadToken(mac)
	cfg := storage.Config{}
	// 空间对应的机房
	cfg.Zone = &storage.ZoneHuadong
	// 是否使用https域名
	cfg.UseHTTPS = true
	// 上传是否使用CDN上传加速
	cfg.UseCdnDomains = false
	// 构建表单上传的对象
	formUploader := storage.NewFormUploader(&cfg)
	ret := storage.PutRet{}
	// 可选配置
	putExtra := storage.PutExtra{
		Params: map[string]string{
			"x:name": "github logo",
		},
	}
	err := formUploader.PutFile(context.Background(), &ret, upToken, key, localFile, &putExtra)
	if err != nil {
		return err
	}
	fmt.Println(ret.Key, ret.Hash)

	return nil
}

func (receiver Kodo) Upload() {

}
