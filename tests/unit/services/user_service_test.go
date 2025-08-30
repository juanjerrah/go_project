package services

import (
	"errors"
	"testing"

	"github.com/juanjerrah/go-project/internal/models"
	"github.com/juanjerrah/go-project/internal/service"
	"github.com/juanjerrah/go-project/tests/mocks"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
	"github.com/stretchr/testify/suite"
)

type UserServiceTestSuite struct {
	suite.Suite
	mockRepo *mocks.MockUserRepository
	service  service.UserService
}

func (suite *UserServiceTestSuite) SetupTest() {
	suite.mockRepo = new(mocks.MockUserRepository)
	suite.service = service.NewUserService(suite.mockRepo)
}

func (suite *UserServiceTestSuite) TestCreateUser_Success() {
	// Arrange
	req := models.CreateUserRequest{
		Name:  "John Doe",
		Email: "john@example.com",
	}

	suite.mockRepo.On("Create", mock.AnythingOfType("*models.User")).Return(nil).Run(
		func(args mock.Arguments) {
		user := args.Get(0).(*models.User)
		user.ID = 1 // Simula a geração do ID
	})

	// Act
	res, err := suite.service.CreateUser(&req)

	// Assert
	assert.NoError(suite.T(), err)
	assert.NotNil(suite.T(), res)
	assert.Equal(suite.T(), uint(1), res.ID)
	assert.Equal(suite.T(), req.Name, res.Name)
	assert.Equal(suite.T(), req.Email, res.Email)

	suite.mockRepo.AssertCalled(suite.T(), "Create", mock.AnythingOfType("*models.User"))
	suite.mockRepo.AssertNumberOfCalls(suite.T(), "Create", 1)
}

func (suite *UserServiceTestSuite) TestCreateUser_Failure() {
	// Arrange
	req := models.CreateUserRequest{
		Name:  "John Doe",
		Email: "john@example.com",
	}

	suite.mockRepo.On("Create", mock.AnythingOfType("*models.User")).Return(errors.New("error creating user"))

	// Act
	res, err := suite.service.CreateUser(&req)

	// Assert
	assert.Error(suite.T(), err)
	assert.Nil(suite.T(), res)

	suite.mockRepo.AssertCalled(suite.T(), "Create", mock.AnythingOfType("*models.User"))
	suite.mockRepo.AssertNumberOfCalls(suite.T(), "Create", 1)
}

func TestUserServiceTestSuite(t *testing.T) {
    suite.Run(t, new(UserServiceTestSuite))
}
