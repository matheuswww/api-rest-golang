package service

import (
	"testing"

	"github.com/matheuswww/api-rest-golang/src/configuration/rest_err"
	"github.com/matheuswww/api-rest-golang/src/model"
	"github.com/matheuswww/api-rest-golang/src/test/mocks"
	"github.com/stretchr/testify/assert"
	"go.uber.org/mock/gomock"
)

func TestUserDomainService_UpdateUser(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	repository := mocks.NewMockUserRepository(ctrl)
	service := NewUserDomainService(repository)
	
	t.Run("Success",func(t *testing.T) {
		userDomain := model.NewUserDomain("test@test.com","test","test",uint8(40),1)
		repository.EXPECT().UpdateUser("1",userDomain).Return(nil)

		err := service.UpdateUser("1",userDomain)

		assert.Nil(t, err)
	})

	t.Run("return_error_user_not_found",func(t *testing.T) {
		userDomain := model.NewUserDomain("test@test.com","test","test",uint8(40),1)
		repository.EXPECT().UpdateUser("1",userDomain).Return(rest_err.NewNotFoundError("user not found"))

		err := service.UpdateUser("1",userDomain)

		assert.NotNil(t, err)
		assert.EqualValues(t,err.Message,"user not found")
	})
}