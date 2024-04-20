package authenticationDto

import (
	"github.com/google/uuid"
	"time"
)

type (
	RegisterReq struct {
		Email    string `json:"email"`
		Password string `json:"password"`
	}

	Register struct {
		UsersID   uuid.UUID `json:"users_id"`
		Email     string    `json:"email"`
		Password  string    `json:"password"`
		Roles     string    `json:"roles"`
		CreatedAt time.Time `json:"created_at"`
		UpdatedAt time.Time `json:"updated_at"`
	}

	RegisterRes struct {
		UsersID   uuid.UUID `json:"users_id"`
		Email     string    `json:"email"`
		CreatedAt time.Time `json:"created_at"`
	}

	LoginReq struct {
		Email    string `json:"email"`
		Password string `json:"password"`
	}
)
