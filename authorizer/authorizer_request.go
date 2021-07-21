package main

import (
	"errors"
	"github.com/aws/aws-lambda-go/events"
	"strings"
)

type AuthorizerV2Requests struct {
	Version               string                                `json:"version"`
	Type                  string                                `json:"type"`
	RouteARN              string                                `json:"routeArn"`
	IdentitySource        []string                              `json:"identitySource"`
	RouteKey              string                                `json:"routeKey"`
	RawPath               string                                `json:"rawPath"`
	RawQueryString        string                                `json:"rawQueryString"`
	Cookies               []string                              `json:"cookies"`
	Headers               map[string]string                     `json:"headers"`
	QueryStringParameters map[string]string                     `json:"queryStringParameters"`
	RequestContext        events.APIGatewayV2HTTPRequestContext `json:"requestContext"`
}

func (r AuthorizerV2Requests) GetJWTToken() (string, error) {
	if _, ok := r.Headers["authorization"]; !ok {
		return "", errors.New("request does not contain authorizazion header")
	}

	// remove Bearer scheme from token
	header := r.Headers["authorization"]
	header = strings.Replace(header, "Bearer", "", -1)
	header = strings.TrimSpace(header)
	return header, nil
}
