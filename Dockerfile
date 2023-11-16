# syntax=docker/dockerfile:1

# Build the application from source
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


FROM scratch AS server
COPY --from=build-server /bin/nico /bin/
EXPOSE 8848

ENTRYPOINT [ "/bin/nico" ]