package trace

import (
	"gin-learn-todo/pkg/log"
	"github.com/opentracing/opentracing-go"
	zkOt "github.com/openzipkin-contrib/zipkin-go-opentracing"
	"github.com/openzipkin/zipkin-go"
	zkHttp "github.com/openzipkin/zipkin-go/reporter/http"
	"time"
)

const (
	ServerName         = "zipkin_grpc_server"
	ZKHttpEndPoint     = "http://127.0.0.1:9411/api/v1/spans"
	ZKRecorderHostPort = "127.0.0.1:80"
)

func NewTrace() (zkTracer opentracing.Tracer, err error) {
	reporter := zkHttp.NewReporter(ZKHttpEndPoint, zkHttp.Timeout(time.Duration(5*time.Second)))
	defer reporter.Close()

	// create our local service endpoint
	endpoint, err := zipkin.NewEndpoint(ServerName, ZKRecorderHostPort)
	if err != nil {
		log.Sugar().Fatalf("unable to create local endpoint: %+v\n", err)
		return nil, err
	}

	// initialize our tracer
	nativeTracer, err := zipkin.NewTracer(reporter, zipkin.WithLocalEndpoint(endpoint), zipkin.WithTraceID128Bit(true))
	if err != nil {
		log.Sugar().Fatalf("unable to create tracer: %+v\n", err)
		return nil, err
	}

	// use zipkin-go-opentracing to wrap our tracer
	zkTracer = zkOt.Wrap(nativeTracer)

	// optionally set as Global OpenTracing tracer instance
	opentracing.InitGlobalTracer(zkTracer)
	log.Sugar().Info("ZipKin init success")
	return zkTracer, nil
}
