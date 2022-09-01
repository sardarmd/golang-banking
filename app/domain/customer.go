package domain

type Customer struct {
	Id          string
	Name        string
	City        string
	Zip_code    string
	DateofBirth string
	Status      string
}

type CustomerRepository interface {
	FindAll() ([]Customer, error)
}

type CustomerRepositoryStub struct {
	customers []Customer
}

func (s CustomerRepositoryStub) FindAll() ([]Customer, error) {
	return s.customers, nil
}

func NewCustomerRespositoryStub() CustomerRepositoryStub {
	customers := []Customer{{"1001", "Sardar", "Bangalore", "560067", "25-08-1987", "1"}, {"1002", "Sadia", "Kanpu", "201301", "25-08-1987", "1"}}

	return CustomerRepositoryStub{customers}
}
