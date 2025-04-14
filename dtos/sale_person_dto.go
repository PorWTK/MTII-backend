package dtos

type (
	SalePerson struct {
		Id   int    `json:"id"`
		Name string `json:"name"`
	}

	SalePersonRequest struct {
		Name string `json:"name" binding:"required"`
	}

	SalePersonResponse struct {
		Id int `json:"id"`
	}
)
