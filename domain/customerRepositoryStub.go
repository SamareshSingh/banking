package domain

// Create Mock Adapter. Used to test the expectation
type CustomerRepositoryStub struct {
	customers []Customer
}

// Implement FindAll(). If looks like a duck and quack like a duck then it is a duck
func (s CustomerRepositoryStub) FindAll() ([]Customer, error) {
	return s.customers, nil
}

// Helper function for creating new customer repository stub
func NewCustomerRepositoryStub() CustomerRepositoryStub {
	customers := []Customer{
		{Id: "1001", Name: "Ashish", City: "Austin", Zipcode: "78787", DateofBirth: "2000-01-01", Status: "1"},
		{Id: "1002", Name: "Samaresh", City: "Denver", Zipcode: "67652", DateofBirth: "2010-06-11", Status: "1"},
	}

	return CustomerRepositoryStub{customers}
}
