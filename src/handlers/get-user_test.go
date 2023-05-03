package handlers_test

import (
	"context"
	"encoding/json"
	"errors"
	"testing"

	"github.com/aws/aws-lambda-go/events"
	"test-go-lambda/handlers"
	"test-go-lambda/models"
	"test-go-lambda/services"
)

type mockUserRepository struct{}

func (repo mockUserRepository) FindByID(id string) (*models.User, error) {
	if id == "1234" {
		return &models.User{
			ID:       "1234",
			Username: "testuser",
		}, nil
	}
	return nil, errors.New("User not found")
}

func TestGetUser(t *testing.T) {
	ctx := context.Background()

	t.Run("Get user by ID", func(t *testing.T) {
		request := events.APIGatewayProxyRequest{
			PathParameters: map[string]string{
				"id": "1234",
			},
		}
		repo := mockUserRepository{}
		service := services.NewUserService(repo)

		response, err := handlers.GetUser(ctx, request, service)
		if err != nil {
			t.Errorf("unexpected error: %v", err)
		}

		expectedBody, _ := json.Marshal(&models.User{
			ID:       "1234",
			Username: "testuser",
		})
		expectedResponse := events.APIGatewayProxyResponse{
			StatusCode: 200,
			Body:       string(expectedBody),
		}

		if response.StatusCode != expectedResponse.StatusCode {
			t.Errorf("status code is %v, want %v", response.StatusCode, expectedResponse.StatusCode)
		}
		if response.Body != expectedResponse.Body {
			t.Errorf("response body is %v, want %v", response.Body, expectedResponse.Body)
		}
	})

	t.Run("User not found", func(t *testing.T) {
		request := events.APIGatewayProxyRequest{
			PathParameters: map[string]string{
				"id": "5678",
			},
		}
		repo := mockUserRepository{}
		service := services.NewUserService(repo)

		_, err := handlers.GetUser(ctx, request, service)
		if err == nil {
			t.Error("expected error, but not occurred")
		}
	})
}
