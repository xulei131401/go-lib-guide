# go-lib-guide
练习go第三方库

## 工作空间
### db
数据库相关

## 创建工作空间
```shell
cd ./basic && \
go work init
```

## 创建module
```shell
mkdir xxx && \
go mod init github.com/zz-guide/go-lib-guide/xxx/xxx
```

## module 加入工作空间
```shell
go work use ./xxx
```