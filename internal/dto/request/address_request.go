package request

type CreateAddressRequest struct {
	RecipientName string `json:"recipient_name"`
	Phone         string `json:"phone"`
	Address       string `json:"address"`
	City          string `json:"city"`
	PostalCode    string `json:"postal_code"`
}

type UpdateAddressRequest struct {
	RecipientName string `json:"recipient_name"`
	Phone         string `json:"phone"`
	Address       string `json:"address"`
	City          string `json:"city"`
	PostalCode    string `json:"postal_code"`
}
