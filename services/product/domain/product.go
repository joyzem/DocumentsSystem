package domain

type Product struct {
	Id     int    `json:"id"`
	Name   string `json:"name"`
	Price  int    `json:"price"`
	UnitId int    `json:"unit_id"`
}
