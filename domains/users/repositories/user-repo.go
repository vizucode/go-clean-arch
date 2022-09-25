package userrepo

import (
	entity "echo-boilerplate/domains/users/entities"
	model "echo-boilerplate/domains/users/models"
	"echo-boilerplate/exceptions"
)

type userRepo struct {
	DB []model.User
}

func New(db []model.User) *userRepo {
	return &userRepo{
		DB: db,
	}
}

func (r *userRepo) Store(userEntity entity.UserEntity) {
	id := len(r.DB) + 1
	userModel := model.EntityToModel(userEntity)
	userModel.ID = uint(id)
	r.DB = append(r.DB, userModel)
}

func (r *userRepo) Update(userEntity entity.UserEntity) {
	userIndex := -1
	userModel := model.EntityToModel(userEntity)

	for index, product := range r.DB {
		if product.ID == userModel.ID {
			userIndex = index
		}
	}

	if userIndex == -1 {
		panic(exceptions.NotFoundErrorMsg)
	}

	r.DB[userIndex] = userModel
}

func (r *userRepo) Delete(userEntity entity.UserEntity) {
	userIndex := -1

	for index, product := range r.DB {
		if product.ID == userEntity.UID {
			userIndex = index
		}
	}

	if userIndex == -1 {
		panic(exceptions.NotFoundErrorMsg)
	}

	userArr1 := r.DB[0:userIndex]
	userArr2 := r.DB[userIndex+1:]

	r.DB = append(userArr1, userArr2...)
}

func (r *userRepo) FindById(userEntity entity.UserEntity) entity.UserEntity {
	userIndex := -1

	for index, product := range r.DB {
		if product.ID == userEntity.UID {
			userIndex = index
		}
	}

	if userIndex == -1 {
		panic(exceptions.NotFoundErrorMsg)
	}

	return model.ModelToEntity(r.DB[userIndex])
}
