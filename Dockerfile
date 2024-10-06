FROM golang:1.23 AS build-stage

# Destination for Copy
WORKDIR /app

#Copy Go Deps & Download
COPY go.mod ./
RUN go mod download


#Copy Source Code & Build
COPY *.go .
RUN CGO_ENABLED=0 GOOS=linux go build -o auth

#Build
RUN CGO_ENABLED=0 GOOS=linux go build


FROM alpine:latest AS release-stage

WORKDIR /

COPY --from=build-stage /app/auth /auth

EXPOSE 20000

CMD ["/auth"]
