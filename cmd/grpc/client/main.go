package main

import (
	"context"
	"evento-standalone/internal/grpc/domain"
	"evento-standalone/internal/grpc/service"
	"fmt"

	"google.golang.org/grpc"
)

func main() {
	serverAddress := "localhost:7000"

	conn, e := grpc.Dial(serverAddress, grpc.WithInsecure())

	if e != nil {
		panic(e)
	}
	defer conn.Close()

	client := service.NewEventServiceClient(conn)

	for i := range [10]int{} {
		eventModel := domain.Event{
			Id:          int64(i),
			Name:        "TurboLove",
			Description: "Hello mazafaka",
		}

		if responseMessage, e := client.Add(context.Background(), &eventModel); e != nil {
			panic(fmt.Sprintf("Was not able to insert Record %v", e))
		} else {
			fmt.Println("Record Inserted..")
			fmt.Println(responseMessage)
			fmt.Println("=============================")
		}
	}
}
