FROM golang:1.21-alpine AS builder

WORKDIR /build 

# copy and download dependency using go mod 

COPY go.mod go.sum ./
RUN go mod download

# copy the code into the container

COPY . .

