package service

import (
	"strconv"
	"testing"

	"github.com/matheuswww/api-rest-golang/src/configuration/rest_err"
	"github.com/matheuswww/api-rest-golang/src/model"
	"github.com/matheuswww/api-rest-golang/src/test/mocks"
	"github.com/stretchr/testify/assert"
	"go.uber.org/mock/gomock"
)


func TestUserDomainService_FindUserByIDServices(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	repository := mocks.NewMockUserRepository(ctrl)
	service := NewUserDomainService(repository)
	
	t.Run("Success",func(t *testing.T) {
		id := 1
		userDomain := model.NewUserDomain("test@test.com","test","test",uint8(40),0)	
		repository.EXPECT().FindUser("id",strconv.Itoa(id)).Return(userDomain,nil)
		userDomain.SetId(uint(id))	

		userDomainReturn,err := service.FindUser("id",strconv.Itoa(id))

		assert.Nil(t, err)
		assert.EqualValues(t, userDomainReturn.GetId(),uint(id))
		assert.EqualValues(t, userDomainReturn.GetEmail(), "test@test.com")
		assert.EqualValues(t, userDomainReturn.GetName(), "test")
		assert.EqualValues(t, userDomainReturn.GetPassword(), "test")
		assert.EqualValues(t, userDomainReturn.GetAge(), uint8(40))
	})

	t.Run("return_error_from_user_not_found",func(t *testing.T) {
		id := 1
		repository.EXPECT().FindUser("id",strconv.Itoa(id)).Return(nil,rest_err.NewNotFoundError("user not found"))

		userDomainReturn,err := service.FindUser("id",strconv.Itoa(id))

		assert.Nil(t,userDomainReturn)
		assert.NotNil(t,err)
		assert.EqualValues(t,err.Message,"user not found")
	})
}

func TestUserDomainService_FindUserByEmailServices(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	repository := mocks.NewMockUserRepository(ctrl)
	service := NewUserDomainService(repository)
	
	t.Run("Success",func(t *testing.T) {
		email := "test@test.com"
		userDomain := model.NewUserDomain("test@test.com","test","test",uint8(40),0)
		repository.EXPECT().FindUser("email",email).Return(userDomain,nil)
		userDomain.SetId(1)	

		userDomainReturn,err := service.FindUser("email",email)

		assert.Nil(t, err)
		assert.EqualValues(t, userDomainReturn.GetId(),1)
		assert.EqualValues(t, userDomainReturn.GetEmail(), email)
		assert.EqualValues(t, userDomainReturn.GetName(), "test")
		assert.EqualValues(t, userDomainReturn.GetPassword(), "test")
		assert.EqualValues(t, userDomainReturn.GetAge(), uint8(40))
	})

	t.Run("return_error_from_user_not_found",func(t *testing.T) {
		email := "test@test.com"
		repository.EXPECT().FindUser("email",email).Return(nil,rest_err.NewNotFoundError("user not found"))

		userDomainReturn,err := service.FindUser("email",email)

		assert.Nil(t,userDomainReturn)
		assert.NotNil(t,err)
		assert.EqualValues(t,err.Message,"user not found")
	})
}

func TestUserDomainService_FindUserByEmailAndPassword(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	repository := mocks.NewMockUserRepository(ctrl)
	service := &userDomainService{repository}
	
	t.Run("Success",func(t *testing.T) {
		email := "test@test.com"
		password := "test"
		userDomain := model.NewUserDomain("test@test.com","test","test",uint8(40),1)
		repository.EXPECT().FindUserByEmailAndPassword(email,password).Return(userDomain,nil)

		userDomainReturn,err := service.findUserByEmailAndPassword(email,password)

		assert.Nil(t, err)
		assert.EqualValues(t, userDomainReturn.GetId(),1)
		assert.EqualValues(t, userDomainReturn.GetEmail(), email)
		assert.EqualValues(t, userDomainReturn.GetName(), "test")
		assert.EqualValues(t, userDomainReturn.GetPassword(), "test")
		assert.EqualValues(t, userDomainReturn.GetAge(), uint8(40))
	})

	t.Run("return_error_from_user_not_found",func(t *testing.T) {
		email := "test@test.com"
		password := "test"
		repository.EXPECT().FindUserByEmailAndPassword(email,password).Return(nil,rest_err.NewNotFoundError("user not found"))

		userDomainReturn,err := service.findUserByEmailAndPassword(email,password)

		assert.Nil(t,userDomainReturn)
		assert.NotNil(t,err)
		assert.EqualValues(t,err.Message,"user not found")
	})
}