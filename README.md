# go-object-storage

[![Godoc](http://img.shields.io/badge/godoc-reference-blue.svg?style=flat)](https://godoc.org/github.com/qwxingzhe/go-object-storage)
[![Build Status](https://travis-ci.org/qwxingzhe/go-object-storage.svg)](https://travis-ci.org/qwxingzhe/go-object-storage)
[![GitHub release](http://img.shields.io/github/release/qwxingzhe/go-object-storage.svg?style=flat-square)](release)
[![MIT License](http://img.shields.io/badge/license-MIT-blue.svg?style=flat-square)](license)
[![Commit activity](https://img.shields.io/github/commit-activity/m/qwxingzhe/go-object-storage)](https://github.com/qwxingzhe/go-object-storage/graphs/commit-activity)
[![Average time to resolve an issue](http://isitmaintained.com/badge/resolution/qwxingzhe/go-object-storage.svg)](http://isitmaintained.com/project/qwxingzhe/go-object-storage "Average time to resolve an issue")
[![Percentage of issues still open](http://isitmaintained.com/badge/open/qwxingzhe/go-object-storage.svg)](http://isitmaintained.com/project/qwxingzhe/go-object-storage "Percentage of issues still open")



<p align="center">go-object-storage 是一个用go开发的针对主流云服务商对象存储SDK集成包。其实现了文件路径自动生产，不同存储引擎简单配置即可使用，欢迎使用。</p>



# 安装

```shell
$ go get -u github.com/qwxingzhe/go-object-storage
```

# 使用指南



~~~
package main

import (
	"fmt"
	objectstorage "github.com/qwxingzhe/go-object-storage"
	objectstoragedrives "github.com/qwxingzhe/go-object-storage/drives"
	"github.com/qwxingzhe/go-package-demo/config"
)

// 配置存储引擎
func GetObjectStorage(isAutomaticProductionPath bool) *objectstorage.ObjectStorage {
	return &objectstorage.ObjectStorage{
		Drive: objectstoragedrives.Kodo{
			AccessKey: config.KodoConfig.AccessKey,
			SecretKey: config.KodoConfig.SecretKey,
			Bucket:    config.KodoConfig.Bucket,
		},
		IsAutomaticProductionPath: isAutomaticProductionPath,
		FilePathPrefix:            "test/",
		IsAppendExt:               false,
		BaseUrl:                   "http://qynr9haq9.hd-bkt.clouddn.com/",
	}
}

func main() {
    // 转存网络文件
	url := "https://tenfei03.cfp.cn/creative/vcg/veer/800/new/VCG41N813911068.jpg"
	fileInfo, _ := GetObjectStorage(true).PutNetFile(url)
	fmt.Println(fileInfo)
}
~~~

上传指定内容，指定存储路径

~~~
localFile := "111111111111111"
key := "temp/11111111111.txt"
fileInfo, err := GetObjectStorage(false).SetFilePath(key).PutStr(localFile)
if err != nil {
    t.Errorf(err.Error())
}
fmt.Println(fileInfo)
~~~

上传本地文件

~~~
localFile := "./golang.jpg"
fileInfo, err := GetObjectStorage(true).PutFile(localFile)
if err != nil {
    t.Errorf(err.Error())
}
fmt.Println(fileInfo)
~~~

## 参数说明

#### ObjectStorage 

| 参数                      | 说明                   |
| ------------------------- | ---------------------- |
| Drive                     | 对象存储驱动           |
| IsAutomaticProductionPath | 是否自动生产文件路径   |
| FilePathPrefix            | 指定文件存储前缀       |
| IsAppendExt               | 是否自动拼接文件名后缀 |
| BaseUrl                   | 基础url路径            |

## 实现的存储引擎

- [x] 七牛云 kodo
- [ ] 阿里云 oos
- [ ] 腾讯云 cos
- [ ] Amazon S3 

# Enjoy it! :heart:

# 参照

- [kodo Go SDK 官方文档](https://developer.qiniu.com/kodo/1238/go)

# License

MIT

