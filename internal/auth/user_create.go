package auth

import (
	"context"
	"fabricated-calendar/config"
	"fabricated-calendar/internal/database"
)

func CreateUser(cfg config.Config, username string, password string, email string) error {
	hashedPassword, err := HashPassword(password)
	if err != nil {
		return err
	}

	err = CheckEmailValidation(email)
	if err != nil {
		return err
	}

	_, err = cfg.DB.CreateUser(context.Background(), database.CreateUserParams{
		Username:       username,
		HashedPassword: hashedPassword,
		Email:          email,
	})
	if err != nil {
		return err
	}

	return nil
}
