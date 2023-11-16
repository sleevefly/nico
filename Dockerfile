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

RUN CGO_ENABLED=0 GOOS=linux go build -o /bin/nico ./cmd/main.go


EXPOSE 8848

ENTRYPOINT [ "/bin/nico" ]