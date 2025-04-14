package dtos

type (
	Platform struct {
		Id   int    `json:"id"`
		Name string `json:"name"`
	}

	PlatformRequest struct {
		Name string `json:"name" binding:"required"`
	}

	PlatformResponse struct {
		Id int `json:"id"`
	}
)
