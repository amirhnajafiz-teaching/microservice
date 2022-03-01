# syntax=docker/dockerfile:1

FROM golang:1.16-alpine

# switch to root
USER root

# set work directory
WORKDIR /app

# copy mod and sum
COPY go.mod go.sum ./

# download dependencies
RUN go mod download

# copy all project files
COPY . ./

# build the exe file
RUN go build -o /go-micro

# run the file
CMD ["/go-micro"]
