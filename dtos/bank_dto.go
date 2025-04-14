package dtos

type (
	Bank struct {
		Id   int    `json:"id"`
		Name string `json:"name"`
	}

	BankRequest struct {
		Name string `json:"name" binding:"required"`
	}

	BankResponse struct {
		Id int `json:"id"`
	}
)
