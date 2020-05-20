package grpc

import (
	grpc "github.com/go-kit/kit/transport/grpc"
	endpoint "github.com/michaelhhl/nereus/timeservice/pkg/endpoint"
	pb "github.com/michaelhhl/nereus/timeservice/pkg/grpc/pb"
)

// NewGRPCServer makes a set of endpoints available as a gRPC AddServer
type grpcServer struct {
	getTimeByTimeZone grpc.Handler
}

//NewGRPCServer create new grpc server
func NewGRPCServer(endpoints endpoint.Endpoints, options map[string][]grpc.ServerOption) pb.TimeserviceServer {
	return &grpcServer{getTimeByTimeZone: makeGetTimeByTimeZoneHandler(endpoints, options["GetTimeByTimeZone"])}
}
