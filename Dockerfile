FROM docker.io/golang:1.20 AS go-builder
WORKDIR /usr/app/src
COPY go.mod go.sum ./
RUN go mod download && go mod verify
COPY . .
RUN CGO_ENABLED=0 GOOS=linux GOARCH=amd64 go build -o /usr/app/disbot .

FROM docker.io/alpine:3.17 AS final
WORKDIR /app
RUN mkdir config
COPY --from=go-builder /usr/app/disbot /app/disbot
COPY config.yaml ./config
CMD ["/app/disbot", "--config", "./config/config.yaml"]