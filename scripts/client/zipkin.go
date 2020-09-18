package client

import (
	"github.com/opentracing/opentracing-go"
	zKOt "github.com/openzipkin-contrib/zipkin-go-opentracing"
	"github.com/openzipkin/zipkin-go"
	zkHttp "github.com/openzipkin/zipkin-go/reporter/http"
	"log"
)

const (
	ServerName             = "ZpKin_GRpc_Client"
	ZipKinHttpEndPoint     = "http://127.0.0.1:9411/api/v1/spans"
	ZipKinRecorderHostPort = "127.0.0.1:80"
)

func NewTracer() (opentracing.Tracer, error) {

	reporter := zkHttp.NewReporter(ZipKinHttpEndPoint)
	defer reporter.Close()

	// create our local service endpoint
	endpoint, err := zipkin.NewEndpoint(ServerName, ZipKinRecorderHostPort)
	if err != nil {
		log.Fatalf("unable to create local endpoint: %+v\n", err)
	}

	// initialize our tracer
	nativeTracer, err := zipkin.NewTracer(reporter, zipkin.WithLocalEndpoint(endpoint))
	if err != nil {
		log.Fatalf("unable to create tracer: %+v\n", err)
	}

	// use zipkin-go-opentracing to wrap our tracer
	tracer := zKOt.Wrap(nativeTracer)

	// optionally set as Global OpenTracing tracer instance
	opentracing.SetGlobalTracer(tracer)
	return tracer, nil
}
