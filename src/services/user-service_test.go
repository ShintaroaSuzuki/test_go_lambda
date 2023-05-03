package services_test

import (
	"errors"
	"testing"

	"test-go-lambda/models"
	"test-go-lambda/services"
)

type MockUserRepository struct{}

func (m *MockUserRepository) FindByID(id string) (*models.User, error) {
	if id == "1" {
		return &models.User{ID: "1", Username: "test_user_1"}, nil
	} else if id == "2" {
		return &models.User{ID: "2", Username: "test_user_2"}, nil
	} else {
		return nil, errors.New("User not found")
	}
}

func TestGetUser(t *testing.T) {
	// モックを作成
	repo := &MockUserRepository{}

	// テスト対象の関数を実行
	userService := services.NewUserService(repo)
	user, err := userService.GetUserByID("1")

	// 結果を検証
	if err != nil {
		t.Errorf("unexpected error: %v", err)
	}
	if user.ID != "1" {
		t.Errorf("user id is incorrect: expected=%s, actual=%s", "1", user.ID)
	}
	if user.Username != "test_user_1" {
		t.Errorf("username is incorrect: expected=%s, actual=%s", "test_user_1", user.Username)
	}

	// ユーザーが存在しない場合のテスト
	user, err = userService.GetUserByID("3")
	if err == nil {
		t.Error("error is expected but not returned")
	}
	if user != nil {
		t.Errorf("unexpected user: %v", user)
	}
}
