package service

import (
	"context"
	"time"
)

// TimeserviceService describes the service.
type TimeserviceService interface {
	GetTimeByTimeZone(ctx context.Context, timeZone string) string
}

type basicTimeserviceService struct{}

func (b *basicTimeserviceService) GetTimeByTimeZone(ctx context.Context, timeZone string) (theTime string) {
	// TODO implement the business logic of GetTimeByTimeZone
	tm := time.Now()
	switch timeZone {
	case "1":
	case "2":
		tm = tm.Add(time.Hour * 1)
	case "3":
		tm = tm.Add(time.Hour * 2)
	}
	theTime = tm.Format("2006-01-02 15:04:05")
	return
}

// NewBasicTimeserviceService returns a naive, stateless implementation of TimeserviceService.
func NewBasicTimeserviceService() TimeserviceService {
	return &basicTimeserviceService{}
}

// New returns a TimeserviceService with all of the expected middleware wired in.
func New(middlewares []Middleware) TimeserviceService {
	var svc TimeserviceService = NewBasicTimeserviceService()
	for _, m := range middlewares {
		svc = m(svc)
	}
	return svc
}
