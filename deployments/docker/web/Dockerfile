# Telling to use Docker's golang ready image
FROM golang:onbuild
EXPOSE 8080

RUN apt-get update && apt-get install -y mysql-client
RUN go get -u github.com/go-sql-driver/mysql
RUN go get -u github.com/boantp/go-mysql-resto
