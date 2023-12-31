# build stage
FROM golang:1.19-buster AS build

WORKDIR /app

COPY go.mod go.sum ./

COPY ./config/container.env ./.env

RUN go mod download

COPY ./ ./

RUN go build -o ./app ./cmd/basicserver/main.go

# deploy stage
FROM gcr.io/distroless/base-debian11

WORKDIR /app

COPY --from=build ./app ./

EXPOSE 8080

# distroless specific instruction
USER nonroot:nonroot

ENTRYPOINT ["./app"]

