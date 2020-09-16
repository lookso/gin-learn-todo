package zipkin

import (
	"gin-learn-todo/pkg/log"
	"github.com/opentracing/opentracing-go"
	zkOt "github.com/openzipkin-contrib/zipkin-go-opentracing"
	"github.com/openzipkin/zipkin-go"
	zkHttp "github.com/openzipkin/zipkin-go/reporter/http"
)

// 第一步: 开一个全局变量
var ZkTracer opentracing.Tracer

func Init() error {
	// 第二步: 初始化 tracer
	reporter := zkHttp.NewReporter("http://localhost:9411/api/v2/spans")
	defer reporter.Close()
	endpoint, err := zipkin.NewEndpoint("gin-zin-server", "localhost:80")
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
	return nil
}
