package service_test

import (
	"bootcamp/config"
	mock_repository "bootcamp/mocks/repository"
	"bootcamp/model"
	"bootcamp/service"
	"os"
	"testing"

	"github.com/golang/mock/gomock"
	"github.com/stretchr/testify/assert"
)

func init() {
	os.Chdir("..")
	config.Initialize()
}

func TestGetCustomerByUsername(t *testing.T) {
	mockController := gomock.NewController(t)
	mockRepo := mock_repository.NewMockIRepository(mockController)
	entity := &model.Customer{Username: "username", Balance: 100}
	mockRepo.EXPECT().GetCustomerByUsername(entity.Username).Return(entity, nil)
	service := service.NewService(mockRepo)
	entityResult, err := service.GetCustomerByUsername(entity.Username)
	assert.Nil(t, err)
	assert.Equal(t, entity, entityResult)
}

func TestGetAllCustomers(t *testing.T) {
	mockController := gomock.NewController(t)
	mockRepo := mock_repository.NewMockIRepository(mockController)
	list := &[]model.Customer{{Username: "username", Balance: 100}}
	mockRepo.EXPECT().Customers().Return(list)
	service := service.NewService(mockRepo)
	listResult := service.GetAllCustomers()
	assert.Equal(t, list, listResult)
}

func TestCreateCustomer(t *testing.T) {
	t.Run("Nil customer create", func(t *testing.T) {
		mockController := gomock.NewController(t)
		mockRepo := mock_repository.NewMockIRepository(mockController)
		service := service.NewService(mockRepo)
		err := service.CreateCustomer(nil)
		assert.NotNil(t, err)
	})

	t.Run("Customer create with init balance", func(t *testing.T) {
		mockController := gomock.NewController(t)
		mockRepo := mock_repository.NewMockIRepository(mockController)
		entityRaw := model.Customer{Username: "username", Balance: 100}
		entity := model.Customer{Username: "username", Balance: config.Env.InitialBalanceAmount}
		mockRepo.EXPECT().Add(entity).Return()
		service := service.NewService(mockRepo)
		err := service.CreateCustomer(&entityRaw)
		assert.Nil(t, err)
		assert.Equal(t, entity, entityRaw)
	})
}

func TestUpdateCustomer(t *testing.T) {
	t.Run("Nil customer update", func(t *testing.T) {
		mockController := gomock.NewController(t)
		mockRepo := mock_repository.NewMockIRepository(mockController)
		service := service.NewService(mockRepo)
		err := service.UpdateCustomer(nil)
		assert.NotNil(t, err)
	})

	t.Run("Customer update", func(t *testing.T) {
		mockController := gomock.NewController(t)
		mockRepo := mock_repository.NewMockIRepository(mockController)
		entity := model.Customer{Username: "username", Balance: 100}
		mockRepo.EXPECT().Update(entity).Return(nil)
		service := service.NewService(mockRepo)
		err := service.UpdateCustomer(&entity)
		assert.Nil(t, err)
	})
}

func TestIncreaseBalance(t *testing.T) {
	t.Run("Nil customer increase balance", func(t *testing.T) {
		mockController := gomock.NewController(t)
		mockRepo := mock_repository.NewMockIRepository(mockController)
		service := service.NewService(mockRepo)
		err := service.IncreaseBalance(nil, 10)
		assert.NotNil(t, err)
	})

	t.Run("Customer increase balance", func(t *testing.T) {
		mockController := gomock.NewController(t)
		mockRepo := mock_repository.NewMockIRepository(mockController)
		entity := model.Customer{Username: "username", Balance: 100}
		entityResult := model.Customer{Username: "username", Balance: 110}
		mockRepo.EXPECT().Update(entityResult).Return(nil)
		service := service.NewService(mockRepo)
		err := service.IncreaseBalance(&entity, 10)
		assert.Nil(t, err)
		assert.Equal(t, entity, entityResult)
	})
}

func TestDescreaseBalance(t *testing.T) {
	t.Run("Nil customer descrease balance", func(t *testing.T) {
		mockController := gomock.NewController(t)
		mockRepo := mock_repository.NewMockIRepository(mockController)
		service := service.NewService(mockRepo)
		err := service.DescreaseBalance(nil, 10)
		assert.NotNil(t, err)
	})

	t.Run("Check descrease balance", func(t *testing.T) {
		mockController := gomock.NewController(t)
		mockRepo := mock_repository.NewMockIRepository(mockController)
		service := service.NewService(mockRepo)
		entity := &model.Customer{Username: "username", Balance: 100}
		err := service.DescreaseBalance(entity, 1000)
		assert.NotNil(t, err)
	})

	t.Run("Descrease balance", func(t *testing.T) {
		mockController := gomock.NewController(t)
		mockRepo := mock_repository.NewMockIRepository(mockController)
		entity := model.Customer{Username: "username", Balance: 100}
		entityResult := model.Customer{Username: "username", Balance: 90}
		mockRepo.EXPECT().Update(entityResult).Return(nil)
		service := service.NewService(mockRepo)
		err := service.DescreaseBalance(&entity, 10)
		assert.Nil(t, err)
		assert.Equal(t, entity, entityResult)
	})
}
