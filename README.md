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
go-mysql-resto
    |--api                            - REST API Package
        |--handlers.go                - Handler REST API in API Package
        |--model.go                   - Struct and mysql query in API Package
    |--client-grpc                    - Restaurant gRPC
        |--restauranggrpc/main.go     - RPC Method client to server for Restaurant Service
    |--config                         - config mysql, template
       |--db                          - mysql connection
          |--db.go
       |--tpl                         - template gohtml
          |--tpl.go
    |--deployments                    - deploy for docker compose
        |--docker                     - Dockerfile for Web Client at folder web, Dockerfile for MySQL at folder db, and gRPC
           |--db
              |--Dockerfile
           |--grpc
              |--Dockerfile
           |--web
              |--Dockerfile
    |--mysql_init                     - init Table restaurant
        |--restaurant.sql     
    |--public                         - js, css, pics
    |--restaurant                     - proto restaurant for gRPC
    |--server-grpc                    - serve gRPC
    |--server-http                    - serve REST API
    |--templates                      - to store html file for golang *gohtml
    |--web                            - Handler Web Client
    |--apiary.apib                    - API design and documentation
    |--database_design.png            - DB structure and relational
    |--docker-compose.yml             - for docker-compose service and config
    |--main.go                   

  
```

## Setup

**Steps**
1. git clone [git@github.com:boantp/go-mysql-rest-api.git](git@github.com:boantp/go-mysql-rest-api.git)
2. install docker and docker-compose 
3. open terminal and run docker-compose build (service are build), docker-compose up(builds, recreates, attaches to container for service), docker-compose down (stop containers and remove containers etc) See [Docker Documentation](https://docs.docker.com/compose/reference/build/) on how to use it.
4. now your server ready for http:localhost:3000/
