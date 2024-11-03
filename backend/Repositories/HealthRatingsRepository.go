package repositories

import (
	"errors"
)

type HealthRatingsRepository struct {
	restaurantIdToHealthRating map[string]string
}

func NewHealthRatingsRepository(data map[string]string) *HealthRatingsRepository {
	repo := HealthRatingsRepository{
		restaurantIdToHealthRating: data,
	}

	return &repo
}

func (repo *HealthRatingsRepository) Get(restaurantId string) (string, error) {
	rating, exists := repo.restaurantIdToHealthRating[restaurantId]
	if exists {
		return rating, nil
	}

	return "", errors.New("could not find health rating information for the given restaurant ID")
}
