package domain

type Organization struct {
	Id             int    `json:"id"`
	Name           string `json:"name"`
	Address        string `json:"address"`
	AccountId      int    `json:"account_id"`
	Chief          string `json:"chief"`
	FinancialChief string `json:"financial_chief"`
}
