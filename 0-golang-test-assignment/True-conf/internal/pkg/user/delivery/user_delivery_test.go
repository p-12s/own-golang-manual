package delivery

import (
	"encoding/json"
	"errors"
	"github.com/golang/mock/gomock"
	"github.com/labstack/echo/v4"
	"github.com/stretchr/testify/assert"
	"net/http"
	"net/http/httptest"
	"strings"
	"testing"
	"userService/internal/pkg/models"
	mock_repository "userService/internal/pkg/user/repository/mocks"
	"userService/internal/pkg/user/usecase"
)

func TestUserDelivery_CreateUser_Success(t *testing.T) {
	//creating testdata
	user := models.User{
		Id:   1,
		Name: "Jon Snow",
	}

	//creating mocks
	mockCtrl := gomock.NewController(t)
	dbMock := mock_repository.NewMockUserRepositoryInterface(mockCtrl)
	dbMock.EXPECT().CreateUser(models.User{Name: "Jon Snow"}).Return(user.Id, nil)
	UserUsecaseTest := usecase.UserUsecase{Repository: dbMock}
	UserDeliveryTest := UserDelivery{Usecase: UserUsecaseTest}

	//creating server
	e := echo.New()
	req := httptest.NewRequest(http.MethodPost, "/users", strings.NewReader(`{"name":"Jon Snow"}`))
	req.Header.Set(echo.HeaderContentType, echo.MIMEApplicationJSON)
	rec := httptest.NewRecorder()

	//send request
	c := e.NewContext(req, rec)
	err := UserDeliveryTest.CreateUser(c)
	if err != nil {
		t.Fatal("cant send request : " + err.Error())
	}

	// check response
	if http.StatusOK != rec.Code {
		t.Fatal(rec.Code)
	}

	body := rec.Result().Body
	var reqUser models.User
	decoder := json.NewDecoder(body)
	err = decoder.Decode(&reqUser)
	if err != nil {
		t.Error("cant parse body : " + err.Error())
	}
	assert.Equal(t, user, reqUser, "User and User after request not same!")
}

func TestUserDelivery_CreateUser_Fail(t *testing.T) {
	//creating testdata
	user := models.User{
		Id:   1,
		Name: "Jon Snow",
	}

	//creating mocks
	mockCtrl := gomock.NewController(t)
	dbMock := mock_repository.NewMockUserRepositoryInterface(mockCtrl)
	dbMock.EXPECT().CreateUser(models.User{Name: "Jon Snow"}).Return(user.Id, errors.New("cant add post"))
	UserUsecaseTest := usecase.UserUsecase{Repository: dbMock}
	UserDeliveryTest := UserDelivery{Usecase: UserUsecaseTest}

	//creating server
	e := echo.New()
	req := httptest.NewRequest(http.MethodPost, "/users", strings.NewReader(`{"name":"Jon Snow"}`))
	req.Header.Set(echo.HeaderContentType, echo.MIMEApplicationJSON)
	rec := httptest.NewRecorder()

	//send request
	c := e.NewContext(req, rec)
	err := UserDeliveryTest.CreateUser(c)
	if err != nil {
		t.Fatal("cant send request : " + err.Error())
	}

	// check response
	if http.StatusBadRequest != rec.Code {
		t.Fatal(rec.Code)
	}
}

func TestUserDelivery_ChangeUserById_Success(t *testing.T) {
	//creating testdata
	user := models.User{
		Id:   1,
		Name: "New Jon Snow",
	}

	//creating mocks
	mockCtrl := gomock.NewController(t)
	dbMock := mock_repository.NewMockUserRepositoryInterface(mockCtrl)
	dbMock.EXPECT().ChangeUser(user).Return(nil)
	UserUsecaseTest := usecase.UserUsecase{Repository: dbMock}
	UserDeliveryTest := UserDelivery{Usecase: UserUsecaseTest}

	//creating server
	e := echo.New()
	req := httptest.NewRequest(http.MethodPut, "/users", strings.NewReader(`{"name":"New Jon Snow"}`))
	req.Header.Set(echo.HeaderContentType, echo.MIMEApplicationJSON)
	rec := httptest.NewRecorder()

	//send request
	c := e.NewContext(req, rec)
	c.SetParamNames("id")
	c.SetParamValues("1")
	err := UserDeliveryTest.ChangeUserById(c)
	if err != nil {
		t.Fatal("cant send request : " + err.Error())
	}

	// check response
	if http.StatusOK != rec.Code {
		t.Fatal(rec.Code)
	}

	body := rec.Result().Body
	var reqUser models.User
	decoder := json.NewDecoder(body)
	err = decoder.Decode(&reqUser)
	if err != nil {
		t.Error("cant parse body : " + err.Error())
	}
	assert.Equal(t, user, reqUser, "User and User after request not same!")
}

func TestUserDelivery_ChangeUserById_Fail(t *testing.T) {
	//creating mocks
	mockCtrl := gomock.NewController(t)
	dbMock := mock_repository.NewMockUserRepositoryInterface(mockCtrl)
	UserUsecaseTest := usecase.UserUsecase{Repository: dbMock}
	UserDeliveryTest := UserDelivery{Usecase: UserUsecaseTest}

	//creating server
	e := echo.New()
	req := httptest.NewRequest(http.MethodPut, "/users", strings.NewReader(`{"name":"New Jon Snow"}`))
	req.Header.Set(echo.HeaderContentType, echo.MIMEApplicationJSON)
	rec := httptest.NewRecorder()

	//send request
	c := e.NewContext(req, rec)
	err := UserDeliveryTest.ChangeUserById(c)
	if err != nil {
		t.Fatal("cant send request : " + err.Error())
	}

	// check response
	if http.StatusBadRequest != rec.Code {
		t.Fatal(rec.Code)
	}
}

