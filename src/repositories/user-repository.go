package repositories

import (
	"test-go-lambda/models"
)

type UserRepository interface {
	FindByID(id string) (*models.User, error)
}

type UserRepositoryImpl struct{}

func (repo UserRepositoryImpl) FindByID(id string) (*models.User, error) {
	// ここにデータベースからユーザー情報を取得する処理を実装する
	return &models.User{
		ID:       id,
		Username: "test",
	}, nil
}
