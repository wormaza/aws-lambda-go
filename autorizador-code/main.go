package main

import (
	"context"
	"errors"

	"github.com/aws/aws-lambda-go/events"
	"github.com/aws/aws-lambda-go/lambda"
)

// Help function to generate an IAM policy
func generatePolicy(principalId, effect, resource string, custommessage string) events.APIGatewayCustomAuthorizerResponse {
	authResponse := events.APIGatewayCustomAuthorizerResponse{PrincipalID: principalId}

	if effect != "" && resource != "" {
		authResponse.PolicyDocument = events.APIGatewayCustomAuthorizerPolicy{
			Version: "2012-10-17",
			Statement: []events.IAMPolicyStatement{
				{
					Action:   []string{"execute-api:Invoke"},
					Effect:   effect,
					Resource: []string{resource},
				},
			},
		}
	}

	// Optional output with custom properties of the String, Number or Boolean type.
	authResponse.Context = map[string]interface{}{
		"stringKey":  custommessage,
		"numberKey":  123,
		"booleanKey": true,
	}
	return authResponse
}

func handleRequest(ctx context.Context, event events.APIGatewayCustomAuthorizerRequest) (events.APIGatewayCustomAuthorizerResponse, error) {

	token := event.AuthorizationToken

	if len(token) == 0 {
		return events.APIGatewayCustomAuthorizerResponse{}, errors.New("Unauthorized") // Return a 401 Unauthorized response}
	}

	switch token {
	case "token_a_validar_tipo_key":
		return generatePolicy("user", "Allow", event.MethodArn, ""), nil
	default:
		return generatePolicy("user", "Deny", event.MethodArn, "Error: Invalid token"), nil
	}
}

func main() {
	lambda.Start(handleRequest)
}
