package dtos

type (
	PaymentMethod struct {
		Id   int    `json:"id"`
		Name string `json:"name"`
	}

	PaymentMethodRequest struct {
		Name string `json:"name" binding:"required"`
	}

	PaymentMethodResponse struct {
		Id int `json:"id"`
	}
)
