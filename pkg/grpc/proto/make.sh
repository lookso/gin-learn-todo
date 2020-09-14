go get -u github.com/grpc-ecosystem/grpc-gateway/protoc-gen-grpc-gateway
go get -u github.com/grpc-ecosystem/grpc-gateway/protoc-gen-swagger
go get -u github.com/golang/protobuf/protoc-gen-go

# 生成grpc protobuf
protoc --proto_path=../ -I/usr/local/include -I. -I$GOPATH/src -I$GOPATH/pkg/mod/github.com/grpc-ecosystem/grpc-gateway@v1.14.8/third_party/googleapis --go_out=plugins=grpc:. product.proto

# 生成网关proto
protoc --proto_path=../ -I/usr/local/include -I. -I$GOPATH/src -I$GOPATH/pkg/mod/github.com/grpc-ecosystem/grpc-gateway@v1.14.8/third_party/googleapis --grpc-gateway_out=logtostderr=true:. product.proto

#protoc使用
#我们按照惯例执行protoc --help（查看帮助文档），我们抽出几个常用的命令进行讲解
#
#1、-IPATH, --proto_path=PATH：指定import搜索的目录，可指定多个，如果不指定则默认当前工作目录
#2、--go_out：生成golang源文件
#
#参数
#若要将额外的参数传递给插件，可使用从输出目录中分离出来的逗号分隔的参数列表:
#protoc --go_out=plugins=grpc,import_path=mypackage:. *.proto
#import_prefix=xxx：将指定前缀添加到所有import路径的开头
#import_path=foo/bar：如果文件没有声明go_package，则用作包。如果它包含斜杠，那么最右边的斜杠将被忽略。
#plugins=plugin1+plugin2：指定要加载的子插件列表（我们所下载的repo中唯一的插件是grpc）
#Mfoo/bar.proto=quux/shme： M参数，指定.proto文件编译后的包名（foo/bar.proto编译后为包名为quux/shme）
#Grpc支持
#如果proto文件指定了RPC服务，protoc-gen-go可以生成与grpc相兼容的代码，我们仅需要将plugins=grpc参数传递给--go_out，就可以达到这个目的
#protoc --go_out=plugins=grpc:. *.proto
