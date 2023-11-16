# syntax=docker/dockerfile:1

FROM golang:1.21-alpine AS base
ENV CGO_ENABLED 0
ENV GOOS linux
ENV GOPROXY https://goproxy.cn,direct
ENV GO111MODULE on

WORKDIR /app
COPY go.mod go.sum ./
RUN go mod download
ADD . .

FROM base AS build-server

RUN CGO_ENABLED=0 GOOS=linux go build -o /bin/nico ./cmd/main.go


FROM busybox AS server
COPY --from=build-server /bin/nico /bin/
EXPOSE 9002

ENTRYPOINT [ "/bin/nico" ]