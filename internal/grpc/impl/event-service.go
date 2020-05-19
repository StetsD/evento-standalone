package impl

import (
	"context"
	"evento-standalone/internal/grpc/domain"
	"evento-standalone/internal/grpc/service"
	"log"
	"strconv"
)

type EventServiceGrpcImpl struct{}

func NewEventServiceGrpcImpl() *EventServiceGrpcImpl {
	return &EventServiceGrpcImpl{}
}

func (serviceImpl *EventServiceGrpcImpl) Add(ctx context.Context, in *domain.Event) (*service.AddEventResponse, error) {
	log.Println("Receive request for adding repository with id " + strconv.FormatInt(in.Id, 10))
	log.Println("Event persisted to the storage")

	return &service.AddEventResponse{
		AddedEvent: in,
		Error:      nil,
	}, nil
}