func TestUserDelivery_DeleteUser_Success(t *testing.T) {
	//creating mocks
	mockCtrl := gomock.NewController(t)
	dbMock := mock_repository.NewMockUserRepositoryInterface(mockCtrl)
	dbMock.EXPECT().DeleteUser(1).Return(nil)
	UserUsecaseTest := usecase.UserUsecase{Repository: dbMock}
	UserDeliveryTest := UserDelivery{Usecase: UserUsecaseTest}

	//creating server
	e := echo.New()
	req := httptest.NewRequest(http.MethodDelete, "/users", nil)
	req.Header.Set(echo.HeaderContentType, echo.MIMEApplicationJSON)
	rec := httptest.NewRecorder()

	//send request
	c := e.NewContext(req, rec)
	c.SetParamNames("id")
	c.SetParamValues("1")
	err := UserDeliveryTest.DeleteUser(c)
	if err != nil {
		t.Fatal("cant send request : " + err.Error())
	}

	// check response
	if http.StatusOK != rec.Code {
		t.Fatal(rec.Code)
	}
}

func TestUserDelivery_DeleteUser_Fail(t *testing.T) {
	//creating mocks
	mockCtrl := gomock.NewController(t)
	dbMock := mock_repository.NewMockUserRepositoryInterface(mockCtrl)
	UserUsecaseTest := usecase.UserUsecase{Repository: dbMock}
	UserDeliveryTest := UserDelivery{Usecase: UserUsecaseTest}

	//creating server
	e := echo.New()
	req := httptest.NewRequest(http.MethodDelete, "/users", nil)
	req.Header.Set(echo.HeaderContentType, echo.MIMEApplicationJSON)
	rec := httptest.NewRecorder()

	//send request
	c := e.NewContext(req, rec)
	err := UserDeliveryTest.DeleteUser(c)
	if err != nil {
		t.Fatal("cant send request : " + err.Error())
	}

	// check response
	if http.StatusBadRequest != rec.Code {
		t.Fatal(rec.Code)
	}
}

func TestUserDelivery_GetUserById(t *testing.T) {
	//creating testdata
	user := models.User{
		Id:   1,
		Name: "Jon Snow",
	}

	//creating mocks
	mockCtrl := gomock.NewController(t)
	dbMock := mock_repository.NewMockUserRepositoryInterface(mockCtrl)
	dbMock.EXPECT().GetUserById(1).Return(user, nil)
	UserUsecaseTest := usecase.UserUsecase{Repository: dbMock}
	UserDeliveryTest := UserDelivery{Usecase: UserUsecaseTest}

	//creating server
	e := echo.New()
	req := httptest.NewRequest(http.MethodGet, "/users", nil)
	req.Header.Set(echo.HeaderContentType, echo.MIMEApplicationJSON)
	rec := httptest.NewRecorder()

	//send request
	c := e.NewContext(req, rec)
	c.SetParamNames("id")
	c.SetParamValues("1")
	err := UserDeliveryTest.GetUserById(c)
	if err != nil {
		t.Fatal("cant send request : " + err.Error())
	}

	// check response
	if http.StatusOK != rec.Code {
		t.Fatal(rec.Code)
	}

	body := rec.Result().Body
	var reqUser models.User
	decoder := json.NewDecoder(body)
	err = decoder.Decode(&reqUser)
	if err != nil {
		t.Error("cant parse body : " + err.Error())
	}
	assert.Equal(t, user, reqUser, "User and User after request not same!")
}

func TestUserDelivery_GetAllUsers(t *testing.T) {
	//creating testdata

	users := []models.User{
		{
			Id:   1,
			Name: "Jon Snow1",
		},
		{
			Id:   2,
			Name: "Jon Snow2",
		},
		{
			Id:   3,
			Name: "Jon Snow3",
		},
	}

	//creating mocks
	mockCtrl := gomock.NewController(t)
	dbMock := mock_repository.NewMockUserRepositoryInterface(mockCtrl)
	dbMock.EXPECT().GetAllUsers().Return(users, nil)
	UserUsecaseTest := usecase.UserUsecase{Repository: dbMock}
	UserDeliveryTest := UserDelivery{Usecase: UserUsecaseTest}

	//creating server
	e := echo.New()
	req := httptest.NewRequest(http.MethodGet, "/users", nil)
	req.Header.Set(echo.HeaderContentType, echo.MIMEApplicationJSON)
	rec := httptest.NewRecorder()

	//send request
	c := e.NewContext(req, rec)
	err := UserDeliveryTest.GetAllUsers(c)
	if err != nil {
		t.Fatal("cant send request : " + err.Error())
	}

	// check response
	if http.StatusOK != rec.Code {
		t.Fatal(rec.Code)
	}

	body := rec.Result().Body
	var reqUsers []models.User
	decoder := json.NewDecoder(body)
	err = decoder.Decode(&reqUsers)
	if err != nil {
		t.Error("cant parse body : " + err.Error())
	}
	assert.Equal(t, users, reqUsers, "User and User after request not same!")
}
