package repository

import (
	"bootcamp/model"
)

type RepositoryError struct {
	Message string
}

func (err *RepositoryError) Error() string {
	return err.Message
}

type IRepository interface {
	Customers() *[]model.Customer
	GetCustomerByUsername(username string) (*model.Customer, error)
	Add(customer model.Customer)
	Update(customer model.Customer) error
}

type repository struct {
	db []model.Customer
}

func (r *repository) Customers() *[]model.Customer {
	return &r.db
}

func (r *repository) GetCustomerByUsername(username string) (*model.Customer, error) {
	for _, entity := range r.db {
		if entity.Username == username {
			return &entity, nil
		}
	}
	return nil, &RepositoryError{Message: "Customer not found"}
}

func (r *repository) Add(customer model.Customer) {
	r.db = append(r.db, customer)
}

func (r *repository) Update(customer model.Customer) error {
	for i, entity := range r.db {
		if entity.Username == customer.Username {
			r.db[i] = customer
			return nil
		}
	}
	return &RepositoryError{Message: "Customer not found!"}
}

func NewRepository() IRepository {
	return &repository{db: []model.Customer{}}
}
