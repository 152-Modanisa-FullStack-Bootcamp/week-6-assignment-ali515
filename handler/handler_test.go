package handler_test

import (
	"bootcamp/handler"
	mock_service "bootcamp/mocks/service"
	"bootcamp/model"
	"bytes"
	"encoding/json"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/golang/mock/gomock"
	"github.com/stretchr/testify/assert"
)

type ErrorTest struct {
	Message string
}

func (e *ErrorTest) Error() string {
	return e.Message
}

func TestGet(t *testing.T) {
	t.Run("Nil customers", func(t *testing.T) {
		mockController := gomock.NewController(t)
		mockService := mock_service.NewMockIService(mockController)
		mockService.EXPECT().GetCustomerByUsername("username").Return(nil, &ErrorTest{})
		handler := handler.NewHandler(mockService)
		req := httptest.NewRequest(http.MethodGet, "/username", bytes.NewBuffer([]byte("")))
		res := httptest.NewRecorder()
		handler.Get(res, req)
		actual := res.Result().StatusCode
		assert.Equal(t, http.StatusGone, actual)
	})

	t.Run("Customers", func(t *testing.T) {
		mockController := gomock.NewController(t)
		mockService := mock_service.NewMockIService(mockController)
		customer := model.Customer{Username: "username", Balance: 100}
		mockService.EXPECT().GetCustomerByUsername(customer.Username).Return(&customer, nil)
		handler := handler.NewHandler(mockService)
		req := httptest.NewRequest(http.MethodGet, "/username", bytes.NewBuffer([]byte("")))
		res := httptest.NewRecorder()
		handler.Get(res, req)
		customerResult := model.Customer{}
		_ = json.Unmarshal(res.Body.Bytes(), &customerResult)
		assert.Equal(t, customer, customerResult)
	})
}

func TestGetAll(t *testing.T) {
	t.Run("Nil customers", func(t *testing.T) {
		mockController := gomock.NewController(t)
		mockService := mock_service.NewMockIService(mockController)
		mockService.EXPECT().GetAllCustomers().Return(nil)
		handler := handler.NewHandler(mockService)
		req := httptest.NewRequest(http.MethodGet, "/", bytes.NewBuffer([]byte("")))
		res := httptest.NewRecorder()
		handler.GetAll(res, req)
		result := res.Result().StatusCode
		assert.Equal(t, http.StatusNoContent, result)
	})

	t.Run("Customers", func(t *testing.T) {
		mockController := gomock.NewController(t)
		mockService := mock_service.NewMockIService(mockController)
		mockService.EXPECT().GetAllCustomers().Return(&[]model.Customer{{Username: "username", Balance: 100}})
		handler := handler.NewHandler(mockService)
		req := httptest.NewRequest(http.MethodGet, "/", bytes.NewBuffer([]byte("")))
		res := httptest.NewRecorder()
		handler.GetAll(res, req)
		customers, _ := json.Marshal(model.CustomerResponse{model.Customer{Username: "username", Balance: 100}})
		customersResult := res.Body.Bytes()
		assert.Equal(t, customers, customersResult)
	})
}

func TestCreate(t *testing.T) {
	t.Run("Bad request", func(t *testing.T) {
		mockController := gomock.NewController(t)
		mockService := mock_service.NewMockIService(mockController)
		handler := handler.NewHandler(mockService)
		req := httptest.NewRequest(http.MethodPut, "/username", bytes.NewBuffer([]byte("")))
		res := httptest.NewRecorder()
		handler.Create(res, req)
		result := res.Result().StatusCode
		assert.Equal(t, http.StatusBadRequest, result)
	})

	t.Run("Nil customer", func(t *testing.T) {
		mockController := gomock.NewController(t)
		mockService := mock_service.NewMockIService(mockController)
		handler := handler.NewHandler(mockService)
		customer := model.Customer{}
		buf, _ := json.Marshal(customer)
		req := httptest.NewRequest(http.MethodPut, "/username", bytes.NewBuffer(buf))
		res := httptest.NewRecorder()
		handler.Create(res, req)
		result := res.Result().StatusCode
		assert.Equal(t, http.StatusNotAcceptable, result)
	})

	t.Run("Create customer", func(t *testing.T) {
		mockController := gomock.NewController(t)
		mockService := mock_service.NewMockIService(mockController)
		customer := model.Customer{Username: "username"}
		mockService.EXPECT().CreateCustomer(&customer).Return(nil)
		handler := handler.NewHandler(mockService)
		buf, _ := json.Marshal(customer)
		req := httptest.NewRequest(http.MethodPut, "/username", bytes.NewBuffer(buf))
		res := httptest.NewRecorder()
		handler.Create(res, req)
		result := res.Result().StatusCode
		assert.Equal(t, http.StatusCreated, result)
	})
}

func TestUpdate(t *testing.T) {
	t.Run("Bad request", func(t *testing.T) {
		mockController := gomock.NewController(t)
		mockService := mock_service.NewMockIService(mockController)
		handler := handler.NewHandler(mockService)
		req := httptest.NewRequest(http.MethodPost, "/username", bytes.NewBuffer([]byte("")))
		res := httptest.NewRecorder()
		handler.Update(res, req)
		result := res.Result().StatusCode
		assert.Equal(t, http.StatusBadRequest, result)
	})

	t.Run("Nil customer", func(t *testing.T) {
		mockController := gomock.NewController(t)
		mockService := mock_service.NewMockIService(mockController)
		handler := handler.NewHandler(mockService)
		customer := model.Customer{}
		buf, _ := json.Marshal(customer)
		req := httptest.NewRequest(http.MethodPost, "/username", bytes.NewBuffer(buf))
		res := httptest.NewRecorder()
		handler.Update(res, req)
		result := res.Result().StatusCode
		assert.Equal(t, http.StatusNotAcceptable, result)
	})

	t.Run("Not found customer", func(t *testing.T) {
		mockController := gomock.NewController(t)
		mockService := mock_service.NewMockIService(mockController)
		customer := model.Customer{Username: "username"}
		//mockService.EXPECT().UpdateCustomer(&customer).Return(nil)
		mockService.EXPECT().GetCustomerByUsername("username").Return(nil, &ErrorTest{})
		handler := handler.NewHandler(mockService)
		buf, _ := json.Marshal(customer)
		req := httptest.NewRequest(http.MethodPost, "/username", bytes.NewBuffer(buf))
		res := httptest.NewRecorder()
		handler.Update(res, req)
		result := res.Result().StatusCode
		assert.Equal(t, http.StatusGone, result)
	})

	t.Run("Update customer", func(t *testing.T) {
		mockController := gomock.NewController(t)
		mockService := mock_service.NewMockIService(mockController)
		customer := model.Customer{Username: "username"}
		mockService.EXPECT().UpdateCustomer(&customer).Return(nil)
		mockService.EXPECT().GetCustomerByUsername("username").Return(&customer, nil)
		handler := handler.NewHandler(mockService)
		buf, _ := json.Marshal(customer)
		req := httptest.NewRequest(http.MethodPost, "/username", bytes.NewBuffer(buf))
		res := httptest.NewRecorder()
		handler.Update(res, req)
		result := res.Result().StatusCode
		assert.Equal(t, http.StatusAccepted, result)
	})
}
