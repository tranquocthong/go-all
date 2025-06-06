package service

import (
	"testing"
	"unit-testing/repo"

	"github.com/stretchr/testify/suite"
)

// Testify test suite
// A group of related unit test

type UserServiceTestSuite struct {
	suite.Suite
	userRepo    *repo.MockUserRepo
	userService UserSevice
}

// Dùng chung cho all-test
func (suite *UserServiceTestSuite) SetupSuite() {
	suite.userRepo = repo.NewMockUserRepo(suite.T())
	suite.T().Log("SetupSuite")
}

// Finally clean up
func (suite *UserServiceTestSuite) TearDownSuite() {
	suite.userRepo = nil
	suite.T().Log("TearDownSuite")
}

// Dùng riêng từng test
func (suite *UserServiceTestSuite) SetupTest() {
	suite.userService = NewUserSevice(suite.userRepo)
	suite.T().Log("SetupTest")
}

// Finally clean up per test
func (suite *UserServiceTestSuite) TearDownTest() {
	suite.userService = nil
	suite.T().Log("TearDownTest")
}

func (suite *UserServiceTestSuite) TestLogin_OK() {
	suite.userRepo.EXPECT().GetUser("test").Return("token_test", nil)

	token, err := suite.userService.Login("test")
	suite.Nil(err)
	suite.Equal("token_test", token)
}

func (suite *UserServiceTestSuite) TestRegister_OK() {
	suite.userRepo.EXPECT().AddUser("test", "password").Return(nil)
	err := suite.userService.Register("test", "password")
	suite.Nil(err)
}

func (suite *UserServiceTestSuite) TestLogout_OK() {
	suite.userRepo.EXPECT().GetUser("test").Return("token_test", nil)
	err := suite.userService.Logout("test")
	suite.Nil(err)
}

func TestUserServiceSuite(t *testing.T) {
	suite.Run(t, new(UserServiceTestSuite))
}
