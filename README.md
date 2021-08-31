# go-object-storage
对主流对象存储服务商库的二次封装

# DEMO
~~~
# 上传本地文件
os := ObjectStorage{drive: drives.Kodo{
    AccessKey: kodoConfig.AccessKey,
    SecretKey: kodoConfig.SecretKey,
    Bucket:    kodoConfig.Bucket,
    //BasePath:  kodoConfig.BasePath,
}}

localFile := "./golang.jpg"
key := "temp/golang.jpg"
os.PutFile(localFile,key)
~~~