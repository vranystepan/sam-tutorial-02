package main

type AuthorizerV2ResponsePolicyStatement struct {
	Action   string `json:"Action"`
	Effect   string `json:"Effect"`
	Resource []string `json:"Resource"`
}

type AuthorizerV2ResponsePolicyDocument struct {
	Version   string                                `json:"Version"`
	Statement []AuthorizerV2ResponsePolicyStatement `json:"Statement"`
}

type AuthorizerV2Response struct {
	PrincipalID    string                             `json:"principalId"`
	PolicyDocument AuthorizerV2ResponsePolicyDocument `json:"policyDocument"`
	Context        map[string]string                  `json:"context,omitempty"`
}
