package main

func main() {

}

// func main() {
// 	// mock data
// 	restaurantData := map[string]*restaurantv1.RestaurantInfo{
// 		"1": {Id: "1", Name: "Alinea", Address: "123 Chicago St", Description: "New American tasting menus and modernist cuisine with a twist.", UserRating: 5},
// 		"2": {Id: "2", Name: "McDonalds", Address: "456 Random St", Description: "It's McDonalds.", UserRating: 5},
// 		"3": {Id: "3", Name: "Ce Qui N'Existe Pas", Address: "789 Obscure St", Description: "Avant-garde French eats in a random basement that you probably won't be able to find before they give your reservation away.", UserRating: 4},
// 		"4": {Id: "4", Name: "Plates", Address: "100 Promenade Ave", Description: "Gives you food on every manner of serving vessel except for plates.", UserRating: 2},
// 		"5": {Id: "5", Name: "Fruit Butterfly", Address: "120 Main St", Description: "A popular tourist destination in a city known for its food.", UserRating: 3},
// 		"6": {Id: "6", Name: "Ssam", Address: "678 Main St", Description: "Award-winning Korean steakhouse in a modernist setting.", UserRating: 5},
// 		"7": {Id: "7", Name: "Anand", Address: "789 Main St", Description: "Indian-Thai fusion with an avant-garde flair.", UserRating: 5},
// 	}

// 	healthData := map[string]string{
// 		"1": "A",
// 		"2": "A",
// 		"3": "B",
// 		"4": "C",
// 		"5": "A",
// 		"6": "A",
// 		"7": "D",
// 	}

// 	cuisineData := map[string][]string{
// 		"1": {"American"},
// 		"2": {"American"},
// 		"3": {"French"},
// 		"4": {"American"},
// 		"5": {"American", "Italian"},
// 		"6": {"Korean"},
// 		"7": {"Indian", "Thai"},
// 	}

// 	zipcodeData := map[string][]string{
// 		"12345": {"1", "4"},
// 		"34567": {"2", "3"},
// 		"10001": {"5", "6", "7"},
// 	}

// 	restaurantRepo := repositories.NewRestaurantInfoRepository(restaurantData)
// 	healthRepo := repositories.NewHealthRatingsRepository(healthData)
// 	cuisineRepo := repositories.NewRestaurantCuisineRepository(cuisineData)
// 	zipcodeRepo := repositories.NewZipcodeRepository(zipcodeData)

// 	fmt.Println(restaurantRepo.Get("1"))
// 	fmt.Println(healthRepo.Get("1"))
// 	fmt.Println(cuisineRepo.Get([]string{"1"}))
// 	fmt.Println(zipcodeRepo.Get("12345"))
// }
