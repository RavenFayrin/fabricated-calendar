package calendar

import (
	"context"
	"database/sql"
	"fabricated-calendar/config"
	"fabricated-calendar/internal/database"
)

func universeCreate(cfg config.Config, user database.User, name string, description string) (bool, error) {
	desc := sql.NullString{}

	switch description {
	case "":
		desc = sql.NullString{
			String: description,
			Valid:  true,
		}
	default:
		desc = sql.NullString{
			String: "",
			Valid:  false,
		}
	}

	_, err := cfg.DB.CreateUniverse(context.Background(), database.CreateUniverseParams{
		Name:        name,
		Description: desc,
		UserID:      user.ID,
	})
	if err != nil {
		return false, err
	}
	return true, nil
}

func universesGetByUserId(cfg config.Config, user database.User) (universes []database.Universe, err error) {
	universes, err = cfg.DB.GetUniversesByUserId(context.Background(), user.ID)
	if err != nil {
		return nil, err
	}
	return universes, nil
}
