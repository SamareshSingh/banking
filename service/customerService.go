package service

import (
	"github.com/SamareshSingh/banking/domain"
	"github.com/SamareshSingh/banking/errs"
)

// Service interface. Primary Port
type CustomerService interface {
	GetAllCustomer(string) ([]domain.Customer, *errs.AppError)
	GetCustomer(string) (*domain.Customer, *errs.AppError)
}

// Service implementation --> Business Logic. this has dependency on the Repository interface
type DefaultCustomerService struct {
	repo domain.CustomerRepository // injected dependency on the repository
}

// receiver function that will return slice of customer
func (s DefaultCustomerService) GetAllCustomer(status string) ([]domain.Customer, *errs.AppError) {
	if status == "active" {
		status = "1"
	} else if status == "inactive" {
		status = "0"
	} else {
		status = ""
	}
	return s.repo.FindAll(status)
}

func (s DefaultCustomerService) GetCustomer(id string) (*domain.Customer, *errs.AppError) {
	return s.repo.ById(id)
}

// Helper function to instantiate the DefaultCustomerService
// till this point i have connection Service <<interface>> <---> Business Logic <----> Repository <<interface>>
func NewCustomerService(repository domain.CustomerRepository) DefaultCustomerService {
	return DefaultCustomerService{repository}
}
