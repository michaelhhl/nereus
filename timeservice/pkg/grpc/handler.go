package grpc

import (
	"context"
	"fmt"

	grpc "github.com/go-kit/kit/transport/grpc"
	endpoint "github.com/michaelhhl/nereus/timeservice/pkg/endpoint"
	pb "github.com/michaelhhl/nereus/timeservice/pkg/grpc/pb"
	context1 "golang.org/x/net/context"
)

// makeGetTimeByTimeZoneHandler creates the handler logic
func makeGetTimeByTimeZoneHandler(endpoints endpoint.Endpoints, options []grpc.ServerOption) grpc.Handler {
	return grpc.NewServer(endpoints.GetTimeByTimeZoneEndpoint, decodeGetTimeByTimeZoneRequest, encodeGetTimeByTimeZoneResponse, options...)
}

// decodeGetTimeByTimeZoneResponse is a transport/grpc.DecodeRequestFunc that converts a
// gRPC request to a user-domain GetTimeByTimeZone request.
// TODO implement the decoder
func decodeGetTimeByTimeZoneRequest(_ context.Context, r interface{}) (interface{}, error) {
	fmt.Printf("the r type: %v", r)
	// timeZone := r.(string)
	timeZoneRequest := r.(pb.GetTimeByTimeZoneRequest)
	timeZone := timeZoneRequest.TimeZone
	return endpoint.GetTimeByTimeZoneRequest{TimeZone: timeZone}, nil
}

// encodeGetTimeByTimeZoneResponse is a transport/grpc.EncodeResponseFunc that converts
// a user-domain response to a gRPC reply.
// TODO implement the encoder
func encodeGetTimeByTimeZoneResponse(_ context.Context, r interface{}) (interface{}, error) {
	fmt.Printf("the r type: %v", r)
	timeIsResponse := r.(endpoint.GetTimeByTimeZoneResponse)
	timeIs := timeIsResponse.S0
	return pb.GetTimeByTimeZoneReply{TimeIs: timeIs}, nil
}

func (g *grpcServer) GetTimeByTimeZone(ctx context1.Context, req *pb.GetTimeByTimeZoneRequest) (*pb.GetTimeByTimeZoneReply, error) {
	_, rep, err := g.getTimeByTimeZone.ServeGRPC(ctx, req)
	if err != nil {
		return nil, err
	}
	return rep.(*pb.GetTimeByTimeZoneReply), nil
}
