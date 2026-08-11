package auth

import (
	"context"
	"fabricated-calendar/config"
)

func Login(cfg config.Config, username string, password string) (bool, error) {
	user, err := cfg.DB.GetUserByUsername(context.Background(), username)
	if err != nil {
		return false, err
	}
	match, err := CheckPasswordHash(password, user.HashedPassword)
	if err != nil {
		return false, err
	}
	return match, nil
}
