FROM golang:1.21 AS dependencies
WORKDIR /app
COPY go.mod ./
RUN go mod tidy

FROM golang:1.21 AS base
WORKDIR /app
COPY --from=dependencies /app/go.mod ./
COPY . .

FROM base AS development
CMD ["go", "run", "main.go"]

FROM base AS build
RUN CGO_ENABLED=0 GOOS=linux GOARCH=amd64 go build -o /app/app .

FROM debian:bookworm-slim AS production

WORKDIR /root/
COPY --from=build /app/app ./app
EXPOSE 8080
CMD ["./app"]
