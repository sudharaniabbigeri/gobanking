package domain

	type CustomerRepositoryStub struct {
	customers []Customer
   
   }
   
   func (s CustomerRepositoryStub) FindAll() ([]Customer, error) {
	   return s.customers, nil
   }
   
   func NewCustomerRepositoryStub() CustomerRepositoryStub {
	   customers := []Customer {
		   {ID:"1001", Name:"Jhon", City:"Mumbai", Zipcode:"521121", DateofBirth:"01/01/2000", Status: "1"},
		   {ID:"1002", Name:"Jhosna", City:"Delhi", Zipcode:"520001", DateofBirth:"01/01/2001", Status: "1"},
	   }
   
	   return CustomerRepositoryStub{customers}
   }