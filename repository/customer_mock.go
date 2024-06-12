package repository

import "errors"

type customersRepositoryMock struct {
	customers []Customer
}

func NewCustomerRepositoryMock() customersRepositoryMock {
	customers := []Customer{
		{CustomerID: 1001, Name: "Ashish", City: "New Delhi", ZipCode: "xxxx", Status: 200},
		{CustomerID: 1002, Name: "Rob", City: "New Delhi", ZipCode: "xxxx", Status: 200},
	}

	return customersRepositoryMock{customers: customers}
}

func (r customersRepositoryMock) GetAll() ([]Customer, error) {
	return r.customers, nil
}

func (r customersRepositoryMock) GetById(id int) (*Customer, error) {
	for _, customer := range r.customers {
		if customer.CustomerID == id {
			return &customer, nil
		}
	}
	return nil, errors.New("customer not found")
}
