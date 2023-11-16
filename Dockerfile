# syntax=docker/dockerfile:1

# Build the application from source
FROM golang:1.20 AS build-stage

ENV CGO_ENABLED 0
ENV GOOS linux
ENV GOPROXY https://goproxy.cn,direct
ENV GO111MODULE on

WORKDIR /app

COPY go.mod go.sum ./
RUN go mod download

ADD . .

RUN CGO_ENABLED=0 GOOS=linux go build -o bin/nico /cmd/main.go

# Run the tests in the container
FROM build-stage AS run-test-stage
RUN go test -v ./...

# Deploy the application binary into a lean image
FROM centos AS build-release-stage

WORKDIR /

COPY --from=build-stage /app/bin/nico /nico

EXPOSE 8848

USER nonroot:nonroot

ENTRYPOINT ["/nico"]