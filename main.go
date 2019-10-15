package main

import (
	"time"

	"github.com/openzipkin/zipkin-go"
	zipkinhttp "github.com/openzipkin/zipkin-go/reporter/http"
)

func main() {
	rep := zipkinhttp.NewReporter("http://localhost:9411/api/v2/spans")
	defer rep.Close()

	tracer, _ := zipkin.NewTracer(rep)
	sp := tracer.StartSpan("test")
	defer sp.Finish()
	sp.Tag("key", "value")
	time.Sleep(10 * time.Millisecond)
}
