package grpc

import (
	"context"
	"errors"
	grpc "github.com/go-kit/kit/transport/grpc"
	context1 "golang.org/x/net/context"
	endpoint "hhl.com/timeservice/pkg/endpoint"
	pb "hhl.com/timeservice/pkg/grpc/pb"
)

// makeGetTimeByTimeZoneHandler creates the handler logic
func makeGetTimeByTimeZoneHandler(endpoints endpoint.Endpoints, options []grpc.ServerOption) grpc.Handler {
	return grpc.NewServer(endpoints.GetTimeByTimeZoneEndpoint, decodeGetTimeByTimeZoneRequest, encodeGetTimeByTimeZoneResponse, options...)
}

// decodeGetTimeByTimeZoneResponse is a transport/grpc.DecodeRequestFunc that converts a
// gRPC request to a user-domain GetTimeByTimeZone request.
// TODO implement the decoder
func decodeGetTimeByTimeZoneRequest(_ context.Context, r interface{}) (interface{}, error) {
	return nil, errors.New("'Timeservice' Decoder is not impelemented")
}

// encodeGetTimeByTimeZoneResponse is a transport/grpc.EncodeResponseFunc that converts
// a user-domain response to a gRPC reply.
// TODO implement the encoder
func encodeGetTimeByTimeZoneResponse(_ context.Context, r interface{}) (interface{}, error) {
	return nil, errors.New("'Timeservice' Encoder is not impelemented")
}
func (g *grpcServer) GetTimeByTimeZone(ctx context1.Context, req *pb.GetTimeByTimeZoneRequest) (*pb.GetTimeByTimeZoneReply, error) {
	_, rep, err := g.getTimeByTimeZone.ServeGRPC(ctx, req)
	if err != nil {
		return nil, err
	}
	return rep.(*pb.GetTimeByTimeZoneReply), nil
}
