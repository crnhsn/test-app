package repositories

import (
	"errors"
)

type RestaurantCuisineRepository struct {
	restaurantIdToCuisines map[string][]string
}

func NewRestaurantCuisineRepository(data map[string][]string) *RestaurantCuisineRepository {

	var repo *RestaurantCuisineRepository

	repo = &RestaurantCuisineRepository{
		restaurantIdToCuisines: data,
	}

	return repo

}

// get a map of cuisine to the restaurant IDs for that cuisine
func (repo *RestaurantCuisineRepository) Get(restaurantIds []string) (map[string][]string, error) {

	cuisineToRestaurantIds := make(map[string][]string)

	for _, restaurantId := range restaurantIds {

		cuisines, idExists := repo.restaurantIdToCuisines[restaurantId] // get the cuisine(s) associated with the restaurant

		if idExists {
			for _, cuisine := range cuisines { // for every cuisine, map the id to the cuisine

				cuisineToRestaurantIds[cuisine] = append(cuisineToRestaurantIds[cuisine], restaurantId)

			}
		}
	}

	if len(cuisineToRestaurantIds) == 0 {
		return nil, errors.New("no valid restaurant IDs found")
	}

	return cuisineToRestaurantIds, nil

}
