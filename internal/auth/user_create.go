package auth

import (
	"context"
	"fabricated-calendar/config"
	"fabricated-calendar/internal/database"
)

func CreateUser(cfg config.Config, email string, password string) error {
	hashedPassword, err := HashPassword(password)
	if err != nil {
		return err
	}

	_, err = cfg.DB.CreateUser(context.Background(), database.CreateUserParams{
		Email:          email,
		HashedPassword: hashedPassword,
	})
	if err != nil {
		return err
	}

	return nil
}
