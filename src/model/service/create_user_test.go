package service

import (
	"testing"

	"github.com/matheuswww/api-rest-golang/src/configuration/rest_err"
	"github.com/matheuswww/api-rest-golang/src/model"
	"github.com/matheuswww/api-rest-golang/src/test/mocks"
	"github.com/stretchr/testify/assert"
	"go.uber.org/mock/gomock"
)

func TestUserDomainService_CreateUserServices(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	repository := mocks.NewMockUserRepository(ctrl)
	service := NewUserDomainService(repository)

	t.Run("Success", func(t *testing.T) {
		id := 1

		userDomain := model.NewUserDomain("test@test.com", "test", "test", 50, 0)
		userDomain.SetId(uint(id))

		repository.EXPECT().CreateUser(userDomain).Return(userDomain, nil)

		user, err := service.CreateUser(userDomain)

		assert.NotNil(t, user)
		assert.Nil(t, err)
		assert.EqualValues(t, userDomain.GetId(), uint(1))
		assert.EqualValues(t, userDomain.GetEmail(), "test@test.com")
		assert.EqualValues(t, userDomain.GetName(), "test")
		assert.EqualValues(t, userDomain.GetAge(), 50)
	})

	t.Run("return_error_duplicated_email", func(t *testing.T) {
		id := 1

		userDomain := model.NewUserDomain("test@test.com", "test", "test", 50, 0)
		userDomain.SetId(uint(id))

		repository.EXPECT().CreateUser(userDomain).Return(nil, rest_err.NewConflictError("Email is already registered in another account"))

		user, err := service.CreateUser(userDomain)

		assert.Nil(t, user)
		assert.NotNil(t, err)
		assert.EqualValues(t, err.Message, "Email is already registered in another account")
	})

	t.Run("database_error", func(t *testing.T) {
		id := 1

		userDomain := model.NewUserDomain("test@test.com", "test", "test", 50, 0)
		userDomain.SetId(uint(id))

		repository.EXPECT().CreateUser(userDomain).Return(nil, rest_err.NewInternalServerError("database error"))

		user, err := service.CreateUser(userDomain)

		assert.Nil(t, user)
		assert.NotNil(t, err)
		assert.EqualValues(t, err.Message, "database error")
	})
}
