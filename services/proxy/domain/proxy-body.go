package domain

// Тело доверенности. Каждая строка с товаром представлена этой структурой
type ProxyBodyItem struct {
	Id            int `json:"id"`
	ProductId     int `json:"product_id"`
	ProxyId       int `json:"proxy_id"`
	ProductAmount int `json:"product_amount"`
}
