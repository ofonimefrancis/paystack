package paystack

func (ps PaystackApi) GetCutomers()                           {}
func (ps PaystackApi) GetCustomerById(id int32)               {}
func (ps PaystackApi) GetCustomerByCode(customer_code string) {}

func (ps PaystackApi) CreateCustomer()          {}
func (ps PaystackApi) UpdateCustomer()          {}
func (ps PaystackApi) DeactivateAuthorization() {}
