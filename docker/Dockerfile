#FROM golang:alpine AS builder
#WORKDIR /app
#COPY . .
#RUN go mod download
#RUN go build -o ./go-demo ./server.go
#
#
#FROM alpine:latest AS runner
#WORKDIR /app
#COPY --from=builder /app/go-demo .
#COPY .env .
#EXPOSE 8000
#ENTRYPOINT ["./go-demo"]

FROM golang:1.19-alpine

RUN apk update && apk upgrade && apk add gcc

WORKDIR /app

RUN go install github.com/cosmtrek/air@latest

COPY ../go.mod go.sum ./
RUN go mod download

CMD ["air"]
#ENTRYPOINT ["entrypoint.sh"]
