package delivery

import (
	"github.com/labstack/echo/v4"
	"net/http"
	"strconv"
	"userService/internal/pkg/models"
	"userService/internal/pkg/user/usecase"
)

type UserDeliveryInterface interface {
	CreateUser(c echo.Context) error
	GetAllUsers(c echo.Context) error
	GetUserById(c echo.Context) error
	ChangeUserById(c echo.Context) error
	DeleteUser(c echo.Context) error
}

func (u UserDelivery) SetRoutersForUser(e *echo.Echo) {
	e.POST("/users", u.CreateUser)
	e.GET("/users", u.GetAllUsers)
	e.GET("/users/:id", u.GetUserById)
	e.PUT("/users/:id", u.ChangeUserById)
	e.DELETE("/users/:id", u.DeleteUser)
}

type UserDelivery struct {
	Usecase usecase.UserUsecaseInterface
}

func (u UserDelivery) CreateUser(c echo.Context) error {
	reqUser := new(models.User)
	err := c.Bind(&reqUser)
	if err != nil {
		c.Logger().Info("Create User - " + err.Error())
		return c.String(http.StatusBadRequest, "Invalid Input")
	}

	user, err := u.Usecase.CreateUser(*reqUser)
	if err != nil {
		c.Logger().Info("Create User - " + err.Error())
		return c.String(http.StatusBadRequest, "Invalid Input")
	}

	return c.JSON(http.StatusOK, user)
}

func (u UserDelivery) GetAllUsers(c echo.Context) error {
	users, err := u.Usecase.GetAllUsers()

	if err != nil {
		c.Logger().Info("GetAllUser - " + err.Error())
		return c.String(http.StatusBadRequest, "Error")
	}

	return c.JSON(http.StatusOK, users)
}

func (u UserDelivery) GetUserById(c echo.Context) error {
	id := c.Param("id")
	intId, err := strconv.Atoi(id)
	if err != nil {
		return c.String(http.StatusBadRequest, "Invalid Id")
	}

	user, err := u.Usecase.GetUserById(intId)
	if err != nil {
		c.Logger().Info("Get User - " + err.Error())
		return c.String(http.StatusNotFound, "Cant Find User")
	}

	return c.JSON(http.StatusOK, user)
}

func (u UserDelivery) ChangeUserById(c echo.Context) error {
	id := c.Param("id")
	intId, err := strconv.Atoi(id)
	if err != nil {
		return c.String(http.StatusBadRequest, "Invalid Id")
	}

	reqUser := new(models.User)
	err = c.Bind(&reqUser)
	if err != nil {
		c.Logger().Info("Change User - " + err.Error())
		return c.String(http.StatusBadRequest, "Invalid Input")
	}

	user, err := u.Usecase.ChangeUser(*reqUser, intId)
	if err != nil {
		c.Logger().Info("Change User - " + err.Error())
		return c.String(http.StatusBadRequest, "Invalid Input")
	}

	return c.JSON(http.StatusOK, user)
}

func (u UserDelivery) DeleteUser(c echo.Context) error {
	id := c.Param("id")
	intId, err := strconv.Atoi(id)
	if err != nil {
		return c.String(http.StatusBadRequest, "Invalid Id")
	}

	err = u.Usecase.DeleteUser(intId)
	if err != nil {
		c.Logger().Info("Get User - " + err.Error())
		return c.String(http.StatusNotFound, "Cant Find User")
	}

	return c.NoContent(http.StatusOK)
}
