package repositories

import (
	"errors"
	restaurantv1 "test-app/gen/restaurant/v1"
)

type RestaurantInfoRepository struct {
	restaurants map[string]*restaurantv1.RestaurantInfo
}

func NewRestaurantInfoRepository(data map[string]*restaurantv1.RestaurantInfo) *RestaurantInfoRepository {
	return &RestaurantInfoRepository{
		restaurants: data,
	}
}

func (repo *RestaurantInfoRepository) Get(restaurantId string) (*restaurantv1.RestaurantInfo, error) {
	restaurant, exists := repo.restaurants[restaurantId]

	if exists {

		return restaurant, nil

	}

	return nil, errors.New("no restaurant with this id was found")
}
