package view

import (
	"github.com/matheuswww/api-rest-golang/src/controller/model/response"
	"github.com/matheuswww/api-rest-golang/src/model"
)

func ConvertDomainToResponse(userDomain model.UserDomainInterface) response.UserResponse{
	return response.UserResponse{
		Name: userDomain.GetName(),
		Email: userDomain.GetEmail(),
		Age: userDomain.GetAge(),
		Id: userDomain.GetId(),
	}
}