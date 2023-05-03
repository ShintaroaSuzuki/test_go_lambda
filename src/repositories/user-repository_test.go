package repositories_test

import (
	"testing"

	"test-go-lambda/models"
	"test-go-lambda/repositories"
)

func TestUserRepository_FindById(t *testing.T) {
	// テスト用のIDを指定
	id := "1"

	// UserRepositoryのインスタンスを生成
	repo := &repositories.UserRepositoryImpl{}

	// UserRepositoryのFindByIdメソッドを呼び出し
	user, err := repo.FindByID(id)

	// 期待するユーザー情報
	expectedUser := &models.User{ID: id, Username: "test"}

	// エラーが発生した場合には、テストを失敗させる
	if err != nil {
		t.Fatalf("failed to execute FindById: %v", err)
	}

	// 実際のユーザー情報が期待する値と異なる場合には、テストを失敗させる
	if user.ID != expectedUser.ID {
		t.Errorf("unexpected user id: got %s, want %s", user.ID, expectedUser.ID)
	}
	if user.Username != expectedUser.Username {
		t.Errorf("unexpected username: got %s, want %s", user.Username, expectedUser.Username)
	}
}
