package userhandler

import (
	entity "echo-boilerplate/domains/users/entities"
)

type Response struct {
	UID   uint   `json:"id"`
	Name  string `json:"name"`
	Email string `json:"email"`
}

func EntityToResponse(userEntity entity.UserEntity) Response {
	return Response{
		UID:   userEntity.UID,
		Name:  userEntity.Name,
		Email: userEntity.Email,
	}
}
