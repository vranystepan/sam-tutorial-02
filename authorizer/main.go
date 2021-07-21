package main

import (
	"errors"
	"fmt"
	"github.com/aws/aws-lambda-go/lambda"
	"github.com/golang-jwt/jwt"
	"log"
)

var (
	signingSecret = []byte("heslojeveslo")
)

func handler(request AuthorizerV2Requests) (AuthorizerV2Response, error) {

	// extract token from request
	tokenString, err := request.GetJWTToken()
	if err != nil {
		log.Printf("could not extract token: %s", err)
		return generateAuthorizerV2Response(request.RouteARN, "deny"), nil
	}

	// parse token
	tokenJWT, err := parse(tokenString, signingSecret)
	if err != nil {
		log.Printf("could not parse token: %s", err)
		return generateAuthorizerV2Response(request.RouteARN, "deny"), nil
	}

	// check if token is valid
	if !tokenJWT.Valid {
		log.Printf("token '%s' is not valid", tokenString)
		return generateAuthorizerV2Response(request.RouteARN, "deny"), nil
	}

	// get claims
	claims, ok := tokenJWT.Claims.(jwt.MapClaims)
	if ! ok {
		log.Printf("could not get claims")
		return generateAuthorizerV2Response(request.RouteARN, "deny"), nil
	}

	// get the claim
	isAdmin, ok := claims["https://vrany.dev/jwt_claims/is_admin"]
	if ! ok {
		log.Printf("could not get is_admin claim")
		return generateAuthorizerV2Response(request.RouteARN, "deny"), nil
	}

	// validata the claim
	if ! isAdmin.(bool) {
		log.Printf("is_admin is set to false")
		return generateAuthorizerV2Response(request.RouteARN, "deny"), nil
	}

	return generateAuthorizerV2Response(request.RouteARN, "allow"), nil
}

func main() {
	lambda.Start(handler)
}

func parse(tokenString string, secret []byte) (*jwt.Token, error) {
	return jwt.Parse(tokenString, func(token *jwt.Token) (interface{}, error) {
		if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
			return nil, errors.New(fmt.Sprintf("Unexpected signing method: %v", token.Header["alg"]))
		}
		return []byte(secret), nil
	})
}

func generateAuthorizerV2Response(resource string, effect string) AuthorizerV2Response {
	return AuthorizerV2Response{
		PrincipalID: "apigateway.amazonaws.com",
		PolicyDocument: AuthorizerV2ResponsePolicyDocument{
			Version: "2012-10-17",
			Statement: []AuthorizerV2ResponsePolicyStatement{
				AuthorizerV2ResponsePolicyStatement{
					Action: "execute-api:Invoke",
					Effect: effect,
					Resource: []string{resource},
				},
			},
		},
	}
}