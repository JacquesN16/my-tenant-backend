package models

type Tenant struct {
	FirstName string  `json:"firstName"`
	LastName  string  `json:"lastName"`
	Email     string  `json:"email"`
	StartDate int64   `json:"startDate"`
	Rent      float64 `json:"rent"`
	Charge    float64 `json:"charge"`
}
