package domain

	type Employee struct {
		Id                  int    `json:"id"`
		FirstName           string `json:"first_name"`
		LastName            string `json:"last_name"`
		MiddleName          string `json:"middle_name"`
		Post                string `json:"post"`
		PassportSeries      string `json:"passport_series"`
		PassportNumber      string `json:"passport_number"`
		PassportIssuedBy    string `json:"passport_issued_by"`
		PassportDateOfIssue string `json:"passport_date_of_issue"`
	}
