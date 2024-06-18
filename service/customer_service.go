package service

import (
	"database/sql"

	"github.com/Aritiaya50217/CodeBangkok/errs"
	"github.com/Aritiaya50217/CodeBangkok/logs"
	"github.com/Aritiaya50217/CodeBangkok/repository"
)

type customerService struct {
	custRepo repository.CustomerRepository
}

func NewCustomerService(cusRepo repository.CustomerRepository) customerService {
	return customerService{custRepo: cusRepo}
}

func (s customerService) GetCustomers() ([]CustomerResponse, error) {
	customers, err := s.custRepo.GetAll()
	if err != nil {
		logs.Error(err)
		return nil, errs.NewUnexpectedError()
	}
	customersRes := []CustomerResponse{}
	for _, customer := range customers {
		customerResp := CustomerResponse{
			CustomerID: customer.CustomerID,
			Name:       customer.Name,
			Status:     customer.Status,
		}
		customersRes = append(customersRes, customerResp)
	}
	return customersRes, nil
}

func (s customerService) GetCustomer(id int) (*CustomerResponse, error) {
	customer, err := s.custRepo.GetById(id)
	if err != nil {
		if err == sql.ErrNoRows {
			return nil, errs.NewNotFoundError("customer not found")
		}
		logs.Error(err)
		return nil, errs.NewUnexpectedError()
	}
	custResponse := CustomerResponse{
		CustomerID: customer.CustomerID,
		Name:       customer.Name,
		Status:     customer.Status,
	}
	return &custResponse, nil
}
