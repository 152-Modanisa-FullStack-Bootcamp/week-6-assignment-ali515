package controller_test

import (
	"bootcamp/config"
	"bootcamp/controller"
	mock_handler "bootcamp/mocks/handler"
	"bytes"
	"net/http"
	"net/http/httptest"
	"os"
	"testing"

	"github.com/golang/mock/gomock"
	"github.com/stretchr/testify/assert"
)

func init() {
	os.Chdir("..")
	config.Initialize()
}

func TestRouting(t *testing.T) {
	t.Run("Method not allowed for get all", func(t *testing.T) {
		mockController := gomock.NewController(t)
		mockHandler := mock_handler.NewMockIHandler(mockController)
		controller := controller.NewController(mockHandler)
		req := httptest.NewRequest(http.MethodPost, "/", bytes.NewBuffer([]byte("")))
		res := httptest.NewRecorder()
		controller.Routing(res, req)
		result := res.Result().StatusCode
		assert.Equal(t, http.StatusMethodNotAllowed, result)
	})

	t.Run("Get all", func(t *testing.T) {
		mockController := gomock.NewController(t)
		mockHandler := mock_handler.NewMockIHandler(mockController)
		req := httptest.NewRequest(http.MethodGet, "/", bytes.NewBuffer([]byte("")))
		res := httptest.NewRecorder()
		mockHandler.EXPECT().GetAll(res, req)
		controller := controller.NewController(mockHandler)
		controller.Routing(res, req)
		result := res.Result().StatusCode
		assert.Equal(t, http.StatusOK, result)
	})

	t.Run("Not found", func(t *testing.T) {
		mockController := gomock.NewController(t)
		mockHandler := mock_handler.NewMockIHandler(mockController)
		req := httptest.NewRequest(http.MethodGet, "/username/test", bytes.NewBuffer([]byte("")))
		res := httptest.NewRecorder()
		controller := controller.NewController(mockHandler)
		controller.Routing(res, req)
		result := res.Result().StatusCode
		assert.Equal(t, http.StatusNotFound, result)
	})

	t.Run("Get", func(t *testing.T) {
		mockController := gomock.NewController(t)
		mockHandler := mock_handler.NewMockIHandler(mockController)
		req := httptest.NewRequest(http.MethodGet, "/username", bytes.NewBuffer([]byte("")))
		res := httptest.NewRecorder()
		mockHandler.EXPECT().Get(res, req)
		controller := controller.NewController(mockHandler)
		controller.Routing(res, req)
		result := res.Result().StatusCode
		assert.Equal(t, http.StatusOK, result)
	})

	t.Run("Create", func(t *testing.T) {
		mockController := gomock.NewController(t)
		mockHandler := mock_handler.NewMockIHandler(mockController)
		req := httptest.NewRequest(http.MethodPut, "/username", bytes.NewBuffer([]byte("")))
		res := httptest.NewRecorder()
		mockHandler.EXPECT().Create(res, req)
		controller := controller.NewController(mockHandler)
		controller.Routing(res, req)
		result := res.Result().StatusCode
		assert.Equal(t, http.StatusOK, result)
	})

	t.Run("Update", func(t *testing.T) {
		mockController := gomock.NewController(t)
		mockHandler := mock_handler.NewMockIHandler(mockController)
		req := httptest.NewRequest(http.MethodPost, "/username", bytes.NewBuffer([]byte("")))
		res := httptest.NewRecorder()
		mockHandler.EXPECT().Update(res, req)
		controller := controller.NewController(mockHandler)
		controller.Routing(res, req)
		result := res.Result().StatusCode
		assert.Equal(t, http.StatusOK, result)
	})
}
