package domain

type ProxyHeader struct {
	Id             int    `json:"id"`
	OrganizationId int    `json:"organization_id"`
	CustomerId     int    `json:"customer_id"`
	EmployeeId     int    `json:"employee_id"`
	DateOfIssue    string `json:"date_of_issue"`
	IsValidUntil   string `json:"is_valid_until"`
}
