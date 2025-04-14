package dtos

type (
	Receiver struct {
		Id         int    `json:"id"`
		Name       string `json:"name"`
		Address    string `json:"address"`
		Email      string `json:"email"`
		Phone      string `json:"phone"`
		TaxPayerId string `json:"tax_payer_id"`
	}

	CreateReceiverRequest struct {
		Name       string `json:"name" binding:"required"`
		Address    string `json:"address" binding:"required"`
		Email      string `json:"email" binding:"required"`
		Phone      string `json:"phone" binding:"required"`
		TaxPayerId string `json:"tax_payer_id" binding:"required"`
	}

	UpdateReceiverRequest struct {
		Name       string `json:"name"`
		Address    string `json:"address"`
		Email      string `json:"email"`
		Phone      string `json:"phone"`
		TaxPayerId string `json:"tax_payer_id"`
	}

	ReceiverResponse struct {
		Id int `json:"id"`
	}
)
