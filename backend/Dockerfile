FROM golang:1.22.3-alpine

WORKDIR /app

COPY . .
COPY .env .

# Download all the dependencies
RUN go get -d -v ./...


# Build the Go app
RUN go build -o /be

EXPOSE 8080

CMD [ "/be" ]
