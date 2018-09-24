package main

import (
	"net/http"

	"github.com/boantp/go-mysql-resto/web"
)

func main() {
	//GET Web Controller instance
	web := web.NewWebController()

	//View GUI input given location
	http.HandleFunc("/", web.NearbyRestaurant)
	http.HandleFunc("/restaurant", web.DetailRestaurant)
	http.HandleFunc("/reservation", web.ReservationRestaurantProcess)

	// add route to serve pictures, css, js
	http.Handle("/public/", http.StripPrefix("/public", http.FileServer(http.Dir("./public"))))
	http.Handle("/favicon.ico", http.NotFoundHandler())
	http.ListenAndServe(":4000", nil)
}
