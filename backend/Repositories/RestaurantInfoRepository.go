package repositories

import (
	"errors"
	"test-app/backend/models"
)

type RestaurantInfoRepository struct {
	restaurantData map[string]*models.RestaurantInfo
}

func NewRestaurantInfoRepository(data map[string]*models.RestaurantInfo) *RestaurantInfoRepository {
	return &RestaurantInfoRepository{
		restaurantData: data,
	}
}

func (repo *RestaurantInfoRepository) Get(restaurantId string) (*models.RestaurantInfo, error) {
	restaurant, exists := repo.restaurantData[restaurantId]

	if exists {
		return restaurant, nil
	}

	return nil, errors.New("restaurant not found")
}
