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

func (repo *ZipcodeRepository) GetFromMany(zipcodes []string) (map[string]bool, error) {
	restaurantIdsFromZipcodes := make(map[string]bool) // simulate a set

	for _, zipcode := range zipcodes {

		restaurantIds, err := repo.Get(zipcode)

		// keep going if nothing found or there was some issue -
		// in the case of multiple zipcodes, we want to be greedy and get any restaurants
		// that match criteria and are okay if some error (invalid zipcode, internal repo error, etc.)
		// results in a single zipcode's not necessarily getting resolved
		// better to return some results to the user if possible than to return nothing

		// if there's only a single zipcode and it errors out, then the caller should handle that

		// in a more developed version of this, we'd use a proximity based location service
		// instead of exact zipcode match, which would alleviate this issue,
		// but for small test app purposes this is fine

		if err != nil {
			continue
		}

		for _, restaurantId := range restaurantIds {
			restaurantIdsFromZipcodes[restaurantId] = true
		}

	}

	return restaurantIdsFromZipcodes, nil
}
