package api

import (
	"log"

	"github.com/boantp/go-mysql-resto/client-grpc/restaurantgrpc"
	"github.com/boantp/go-mysql-resto/config/db"
	rs "github.com/boantp/go-mysql-resto/restaurant"
	"google.golang.org/grpc"
)

// address for grpc server call in model
const (
	address = "rest.dev:50051"
)

type Restaurant struct {
	RestaurantId          int32         `json:"restaurant_id"`
	RestaurantName        string        `json:"restaurant_name"`
	RestaurantUrl         string        `json:"restaurant_url"`
	RestaurantDescription string        `json:"restaurant_description"`
	RestaurantAddress     string        `json:"restaurant_address"`
	RestaurantPhone       string        `json:"restaurant_phone"`
	RestaurantLocation    string        `json:"restaurant_location"`
	RestaurantCuisines    Cuisines      `json:"restaurant_cuisines"`
	RestaurantOperational []Operational `json:"restaurant_operationals"`
	Restaurantlatitude    float64       `json:"restaurant_latitude"`
	RestaurantLongitude   float64       `json:"restaurant_longitude"`
	RestaurantImage       string        `json:"restaurant_image"`
	RadiusKm              float64       `json:"radius_km"`
}

type Cuisines struct {
	CuisinesId   int32  `json:"cuisines_id"`
	CuisinesName string `json:"restaurant_cuisines_name"`
}

type Operational struct {
	OperationalId           int32  `json:"operational_id"`
	OperationalRestaurantId int32  `json:"operational_restaurant_id"`
	OperationalDay          string `json:"operational_day"`
	OperationalOpenHour     string `json:"operational_open_hour"`
	OperationalClosedHour   string `json:"operational_closed_hour"`
}

type Reservation struct {
	ReservationId         int        `json:"reservation_id"`
	ReservationRestaurant Restaurant `json:"reservation_restaurant"`
	ReservationCode       string     `json:"reservation_code"`
	ReservationTotalGuest string     `json:"reservation_total_guest"`
	ReservationDatetime   string     `json:"reservation_datetime"`
	ReservationCustomer   Customer   `json:"reservation_customer"`
}

type Customer struct {
	CustomerId    int    `json:"customer_id"`
	CustomerName  string `json:"customer_name"`
	CustomerPhone string `json:"customer_phone"`
}

type ReqNearbyResto struct {
	CustomerId int     `json:"customer_id"`
	Latitude   float64 `json:"latitude"`
	Longitude  float64 `json:"longitude"`
}

type NearbyRestaurantResp struct {
	RespCode string       `json:"response:code"`
	RespDesc string       `json:"response_description"`
	Data     []Restaurant `json:"data"`
	Location string       `json:"location"`
}

type DetailRestaurantResp struct {
	RespCode string     `json:"response:code"`
	RespDesc string     `json:"response_description"`
	Data     Restaurant `json:"data"`
	Location string     `json:"location"`
}

type ReqReservation struct {
	ReservationCustomerId     int    `json:"reservation_customer_id"`
	ReservationCustomerName   string `json:"reservation_customer_name"`
	ReservationCustomerPhone  string `json:"reservation_customer_phone"`
	ReservationRestaurantId   int32  `json:"reservation_restaurant_id"`
	ReservationRestaurantName string `json:"reservation_restaurant_name"`
	ReservationTotalGuest     string `json:"reservation_total_guest"`
	ReservationDateTime       string `json:"reservation_datetime"`
}

type ReservationResp struct {
	RespCode string      `json:"response:code"`
	RespDesc string      `json:"response_description"`
	Data     Reservation `json:"data"`
}

