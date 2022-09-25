package userentity

type IrepoUser interface {
	Store(userEntity UserEntity)
	Update(userEntity UserEntity)
	Delete(userEntity UserEntity)
	FindById(userEntity UserEntity) UserEntity
}

type IusecaseUser interface {
	Store(userEntity UserEntity)
	Update(userEntity UserEntity)
	Delete(userEntity UserEntity)
	GetById(userEntity UserEntity) UserEntity
}
