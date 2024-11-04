import { useEffect, useState } from 'react';
import { createClient } from "@connectrpc/connect";
import { createConnectTransport } from "@connectrpc/connect-web";

import { RestaurantService } from './gen/restaurant/v1/restaurant_connect';
import { GetRestaurantsRequest, RestaurantInfo } from './gen/restaurant/v1/restaurant_pb';



function App() {
  const [count, setCount] = useState(0)

  const [restaurants, setRestaurants] = useState<RestaurantInfo[]>([]);
  const [error, setError] = useState<string>("");

  const client = createClient(
    RestaurantService,
    createConnectTransport({
      baseUrl: "http://localhost:8080/",
    }),
  );

  const getRestaurants = async () => {
    const request = new GetRestaurantsRequest({
      zipcodes: ["10001"],  
      cuisines: ["Italian"], 
      healthRating: "A",    
    });

    try {
      const response = await client.getRestaurants(request);
      setRestaurants(response.restaurants);
    } catch (error) {
      setError("Error fetching restaurants: " + error);
      console.error("Error getting restaurants:", error);
    }
  };

  useEffect(() => {
    getRestaurants()
  }, []); 

  return (
   <div>
    hello world
   </div>
  )
}

export default App
