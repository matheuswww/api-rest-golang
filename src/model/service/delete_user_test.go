package service

import (
	"testing"

	"github.com/matheuswww/api-rest-golang/src/configuration/rest_err"
	"github.com/matheuswww/api-rest-golang/src/test/mocks"
	"github.com/stretchr/testify/assert"
	"go.uber.org/mock/gomock"
)

func TestUserDomainService_DeleteUser(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	repository := mocks.NewMockUserRepository(ctrl)
	service := NewUserDomainService(repository)

	t.Run("Success", func(t *testing.T) {
		repository.EXPECT().DeleteUser("1").Return(nil)

		err := service.DeleteUser("1")

		assert.Nil(t, err)
	})

	t.Run("return_error_user_not_found", func(t *testing.T) {
		repository.EXPECT().DeleteUser("1").Return(rest_err.NewNotFoundError("user not found"))

		err := service.DeleteUser("1")

		assert.NotNil(t, err)
		assert.EqualValues(t, err.Message, "user not found")
	})
}
