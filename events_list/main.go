package main

import (
	"github.com/aws/aws-lambda-go/events"
	"github.com/aws/aws-lambda-go/lambda"
)

func handler(request events.APIGatewayV2HTTPRequest) (events.APIGatewayV2HTTPResponse, error) {

	return events.APIGatewayV2HTTPResponse{
		StatusCode: 201,
		Body: "I'm alive",
	}, nil
}

func main() {
	lambda.Start(handler)
}
