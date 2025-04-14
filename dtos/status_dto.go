package dtos

type (
	Status struct {
		Id   int    `json:"id"`
		Name string `json:"name"`
	}

	StatusRequest struct {
		Name string `json:"name" binding:"required"`
	}

	StatusResponse struct {
		Id int `json:"id"`
	}
)
