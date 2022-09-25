package userhandler

import (
	entity "echo-boilerplate/domains/users/entities"
	"echo-boilerplate/exceptions"
	"echo-boilerplate/utils/helpers"
	"errors"
	"net/http"
	"strconv"

	"github.com/labstack/echo/v4"
)

type userHandler struct {
	Usecase entity.IusecaseUser
}

func New(usecase entity.IusecaseUser) *userHandler {
	return &userHandler{
		Usecase: usecase,
	}
}

func (h *userHandler) Store(c echo.Context) error {
	request := Request{}
	err := c.Bind(&request)
	if err != nil {
		return errors.New(exceptions.BadRequestErrorMsg)
	}

	err = c.Validate(&request)
	if err != nil {
		return err
	}

	h.Usecase.Store(RequestToEntity(request))

	return c.JSON(http.StatusOK, helpers.SuccessGetResponse("success create data"))
}

func (h *userHandler) Update(c echo.Context) error {
	id, err := strconv.Atoi(c.Param("id"))
	request := Request{}
	if err != nil {
		return errors.New(exceptions.BadRequestErrorMsg)
	}

	err = c.Bind(&request)
	if err != nil {
		return errors.New(exceptions.BadRequestErrorMsg)
	}

	err = c.Validate(request)
	if err != nil {
		return err
	}

	userEntity := RequestToEntity(request)
	userEntity.UID = uint(id)

	h.Usecase.Update(userEntity)

	return c.JSON(http.StatusOK, helpers.SuccessGetResponse("success update data"))
}

func (h *userHandler) Delete(c echo.Context) error {
	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		return errors.New(exceptions.BadRequestErrorMsg)
	}

	h.Usecase.Delete(entity.UserEntity{
		UID: uint(id),
	})

	return c.JSON(http.StatusOK, helpers.SuccessGetResponse("success delete data"))
}

func (h *userHandler) GetById(c echo.Context) error {
	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		return errors.New(exceptions.BadRequestErrorMsg)
	}

	user := h.Usecase.GetById(entity.UserEntity{
		UID: uint(id),
	})

	return c.JSON(http.StatusOK, helpers.SuccessGetResponseData("success get data", EntityToResponse(user)))
}
