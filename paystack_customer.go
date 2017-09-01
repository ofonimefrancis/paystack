package paystack

type Customer struct {
	Integration  int32        `json:"integration"`
	FirstName    string       `json:"first_name"`
	LastName     string       `json:"last_name"`
	Email        string       `json:"email"`
	Phone        string       `json:"phone"` //Assume its a string
	Metadata     CustomerMeta `json:"metadata"`
	Domain       string       `json:"domain"`
	CustomerCode string       `json:"customer_code"`
	Id           int32        `json:"id"`
	CreatedAt    string       `json:"createdAt"`
	UpdatedAt    string       `json:"updatedAt"`
}

type CustomerMeta struct {
	Photos []MetaPhoto `json:"photos"`
}

type MetaPhoto struct {
	Type      string `json:"type"`
	TypeId    string `json:"typeId"`
	TypeName  string `json:"typeName"`
	Url       string `json:"url"`
	IsPrimary bool   `json:"isPrimary"`
}

type ResponseMeta struct {
	Total     int32 `json:"total"`
	Skipped   int32 `json:"skipped"`
	PerPage   int32 `json:"perPage"`
	Page      int32 `json:"page"`
	PageCount int32 `json:"pageCount"`
}

type CustomerResponse struct {
	Status  bool         `json:"status"`
	Message string       `json:"message"`
	Data    []Customer   `json:"data"`
	Meta    ResponseMeta `json:"meta"`
}
