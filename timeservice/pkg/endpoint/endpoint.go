package endpoint

import (
	"context"

	endpoint "github.com/go-kit/kit/endpoint"
	service "github.com/michaelhhl/nereus/timeservice/pkg/service"
)

// GetTimeByTimeZoneRequest collects the request parameters for the GetTimeByTimeZone method.
type GetTimeByTimeZoneRequest struct {
	TimeZone string `json:"time_zone"`
}

// GetTimeByTimeZoneResponse collects the response parameters for the GetTimeByTimeZone method.
type GetTimeByTimeZoneResponse struct {
	S0 string `json:"s0"`
}

// MakeGetTimeByTimeZoneEndpoint returns an endpoint that invokes GetTimeByTimeZone on the service.
func MakeGetTimeByTimeZoneEndpoint(s service.TimeserviceService) endpoint.Endpoint {
	return func(ctx context.Context, request interface{}) (interface{}, error) {
		req := request.(GetTimeByTimeZoneRequest)
		s0 := s.GetTimeByTimeZone(ctx, req.TimeZone)
		return GetTimeByTimeZoneResponse{S0: s0}, nil
	}
}

// GetTimeByTimeZone implements Service. Primarily useful in a client.
func (e Endpoints) GetTimeByTimeZone(ctx context.Context, timeZone string) (s0 string) {
	request := GetTimeByTimeZoneRequest{TimeZone: timeZone}
	response, err := e.GetTimeByTimeZoneEndpoint(ctx, request)
	if err != nil {
		return
	}
	return response.(GetTimeByTimeZoneResponse).S0
}
