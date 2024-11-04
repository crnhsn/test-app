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
      zipcodes: ["10001", "12345", "34567"],  
      cuisines: [], 
      healthRating: "D",    
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
   <ul>
        {restaurants.map(({ id, name, healthRating, description, userRating }) => (
          <li key={id}>
            {name} 
            <p>
            User Rating: {userRating}
            </p>
           
            <p>
            {description}
            </p>
            
            <p>
              Health Rating: {healthRating}
            </p>
            
          </li>
        ))}
      </ul>
   </div>
  )
}

export default App
