package trace

import (
	"gin-learn-todo/pkg/log"
	"github.com/opentracing/opentracing-go"
	zkOt "github.com/openzipkin-contrib/zipkin-go-opentracing"
	"github.com/openzipkin/zipkin-go"
	zkHttp "github.com/openzipkin/zipkin-go/reporter/http"
)

// 第一步: 开一个全局变量
var ZkTracer opentracing.Tracer

const (
	ServerName             = "ZpKin_Server"
	ZipKinHttpEndPoint     = "http://127.0.0.1:9411/api/v1/spans"
	ZipKinRecorderHostPort = "127.0.0.1:80"
)

func NewTrace() error {
	// 第二步: 初始化 tracer
	reporter := zkHttp.NewReporter(ZipKinHttpEndPoint)
	defer reporter.Close()
	endpoint, err := zipkin.NewEndpoint(ServerName, ZipKinRecorderHostPort)
	if err != nil {
		log.Sugar().Fatalf("unable to create local endpoint: %+v\n", err)
		return err
	}
	nativeTracer, err := zipkin.NewTracer(reporter, zipkin.WithLocalEndpoint(endpoint))
	if err != nil {
		log.Sugar().Fatalf("unable to create tracer: %+v\n", err)
		return err
	}
	ZkTracer = zkOt.Wrap(nativeTracer)
	opentracing.SetGlobalTracer(ZkTracer)
	log.Sugar().Info("ZipKin init success")
	return nil
}
