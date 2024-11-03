package repositories

import (
	"errors"
)

type ZipcodeRepository struct {
	zipCodeToRestaurants map[string][]string
}

func NewZipcodeRepository(data map[string][]string) *ZipcodeRepository {

	repo := &ZipcodeRepository{
		zipCodeToRestaurants: data,
	}

	return repo

}

func (repo *ZipcodeRepository) Get(zipcode string) ([]string, error) {

	restaurantIds, exists := repo.zipCodeToRestaurants[zipcode]

	if exists {
		return restaurantIds, nil
	}

	return nil, errors.New("zip code not found in data") // todo: put all of this in an ErrorCodes object
}
