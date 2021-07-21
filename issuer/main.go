package main

import (
	"log"
	"time"

	"github.com/aws/aws-lambda-go/events"
	"github.com/aws/aws-lambda-go/lambda"
	"github.com/golang-jwt/jwt"
	"github.com/google/uuid"
)

var (
	signingSecret = []byte("heslojeveslo")
)

func handler(request events.APIGatewayV2HTTPRequest) (events.APIGatewayV2HTTPResponse, error) {
	tokenString, err := token(true, signingSecret)
	if err != nil {
		log.Printf("could not obtain token: %s", err)
		return events.APIGatewayV2HTTPResponse{}, err
	}

	return events.APIGatewayV2HTTPResponse{
		StatusCode: 201,
		Body: tokenString,
	}, nil
}

func main() {
	lambda.Start(handler)
}

func token(isAdmin bool, secret []byte) (string, error) {
	return jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.MapClaims{
		"aud": "api.vrany.dev",
		"exp": time.Now().Add(60 * time.Minute).Unix(),
		"jti": uuid.New().String(),
		"iat": time.Now().Unix(),
		"iss": "vrany.dev",
		"https://vrany.dev/jwt_claims/is_admin": isAdmin,
	}).SignedString(secret)
}