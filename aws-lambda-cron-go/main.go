package main

import (
	"context"
	"log"

	"github.com/aws/aws-lambda-go/lambda"
)

type Event struct {
	ID     string `json:"id,omitempty"`
	Action string `json:"action,omitempty"`
}

func HandleRequest(ctx context.Context, event Event) error {
	log.Println("Context: ", ctx)
	log.Println("Event received: ", event)

	return nil
}

func main() {
	lambda.Start(HandleRequest)
}
