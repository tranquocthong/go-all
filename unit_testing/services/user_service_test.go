package service

import (
	"testing"
	"unit-testing/repo"

	"github.com/stretchr/testify/assert"
)

// Unit tesing with mock
func TestUserServiceImp_Login(t *testing.T) {
	mockUserRepo := repo.NewMockUserRepo(t)
	mockUserRepo.EXPECT().GetUser("test").Return("token_for_test", nil)

	userService := NewUserSevice(mockUserRepo)

	token, err := userService.Login("test")

	assert.Nil(t, err)
	assert.Equal(t, "token_for_test", token)
}
