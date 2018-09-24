package main

import (
	"log"
	"net"
	"strings"

	"golang.org/x/net/context"
	"google.golang.org/grpc"

	rs "github.com/boantp/go-mysql-resto/restaurant"
)

const (
	port = ":50051"
)

// server is used to implement restaurant.RestaurantServer.
type server struct {
	savedRestaurants []*rs.RestaurantRequest
}

// CreateRestaurant creates a new Restaurant
func (s *server) CreateRestaurant(ctx context.Context, in *rs.RestaurantRequest) (*rs.RestaurantResponse, error) {
	if len(s.savedRestaurants) != 0 {
		i := 0
		for _, resto := range s.savedRestaurants {
			if strings.Contains(resto.RestaurantName, in.RestaurantName) {
				i++
			}
		}
		if i == 0 {
			s.savedRestaurants = append(s.savedRestaurants, in)
		}
	} else {
		s.savedRestaurants = append(s.savedRestaurants, in)
	}
	return &rs.RestaurantResponse{RestaurantId: in.RestaurantId, Success: true}, nil
}

// GetRestaurants returns all restaurants by given filter
func (s *server) GetRestaurants(filter *rs.RestaurantFilter, stream rs.Restaurant_GetRestaurantsServer) error {
	for _, resto := range s.savedRestaurants {
		if filter.Keyword != "" {
			if !strings.Contains(resto.RestaurantUrl, filter.Keyword) {
				continue
			}
		}
		if err := stream.Send(resto); err != nil {
			return err
		}
	}
	return nil
}

func main() {
	lis, err := net.Listen("tcp", port)
	if err != nil {
		log.Fatalf("failed to listen: %v", err)
	}
	// Creates a new gRPC server
	s := grpc.NewServer()
	rs.RegisterRestaurantServer(s, &server{})
	s.Serve(lis)
}
