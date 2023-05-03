package handlers

import (
	"context"
	"encoding/json"

	"github.com/aws/aws-lambda-go/events"
	"test-go-lambda/services"
)

func GetUser(ctx context.Context, request events.APIGatewayProxyRequest, userService services.UserService) (*events.APIGatewayProxyResponse, error) {
	id := request.PathParameters["id"]

	user, err := userService.GetUserByID(id)
	if err != nil {
		return nil, err
	}

	userJSON, _ := json.Marshal(user)

	return &events.APIGatewayProxyResponse{
		StatusCode: 200,
		Body:       string(userJSON),
	}, nil
}
