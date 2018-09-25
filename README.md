# go-mysql-resto
a set of API for restaurant, gRPC Service, and web client communication with gRPC and API. Tech Stack : Golang, MySQL, HTML, Javascript with REST API and gRPC.

This API serve 3 routes at Nearby Restaurant [/restaurant], Restaurant Detail [/restaurant/{restaurant_id}], Reservation Restaurant [//restaurant/reservation] .
- You can perform fetch all nearby restaurant [POST] from Nearby Restaurant Collection
- You can get detail restaurant [GET] from Restaurant Detail Collection
- And you can make reservation [POST] from Reservation Restaurant.

The gRPC Service serve : Create Restaurant, Get Restaurant

So in the web client when user search restaurant with input location it will call Nearby Restaurant API and serve gRPC create restaurant. and then customer can click restaurant to view restaurant detail with data from gRPC get restaurant. In the reservation client will call Reservation Restaurant API.

The web client show some GUI for Search Restaurant, Restaurant Detail, Reservation Response

See [API Documentation](https://github.com/boantp/go-mysql-resto/blob/master/apiary.apib) on how to use it.

## Directory Structure
```
go-mysql-rest-api
    |--config                   - to initialize template and database
        |--db.go                - for initialize mysql database connection
        |--tpl.go               - for template view configuration
    |--controllers              - to store package controllers
        |--cart.go              - to handle Cart Collection [/cart]
        |--order.go             - to handle Order Collection [/order/{store_id}]
        |--tax_code.go          - to handle Tax Code Collection [/tax_code]
        |--web.go               - to handle Web/Frontend for GUI add cart["/"], GUI view bill ["order_view/:store_id]
    |--docker                   - Dockerfile for Golang at folder web, Dockerfile for MySQL at folder db
        |--db
            |--Dockerfile
        |--web
            |--Dockerfile
    |--models                    - to store package models for object and mysql query
        |--order_details.go      - for table order_details
        |--orders.go             - for total amount, total tax amount, grand total object, and for table orders
        |--tax_code.go           - for table tax_code
    |--mysql_init                - init Table shopee
        |--shopee.sql
    |--templates                 - to store html file for golang *gohtml
    |--apiary.apib               - json file docs from APIARY for API DOCS
    |--database_design.png       - DB structure and explanation
    |--docker-compose.yml        - for docker-compose service and config
    |--main.go                   

  
```

## Setup

**Steps**
1. git clone [git@github.com:boantp/go-mysql-rest-api.git](git@github.com:boantp/go-mysql-rest-api.git)
2. install docker and docker-compose 
3. open terminal and run docker-compose build (service are build), docker-compose up(builds, recreates, attaches to container for service), docker-compose down (stop containers and remove containers etc) See [Docker Documentation](https://docs.docker.com/compose/reference/build/) on how to use it.
4. now your server ready for http:localhost:3000/
