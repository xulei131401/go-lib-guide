1.goreleaser init
2.goreleaser build --snapshot --skip-publish --rm-dist
3.goreleaser check 验证yml正确性
4.goreleaser build --single-target
5.https://goreleaser.com/quick-start/
6.http://www.idmiss.com/701
7.GOOS=darwin GOARCH=amd64 go build main.go
8.go官方文档支持哪些os和架构,https://golang.google.cn/doc/install/source#introduction
9.go tool dist list
10.go条件编译指令https://blog.csdn.net/puss0/article/details/124058910