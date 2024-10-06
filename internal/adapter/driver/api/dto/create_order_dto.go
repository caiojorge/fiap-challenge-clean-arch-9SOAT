package dto

type CreateOrderDTO struct {
	Items       []*CreateOrderItemDTO `json:"items"`
	CustomerCPF string                `json:"cpf"`
}

type CreateOrderItemDTO struct {
	ProductID string `json:"productid"`
	Quantity  int    `json:"quantity"`
}
