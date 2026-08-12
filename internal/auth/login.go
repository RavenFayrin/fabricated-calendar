package auth

import (
	"context"
	"fabricated-calendar/config"
	"fabricated-calendar/internal/database"
)

func Login(cfg config.Config, username string, password string) (database.User, error) {
	user, err := cfg.DB.GetUserByUsername(context.Background(), username)
	if err != nil {
		return database.User{}, err
	}
	_, err = CheckPasswordHash(password, user.HashedPassword)
	if err != nil {
		return database.User{}, err
	}
	return user, nil
}