// Query for Nearby restaurant and create gRPC restaurant
func GetNearbyRestaurant(lat, lnt float64) []Restaurant {
	rows, err := db.DB.Query("SELECT a.restaurant_id, a.restaurant_name, a.restaurant_url, a.restaurant_description, a.restaurant_address, a.restaurant_phone, a.restaurant_location, b.cuisines_id, b.cuisines_name, a.restaurant_latitude, a.restaurant_longitude, a.restaurant_image, ROUND( ? * ACOS( SIN(? * PI() / 180) * SIN(a.restaurant_latitude * PI() / 180) + COS(? * PI() / 180) * COS(a.restaurant_latitude * PI() / 180) * COS( (a.restaurant_longitude * PI() / 180) - (? * PI() / 180) ) ), 1 ) AS distance FROM restaurant AS a LEFT JOIN cuisines AS b ON a.`restaurant_cuisines_id` = b.`cuisines_id` ORDER BY distance ASC", 6371, lat, lat, lnt)
	checkErr(err)
	defer rows.Close()

	// Set up connection to the gRPC server
	conn, err := grpc.Dial(address, grpc.WithInsecure())
	if err != nil {
		log.Fatalf("did not connect: %v", err)
	}
	defer conn.Close()

	// Creates a new RestaurantClient
	client := rs.NewRestaurantClient(conn)

	restaurants := make([]Restaurant, 0)
	for rows.Next() {
		restaurant := Restaurant{}
		err := rows.Scan(&restaurant.RestaurantId, &restaurant.RestaurantName, &restaurant.RestaurantUrl, &restaurant.RestaurantDescription, &restaurant.RestaurantAddress, &restaurant.RestaurantPhone, &restaurant.RestaurantLocation, &restaurant.RestaurantCuisines.CuisinesId, &restaurant.RestaurantCuisines.CuisinesName, &restaurant.Restaurantlatitude, &restaurant.RestaurantLongitude, &restaurant.RestaurantImage, &restaurant.RadiusKm)
		checkErr(err)

		//get operational
		restId := restaurant.RestaurantId
		op, req := GetOperational(restId)
		restaurant.RestaurantOperational = op

		//Create grpc data for CreateRestaurant
		resto := &rs.RestaurantRequest{
			RestaurantId:          restaurant.RestaurantId,
			RestaurantName:        restaurant.RestaurantName,
			RestaurantUrl:         restaurant.RestaurantUrl,
			RestaurantDescription: restaurant.RestaurantDescription,
			RestaurantAddress:     restaurant.RestaurantAddress,
			RestaurantPhone:       restaurant.RestaurantPhone,
			RestaurantLocation:    restaurant.RestaurantLocation,
			RestaurantImage:       restaurant.RestaurantImage,
			Operationals:          req,
		}

		// Create a new restaurant
		restaurantgrpc.CreateRestaurant(client, resto)

		restaurants = append(restaurants, restaurant)
	}
	return restaurants
}

// Query for get detail restaurant, Call to db not used in web only for api cause we call gRPC service
func GetDetailRestaurant(restId int32) Restaurant {
	rows := db.DB.QueryRow("SELECT a.restaurant_id, a.restaurant_name, a.restaurant_description, a.restaurant_address, a.restaurant_phone, a.restaurant_location, b.cuisines_id, b.cuisines_name, a.restaurant_latitude, a.restaurant_longitude, a.restaurant_image FROM restaurant AS a LEFT JOIN cuisines AS b ON a.`restaurant_cuisines_id` = b.`cuisines_id` WHERE a.`restaurant_id` =?", restId)

	restaurant := Restaurant{}
	err := rows.Scan(&restaurant.RestaurantId, &restaurant.RestaurantName, &restaurant.RestaurantDescription, &restaurant.RestaurantAddress, &restaurant.RestaurantPhone, &restaurant.RestaurantLocation, &restaurant.RestaurantCuisines.CuisinesId, &restaurant.RestaurantCuisines.CuisinesName, &restaurant.Restaurantlatitude, &restaurant.RestaurantLongitude, &restaurant.RestaurantImage)
	checkErr(err)

	//get operational
	restId = restaurant.RestaurantId
	op, _ := GetOperational(restId)
	restaurant.RestaurantOperational = op

	return restaurant
}

//Query for get operational restaurant
func GetOperational(restId int32) ([]Operational, []*rs.RestaurantRequest_Operational) {
	rows, err := db.DB.Query("SELECT operational_id, operational_day, operational_open_hour, operational_closed_hour FROM operational WHERE operational_restaurant_id =?", restId)
	checkErr(err)
	defer rows.Close()

	requestOperational := make([]*rs.RestaurantRequest_Operational, 0)
	operationals := make([]Operational, 0)
	for rows.Next() {
		//operational for struct Operational
		operational := Operational{}
		err := rows.Scan(&operational.OperationalId, &operational.OperationalDay, &operational.OperationalOpenHour, &operational.OperationalClosedHour)
		checkErr(err)

		operationals = append(operationals, operational)

		// Operational for struct RestaurantRequest_Operational gRPC
		reqOpertional := &rs.RestaurantRequest_Operational{}
		errs := rows.Scan(&reqOpertional.OperationalId, &reqOpertional.OperationalDay, &reqOpertional.OperationalOpenHour, &reqOpertional.OperationalClosedHour)
		checkErr(errs)

		requestOperational = append(requestOperational, reqOpertional)
	}

	return operationals, requestOperational
}

//Query to make reservation
func CreateReservationRestaurant(rs Reservation) Reservation {
	sqlStr := "INSERT INTO reservation(reservation_restaurant_id, reservation_code, reservation_total_guest, reservation_datetime, reservation_customer_id, reservation_customer_name, reservation_customer_phone) VALUES (?, ?, ?, ?, ?, ?, ?)"
	stmt, err := db.DB.Prepare(sqlStr)
	checkErr(err)

	_, err = stmt.Exec(rs.ReservationRestaurant.RestaurantId, rs.ReservationCode, rs.ReservationTotalGuest, rs.ReservationDatetime, rs.ReservationCustomer.CustomerId, rs.ReservationCustomer.CustomerName, rs.ReservationCustomer.CustomerPhone)
	checkErr(err)

	return rs
}

func checkErr(err error) {
	if err != nil {
		panic(err)
	}
}
