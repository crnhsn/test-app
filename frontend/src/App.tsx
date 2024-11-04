import { useEffect, useState } from 'react';
import { createClient } from "@connectrpc/connect";
import { createConnectTransport } from "@connectrpc/connect-web";

import { RestaurantService } from './gen/restaurant/v1/restaurant_connect';
import { GetRestaurantsRequest } from './gen/restaurant/v1/restaurant_pb';



function App() {
  const [count, setCount] = useState(0)

  // const transport = createConnectTransport({
  //   baseUrl: "https://demo.connectrpc.com",
  // });

  const client = createClient(
    RestaurantService,
    createConnectTransport({
      baseUrl: "http://localhost:8080/",
    }),
  );

  return (
   <div>
    hello world
   </div>
  )
}

export default App
