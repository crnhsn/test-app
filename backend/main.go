package main

import (
	"context"
	"errors"
	"net/http"

	"connectrpc.com/connect"
	"golang.org/x/net/http2"
	"golang.org/x/net/http2/h2c"

	"test-app/backend/repositories"
	"test-app/backend/services"

	restaurantv1 "test-app/gen/restaurant/v1" // generated by protoc-gen-go
	"test-app/gen/restaurant/v1/restaurantv1connect"
)

type RestaurantService struct {
	zipcodeDataRepo        *repositories.ZipcodeRepository
	healthRatingsDataRepo  *repositories.HealthRatingsRepository
	cuisinesDataRepo       *repositories.RestaurantCuisineRepository
	restaurantInfoDataRepo *repositories.RestaurantInfoRepository
	locationInferrer       *services.MockZipcodeInferrer
	healthRatingComparer   *services.LetterHealthRatingComparer
}

func (restaurantService *RestaurantService) GetRestaurants(ctx context.Context,
	req *connect.Request[restaurantv1.GetRestaurantsRequest]) (*connect.Response[restaurantv1.GetRestaurantsResponse], error) {

	var requestedLocations []string
	var cuisines []string = req.Msg.Cuisines
	var minimumHealthRating string = req.Msg.HealthRating

	// try to get restaurants that match the zipcodes the user provided
	requestedLocations = req.Msg.Zipcodes
	restaurantIdsNearLocation, _ := restaurantService.zipcodeDataRepo.GetFromMany(requestedLocations)

	// user either provided no zipcodes, or their zipcodes were invalid / didn't work for some reason
	// so try a default set of zipcodes by attempting to infer location based on request
	// todo: improve err handling here and/or validation before here so we know whether the issue was that the
	// zipcodes were empty, invalid, or there was some internal repo error, etc.
	if len(restaurantIdsNearLocation) == 0 {
		inferredLocation, err := restaurantService.locationInferrer.InferLocation(ctx)
		if err != nil {
			return nil, errors.New("could not infer default location, and user either did not provide any zipcodes, or provided zipcodes that could not be resolved")
		}

		restaurantIdsNearLocation, _ = restaurantService.zipcodeDataRepo.GetFromMany(inferredLocation)
	}

	// if we still haven't found something, there's some issue -
	// our location inference service and location data repo should provide a strong guarantee of data return
	// in the case of defaults, so no data by this point might indicate a broader problem.
	// in an actual app of this kind we'd probably use proximity based lookups instead of
	// specific zipcodes, which should mitigate the problem: a proximity based lookup
	// on a sufficiently broad location (e.g., "New York, NY") should have high probability
	// of returning some non-empty set of data
	if len(restaurantIdsNearLocation) == 0 {
		return nil, errors.New("neither the user provided zipcodes nor the default zipcodes have any restaurants associated with them")
	}

	// get the set of location-filtered restaurant IDs as a slice
	var restaurantIdsNearLocationSlice []string
	for restaurantId := range restaurantIdsNearLocation {
		restaurantIdsNearLocationSlice = append(restaurantIdsNearLocationSlice, restaurantId)
	}

	var filteredRestaurantIds []string

	// apply cuisine-based filtering if the user specified it
	if len(cuisines) > 0 {
		cuisineToIds, _ := restaurantService.cuisinesDataRepo.Get(restaurantIdsNearLocationSlice)

		for _, cuisine := range cuisines {

			restaurantIds, exists := cuisineToIds[cuisine]

			if exists {
				filteredRestaurantIds = append(filteredRestaurantIds, restaurantIds...)
			}
		}

	} else {
		filteredRestaurantIds = restaurantIdsNearLocationSlice
	}

	if len(filteredRestaurantIds) == 0 {
		return nil, errors.New("no restaurants match the provided cuisine criteria")
	}

	var healthRatingFilteredIds []string

	// apply health rating based filtering if the user specified it
	if minimumHealthRating != "" {

		for _, restaurantId := range filteredRestaurantIds {
			healthRating, err := restaurantService.healthRatingsDataRepo.Get(restaurantId)
			if err != nil {
				continue // skip if there's some error getting health rating - we want to be greedy here for similar reason as above
			}

			meetsMinimumHealthRating, err := restaurantService.healthRatingComparer.IsBetterOrEqual(healthRating, minimumHealthRating)
			if err != nil {
				continue // skip for same reason as above
			}

			if meetsMinimumHealthRating {
				healthRatingFilteredIds = append(healthRatingFilteredIds, restaurantId)
			}
		}

	}

	if len(healthRatingFilteredIds) == 0 {
		return nil, errors.New("no restaurants found that match the health criteria - try broadening")
	}

	var restaurantInfos []*restaurantv1.RestaurantInfo
	for _, restaurantId := range healthRatingFilteredIds {
		restaurantInfo, err := restaurantService.restaurantInfoDataRepo.Get(restaurantId)
		if err == nil {
			restaurantInfos = append(restaurantInfos, restaurantInfo)
		}
	}

	responseObject := &restaurantv1.GetRestaurantsResponse{
		Restaurants: restaurantInfos,
	}

	response := connect.NewResponse(responseObject)

	// if we get to this point, we've input some non empty set of IDs into our restaurant info repo because
	// empty IDs that occur as a result of filtering get handled by the error handling before this point
	// so at this point there are two states possible: the response object has some restaurant info data,
	// or the object has no restaurant info data (because there was an error in the restaurant info repo, etc.)
	// either way - return to the client to handle

	return response, nil

}

