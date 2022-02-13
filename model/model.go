package model

type Customer struct {
	Username string  `json:"username"`
	Balance  float64 `json:"balance"`
	Name     string  `json:"name"`
	Surname  string  `json:"surname"`
}

type CustomerResponse []Customer
