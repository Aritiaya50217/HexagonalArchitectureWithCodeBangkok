package service

import (
	"log"

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
		log.Println(err)
		return nil, err
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
		log.Println(err)
		return nil, err
	}
	custResponse := CustomerResponse{
		CustomerID: customer.CustomerID,
		Name:       customer.Name,
		Status:     customer.Status,
	}
	return &custResponse, nil
}
