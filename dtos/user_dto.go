package dtos

import (
	"time"
)

type (
	LoginRequest struct {
		Username string `json:"username" binding:"required" form:"username"`
		Password string `json:"password" binding:"required" form:"password"`
	}

	LoginResponse struct {
		Token     string    `json:"token"`
		ExpiresAt time.Time `json:"expires_at"`
	}
)
