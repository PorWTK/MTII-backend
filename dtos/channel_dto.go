package dtos

type (
	Channel struct {
		Id   int    `json:"id"`
		Name string `json:"name"`
	}

	ChannelRequest struct {
		Name string `json:"name" binding:"required"`
	}

	ChannelResponse struct {
		Id int `json:"id"`
	}
)
