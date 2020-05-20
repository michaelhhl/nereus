package service

import (
	endpoint1 "github.com/go-kit/kit/endpoint"
	log "github.com/go-kit/kit/log"
	opentracing "github.com/go-kit/kit/tracing/opentracing"
	grpc "github.com/go-kit/kit/transport/grpc"
	endpoint "github.com/michaelhhl/nereus/timeservice/pkg/endpoint"
	group "github.com/oklog/oklog/pkg/group"
	opentracinggo "github.com/opentracing/opentracing-go"
)

func createService(endpoints endpoint.Endpoints) (g *group.Group) {
	g = &group.Group{}
	initGRPCHandler(endpoints, g)
	return g
}
func defaultGRPCOptions(logger log.Logger, tracer opentracinggo.Tracer) map[string][]grpc.ServerOption {
	options := map[string][]grpc.ServerOption{"GetTimeByTimeZone": {grpc.ServerErrorLogger(logger), grpc.ServerBefore(opentracing.GRPCToContext(tracer, "GetTimeByTimeZone", logger))}}
	return options
}
func addEndpointMiddlewareToAllMethods(mw map[string][]endpoint1.Middleware, m endpoint1.Middleware) {
	methods := []string{"GetTimeByTimeZone"}
	for _, v := range methods {
		mw[v] = append(mw[v], m)
	}
}
