package service

import (
	"bootcamp/config"
	"bootcamp/model"
	"bootcamp/repository"
)

type ServiceError struct {
	Message string
}

func (err *ServiceError) Error() string {
	return err.Message
}

type IService interface {
	GetCustomerByUsername(username string) (*model.Customer, error)
	GetAllCustomers() *[]model.Customer
	CreateCustomer(customer *model.Customer) error
	UpdateCustomer(customer *model.Customer) error
	IncreaseBalance(customer *model.Customer, amount float64) error
	DescreaseBalance(customer *model.Customer, amount float64) error
}

type service struct {
	repo repository.IRepository
}

func NewService(repo repository.IRepository) IService {
	return &service{repo: repo}
}

func (s *service) GetCustomerByUsername(username string) (*model.Customer, error) {
	return s.repo.GetCustomerByUsername(username)
}

func (s *service) GetAllCustomers() *[]model.Customer {
	return s.repo.Customers()
}

func (s *service) CreateCustomer(customer *model.Customer) error {
	if customer == nil {
		return &ServiceError{Message: "Customer is nil"}
	}
	customer.Balance = config.Env.InitialBalanceAmount
	s.repo.Add(*customer)
	return nil
}

func (s *service) UpdateCustomer(customer *model.Customer) error {
	if customer == nil {
		return &ServiceError{Message: "Customer is nil"}
	}
	return s.repo.Update(*customer)
}

func (s *service) IncreaseBalance(customer *model.Customer, amount float64) error {
	if customer == nil {
		return &ServiceError{Message: "Customer is nil"}
	}
	customer.Balance += amount
	return s.repo.Update(*customer)
}

func (s *service) DescreaseBalance(customer *model.Customer, amount float64) error {
	if customer == nil {
		return &ServiceError{Message: "Customer is nil"}
	}
	balance := customer.Balance - amount
	if config.Env.MinumumBalanceAmount > balance {
		return &ServiceError{Message: "Minumum Balance Amount Error"}
	}
	customer.Balance = balance
	return s.repo.Update(*customer)
}
