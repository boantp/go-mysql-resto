package web

import (
	"bytes"
	"encoding/json"
	_ "io"
	"io/ioutil"
	"log"
	"net/http"
	"strconv"

	"github.com/boantp/go-mysql-resto/api"
	"github.com/boantp/go-mysql-resto/config/tpl"
	"google.golang.org/grpc"

	"github.com/boantp/go-mysql-resto/client-grpc/restaurantgrpc"
	rs "github.com/boantp/go-mysql-resto/restaurant"
)

const (
	address = "rpc.dev:50051"
)

type (
	//WebController represents the controller for operating on the Web resource
	WebController struct{}
)

func NewWebController() *WebController {
	return &WebController{}
}

type BaseUrl string

const (
	development BaseUrl = "http://api.dev:3000/"
	staging     BaseUrl = "http://localhost:3000/"
	production  BaseUrl = "http://localhost:3000/"
)

func (wc WebController) NearbyRestaurant(w http.ResponseWriter, r *http.Request) {
	resto := api.NearbyRestaurantResp{}
	//When method post
	if r.Method == http.MethodPost {
		//get from values
		formLatitude := r.FormValue("latitude")
		formLongitude := r.FormValue("longitude")
		formCustomerId := r.FormValue("customer_id")
		formLocation := r.FormValue("city")

		//string to int
		customerId, err := strconv.Atoi(formCustomerId)
		if err != nil {
			log.Fatal(err)
		}
		//convert form value
		fLat, err := strconv.ParseFloat(formLatitude, 32)
		if err != nil {
			log.Fatal(err)
		}
		//convert form value
		fLnt, err := strconv.ParseFloat(formLongitude, 32)
		if err != nil {
			log.Fatal(err)
		}
		//Prepare data request for API
		searchResto := api.ReqNearbyResto{}
		searchResto.CustomerId = customerId
		searchResto.Latitude = float64(fLat)
		searchResto.Longitude = float64(fLnt)

		//Marshal provided interface into JSON structure
		uj, _ := json.Marshal(searchResto)

		//POSTURI
		var buffer bytes.Buffer
		buffer.WriteString(string(development))
		buffer.WriteString("restaurant")
		url := buffer.String()
		log.Println("url")
		//POST
		req, err := http.NewRequest("POST", url, bytes.NewBuffer(uj))
		req.Header.Set("X-Custom-Header", "myvalue")
		req.Header.Set("Content-Type", "application/json")

		client := &http.Client{}
		resp, err := client.Do(req)
		if err != nil {
			panic(err)
		}
		defer resp.Body.Close()
		body, _ := ioutil.ReadAll(resp.Body)

		jsonErr := json.Unmarshal(body, &resto)

		//Data for location result
		resto.Location = formLocation
		if jsonErr != nil {
			log.Fatal(jsonErr)
		}
	}

	tpl.TPL.ExecuteTemplate(w, "search.gohtml", resto)
}

func (wc WebController) DetailRestaurant(w http.ResponseWriter, r *http.Request) {
	//Grab Restaurant URL
	keys, _ := r.URL.Query()["key"]
	key := keys[0]
	//to improve google indexing use url name static and for search gRPC getrestaurants
	restaurantUrl := string(key)

	// Set up connection to the gRPC server
	conn, err := grpc.Dial(address, grpc.WithInsecure())
	if err != nil {
		log.Fatalf("did not connect: %v", err)
	}
	defer conn.Close()

	// Creates a new RestaurantClient
	client := rs.NewRestaurantClient(conn)

	filter := &rs.RestaurantFilter{Keyword: restaurantUrl}
	//Filter with restauran url for gRPC getrestaurants
	detail := restaurantgrpc.GetRestaurants(client, filter)

	tpl.TPL.ExecuteTemplate(w, "restaurant.gohtml", detail)
}

func (wc WebController) ReservationRestaurantProcess(w http.ResponseWriter, r *http.Request) {
	//Code Here take post input web
	formCustomerId := r.FormValue("customer_id")
	formRestaurantId := r.FormValue("restaurant_id")
	formRestaurantName := r.FormValue("restaurant_name")
	formCustomerName := r.FormValue("customer_name")
	formCustomerPhone := r.FormValue("customer_phone")
	formTotalGuest := r.FormValue("total_guest")
	formDate := r.FormValue("date")
	formTime := r.FormValue("time")
	inputTime := formTime[:len(formTime)-3]

	//string to int
	customerId, err := strconv.Atoi(formCustomerId)
	if err != nil {
		log.Fatal(err)
	}
	restaurantId, err := strconv.Atoi(formRestaurantId)
	if err != nil {
		log.Fatal(err)
	}

	//Prepare data for request reservation
	reserve := api.ReqReservation{}
	reserve.ReservationCustomerId = customerId
	reserve.ReservationCustomerName = formCustomerName
	reserve.ReservationCustomerPhone = formCustomerPhone
	reserve.ReservationRestaurantId = int32(restaurantId)
	reserve.ReservationRestaurantName = formRestaurantName
	reserve.ReservationTotalGuest = formTotalGuest
	reserve.ReservationDateTime = formDate + " " + inputTime + ":00"

	//Marshal provided interface into JSON structure
	uj, _ := json.Marshal(reserve)

	//POSTURI
	var buffer bytes.Buffer
	buffer.WriteString(string(development))
	buffer.WriteString("restaurant/reservation")
	url := buffer.String()

	//POST
	req, err := http.NewRequest("POST", url, bytes.NewBuffer(uj))
	req.Header.Set("X-Custom-Header", "myvalue")
	req.Header.Set("Content-Type", "application/json")

	client := &http.Client{}
	resp, err := client.Do(req)
	if err != nil {
		panic(err)
	}
	defer resp.Body.Close()
	body, _ := ioutil.ReadAll(resp.Body)

	//Response data for final reservation
	reserveResp := api.ReservationResp{}
	jsonErr := json.Unmarshal(body, &reserveResp)
	if jsonErr != nil {
		log.Fatal(jsonErr)
	}

	tpl.TPL.ExecuteTemplate(w, "reserved.gohtml", reserveResp)
}
