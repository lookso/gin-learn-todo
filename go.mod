module gin-learn-todo

go 1.14

require (
	github.com/BurntSushi/toml v0.3.1
	github.com/alecthomas/template v0.0.0-20190718012654-fb15b899a751
	github.com/coreos/etcd v3.3.25+incompatible // indirect
	github.com/coreos/go-systemd v0.0.0-20191104093116-d3cd4ed1dbcf // indirect
	github.com/coreos/pkg v0.0.0-20180928190104-399ea9e2e55f // indirect
	github.com/dgrijalva/jwt-go v3.2.0+incompatible
	github.com/getsentry/sentry-go v0.7.0
	github.com/gin-gonic/gin v1.6.3
	github.com/go-redis/redis v6.15.9+incompatible
	github.com/golang/protobuf v1.4.2
	github.com/google/gops v0.3.12
	github.com/google/uuid v1.1.2
	github.com/grpc-ecosystem/grpc-gateway v1.14.8
	github.com/jinzhu/gorm v1.9.16
	github.com/prometheus/client_golang v1.7.1
	github.com/swaggo/gin-swagger v1.2.0
	github.com/swaggo/swag v1.6.7
	go.etcd.io/etcd v3.3.25+incompatible
	go.uber.org/zap v1.16.0
	golang.org/x/net v0.0.0-20200324143707-d3edc9973b7e
	google.golang.org/genproto v0.0.0-20200526211855-cb27e3aa2013
	google.golang.org/grpc v1.31.1
	google.golang.org/protobuf v1.25.0
	gopkg.in/natefinch/lumberjack.v2 v2.0.0
)

replace google.golang.org/grpc v1.31.1 => google.golang.org/grpc v1.27.0
