package api

import (
	"crypto/rand"
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"strconv"

	"github.com/julienschmidt/httprouter"
)

type (
	//ApiController represents the controller for operating on the API resource
	ApiController struct{}
)

func NewApiController() *ApiController {
	return &ApiController{}
}

func (ac ApiController) GetNearbyRestaurant(w http.ResponseWriter, r *http.Request, p httprouter.Params) {
	//Stub an Request Nearby Restaurant to be populated from the body
	u := ReqNearbyResto{}

	//Populate the restaurant request data
	json.NewDecoder(r.Body).Decode(&u)

	//Grab for lat and lnt
	lat := u.Latitude
	lnt := u.Longitude

	//Get Nearby Restaurant Query
	restoQuery := GetNearbyRestaurant(lat, lnt)

	//define resp
	d := NearbyRestaurantResp{"1", "success", restoQuery, ""}
	//Marshal provided interface into JSON structure
	uj, _ := json.Marshal(d)

	//Write content-type, status code, data
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(200)
	fmt.Fprintf(w, "%s", uj)
}

func (ac ApiController) GetDetailRestaurant(w http.ResponseWriter, r *http.Request, p httprouter.Params) {
	//Grab restaurant ID
	i := p.ByName("restaurant_id")
	restoId, _ := strconv.Atoi(i)

	//Query into DB
	detailResto := GetDetailRestaurant(int32(restoId))

	//define response
	d := DetailRestaurantResp{"1", "success", detailResto, ""}
	//Marshal provided interface into JSON structure
	uj, _ := json.Marshal(d)

	//Write content-type, status code, data
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(200)
	fmt.Fprintf(w, "%s", uj)
}

func (ac ApiController) CreateReservationRestaurant(w http.ResponseWriter, r *http.Request, p httprouter.Params) {
	//Stub an Restaurant to be populated from the body
	u := ReqReservation{}

	//Populate the restaurant data
	json.NewDecoder(r.Body).Decode(&u)

	//Prepare Insert Data Reservation
	rs := Reservation{}
	rs.ReservationCode = EncodeToString(8)
	rs.ReservationCustomer.CustomerName = u.ReservationCustomerName
	rs.ReservationCustomer.CustomerId = u.ReservationCustomerId
	rs.ReservationCustomer.CustomerPhone = u.ReservationCustomerPhone
	rs.ReservationDatetime = u.ReservationDateTime
	rs.ReservationRestaurant.RestaurantName = u.ReservationRestaurantName
	rs.ReservationRestaurant.RestaurantId = u.ReservationRestaurantId
	rs.ReservationTotalGuest = u.ReservationTotalGuest

	//Insert data
	resultQuery := CreateReservationRestaurant(rs)

	//define response
	d := ReservationResp{"1", "success", resultQuery}

	//Marshal provided interface into JSON structure
	uj, _ := json.Marshal(d)

	//Write content-type, status code, data
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(200)
	fmt.Fprintf(w, "%s", uj)
}

func EncodeToString(max int) string {
	var table = [...]byte{'1', '2', '3', '4', '5', '6', '7', '8', '9'}
	b := make([]byte, max)
	n, err := io.ReadAtLeast(rand.Reader, b, max)
	if n != max {
		panic(err)
	}
	for i := 0; i < len(b); i++ {
		b[i] = table[int(b[i])%len(table)]
	}
	return string(b)
}
