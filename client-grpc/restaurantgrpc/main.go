package restaurantgrpc

import (
	"log"

	"golang.org/x/net/context"

	rs "github.com/boantp/go-mysql-resto/restaurant"
)

const (
	address = "localhost:50051"
)

// createRestaurant calls the RPC method createRestaurant of RestaurantServer
func CreateRestaurant(client rs.RestaurantClient, restaurant *rs.RestaurantRequest) {
	resp, err := client.CreateRestaurant(context.Background(), restaurant)
	if err != nil {
		log.Fatalf("Could not create Restaurant: %v", err)
	}
	if resp.Success {
		log.Printf("A new Restaurant has been added with id: %d", resp.RestaurantId)
	}
}

// getRestaurants calls the RPC method getRestaurants of RestaurantServer
func GetRestaurants(client rs.RestaurantClient, filter *rs.RestaurantFilter) *rs.RestaurantRequest {
	// calling the streaming API
	stream, err := client.GetRestaurants(context.Background(), filter)
	if err != nil {
		log.Fatalf("Error on get restaurants: %v", err)
	}

	restaurant, _ := stream.Recv()
	return restaurant
}
