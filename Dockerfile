FROM docker.io/golang:1.25 AS go-builder
WORKDIR /usr/app/src
COPY go.mod go.sum ./
RUN go mod download && go mod verify
COPY . .
RUN CGO_ENABLED=0 GOOS=linux GOARCH=amd64 go build -o /usr/app/disbot .

FROM docker.io/alpine:3.23 AS final
WORKDIR /app
RUN mkdir config
COPY --from=go-builder /usr/app/disbot /app/disbot
CMD ["/app/disbot", "--config", "/app/config/config.yaml"]