func main() {
	// Mock data
	restaurantData := map[string]*restaurantv1.RestaurantInfo{
		"1": {Id: "1", Name: "Alinea", Address: "123 Chicago St", Description: "New American tasting menus and modernist cuisine with a twist.", UserRating: 5},
		"2": {Id: "2", Name: "McDonalds", Address: "456 Random St", Description: "It's McDonalds.", UserRating: 5},
		"3": {Id: "3", Name: "Ce Qui N'Existe Pas", Address: "789 Obscure St", Description: "Avant-garde French eats in a random basement that you probably won't be able to find before they give your reservation away.", UserRating: 4},
		"4": {Id: "4", Name: "Plates", Address: "100 Promenade Ave", Description: "Gives you food on every manner of serving vessel except for plates.", UserRating: 2},
		"5": {Id: "5", Name: "Fruit Butterfly", Address: "120 Main St", Description: "A popular tourist destination in a city known for its food.", UserRating: 3},
		"6": {Id: "6", Name: "Ssam", Address: "678 Main St", Description: "Award-winning Korean steakhouse in a modernist setting.", UserRating: 5},
		"7": {Id: "7", Name: "Anand", Address: "789 Main St", Description: "Indian-Thai fusion with an avant-garde flair.", UserRating: 5},
	}

	healthData := map[string]string{
		"1": "A",
		"2": "A",
		"3": "B",
		"4": "C",
		"5": "A",
		"6": "A",
		"7": "D",
	}

	cuisineData := map[string][]string{
		"1": {"American"},
		"2": {"American"},
		"3": {"French"},
		"4": {"American"},
		"5": {"American", "Italian"},
		"6": {"Korean"},
		"7": {"Indian", "Thai"},
	}

	zipcodeData := map[string][]string{
		"12345": {"1", "4"},
		"34567": {"2", "3"},
		"10001": {"5", "6", "7"},
	}

	restaurantRepo := repositories.NewRestaurantInfoRepository(restaurantData)
	healthRepo := repositories.NewHealthRatingsRepository(healthData)
	cuisineRepo := repositories.NewRestaurantCuisineRepository(cuisineData)
	zipcodeRepo := repositories.NewZipcodeRepository(zipcodeData)

	locationInferrer := services.NewMockZipcodeInferrer([]string{"10001"})
	healthRatingComparer := services.NewLetterHealthRatingComparer()

	restaurantService := &RestaurantService{
		zipcodeDataRepo:        zipcodeRepo,
		healthRatingsDataRepo:  healthRepo,
		cuisinesDataRepo:       cuisineRepo,
		restaurantInfoDataRepo: restaurantRepo,
		locationInferrer:       locationInferrer,
		healthRatingComparer:   healthRatingComparer,
	}

	mux := http.NewServeMux()
	path, handler := restaurantv1connect.NewRestaurantServiceHandler(restaurantService)
	mux.Handle(path, handler)

	http.ListenAndServe(
		"localhost:8080",
		// Use h2c so we can serve HTTP/2 without TLS.
		h2c.NewHandler(mux, &http2.Server{}),
	)

}
