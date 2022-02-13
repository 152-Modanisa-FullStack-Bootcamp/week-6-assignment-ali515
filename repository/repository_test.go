package repository

import (
	"bootcamp/model"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestCustomers(t *testing.T) {
	repo := NewRepository()
	customers := &[]model.Customer{}
	customersResult := repo.Customers()
	assert.Equal(t, customers, customersResult)
}

func TestGetCustomerByUsername(t *testing.T) {
	t.Run("Customer not found", func(t *testing.T) {
		db := []model.Customer{}
		repo := repository{db: db}
		customer, err := repo.GetCustomerByUsername("")
		assert.NotNil(t, err)
		assert.Nil(t, customer)
	})

	t.Run("Customer find", func(t *testing.T) {
		entity := model.Customer{Username: "username", Balance: 100}
		db := []model.Customer{entity}
		repo := repository{db: db}
		customer, err := repo.GetCustomerByUsername(entity.Username)
		assert.Equal(t, &entity, customer)
		assert.Nil(t, err)
	})
}

func TestAdd(t *testing.T) {
	t.Run("Add customer", func(t *testing.T) {
		entity := model.Customer{Username: "username", Balance: 100}
		db := []model.Customer{}
		repo := repository{db: db}
		repo.Add(entity)
		assert.Equal(t, 1, len(repo.db))
	})

	/*
		t.Run("Check Added customer initial balance", func(t *testing.T) {
			db := []model.Customer{}
			repo := repository{db: db}
			repo.Add(model.Customer{Username: "username", Balance: 100})
			entity := repo.db[0]
			assert.Equal(t, float64(0), entity.Balance)
		})
	*/
}

func TestUpdate(t *testing.T) {
	t.Run("Customer update", func(t *testing.T) {
		entity := model.Customer{Username: "username", Balance: 100}
		db := []model.Customer{entity}
		repo := repository{db: db}
		entity.Balance = 10
		err := repo.Update(entity)
		customer := repo.db[0]
		assert.Equal(t, entity, customer)
		assert.Nil(t, err)
	})

	t.Run("Not found customer", func(t *testing.T) {
		entity := model.Customer{Username: "username", Balance: 100}
		db := []model.Customer{}
		repo := repository{db: db}
		entity.Balance = 10
		err := repo.Update(entity)
		assert.NotNil(t, err)
	})
}
