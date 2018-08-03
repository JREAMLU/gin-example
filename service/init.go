package service

import (
	"github.com/JREAMLU/j-kit/http"
	microClient "github.com/micro/go-micro/client"
	opentracing "github.com/opentracing/opentracing-go"
)

const (
	// GreeterClient client
	GreeterClient = "go.micro.srv.greeter"
)

var (
	httpClient *http.Requests
)

// InitMicroClient init micro client
func InitMicroClient(c microClient.Client) {
}

// InitHTTPClient init http client
func InitHTTPClient(tracer opentracing.Tracer) {
	httpClient = http.NewRequests(tracer)
}
