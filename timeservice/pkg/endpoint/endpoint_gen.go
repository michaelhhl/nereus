package endpoint

import (
	endpoint "github.com/go-kit/kit/endpoint"
	service "github.com/michaelhhl/nereus/timeservice/pkg/service"
)

// Endpoints collects all of the endpoints that compose a profile service. It's
// meant to be used as a helper struct, to collect all of the endpoints into a
// single parameter.
type Endpoints struct {
	GetTimeByTimeZoneEndpoint endpoint.Endpoint
}

// New returns a Endpoints struct that wraps the provided service, and wires in all of the
// expected endpoint middlewares
func New(s service.TimeserviceService, mdw map[string][]endpoint.Middleware) Endpoints {
	eps := Endpoints{GetTimeByTimeZoneEndpoint: MakeGetTimeByTimeZoneEndpoint(s)}
	for _, m := range mdw["GetTimeByTimeZone"] {
		eps.GetTimeByTimeZoneEndpoint = m(eps.GetTimeByTimeZoneEndpoint)
	}
	return eps
}
