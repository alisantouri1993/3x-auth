FROM golang:1.23

# Destination for Copy
WORKDIR /app

#Copy
COPY go.mod auth.go ./

#Build
RUN CGO_ENABLED=0 GOOS=linux go build

EXPOSE 20000

CMD ["/auth"]
