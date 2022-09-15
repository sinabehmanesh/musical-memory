FROM golang:1.16-alpine as build

WORKDIR /app

COPY go.mod .
COPY main.go .


RUN go build main.go

FROM alpine

WORKDIR /app

COPY --from=build /app/main .

EXPOSE 8080

CMD ["./main"]