FROM golang:1.21
WORKDIR /app
COPY . .
RUN go build -o app .
CMD ["./app"]
