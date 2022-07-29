package httptypes

type Employees struct {
	Employees []Employee `json:"employees"`
}

type Employee struct {
	UserID       string `json:"userId"`
	EmailAddress string `json:"emailAddress"`
	FirstName    string `json:"firstName"`
	JobTitle     string `json:"jobTitle"`
	LastName     string `json:"lastName"`
	Region       string `json:"region"`
}
