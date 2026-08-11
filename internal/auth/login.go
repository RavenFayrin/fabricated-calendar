package auth

import (
	"context"
	"errors"
	"fabricated-calendar/config"
)

func Login(cfg config.Config, email string, password string) error {
	user, err := cfg.DB.GetUserByEmail(context.Background(), email)
	if err != nil {
		return err
	}
	match, err := CheckPasswordHash(password, user.HashedPassword)
	if err != nil {
		return err
	}
	if match == true {
		return nil
	}
	return errors.New("Incorrect Password")
}
